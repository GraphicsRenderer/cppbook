package args

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Verbose        = kingpin.Flag("verbose", "Verbose mode.").Short('v').Default("false").Bool()
	CompileCommand = kingpin.Flag("compile-command", "The compile command.").Short('c').Default("g++ -Wall -O3").String()
	InputFile      = kingpin.Arg("file", "The input markdown file").Required().String()
)

func Parse() {
	kingpin.Parse()
}
