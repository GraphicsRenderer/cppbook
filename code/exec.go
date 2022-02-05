package code

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func ExecCmd(command string) error {
	var re = regexp.MustCompile(`\s+`)
	command = re.ReplaceAllString(command, " ")
	chunks := strings.Split(command, " ")
	if len(chunks) < 1 {
		return errors.New("empty command is not allowed")
	}
	color.Set(color.FgMagenta)
	fmt.Println(strings.Join(chunks, " "))
	color.Unset()
	args := []string{}
	for i := 1; i < len(chunks); i++ {
		args = append(args, chunks[i])
	}
	cmd := exec.Command(chunks[0], args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
