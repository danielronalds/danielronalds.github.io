package parsing

import (
	"bytes"
	"fmt"
	"os"

	"github.com/danielronalds/danielronalds.github.io/internal/site"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func ParseMarkdown(path string) (site.HTMLContent, error) {
	markdown, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read markdown file %q: %w", path, err)
	}

	var renderedMarkdown bytes.Buffer
	if err := markdownParser.Convert(markdown, &renderedMarkdown); err != nil {
		return "", fmt.Errorf("render markdown file %q: %w", path, err)
	}

	return site.HTMLContent(renderedMarkdown.String()), nil
}

var markdownParser = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithRendererOptions(html.WithUnsafe()),
)
