package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type expandFn func(s string) string

func main() {

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print("could not convert to uint")
	} else {
		times := uint(n)
		input := os.Args[2:]
		s := strings.Join(input, " ")

		out := expand(s, timesX(times))
		fmt.Print(out)
	}
}

func timesX(x uint) expandFn {
	f := func(s string) string {
		var buf bytes.Buffer
		var i uint
		for i = 0; i < x; i++ {
			buf.WriteString(s)
		}
		return buf.String()
	}
	return f
}

func expand(s string, f expandFn) string {

	if len(s) < 1 {
		return s
	}

	elements := strings.Split(s, " ")

	for i := range elements {
		if strings.HasPrefix(elements[i], "$") {

			elements[i] = f(elements[i][1:])

		}
	}
	result := strings.Join(elements, " ")
	return strings.TrimSpace(result)
}
