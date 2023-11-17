package goponent

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

var _ Executor = HttpRequestExecutor{}

type HttpRequestExecutor struct {
	Method      string
	ContentType string
	Path        string
	PathFunc    func(ctx *Context) string
	Body        []byte
}

func (h HttpRequestExecutor) Execute(t *testing.T, context *Context, stepContext *Context) error {
	baseUrl, ok := ContextGet[string](context, "baseUrl")
	if !ok {
		t.Error("no baseUrl in context")
		return errors.New("no baseUrl in context")
	}

	path := h.Path
	if h.PathFunc != nil {
		path = h.PathFunc(context)
	}

	req, err := http.NewRequest(h.Method, baseUrl+path, bytes.NewReader(h.Body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", h.ContentType)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	ContextSet(stepContext, "response", res)
	return nil
}
