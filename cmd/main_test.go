package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRender(t *testing.T) {
	input := strings.Join([]string{
		"# Hello",
		"",
		"Some **strong** text and ~~old text~~.",
		"",
		"| Name | Value |",
		"| ---- | ----- |",
		"| ox   | works |",
		"",
	}, "\n")

	var output bytes.Buffer
	if err := render(strings.NewReader(input), &output); err != nil {
		t.Fatalf("render() error = %v", err)
	}

	page := output.String()
	for _, want := range []string{
		"<!doctype html>",
		"<h1>Hello</h1>",
		"<strong>strong</strong>",
		"<del>old text</del>",
		"<table>",
		"<td>works</td>",
	} {
		if !strings.Contains(page, want) {
			t.Errorf("render() output does not contain %q:\n%s", want, page)
		}
	}

	if strings.Contains(page, "&lt;h1&gt;") {
		t.Errorf("render() escaped Goldmark's HTML:\n%s", page)
	}
}

func TestRenderOmitsRawHTML(t *testing.T) {
	input := "<script>alert('not previewed')</script>\n\n# Safe text\n"

	var output bytes.Buffer
	if err := render(strings.NewReader(input), &output); err != nil {
		t.Fatalf("render() error = %v", err)
	}

	page := output.String()
	if strings.Contains(page, "<script>") {
		t.Errorf("render() included raw script HTML:\n%s", page)
	}
	if !strings.Contains(page, "<h1>Safe text</h1>") {
		t.Errorf("render() did not render Markdown around raw HTML:\n%s", page)
	}
}
