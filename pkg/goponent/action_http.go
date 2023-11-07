package goponent

import (
	"bytes"
	"errors"
	"net/http"
	"testing"
)

var _ Action = HttpRequestAction{}

type HttpRequestAction struct {
	Method      string
	ContentType string
	Path        string
	Body        []byte
}

func (h HttpRequestAction) Execute(t *testing.T, context *Context, stepContext *Context) error {
	baseUrl, ok := ContextGet[string](context, "baseUrl")
	if !ok {
		t.Error("no baseUrl in context")
		return errors.New("no baseUrl in context")
	}
	req, err := http.NewRequest(h.Method, baseUrl+h.Path, bytes.NewReader(h.Body))
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
