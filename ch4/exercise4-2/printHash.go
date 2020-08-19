package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	var s384 = flag.Bool("sha384", false, "generate sha384 hash")
	var s512 = flag.Bool("sha512", false, "generate sha512 hash")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input: ")
	input, _ := reader.ReadString('\n')

	if *s384 {
		fmt.Println("Converted to SHA384")
		fmt.Println(sha512.Sum384([]byte(input)))
	} else if *s512 {
		fmt.Println("Converted to SHA512")
		fmt.Println(sha512.Sum512([]byte(input)))

	} else {
		fmt.Println("Converted to SHA256")
		fmt.Println(sha256.Sum256([]byte(input)))
	}
}
