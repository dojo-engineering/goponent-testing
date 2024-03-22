package goponent

import (
	"errors"
	"io"
	"net/http"
	"testing"
)

var _ Asserter = HttpResponseAsserter{}

type HttpResponseAsserter struct {
	ExpectedBody       string
	ExpectedStatusCode int
}

func (h HttpResponseAsserter) Assert(t *testing.T, context *Context, stepContext *Context) error {
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

var _ Asserter = HttpResponseStatusCodeAsserter{}

type HttpResponseStatusCodeAsserter struct {
	ExpectedStatusCode int
}

func (h HttpResponseStatusCodeAsserter) Assert(t *testing.T, context *Context, stepContext *Context) error {
	res, ok := ContextGet[*http.Response](stepContext, "response")
	if !ok {
		t.Error("no response in context")
		return errors.New("no response in context")
	}

	AssertEqual(t, h.ExpectedStatusCode, res.StatusCode)
	return nil
}
