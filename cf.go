package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	converter "./exercise2-2"
)

func main() {

	if len(os.Args[1:]) > 0 {
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
		fmt.Println("Enter numbers separated by a space: ")

		inputReader := bufio.NewScanner(os.Stdin)

		inputReader.Scan()
		input := inputReader.Text()
		values := strings.Split(input, " ")

		for i := 0; i < len(values); i++ {
			y, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Println("ERROR:", err)
				fmt.Print(values[i], "was not a valid number!")
			} else {
				convert(y)
			}

		}

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
