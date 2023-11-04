package goponent

import (
	"io"
	"strings"
)

var _ RequestGenerator = StringRequestGenerator{}

type StringRequestGenerator struct {
	Method string
	Path   string
}

func (s StringRequestGenerator) GetBody() io.Reader {
	return strings.NewReader("")
}

func (s StringRequestGenerator) GetPath() string {
	return s.Path
}

func (s StringRequestGenerator) GetContentType() string {
	return "text/plain"
}
