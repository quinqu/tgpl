package main

import (
	"fmt"
	"os"
	"strconv"

	converter "./exercise2-2"
)

func main() {

	if len(os.Args) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprint(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}

			convert(t)

		}
	} else {

		//use standard input

	}

}

func convert(input float64) {

	//temp
	f := converter.Fahrenheit(input)
	c := converter.Celsius(input)
	fmt.Printf("%s = %s, %s = %s \n", f, converter.FToC(f), c, converter.CToF(c))

	//weight

	pounds := converter.Pounds(input)
	kilograms := converter.Kilograms(input)
	fmt.Printf("%s = %s, %s = %s \n", pounds, converter.PoundsToKilograms(pounds), kilograms, converter.KilogramsToPounds(kilograms))

	//length
	feet := converter.Feet(input)
	meters := converter.Meters(input)
	fmt.Printf("%s = %s, %s = %s \n", feet, converter.FeetToMeters(feet), meters, converter.MetersToFeet(meters))

}
