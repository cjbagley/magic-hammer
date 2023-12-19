package helpers

func IsValidCrf(crf int) bool {
	return crf >= 0 && crf <= 63
}

func IsValidTimeUnit(unit int) bool {
	return unit >= 0 && unit <= 59
}

func IsValidString(str string) bool {
	return str != ""
}
