package goponent

import (
	"encoding/json"
	"testing"
)

var _ Action = JsonRequestAction[string]{}

type JsonRequestAction[T any] struct {
	Method string
	Path   string
	Body   T
}

func (j JsonRequestAction[T]) Execute(t *testing.T, context *Context, stepContext *Context) error {
	b, err := json.Marshal(j.Body)
	if err != nil {
		return err
	}
	httpAction := HttpRequestAction{
		Method:      j.Method,
		ContentType: "application/json",
		Path:        j.Path,
		Body:        b,
	}
	return httpAction.Execute(t, context, stepContext)

}
