package main

import (
	"fmt"

	tempconv "./exercise2-1"
)

func main() {
	// 2.1
	c := tempconv.Celsius(0.0)
	fmt.Println(tempconv.CToK(c))
	f := tempconv.Fahrenheit(100.0)
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.FToK(f))

}

