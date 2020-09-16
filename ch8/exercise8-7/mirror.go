package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

type Page struct {
	links []string
}

func (p *Page) listUrls() []string {
	return p.links
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func save(urlName string, p *Page) error {
	r, err := url.Parse(urlName)
	if err != nil {
		return err
	}

	path := r.Path
	if path == "" {
		path = "index.html"
	}

	path = filepath.Join("./output", path)

	dir := filepath.Dir(path)

	os.MkdirAll(dir, 0700)

	file, err := os.Create(path)

	if err != nil {
		return err
	}
	defer file.Close()
	for _, link := range p.links {
		_, err := fmt.Fprintln(file, link)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadPage(link string) (*Page, error) {

	r, err := url.Parse(link)

	if err != nil {
		return nil, err
	}
	domain := r.Hostname()
	list, err := links.Extract(link)
	var validList []string
	if err != nil {
		return nil, err
	}

	for _, link := range list {
		r, err := url.Parse(link)
		if err != nil {
			return nil, err
		}
		if domain == r.Hostname() {
			validList = append(validList, link)
		}

	}
	p := &Page{links: validList}

	err = save(link, p)

	if err != nil {
		return nil, err
	}
	return &Page{links: validList}, nil
}

type Result struct {
	curr string
	list []string
	e    error
}

func main() {
	input := os.Args[1]
	err := os.MkdirAll("./output", 0700)
	if err != nil {
		log.Fatal("Could not create directory")
	}
	worklist := make(chan string)
	results := make(chan Result)

	for i := 0; i < 10; i++ {
		go func() {
			for url := range worklist {
				fmt.Println("downloading page: ", url)
				page, err := downloadPage(url)
				var urls []string
				if err == nil {
					urls = page.listUrls()
				}
				results <- Result{curr: url, list: urls, e: err}
			}
		}()
	}

	seenList := make(map[string]bool)
	list := []string{}
	nextUrl := input
	worklist2 := worklist
	inProgress := []string{nextUrl}

	for {
		select {
		case res := <-results:
			for i, url := range inProgress {

				if url == res.curr {
					inProgress = append(inProgress[:i], inProgress[i+1:]...)
				}
			}
			if res.e != nil {
				log.Println(res.e)
				continue
			}
			for _, url := range res.list {

				if !seenList[url] {
					fmt.Printf("adding %v to list \n", url)
					seenList[url] = true
					list = append(list, url)
				}
			}

			if len(list) > 0 {
				if nextUrl == "" {
					nextUrl = list[0]
					worklist2 = worklist
					list = list[1:]
					inProgress = append(inProgress, nextUrl)
				}
			} else {
				worklist2 = nil
				nextUrl = ""
			}
		case worklist2 <- nextUrl:
			if len(list) > 0 {
				nextUrl = list[0]
				worklist2 = worklist
				list = list[1:]
				inProgress = append(inProgress, nextUrl)
			} else {
				worklist2 = nil
				nextUrl = ""
			}

		}

		if len(list) == 0 && len(inProgress) == 0 {
			return
		}

	}

}
