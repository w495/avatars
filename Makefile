
all: install build

install:
	go get -v ./...

run:
	go run main.go

build:
	go run main.go