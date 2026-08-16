package parsing

import (
	"errors"
	"fmt"
)

type HtmlElement interface {
	Print() string
}

func ParseMarkdown(path string) ([]HtmlElement, error) {
	return nil, errors.New("not implemented")
}

type Header struct {
	level int
	content string
}

func (e Header) Print() string {
	return fmt.Sprintf("<%v class='markdown-h%v'>%s</%v>", e.level, e.level, e.content, e.level)
}
