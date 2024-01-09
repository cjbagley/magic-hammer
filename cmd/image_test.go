package cmd

import "testing"

func TestImageValidateFlags(t *testing.T) {
	var err error

	cmd1 := NewImageCommand()
	cmd1.Init([]string{"-tp", "101"})
	err = cmd1.ValidateFlags()
	if err == nil {
		t.Error("ImageCommand.validateFlags allowed thumbnail percent > 101%.")
	}

	cmd2 := NewImageCommand()
	cmd2.Init([]string{"-f", ""})
	err = cmd2.ValidateFlags()
	if err == nil {
		t.Errorf("ImageCommand.validateFlags allowed empty input filename.")
	}

	cmd3 := NewImageCommand()
	cmd3.Init([]string{"-q", "200"})
	err = cmd3.ValidateFlags()
	if err == nil {
		t.Errorf("ImageCommand.validateFlags allowed image quality > 100.")
	}

	cmd4 := NewImageCommand()
	cmd4.Init([]string{})
	err = cmd4.ValidateFlags()
	if err != nil {
		t.Errorf("ImageCommand.validateFlags returned error when no arguments present.")
	}

	cmd5 := NewImageCommand()
	cmd5.Init([]string{"-tp", "40", "-q", "1"})
	err = cmd5.ValidateFlags()
	if err != nil {
		t.Errorf("ImageCommand.validateFlags returned error when valid arguments present.")
	}
}
