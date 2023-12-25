package cmd

import "testing"

func TestGetThumbnailPercent(t *testing.T) {
	cmd1 := NewImageCommand()
	cmd1.Init([]string{"-tp", "0"})

	// If zero passed, do not set
	if toBeSet, _ := cmd1.GetThumbnailPercent(); toBeSet == true {
		t.Errorf("GetThumbnailPercent.toBeSet = %v, expected %v", toBeSet, false)
	}

	// If passed, set
	cmd2 := NewImageCommand()
	cmd2.Init([]string{"-tp", "20"})
	toBeSet, percentage := cmd2.GetThumbnailPercent()
	if toBeSet == false {
		t.Errorf("GetThumbnailPercent.toBeSet = %v, expected %v", toBeSet, true)
	}
	if percentage != "20%" {
		t.Errorf("GetThumbnailPercent.percentage = %v, expected %v", percentage, "30%")
	}

	// If not passed, set default
	defaultPercentage := "70%"
	cmd3 := NewImageCommand()
	cmd3.Init([]string{})
	toBeSet, percentage = cmd3.GetThumbnailPercent()
	if toBeSet == false {
		t.Errorf("GetThumbnailPercent.toBeSet = %v, expected %v", toBeSet, true)
	}
	if percentage != defaultPercentage {
		t.Errorf("GetThumbnailPercent.percentage = %v, expected %v", percentage, defaultPercentage)
	}
}
