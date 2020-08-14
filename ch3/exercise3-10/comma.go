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
	byteSlice := []byte(s)
	var j = 0
	for i := 3; i < len(byteSlice); i += 2 {
		buf.WriteString(s[j : i-1])
		if i-1 < len(byteSlice) && i-2 < len(byteSlice) {
			buf.WriteByte(',')
		}
		j = i - 1
	}

	if j + 1 < len(byteSlice) {
		buf.WriteString(s[j:])
	}

	return buf.String()
}
