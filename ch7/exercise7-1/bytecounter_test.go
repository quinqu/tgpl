package main

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input string
		want  WordCounter
	}{
		{"hello world", 2},
		{"Using the idea's from bytecounter", 5},
		{"implement counters for words", 4},
		{"", 0},
	}
	var wc WordCounter

	for _, test := range tests {
		t.Run(fmt.Sprintf("Words: %v Count: %v", test.input, test.want), func(t *testing.T) {
			wordCount, err := wc.Write([]byte(test.input))
			if err != nil {
				t.Fatal("input err")
			}

			if wordCount != int(test.want) {
				t.Errorf("got %v; want %v", wordCount, test.want)

			}

		})
	}

}

func TestLineCounter(t *testing.T) {

	tests := []struct {
		input string
		want  LineCounter
	}{
		{"hello \n world", 2},
		{"Using the idea's\n from \n byte\ncounter\n", 4},
		{"implement\n counters\n for words", 3},
		{"", 0},
	}
	var lc LineCounter

	for _, test := range tests {
		t.Run(fmt.Sprintf("Lines: %v Count: %v", test.input, test.want), func(t *testing.T) {
			lineCount, err := lc.Write([]byte(test.input))
			if err != nil {
				t.Fatal("input err")
			}

			if lineCount != int(test.want) {
				t.Errorf("got %v; want %v", lineCount, test.want)

			}

		})
	}

}
