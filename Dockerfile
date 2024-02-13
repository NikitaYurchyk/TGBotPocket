FROM golang:1.21.6-alpine AS builder

COPY . /github.com/NikitaYurchyk/TGPocket/
WORKDIR /github.com/NikitaYurchyk/TGPocket/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/NikitaYurchyk/TGPocket/bin/bot .

EXPOSE 80

CMD ["./bot"]

