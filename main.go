package main

import (
	"cppbook/args"
	"cppbook/code"
	"log"
)

func main() {
	args.Parse()
	err := code.Execute(*args.InputFile)
	if err != nil {
		log.Fatalln(err)
	}
}
