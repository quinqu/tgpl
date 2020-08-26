package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int
type LineCounter int
type WordCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	var count int
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var count int

	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += LineCounter(count)

	return count, nil
}

func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	// word & line counter
	var wc WordCounter
	fmt.Println(wc.Write([]byte("hello world, this is Go")))

	var lc LineCounter
	fmt.Println(lc.Write([]byte("hello\n world\n this\n is\n go")))

}
