package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	title string
	host  string
}

func main() {
	clockServers := []clock{}
	times := os.Args[1:]
	for _, cl := range times {
		attrs := strings.Split(cl, "=")
		newClock := clock{attrs[0], attrs[1]}
		clockServers = append(clockServers, newClock)
	}
	for _, c := range clockServers {
		conn, err := net.Dial("tcp", c.host) //connect to server that was created by clock
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close() // wait until goroutine has finished
		go c.printTimes(conn)
	}

	for {
		time.Sleep(time.Second * 30)
	}
}

func (c clock) printTimes(conn io.Reader) {
	out := os.Stdout
	scanner := bufio.NewScanner(conn) //scans the in the output from server in clock
	for scanner.Scan() {
		output, err := out.WriteString(c.title + ": " + scanner.Text())
		if err != nil {
			log.Fatal("Could not write time: ", err)
		}
		log.Println(output)

	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
