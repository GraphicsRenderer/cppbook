package code

import "fmt"

type CppFile struct {
	Name   string
	Blocks []*CppBlock
}

func (c *CppFile) Execute() error {
	fmt.Println(c.Name)
	return nil
}
