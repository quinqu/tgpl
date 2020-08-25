package expand

import (
	"fmt"
	"os"
	"testing"
)

func TestExpand(t *testing.T) {
	os.Args = []string{"idk", "so", "random"}

	var tests = []struct {
		input  string
		output string
	}{
		{"$foo $bar $hello$hi", "fooisr barisr hello$hiisr"},
		{"testing $123 testing$", "testing 123isr testing$"},
		{"go $go $go asd $fghjkl", "go goisr goisr asd fghjklisr"},
		{"$ $ $", "isr isr isr"},
		{"", ""},
		{"$helloworld", "helloworldisr"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS: %v", test.output), func(t *testing.T) {
		if got := expand(test.input, f); got != test.output {
			t.Errorf("Expected: %q = %v", test.input, test.output)
			t.Errorf("Actual: %q = %v", test.input, got)
		}
	})
	}
}

func TestExpandNoArgs(t *testing.T) {
	os.Args = []string{}

	var tests = []struct {
		input  string
		output string
	}{
		{"$hello world", "hello world"},
		{"testing $testing$", "testing testing$"},
		{"go $go $go", "go go go"},
		{"$", ""},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS: %v from input: %v", test.output, test.input), func(t *testing.T) {
			if got := expand(test.input, f); got != test.output {
				t.Errorf("Expected: %q = %v", test.input, test.output)
				t.Errorf("Actual: %q = %v", test.input, got)
			}
		})
	}
}
