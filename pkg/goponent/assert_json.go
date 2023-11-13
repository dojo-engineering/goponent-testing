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
	t.Logf("response body: %s", string(b))
	err = json.Unmarshal(b, &actualBody)
	if err != nil {
		t.Error(err)
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
