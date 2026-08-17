package parsing

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseMarkdownRendersMarkdownDocument(t *testing.T) {
	path := filepath.Join(t.TempDir(), "document.md")
	writeMarkdownFile(t, path, `# Hello

A paragraph with a [link](https://example.com).

- First
- Second
`)

	content, err := ParseMarkdown(path)
	if err != nil {
		t.Fatalf("ParseMarkdown() error = %v", err)
	}

	expectedFragments := []string{
		"<h1>Hello</h1>",
		`<p>A paragraph with a <a href="https://example.com">link</a>.</p>`,
		"<ul>",
		"<li>First</li>",
		"<li>Second</li>",
	}
	for _, expectedFragment := range expectedFragments {
		if !strings.Contains(string(content), expectedFragment) {
			t.Errorf("ParseMarkdown() = %q, want it to contain %q", content, expectedFragment)
		}
	}
}

func TestParseMarkdownEnablesGitHubFlavouredMarkdown(t *testing.T) {
	path := filepath.Join(t.TempDir(), "document.md")
	writeMarkdownFile(t, path, `~~Removed~~

https://example.com

- [x] Complete

| Name | Value |
| --- | --- |
| First | Second |
`)

	content, err := ParseMarkdown(path)
	if err != nil {
		t.Fatalf("ParseMarkdown() error = %v", err)
	}

	expectedFragments := map[string]string{
		"strikethrough":  "<del>Removed</del>",
		"automatic link": `<a href="https://example.com">https://example.com</a>`,
		"task list":      `type="checkbox"`,
		"table":          "<table>",
	}
	for feature, expectedFragment := range expectedFragments {
		if !strings.Contains(string(content), expectedFragment) {
			t.Errorf("ParseMarkdown() did not render %s; output = %q", feature, content)
		}
	}
}

func TestParseMarkdownAllowsRawHTML(t *testing.T) {
	path := filepath.Join(t.TempDir(), "document.md")
	rawHTML := `<section data-kind="trusted">Trusted content</section>`
	writeMarkdownFile(t, path, rawHTML)

	content, err := ParseMarkdown(path)
	if err != nil {
		t.Fatalf("ParseMarkdown() error = %v", err)
	}
	if !strings.Contains(string(content), rawHTML) {
		t.Errorf("ParseMarkdown() = %q, want it to contain raw HTML %q", content, rawHTML)
	}
}

func TestParseMarkdownReturnsEmptyContentForEmptyDocument(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.md")
	writeMarkdownFile(t, path, "")

	content, err := ParseMarkdown(path)
	if err != nil {
		t.Fatalf("ParseMarkdown() error = %v", err)
	}
	if content != "" {
		t.Errorf("ParseMarkdown() = %q, want empty content", content)
	}
}

func TestParseMarkdownReturnsFileReadingErrors(t *testing.T) {
	t.Run("missing file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.md")

		content, err := ParseMarkdown(path)
		if err == nil {
			t.Fatal("ParseMarkdown() error = nil, want file reading error")
		}
		if content != "" {
			t.Errorf("ParseMarkdown() = %q, want empty content", content)
		}
		if !errors.Is(err, os.ErrNotExist) {
			t.Errorf("ParseMarkdown() error = %v, want os.ErrNotExist", err)
		}
		if !strings.Contains(err.Error(), path) || !strings.Contains(err.Error(), "read markdown file") {
			t.Errorf("ParseMarkdown() error = %q, want path and reading context", err)
		}
	})

	t.Run("directory path", func(t *testing.T) {
		path := t.TempDir()

		content, err := ParseMarkdown(path)
		if err == nil {
			t.Fatal("ParseMarkdown() error = nil, want file reading error")
		}
		if content != "" {
			t.Errorf("ParseMarkdown() = %q, want empty content", content)
		}
		if !strings.Contains(err.Error(), path) || !strings.Contains(err.Error(), "read markdown file") {
			t.Errorf("ParseMarkdown() error = %q, want path and reading context", err)
		}
	})
}

func writeMarkdownFile(t *testing.T, path, content string) {
	t.Helper()

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("os.WriteFile() error = %v", err)
	}
}
