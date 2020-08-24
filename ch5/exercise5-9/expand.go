package main 


import (
	"strings"
	"fmt"
)
const (
	FOO = "$foo"
)

func main() {
	fmt.Println(expand("$foo $foo $fo", func(x string)string {
		return "random"
	})) 

	fmt.Println(expand("", func(x string)string {
		return "random"
	}))
	
	fmt.Println(expand("$hello, $foo", func(x string)string {
		return "hello"
	}))

}

func expand(s string, f func(string) string) string {
	if len(s) < 1 {
		return s 
	}

	elements := strings.Split(s, " ")

	for i := range elements {
		if elements[i] == FOO {
			elements[i] = f(elements[i][1:])
		}
	}

	return strings.Join(elements, " ")
}

