package goponent

import (
	"net/http"
	"strings"
	"testing"

	"github.com/h2non/gock"
)

var _ Setup = SetupHttpDependencyAction{}

type SetupHttpDependencyAction struct {
	Method     string
	Body       string
	StatusCode int
	Host       string
	Path       string
	Client     *http.Client
}

func (a SetupHttpDependencyAction) Setup(t *testing.T, context *Context, stepContext *Context) error {
	if a.Client != nil {
		gock.InterceptClient(a.Client)
	}
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
