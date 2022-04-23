.PHONY: generate clean run help

generate: clean
	go generate ./...

build: clean
	go build -o cobraslash main.go

clean:
	go clean

all: generate
