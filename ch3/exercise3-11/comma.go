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

	match, _ := regexp.MatchString("\\.", s)
	signMatch, _ := regexp.MatchString("[+-]", s)
	var splitString []string
	splitString = strings.Split(s, ".")
	var buf bytes.Buffer

	if (signMatch && len(splitString[0][1:]) <= 3) || len(splitString[0]) <= 3{
		return s
	} 

	if match && signMatch {
		sign := splitString[0][0]
		buf.WriteByte(sign)
		buf = insertCommas(buf, splitString[0][1:])
		buf.WriteString("." + splitString[1])
		return buf.String()

	} else if match == false && signMatch {
		sign := splitString[0][0]
		buf.WriteByte(sign)
		buf = insertCommas(buf, splitString[0][1:])
		return buf.String()

	} else if match {
		buf = insertCommas(buf, splitString[0])
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
