package cmd

import (
	"testing"
)

func TestNewVideoCommand(t *testing.T) {
	cmd := NewVideoCommand()
	cmd.Init([]string{"-fs", "20"})

	if cmd.fromSeconds != 20 {
		t.Errorf("VideoCommand.fromSeconds = %v, expected %v", cmd.fromSeconds, 0)
	}

	if cmd.fromMinutes != 0 {
		t.Errorf("VideoCommand.fromMinutes = %v, expected %v", cmd.fromSeconds, 0)
	}

	if cmd.toSeconds != 0 {
		t.Errorf("VideoCommand.toSeconds = %v, expected %v", cmd.fromSeconds, 0)
	}
}
