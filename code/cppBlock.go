package code

import (
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark/ast"
)

type CppBlock struct {
	mdFilepath      string
	mdSource        []byte
	fencedCodeBlock *ast.FencedCodeBlock
}

func (b *CppBlock) Language() string {
	if b.fencedCodeBlock == nil || b.fencedCodeBlock.Info == nil {
		return ""
	}
	hint := string(b.fencedCodeBlock.Info.Text(b.mdSource))
	chunks := strings.Split(hint, ",")
	if len(chunks) < 1 {
		return ""
	}
	return strings.ToLower(strings.TrimSpace(chunks[0]))
}

func (b *CppBlock) Filepath() string {
	if b.fencedCodeBlock == nil || b.fencedCodeBlock.Info == nil {
		return ""
	}
	hint := string(b.fencedCodeBlock.Info.Text(b.mdSource))
	chunks := strings.Split(hint, ",")
	if len(chunks) < 2 {
		return ""
	}
	return strings.ReplaceAll(b.mdFilepath, filepath.Base(b.mdFilepath), strings.TrimSpace(chunks[1]))
}

func (b *CppBlock) Code() string {
	code := []string{}
	for i := 0; i < b.fencedCodeBlock.Lines().Len(); i++ {
		line := b.fencedCodeBlock.Lines().At(i)
		code = append(code, string(line.Value(b.mdSource)))
	}
	return strings.Join(code, "")
}
