package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"./links"
)

var depth = flag.Int("depth", 1, "Only URLs reachable by at most depth links will be fetched")

func crawl(url string, n int, wg *sync.WaitGroup) ([]string, int) {
	defer wg.Done()
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
	var wg sync.WaitGroup
	// Add command-line arguments to worklist.
	go func() {
		worklist <- flag.Args()
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {

			for link := range unseenLinks {
				wg.Add(1)
				foundLinks, newN := crawl(link, n, &wg)
				n = newN
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)

	for n < *depth {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

	}
	wg.Wait()

}
