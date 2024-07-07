package main

import (
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
)

func mdToHTML(files []string) ([]byte, []byte) {
	// read and store markdown file in 'mdContent'
	var contents []byte
	for i := range files {
		readData, err := os.ReadFile(files[i])
		if err != nil {
			fmt.Println(err)
		} else {
			contents = append(contents, readData...)
		}
	}

	// Create a parser and parse the md stored as bytes
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(contents)

	// Create a html renderer and from the parsed markdown, render html
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer), contents
}
