package code

import (
	"fmt"
	"strings"
)

type CppFile struct {
	Filepath string
	Blocks   []*CppBlock
}

func (c *CppFile) Execute() error {
	fmt.Println(c.Filepath)
	return nil
}

func (c *CppFile) Code() string {
	code := []string{}
	for _, block := range c.Blocks {
		code = append(code, block.Code())
	}
	return strings.Join(code, "\n")
}
