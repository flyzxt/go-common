package common

import (
	"os/exec"
)

func execShellCommand(command string) (bool, error) {
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return true, err
}
