package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"./links"
)

var maxDepth = flag.Int("depth", 2, "Only URLs reachable by at most depth links will be fetched")
var seen = make(map[string]bool)
var mux = sync.Mutex{}
var sem = make(chan struct{}, 20)

// semaphore is accessed by 20 processes at a time
// constrains access to at most 20 routines (holds 20 structs)
// used to bound concurrency

func crawl(url string, currDepth int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(url)

	if currDepth >= *maxDepth {
		return
	}
	// pg 229
	sem <- struct{}{}
	list, err := links.Extract(url)
	<-sem

	if err != nil {
		log.Print(err)
	}

	for _, l := range list {
		mux.Lock()
		if seen[l] {
			mux.Unlock()
			continue
		}
		seen[l] = true
		mux.Unlock()
		wg.Add(1)
		go crawl(l, currDepth+1, wg)
	}

}

func main() {
	flag.Parse()
	wg := &sync.WaitGroup{}
	//wg - wait until all go routines are finished to move on
	wg.Add(1)
	go crawl(os.Args[len(os.Args)-1], 0, wg)
	wg.Wait()
}
