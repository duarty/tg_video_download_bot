package tests

import (
	"testing"

	"dev.azure/duarty/tg_bot/utils"
)

func TestFormatTwitterURL(t *testing.T) {
	testCase := []struct {
		url      string
		expected string
	}{
		{"https://twitter.com/i/status/1666098233427529729", "https://twitter.com/i/status/1666098233427529729"},
		{"https://twitter.com/AnimalBeingBro5/status/1662642435963637760?t=O--xOI7bVtQolxV4gaBLlQ&s=19", "https://twitter.com/AnimalBeingBro5/status/1662642435963637760"},
		{"asdjhasjdbaksbnfasf", ""},
	}

	for _, test := range testCase {
		if utils.FormatTwitterURL(test.url) != test.expected {
			t.Error("FAIL", testCase)
		}

	}
}

func TestIsValidTwitterURL(t *testing.T) {
	testCase := []struct {
		url      string
		expected bool
	}{
		{"https://twitter.com/i/status/1666098233427529729", true},
		{"https://twitter.com/AnimalBeingBro5/status/1662642435963637760?t=O--xOI7bVtQolxV4gaBLlQ&s=19", true},
		{"#$%¨&*()", false},
	}

	for _, test := range testCase {
		isValid, _ := utils.IsValidTwitterURL(test.url)
		if isValid != test.expected {
			t.Error("FAIL", testCase)
		}
	}
}
