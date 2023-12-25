package cmd

import (
	"testing"
)

// Test a few arguments are being set, some set by default, others by user above
// Not sure how useful this test is, but this is more for practice
func TestNewVideoCommandArgs(t *testing.T) {
	cmd := NewVideoCommand()
	cmd.Init([]string{"-fs", "20", "-o", "processed-file"})

	if cmd.fromSeconds != 20 {
		t.Errorf("VideoCommand.fromSeconds = %v, expected %v", cmd.fromSeconds, 0)
	}

	if cmd.fromMinutes != 0 {
		t.Errorf("VideoCommand.fromMinutes = %v, expected %v", cmd.fromSeconds, 0)
	}

	if cmd.toSeconds != 0 {
		t.Errorf("VideoCommand.toSeconds = %v, expected %v", cmd.fromSeconds, 0)
	}

	if cmd.outputName != "processed-file" {
		t.Errorf("VideoCommand.output = %v, expected %v", cmd.outputName, "processed-file")
	}
}

func TestVideoValidateFlags(t *testing.T) {
	var err error

	cmd1 := NewVideoCommand()
	cmd1.Init([]string{"-fs", "61"})
	err = cmd1.ValidateFlags()
	if err == nil {
		t.Errorf("VideoCommand.validateFlags allowed 61 seconds.")
	}

	cmd2 := NewVideoCommand()
	cmd2.Init([]string{"-o", ""})
	err = cmd2.ValidateFlags()
	if err == nil {
		t.Errorf("VideoCommand.validateFlags allowed empty output.")
	}

	cmd3 := NewVideoCommand()
	cmd3.Init([]string{"-crf", "2000"})
	err = cmd3.ValidateFlags()
	if err == nil {
		t.Errorf("VideoCommand.validateFlags allowed incorrect crf.")
	}

	cmd4 := NewVideoCommand()
	cmd4.Init([]string{})
	err = cmd4.ValidateFlags()
	if err != nil {
		t.Errorf("VideoCommand.validateFlags returned error when no arguments present.")
	}

	cmd5 := NewVideoCommand()
	cmd5.Init([]string{"-o", "new-output-name", "ts", "40", "tm", "1"})
	err = cmd5.ValidateFlags()
	if err != nil {
		t.Errorf("VideoCommand.validateFlags returned error when valid arguments present.")
	}
}
