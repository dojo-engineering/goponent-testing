package goponent

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

var _ Action = JsonResponseAsserter[string]{}

type JsonResponseAsserter[T any] struct {
	ExpectedBody       T
	ExpectedStatusCode int
}

func (j JsonResponseAsserter[T]) Execute(t *testing.T, context *Context, stepContext *Context) error {
	res, ok := ContextGet[*http.Response](stepContext, "response")
	if !ok {
		t.Error("no response in context")
		return errors.New("no response in context")
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return err
	}

	var actualBody T
	err = json.Unmarshal(b, &actualBody)
	if err != nil {
		t.Error(err)
		return err
	}

	AssertEqual(t, j.ExpectedBody, actualBody)
	AssertEqual(t, j.ExpectedStatusCode, res.StatusCode)

	return nil
}
