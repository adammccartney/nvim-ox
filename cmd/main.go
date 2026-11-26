// Package main implements ox's Markdown-to-HTML renderer.
package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

//go:embed templates/index.html
var templateFiles embed.FS

var (
	pageTemplate = template.Must(template.ParseFS(templateFiles, "templates/index.html"))
	markdown     = goldmark.New(goldmark.WithExtensions(extension.GFM))
)

type pageData struct {
	Content template.HTML
}

func render(source io.Reader, destination io.Writer) error {
	input, err := io.ReadAll(source)
	if err != nil {
		return fmt.Errorf("read Markdown: %w", err)
	}

	var content bytes.Buffer
	if err := markdown.Convert(input, &content); err != nil {
		return fmt.Errorf("convert Markdown: %w", err)
	}

	data := pageData{
		// Goldmark's default renderer omits raw HTML instead of passing it through.
		Content: template.HTML(content.String()),
	}
	if err := pageTemplate.ExecuteTemplate(destination, "index.html", data); err != nil {
		return fmt.Errorf("execute page template: %w", err)
	}

	return nil
}

func main() {
	if err := render(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "ox: %v\n", err)
		os.Exit(1)
	}
}
