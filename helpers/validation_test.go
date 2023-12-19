package helpers

import (
	"testing"
)

func TestIsValidCrf(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{input: 30, expected: true},
		{input: -30, expected: false},
		{input: 0, expected: true},
		{input: 64, expected: false},
		{input: 63, expected: true},
		{input: 100, expected: false},
	}

	for _, d := range testData {
		result := IsValidCrf(d.input)
		if result != d.expected {
			t.Errorf("IsValidCrf(%d) = %v, expected %v", d.input, result, d.expected)
		}
	}
}

func TestIsValidTimeUnit(t *testing.T) {
	testData := []struct {
		input    int
		expected bool
	}{
		{input: 30, expected: true},
		{input: -30, expected: false},
		{input: 0, expected: true},
		{input: 60, expected: false},
		{input: 53, expected: true},
		{input: 100, expected: false},
	}

	for _, d := range testData {
		result := IsValidTimeUnit(d.input)
		if result != d.expected {
			t.Errorf("IsValidTimeUnit(%d) = %v, expected %v", d.input, result, d.expected)
		}
	}
}

func TestIsValidString(t *testing.T) {
	testData := []struct {
		input    string
		expected bool
	}{
		{input: "test", expected: true},
		{input: "asdasfasdasd", expected: true},
		{input: "", expected: false},
	}

	for _, d := range testData {
		result := IsValidString(d.input)
		if result != d.expected {
			t.Errorf("IsValidString(%s) = %v, expected %v", d.input, result, d.expected)
		}
	}
}
