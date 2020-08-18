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

func comma(s string) string {

	pointMatch, err := regexp.MatchString("\\.", s)
	if err != nil {
		return s + ": " + err.Error()
	}
	signMatch, err := regexp.MatchString("[+-]", s)
	if err != nil {
		return s + ": " + err.Error()
	}

	var splitString []string
	splitString = strings.Split(s, ".")
	var buf bytes.Buffer

	if (signMatch && len(splitString[0][1:]) <= 3) || len(splitString[0]) <= 3 {
		return s
	}

	if signMatch {
		sign := s[0]
		buf.WriteByte(sign)
		s = splitString[0][1:]

		if pointMatch {
			buf = insertCommas(buf, s)
			buf.WriteString("." + splitString[1])
			return buf.String()
		} else {
			buf = insertCommas(buf, s)
			return buf.String()
		}
	} else if pointMatch {
		s = splitString[0]
		buf = insertCommas(buf, s)
		buf.WriteString("." + splitString[1])
		return buf.String()
	}
	buf = insertCommas(buf, s)
	return buf.String()
}

func insertCommas(buf bytes.Buffer, s string) bytes.Buffer {
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
