package main

import (
	"fmt"
	"os"
	"strings"
)

// efficient space complexity
func main() {

	fmt.Println(strings.Join(os.Args[1:], " "))

}
