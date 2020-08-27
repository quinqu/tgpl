// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"fmt"

	"./tempconv"
	"gopkg.in/alecthomas/kingpin.v2"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	kingpin.Parse()
	fmt.Println(*temp)
}
