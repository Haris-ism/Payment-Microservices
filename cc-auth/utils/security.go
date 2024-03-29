package utils

import (
	"bytes"
	"cc-auth/constants"
	"cc-auth/hosts/transaction/models"
	"cc-auth/protogen/merchant"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"

	"io"

	"github.com/goccy/go-json"
)

func serializeStruct(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(input)
	if err != nil {
		return nil, err
	}
	
	return buffer.Bytes(), nil
}

func DecryptTransItemRes(req string)(models.DecTransItem,error){
	res:=models.DecTransItem{}
	chanReq:=make(chan string)

	go DecryptFunc(req,chanReq)
	decrypted:=<-chanReq
	// fmt.Println("decrypted:",decrypted)
	err:=json.Unmarshal([]byte(decrypted),&res)
	if err!=nil{
		fmt.Println("err decrypted:",err)
	}

	return res,nil
}

func EncryptTransItemRes(req models.ResponseItems,codes []string)(models.ResponseItems,error){
	chanID:=make(chan string)
	chanName:=make(chan string)
	chanQuantity:=make(chan string)
	chanCC:=make(chan string)
	chanCode:=make(chan string)

	go EncryptFunc(req.ID,chanID)
	go EncryptFunc(req.Name,chanName)
	go EncryptFunc(req.Quantity,chanQuantity)
	go EncryptFunc(req.CC,chanCC)
	go EncryptArray(codes,chanCode)

	req.ID=<-chanID
	req.Name=<-chanName
	req.Quantity=<-chanQuantity
	req.CC=<-chanCC
	req.Code=<-chanCode

	return req,nil
}

func EncryptArray(arr []string,ch chan string){
	serialized, err := serializeStruct(arr)
	if err != nil {
		fmt.Println("err:",err)
	}
	res,err:=EncryptionAES(string(serialized))
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	ch<-res
}
func DecryptArray(arr string,ch chan []string){
	res,err:=DecryptionAES(arr)
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	var result []string
	buffer := bytes.NewBuffer([]byte(res))
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(&result)
	if err != nil {
		fmt.Println("fail decr",err)
	}

	ch<-result
}

func EncryptTransItem(req models.TransactionItems)(models.DecTransactionItems,error){
	res:=models.DecTransactionItems{}
	// fmt.Println("req:",req)
	bytes,err:=json.Marshal(req)
	if err!=nil{
		return res,errors.New(constants.ERROR_DB)
	}
	chanReq:=make(chan string)

	go EncryptFunc(string(bytes),chanReq)

	encrypted:=<-chanReq
	// fmt.Println("encrypted string:",encrypted)
	res.Req=encrypted

	return res,nil
}

func EncryptTransItemGrpc(req models.TransactionItems)(*merchant.ReqTransItemsModel,error){
	res:=merchant.ReqTransItemsModel{}
	// fmt.Println("req:",req)
	bytes,err:=json.Marshal(req)
	if err!=nil{
		return &res,errors.New(constants.ERROR_DB)
	}
	chanReq:=make(chan string)

	go EncryptFunc(string(bytes),chanReq)

	encrypted:=<-chanReq
	// fmt.Println("encrypted string:",encrypted)
	res.Request=encrypted

	return &res,nil
}
func DecryptTransItem(req models.DecTransactionItems)(models.TransactionItems,error){
	res:=models.TransactionItems{}

	chanReq:=make(chan string)

	go DecryptFunc(req.Req,chanReq)

	decrypted:=<-chanReq
	// fmt.Println("decrypted:",decrypted)
	err:=json.Unmarshal([]byte(decrypted),&res)
	if err!=nil{
		fmt.Println("err decrypted:",err)
	}

	return res,nil
}


func EncryptFunc(input string, ch chan string){
	res,err:=EncryptionAES(input)
	if err!=nil{
		fmt.Println("Encrypt Fails")
	}
	ch<-res
}
func DecryptFunc(input string, ch chan string){
	res,err:=DecryptionAES(input)
	if err!=nil{
		fmt.Println("Decrypt Fails")
	}
	ch<-res
}

func EncryptionAES(input string) (string, error) {
	result:=""
	key, _ := hex.DecodeString(GetEnv("KEY"))

	plaintext:=[]byte(input)

	block, err := aes.NewCipher(key)
	if err != nil {
		return result, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return result, err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptionAES(input string) (string, error) {
	result:=""
	key, _ := hex.DecodeString(GetEnv("KEY"))

	ciphertext, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return result, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return result, err
	}

	if len(ciphertext) < aes.BlockSize {
		return result, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	result=string(ciphertext)

	return result, nil
}