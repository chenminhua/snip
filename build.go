package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func build() {
	loadDocs()
	currentPath, _ := os.Getwd()
	publicPath := filepath.Join(currentPath, "public")
	os.RemoveAll(publicPath)
	for _, path := range docsPaths {
		res, err := os.ReadFile(path)
		if err != nil {
			println("read file %s failed with error %s", path, err)
			continue
		}
		html := mdToHTML(res)
		rfp := publicPath + "/" + filepath.Join(strings.Split(path, "/")[1:]...)
		rfp = strings.TrimSuffix(rfp, filepath.Ext(rfp)) + ".html"
		ensureDirExist(rfp)
		err = os.WriteFile(rfp, html, 0644)
	}
}

func mdToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}

func ensureDirExist(path string) error {
	ppp := strings.Split(path, "/")
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		return os.MkdirAll("/" + filepath.Join(ppp[1:len(ppp)-1]...), 0755)
	}
	return nil 
}
