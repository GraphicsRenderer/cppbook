build:
	go build -o cppbook.exe

fmt:
	go fmt ./...

test: build
	./cppbook.exe ./example.md

.PHONY: build fmt
