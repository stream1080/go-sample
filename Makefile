SHELL := /bin/bash

all: dev

dev: 
	swag init
	go run main.go

