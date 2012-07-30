package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func fileRead(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	for err == nil {
		lines = append(lines, line)
		line, err = r.ReadString('\n')
	}
	if err == io.EOF {
		err = nil
	}
	return
}

type Snippet struct {
	Title    string
	Tags     []string
	Requires []string
	Body     string
}

func newSnippet() *Snippet {
	return &Snippet{
		Tags:     make([]string, 0),
		Requires: make([]string, 0),
	}
}

func main() {
	output := map[string][]*Snippet{}
	snippets, err := ioutil.ReadDir("./snippets")
	if err != nil {
		panic(err)
	}
	for _, f := range snippets {
		snippet := newSnippet()
		lines, err := fileRead("./snippets/" + f.Name())
		if err == nil {
			for _, l := range lines {
				if titleSplit := strings.Split(l, "title:"); len(titleSplit) > 1 {
					// Get title if line contains "title:"
					snippet.Title = strings.TrimSpace(titleSplit[1])
				} else if tagSplit := strings.Split(l, "tags:"); len(tagSplit) > 1 {
					// Get tags if line contains "tags:"
					tagSplitSplit := strings.Split(tagSplit[1], "|")
					for i, t := range tagSplitSplit {
						tagSplitSplit[i] = strings.TrimSpace(t)
					}
					snippet.Tags = tagSplitSplit
				} else if requireSplit := strings.Split(l, "requires:"); len(requireSplit) > 1 {
					fmt.Println(requireSplit)
					requireSplitSplit := strings.Split(requireSplit[1], "|")
					for i, r := range requireSplitSplit {
						requireSplitSplit[i] = strings.TrimSpace(r)
					}
					snippet.Requires = requireSplitSplit
				} else {
					snippet.Body += l
				}
			}
			output["snippets"] = append(output["snippets"], snippet)
		}
	}
	x, e := json.Marshal(output)
	if e == nil {
		fmt.Println(string(x))
	}
}
