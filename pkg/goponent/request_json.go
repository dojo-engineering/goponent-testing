package goponent

import (
	"io"
)

var _ RequestGenerator = JSONRequestGenerator{}

type JSONRequestGenerator struct {
}

func (J JSONRequestGenerator) GetBody() io.Reader {
	//TODO implement me
	panic("implement me")
}

func (J JSONRequestGenerator) GetPath() string {
	//TODO implement me
	panic("implement me")
}

func (J JSONRequestGenerator) GetContentType() string {
	//TODO implement me
	panic("implement me")
}
