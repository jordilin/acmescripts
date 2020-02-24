package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const BASEURL = "https://github.com/trending/"

var projRE = regexp.MustCompile(`<a href="\/[a-zA-z0-9_-]*\/[a-zA-z0-9_-]*"`)

func main() {
	flag.Parse()
	args := flag.Args()
	for _, lang := range args {
		url := BASEURL + lang
		content := get(url)
		sm := projRE.FindAll(content, -1)
		for _, m := range sm {
			// <a href="/golang/go">
			proj := strings.Split(string(m), `"`)
			projPaths := strings.Split(proj[1], `/`)
			// bypass github.com/features/integrations, /about, /site urls
			if projPaths[1] == "features" || projPaths[1] == "about" || projPaths[1] == "site" {
				continue
			}
			fmt.Println("https://github.com" + proj[1])
		}
	}
}

func get(url string) []byte {
	r, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer r.Body.Close()
	data, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal("reading %s: %v", url, e)
	}
	return data
}
