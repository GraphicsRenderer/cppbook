package code

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func ExecCmd(command string) error {
	var re = regexp.MustCompile(`\s+`)
	command = re.ReplaceAllString(command, " ")
	chunks := strings.Split(command, " ")
	if len(chunks) < 1 {
		return errors.New("empty command is not allowed")
	}
	fmt.Println(strings.Join(chunks, " "))
	args := []string{}
	for i := 1; i < len(chunks); i++ {
		args = append(args, chunks[i])
	}
	cmd := exec.Command(chunks[0], args...)
	return cmd.Run()
}
