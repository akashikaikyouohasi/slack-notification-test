.PHONY: build

build:
	cd notification-slack;GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

deploy:
	sam deploy
