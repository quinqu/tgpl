package main

import (
	"fmt"

	tempconv "./exercise2-1"
)

func main() {
	// 2.1
	c := tempconv.Celsius(0.0)
	fmt.Println("celsius to kelvin", tempconv.CToK(c)) // 273.15
	fmt.Println("celsius to fahrenheit", tempconv.CToF(c)) // 32.0
	f := tempconv.Fahrenheit(100.0)
	fmt.Println("fahrenheit to celsuis", tempconv.FToC(f)) // 37.77
	fmt.Println("fahrenheit to kelvin", tempconv.FToK(f)) // 310.92

	k := tempconv.Kelvin(80.0)
	fmt.Println("kelvin to celsius", tempconv.KToC(k)) // -193.15
	fmt.Println("kelvin to fahrenheit", tempconv.KToF(k)) // -315.67


}

