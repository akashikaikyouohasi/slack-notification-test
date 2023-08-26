.PHONY: build

build:
	cd notification-slack;GOOS=linux GOARCH=amd64 go build -o notification-slack 

deploy:
	sam deploy
