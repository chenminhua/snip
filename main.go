package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/antlabs/strsim"
)

var (
	rootPath = "./docs/"
	docsPaths map[string]string
)

func loadDocs(path ...string) {
	docsPaths = map[string]string{}
	filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fileName := info.Name()
		docsPaths[strings.TrimSuffix(fileName, filepath.Ext(fileName))] = path

		return nil
	})
}

func main() {
	paths := os.Args[1:]
	rootPath += filepath.Join(paths...)
	loadDocs()
	build()
	// serve()
	// prompt()	
}

func prompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if text == "list" {
			for fn, fp := range docsPaths {
				println(fn, fp)
			}
			continue
		}
		if text == "reload" {
			loadDocs()
			continue
		}
		
		fp, ok := docsPaths[text];
		if !ok {
			fmt.Printf("file %s not found, do you mean %+v\n", text, findTopNMatch(text, 3))
			continue
		}

		res, err := os.ReadFile(fp)
		if err != nil {
			println("read file %s failed with error %s", fp, err)
			continue
		}

		result := markdown.Render(string(res), 80, 0)
		fmt.Println(string(result))
	}
}

func findTopNMatch(text string, n int) []string {
	if n > len(docsPaths) {
		n = len(docsPaths)
	}
	type fileScore struct {
		fileName string
		score float64
	}
	var fss []fileScore
	for k, _ := range docsPaths {
		fss = append(fss, fileScore{k, strsim.Compare(text, k)})
	}
	sort.Slice(fss, func(i, j int) bool {
		return fss[i].score > fss[j].score
	})
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = fss[i].fileName
	}
	return res
}
