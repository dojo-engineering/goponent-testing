package goponent

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

var _ ContextSetter = ContextSetterJson[string]{}

type ContextSetterJson[T any] struct {
	Properties map[string]func(response T) any
}

func (c ContextSetterJson[T]) SetContext(t *testing.T, context *Context, stepContext *Context) error {
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

	// put another reader in the response body so other actions can read it
	res.Body = io.NopCloser(bytes.NewReader(b))

	var actualBody T
	t.Logf("body: %s", string(b))
	err = json.Unmarshal(b, &actualBody)
	if err != nil {
		t.Error(err)
		return err
	}

	for key, value := range c.Properties {
		ContextSet(context, key, value(actualBody))
	}
	return nil
}
