package util

import (
	"os/exec"
	"strings"
)

type shell struct {
}

var Shell shell

func (shell) Exec(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	out, err := cmd.CombinedOutput()
	return strings.Trim(string(out), "\n"), err
}
