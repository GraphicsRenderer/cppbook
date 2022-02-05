package args

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Verbose   = kingpin.Flag("verbose", "Verbose mode.").Short('v').Default("false").Bool()
	InputFile = kingpin.Arg("file", "The input markdown file").Required().String()
)

func Parse() {
	kingpin.Parse()
}
