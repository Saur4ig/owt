.PHONY: run test lint fmtall all build run down deploy re

lint:
	golangci-lint run

test:
	go test -count=1 ./... -cover

fmtall:
	go fmt ./...

all:
	make fmtall
	make lint
	make test

build:
	docker build -t owt .

run:
	docker-compose up -d owt

down:
	docker-compose rm -sf owt

deploy:
	make build
	make run

re:
	make down
	make deploy