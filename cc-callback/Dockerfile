FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
RUN apk add alpine-sdk
COPY /cc-callback .

# RUN go mod download
RUN go build -o cc-callback

FROM alpine:3.16

WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

COPY --from=builder /app/cc-callback .
COPY --from=builder /app/.env .

EXPOSE 9091 9095

ENTRYPOINT [ "/app/cc-callback" ]
