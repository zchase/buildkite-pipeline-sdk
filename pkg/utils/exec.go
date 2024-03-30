package utils

import "os/exec"

func RunCommand(name string, cmds ...string) error {
	cmd := exec.Command(name, cmds...)
	return cmd.Run()
}
