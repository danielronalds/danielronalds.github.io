package generation

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/danielronalds/danielronalds.github.io/internal/site"
	"github.com/danielronalds/danielronalds.github.io/internal/view"
)

type MarkdownParser func(path string) (site.HTMLContent, error)

func Generate(destDir string, routes []site.Route, parseMarkdown MarkdownParser) error {
	if err := os.RemoveAll(destDir); err != nil {
		return fmt.Errorf("clean destination directory %q: %w", destDir, err)
	}

	for _, route := range routes {
		content, err := parseMarkdown(route.Src)
		if err != nil {
			return fmt.Errorf("parse markdown source %q: %w", route.Src, err)
		}

		var renderedPage bytes.Buffer
		if err := view.Render(&renderedPage, content); err != nil {
			return fmt.Errorf("render source %q for destination %q: %w", route.Src, route.Dest, err)
		}

		destinationDirectory := filepath.Dir(route.Dest)
		if err := os.MkdirAll(destinationDirectory, 0o755); err != nil {
			return fmt.Errorf("create destination directory %q: %w", destinationDirectory, err)
		}

		if err := os.WriteFile(route.Dest, renderedPage.Bytes(), 0o644); err != nil {
			return fmt.Errorf("write generated page %q: %w", route.Dest, err)
		}
	}

	return nil
}
