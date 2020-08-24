package expand

import (
	"strings"
)

func f(x string) string {
	return x + x
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
