package code

import (
	"cppbook/args"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type CppFile struct {
	Filepath string
	Blocks   []*CppBlock
}

func (f *CppFile) Execute() error {
	err := f.SaveAndCompile()
	if err != nil {
		return err
	}
	err = f.Run()
	if err != nil {
		return err
	}
	return nil
}

func (f *CppFile) SaveAndCompile() error {
	err := ioutil.WriteFile(f.Filepath, []byte(f.Code()), 0644)
	if err != nil {
		return err
	}
	command := fmt.Sprintf("%s -o %s %s", *args.CompileCommand, f.OutputPath(), f.Filepath)
	return ExecCmd(command)
}

func (f *CppFile) Run() error {
	return ExecCmd(f.OutputPath())
}

func (f *CppFile) OutputPath() string {
	return strings.ReplaceAll(f.Filepath, filepath.Ext(f.Filepath), ".exe")
}

func (f *CppFile) Code() string {
	code := []string{}
	for _, block := range f.Blocks {
		code = append(code, block.Code())
	}
	return strings.Join(code, "\n")
}
