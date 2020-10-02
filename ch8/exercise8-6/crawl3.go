package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

var maxDepth = flag.Int("depth", 2, "Only URLs reachable by at most depth links will be fetched")

type job struct {
	link  string
	depth int
}

type result struct {
	links []string
	depth int
}

func main() {
	flag.Parse()

	var n int
	var input []string
	if strings.Contains(os.Args[1], "depth") {
		input = os.Args[2:]
	} else {
		input = os.Args[1:]
	}

	worklist := make(chan result) // lists of URLs, may have duplicates
	unseenLinks := make(chan job) // de-duplicated URLs
	seen := make(map[string]bool)

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for job := range unseenLinks {
				foundLinks, err := links.Extract(job.link)
				if err != nil {
					log.Println(err)
				}
				go func(depth int) {
					worklist <- result{links: foundLinks, depth: depth + 1}
				}(job.depth)

			}
		}()
	}

	n++
	// Add command-line arguments to worklist.
	go func() { worklist <- result{links: input, depth: 0} }()

	for ; n > 0; n-- {
		res := <-worklist
		if res.depth >= *maxDepth {
			continue
		}
		for _, l := range res.links {
			if !seen[l] {
				seen[l] = true
				n++
				log.Println(l)
				unseenLinks <- job{link: l, depth: res.depth}
			}
		}

	}
}
