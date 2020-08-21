package main 


import (
	"strings"
	"fmt"
)
const (
	FOO = "$foo"
)

func main() {
	fmt.Println(expand("$foo $foo $fo", f))
}

func expand(s string, f func(string) string) string {
	elements := strings.Split(s, " ")

	for i := range elements {
		if elements[i] == FOO {
			elements[i] = f(elements[i][1:])
		}
	}

	return strings.Join(elements, " ")
}

func f (s string) string {
	return s[0:1]
}