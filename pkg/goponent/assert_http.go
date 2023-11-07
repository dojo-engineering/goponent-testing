package goponent

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

var _ Action = HttpResponseAsserter{}

type HttpResponseAsserter struct {
	ExpectedBody       string
	ExpectedStatusCode int
}

func (h HttpResponseAsserter) Execute(t *testing.T, context *Context, stepContext *Context) error {
	res, ok := ContextGet[*http.Response](stepContext, "response")
	if !ok {
		t.Error("no response in context")
		return errors.New("no response in context")
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	AssertEqual(t, h.ExpectedBody, string(b))
	AssertEqual(t, h.ExpectedStatusCode, res.StatusCode)
	return nil
}
