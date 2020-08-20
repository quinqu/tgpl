package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"encoding/hex"
)

func main() {

	var s384 = flag.Bool("sha384", false, "generate sha384 hash")
	var s512 = flag.Bool("sha512", false, "generate sha512 hash")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input: ")
	input, _ := reader.ReadString('\n')

	var b []byte
	if  *s384{
		fmt.Println("Converted to SHA384")
		conv := sha512.Sum384([]byte(input))
		b = conv[:]
		fmt.Println(hex.EncodeToString(b))
	} else if *s512 {
		fmt.Println("Converted to SHA512")
		conv := sha512.Sum512([]byte(input))
		b = conv[:]
		fmt.Println(hex.EncodeToString(b))
	} else {
		fmt.Println("Converted to SHA256")
		conv := sha256.Sum256([]byte(input))
		b = conv[:]
		fmt.Println(hex.EncodeToString(b))
	}
}
