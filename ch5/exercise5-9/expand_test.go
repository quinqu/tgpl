package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpand(t *testing.T) {

	var tests = []struct {
		input  string
		num    uint
		output string
	}{
		{"hi", 4, "hi"},
		{"$hi", 3, "hihihi"},
		{"go $go $go", 2, "go gogo gogo"},
		{"$$ $", 2, "$$"},
		{"$ruby", 1, "ruby"},
		{"", 1, ""},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("PASS: %v", test.output), func(t *testing.T) {
			if got := expand(test.input, timesX(test.num)); got != test.output {
				assert.Equal(t, got, test.output)
			}
		})
	}
}
