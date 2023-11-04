package goponent

import (
	"io"
)

type RequestGenerator interface {
	GetBody() io.Reader
	GetPath() string
	GetContentType() string
}
