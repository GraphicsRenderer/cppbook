package code

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

func Execute(filename string) error {
	inFile, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	mdSource, err := ioutil.ReadAll(inFile)
	if err != nil {
		return err
	}
	mdFilepath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}
	blocks := findCppBlocks(mdFilepath, mdSource)
	cppFiles := findCppFiles(blocks)
	for _, cppFile := range cppFiles {
		err = cppFile.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func findCppBlocks(mdFileName string, mdSource []byte) []*CppBlock {
	node := goldmark.DefaultParser().Parse(text.NewReader(mdSource))
	blocks := []*CppBlock{}
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if block, ok := n.(*ast.FencedCodeBlock); ok {
			cppBlock := &CppBlock{mdFileName, mdSource, block}
			if cppBlock.Language() == "cpp" || cppBlock.Language() == "c++" {
				blocks = append(blocks, cppBlock)
			}
		}
		return ast.WalkContinue, nil
	})
	return blocks
}

func findCppFiles(blocks []*CppBlock) []*CppFile {
	cppFiles := map[string]*CppFile{}
	for _, block := range blocks {
		filename := block.Filepath()
		_, ok := cppFiles[filename]
		if !ok {
			cppFiles[filename] = &CppFile{filename, []*CppBlock{}}
		}
		cppFiles[filename].Blocks = append(cppFiles[filename].Blocks, block)
	}
	vals := make([]*CppFile, 0, len(cppFiles))
	for k := range cppFiles {
		vals = append(vals, cppFiles[k])
	}
	return vals
}
