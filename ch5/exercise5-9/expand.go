package expand

import (
	"bytes"
	"os"
	"strings"
)

var f = func(x string) string {

	args := os.Args[1:]

	if len(args) == 0 {
		return x
	}
	var buf bytes.Buffer

	for i := 0; i < len(args); i++ {
		buf.WriteByte(args[i][0])
	}
	return x + buf.String()
}

func expand(s string, f func(string) string) string {

	if len(s) < 1 {
		return s
	}

	elements := strings.Split(s, " ")

	for i := range elements {
		if strings.HasPrefix(elements[i], "$") {
			if len(elements[i]) == 1 {
				elements[i] = f(elements[i])
			} else {
				elements[i] = f(elements[i][1:])
			}
		}
	}

	return strings.Join(elements, " ")
}
