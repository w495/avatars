
all: install build

install:
	go get -v ./...

run:
	go run *.go

build:
	go build -o avatars