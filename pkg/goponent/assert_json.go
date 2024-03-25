package goponent

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

var _ Asserter = JsonResponseAsserter[string]{}

type JsonResponseAsserter[T any] struct {
	ExpectedBody       T
	ExpectedBodyFunc   func(ctx *Context) T
	ExpectedStatusCode int
}

func (j JsonResponseAsserter[T]) Assert(t *testing.T, context *Context, stepContext *Context) error {
	actualBody, res, err := extractBody[T](t, stepContext)
	if err != nil {
		return err
	}

	expectedBody := j.ExpectedBody
	if j.ExpectedBodyFunc != nil {
		expectedBody = j.ExpectedBodyFunc(context)
	}

	AssertEqual(t, expectedBody, actualBody)
	AssertEqual(t, j.ExpectedStatusCode, res.StatusCode)

	return nil
}

var _ Asserter = JsonFuncAsserter[string]{}

type JsonFuncAsserter[T any] struct {
	ExpectedFunc       func(ctx *Context, t *testing.T, body T)
	ExpectedStatusCode int
}

func (j JsonFuncAsserter[T]) Assert(t *testing.T, context *Context, stepContext *Context) error {
	actualBody, res, err := extractBody[T](t, stepContext)
	if err != nil {
		return err
	}
	j.ExpectedFunc(context, t, actualBody)
	AssertEqual(t, j.ExpectedStatusCode, res.StatusCode)
	return nil
}

func extractBody[T any](t *testing.T, stepContext *Context) (T, *http.Response, error) {
	var actualBody T
	res, ok := ContextGet[*http.Response](stepContext, "response")
	if !ok {
		t.Error("no response in context")
		return actualBody, res, errors.New("no response in context")
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return actualBody, res, err
	}

	t.Logf("response body: %s", string(b))
	err = json.Unmarshal(b, &actualBody)
	if err != nil {
		t.Error(err)
		return actualBody, res, err
	}
	return actualBody, res, nil
}
