package goponent

import (
	"net/http"
	"strings"
	"testing"

	"github.com/h2non/gock"
)

var _ Setup = ArrangeHttpDependencyAction{}

type ArrangeHttpDependencyAction struct {
	Method     string
	Body       string
	StatusCode int
	Host       string
	Path       string
}

func (a ArrangeHttpDependencyAction) Setup(t *testing.T, context *Context, stepContext *Context) error {
	gock.NetworkingFilter(func(request *http.Request) bool {
		return !strings.Contains(request.URL.Host, a.Host)
	})

	req := gock.New(a.Host).
		Path(a.Path)
	if a.Method != "" {
		req.Method = a.Method
	}

	req.Reply(a.StatusCode).BodyString(a.Body)

	return nil
}
