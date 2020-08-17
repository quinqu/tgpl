package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {

	match, _ := regexp.MatchString("\\.", s) //true or false

	var splitString []string
	splitString = strings.Split(s, ".")

	if len(splitString[0]) <= 3 {
		return s
	}

	if match {
		buf := insertCommas(splitString[0])
		buf.WriteString("." + splitString[1])
		return buf.String()
	}

	buf := insertCommas(s)
	return buf.String()
}

func insertCommas(s string) bytes.Buffer {

	var buf bytes.Buffer
	n := len(s)

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
	return buf
}
