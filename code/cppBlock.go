package code

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

type CppBlock struct {
	mdSource        []byte
	fencedCodeBlock *ast.FencedCodeBlock
}

func (b *CppBlock) Language() string {
	hint := string(b.fencedCodeBlock.Info.Text(b.mdSource))
	chunks := strings.Split(hint, ",")
	if len(chunks) < 1 {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(chunks[0]))
}

func (b *CppBlock) FileName() string {
	hint := string(b.fencedCodeBlock.Info.Text(b.mdSource))
	chunks := strings.Split(hint, ",")
	if len(chunks) < 1 {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(chunks[0]))
}
