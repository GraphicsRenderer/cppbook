build:
	GOOS=linux GOARCH=amd64 go build -o cppbook-linux-amd64
	GOOS=windows GOARCH=amd64 go build -o cppbook-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 go build -o cppbook-darwin-amd64
	go build -o cppbook.exe

fmt:
	go fmt ./...

test: build
	./cppbook.exe ./example.md

.PHONY: build fmt
