.PHONY: build run build-image start-container

.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t tgpocket:v0.1 .

start-container: build-image
	docker run -p 80:80 tgpocket:v0.1