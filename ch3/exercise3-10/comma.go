package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	offset := n % 3
	if offset > 0 {
		buf.WriteString(s[0:offset])
		buf.WriteByte(',')
	}
	s = s[offset:]

	for i := 0; i < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		if i < len(s)-3 {
			buf.WriteByte(',')
		}

	}
	return buf.String()
}
