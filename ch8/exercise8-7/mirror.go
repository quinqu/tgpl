package main

import (
	"bytes"

	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

type Page struct {
	links []string
}

func savePage(resp *http.Response, body io.Reader) error {
	u := resp.Request.URL
	filename := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
		filename = filepath.Join(u.Host, u.Path, "index.html")
	}
	err := os.MkdirAll(filepath.Dir(filename), 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	if body != nil {
		_, err = io.Copy(file, body)
	} else {
		_, err = io.Copy(file, resp.Body)
	}
	if err != nil {
		log.Print("save: ", err)
	}
	err = file.Close()
	if err != nil {
		log.Print("save: ", err)
	}
	return nil
}

func downloadPage(link string) (*Page, error) {
	var validList []string
	fmt.Println(link)
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: %s", link, resp.Status)
	}

	u, err := base.Parse(link)
	if err != nil {
		return nil, err
	}
	if base.Host != u.Host {
		log.Printf("not saving %s: non-local", link)
		return nil, nil
	}

	var body io.Reader
	contentType := resp.Header["Content-Type"]
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
		}
		nodes := linkNodes(doc)
		validList = linkURLs(nodes, u)
		rewriteLinks(nodes, u)
		b := &bytes.Buffer{}
		err = html.Render(b, doc)
		if err != nil {
			log.Printf("render %s: %s", u, err)
		}
		body = b
	}

	err = savePage(resp, body)
	return &Page{links: validList}, nil
}

func rewriteLinks(linkNodes []*html.Node, base *url.URL) {
	for _, n := range linkNodes {
		for i, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil || link.Host != base.Host {
				continue
			}
			link.Scheme = ""
			link.Host = ""
			link.User = nil
			y, err := url.ParseRequestURI(link.String())
			a.Val = absolutePath + "/" + mirror + y.String()[1:] + "index.html"
			n.Attr[i] = a
		}
	}
}

func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}
	}
	forEachNode(n, visitNode, nil)
	return links
}

//chapter 5 code
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func linkURLs(linkNodes []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range linkNodes {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil {
				continue
			}
			if link.Host != base.Host {
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}

var base *url.URL
var absolutePath string
var mirror string

func main() {
	input := os.Args[1]
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	absolutePath = dir
	mirror = strings.Split(input, "//")[1]

	u, err := url.Parse(input)
	if err != nil {
		log.Println(err)
	}

	base = u
	worklist := make(chan []string)
	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				urls, _ := downloadPage(link)
				go func(link string) {
					worklist <- urls.links
				}(link)
			}
		}
	}
}
