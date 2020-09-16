package main

import (
	"flag"
	"fmt"
	"log"

	"./links"
)

var depth = flag.Int("depth", 1, "Only URLs reachable by at most depth links will be fetched")

func crawl(url string, n int) ([]string, int) {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	n++
	return list, n
}

//!+
func main() {
	flag.Parse()
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	var n int
	// Add command-line arguments to worklist.
	go func() {
		worklist <- flag.Args()
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {

			for link := range unseenLinks {

				foundLinks, newN := crawl(link, n)
				n = newN
				go func() {

					worklist <- foundLinks
				}()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)

	for n < *depth {
		list := <-worklist
		for _, link := range list { //here worklist is a channel
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

	}

	if n >= *depth {
		close(unseenLinks)
	}

}
