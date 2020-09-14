package main

import (
	"io"
	"log"
	"net"
	"time"
	"flag"
	
)
//creating servers 

func handleConn(c net.Conn, timezone string) {
	defer c.Close()
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return 
	}
	
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

var p = flag.String("port", "8000", "desired port number")
var TZ = flag.String("timezone", "US/Eastern", "examples: US/Eastern, Asia/Tokyo, Europe/London")
func main() {
	
	flag.Parse()
	host := "localhost:" + *p
	timezone := *TZ
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}

	
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, timezone) // handle connections concurrently
	}
	//!-
}


