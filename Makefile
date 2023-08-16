.PHONY: all dev build test tool clean help

PROJECT=go-sample

all: dev

dev: 
	go run main.go

fmt:
	gofmt -w .

vet:
	go vet ./...

build: 
	go build -o ${PROJECT}

test:
	go test ./...

clean:
	rm -rf ${PROJECT}
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make dev: run go project"
	@echo "make tool: run specified go tool"
	@echo "make test: run go test"
	@echo "make clean: remove object files and cached files"