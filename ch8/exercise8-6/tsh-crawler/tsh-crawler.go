package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

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

var directory = kingpin.Flag("directory", "directory to crawl").Required().String()
var sshProxy = kingpin.Flag("proxy", "host:port the SSH proxy can be accessed at.").Required().String()
var userAndHost = kingpin.Arg("user@host", "user and host of the remote server").Required().String()

func main() {
	kingpin.Parse()
	logrus.SetLevel(logrus.ErrorLevel)
	Crawl([]string{*directory})
}

func parseDirectories(directories string) []string {
	dirs := strings.Split(directories, "\n")
	var newDirs []string

	for i := range dirs {
		if dirs[i] == "" {
			continue

		}
		newDirs = append(newDirs, dirs[i][0:len(dirs[i])-1])
	}
	return newDirs
}

func Crawl(dirs []string) {

	var n int
	input := dirs
	worklist := make(chan result)
	unseenDirs := make(chan job)
	seen := make(map[string]bool)
	var wg sync.WaitGroup

	wg.Add(1)
	for i := 0; i < 20; i++ {
		go func() {
			for job := range unseenDirs {
				foundDirs := Extract(job.dir)

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
				log.Println(l)
				unseenDirs <- job{dir: l, depth: res.depth}
			}

		}
	}

	wg.Done()
	wg.Wait()
}

func Extract(root string) []string {

	userHost := strings.Split(*userAndHost, "@")
	if len(userHost) != 2 {
		log.Fatalf("invalid user and host")
	}
	user := userHost[0]
	host := userHost[1]

	c := teleport.Config{}

	path := teleport.FullProfilePath("")
	c.LoadProfile(path, *sshProxy)
	c.SSHProxyAddr = *sshProxy
	c.Host = host
	c.Username = user
	c.HostLogin = user
	out := &bytes.Buffer{}

	c.Stdout = out
	c.Stderr = ioutil.Discard

	dirLister, err := teleport.NewClient(&c)
	if err != nil {
		log.Fatalf("can't create new client %v", err)
	}

	err = dirLister.SSH(context.TODO(), []string{"cd " + root + " && ls -d */"}, false)

	if err != nil {
		if out.String() == "" {
			return nil
		}
		log.Fatalf("ssh error: %v", err)
	}

	dirs := parseDirectories(out.String())
	for i := range dirs {
		dirs[i] = root + "/" + dirs[i]
	}
	return dirs
}
