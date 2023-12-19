package helpers

import (
	"os"
	"os/exec"
	"testing"
)

func TestExit(t *testing.T) {
	if os.Getenv("TEST_EXIT") == "1" {
		Exit("Test")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestExit")
	cmd.Env = append(cmd.Env, "TEST_EXIT=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}

	t.Errorf("Exit() did not exit")
}
