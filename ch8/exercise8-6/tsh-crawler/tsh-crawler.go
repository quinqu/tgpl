package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"

	teleport "github.com/gravitational/teleport/lib/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

type job struct {
	dir   string
	depth int
}

type result struct {
	dirs  []string
	depth int
}

type Buffer struct {
	b bytes.Buffer
	m sync.Mutex
}

func (b *Buffer) Write(p string) (n int, err error) {
	b.m.Lock()
	defer b.m.Unlock()
	return b.b.WriteString(p)
}

var directory = kingpin.Flag("directory", "directory to crawl").Required().String()
var user = kingpin.Flag("user", "Username is the Teleport account username").String()
var sshProxy = kingpin.Flag("proxy", "host:port the SSH proxy can be accessed at.").Required().String()

func main() {
	kingpin.Parse()

	c := teleport.Config{}

	path := teleport.FullProfilePath("")
	c.LoadProfile(path, *sshProxy)
	c.SSHProxyAddr = *sshProxy

	if *user != "" {
		c.Username = *user
		c.HostLogin = *user
	}
	dirLister, err := teleport.NewClient(&c)
	if err != nil {
		log.Fatalf("can't create new client %v", err)
	}
	newCrawler := crawlerClient{DirLister: dirLister}
	newCrawler.Crawler(*directory, *newCrawler.DirLister)
}

type crawlerClient struct {
	DirLister *teleport.TeleportClient
}

func (tc crawlerClient) Crawler(dir string, client teleport.TeleportClient) {

	out := &bytes.Buffer{}
	client.Stdout = out
	client.Stderr = out

	var n int
	input := []string{*directory}
	worklist := make(chan result)
	unseenDirs := make(chan job)
	seen := make(map[string]bool)
	var wg sync.WaitGroup

	wg.Add(1)
	for i := 0; i < 20; i++ {
		go func() {
			for job := range unseenDirs {
				foundDirs, err := Extract(job.dir)
				if err != nil {
					out.WriteString(err.Error())
				}

				go func(depth int) {
					worklist <- result{dirs: foundDirs, depth: depth + 1}
				}(job.depth)

			}
		}()
	}

	n++

	go func() {
		worklist <- result{dirs: input, depth: 0}
	}()
	for ; n > 0; n-- {
		res := <-worklist
		for _, l := range res.dirs {
			if !seen[l] {
				seen[l] = true
				n++
				_, err := out.WriteString(l)
				if err != nil {
					log.Println(err)
				}
				unseenDirs <- job{dir: l, depth: res.depth}
			}

		}
	}

	wg.Done()
	wg.Wait()
}

func Extract(root string) ([]string, error) {
	var dirs []string

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		path := root + "/" + f.Name()
		segments := strings.Split(path, string(filepath.Separator))
		log.Println(segments)
		if f.IsDir() {
			dirs = append(dirs, path)
		}
	}

	return dirs, nil
}
