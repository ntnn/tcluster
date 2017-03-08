package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

func cmd(cmds []string) error {
	return exec.Command("tmux", cmds...).Run()
}

// Function SendKeys executes tmux send-keys and ends with a newline
func SendKeys(s string) error {
	args := fmt.Sprintf("send-keys %s\n", s)
	return cmd(strings.Split(args, " "))
}
