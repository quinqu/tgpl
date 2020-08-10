package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprint(os.Std.err, "cf: %v\n", err)
			os.Exit
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}
