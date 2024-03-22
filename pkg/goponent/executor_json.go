package goponent

import (
	"encoding/json"
	"testing"
)

var _ Executor = JsonRequestExecutor[string]{}

type JsonRequestExecutor[T any] struct {
	Method   string
	Path     string
	PathFunc func(ctx *Context) string
	Body     T
	BodyFunc func(ctx *Context) T
	Headers  map[string]string
}

func (j JsonRequestExecutor[T]) Execute(t *testing.T, context *Context, stepContext *Context) error {
	body := j.Body
	if j.BodyFunc != nil {
		body = j.BodyFunc(context)
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	path := j.Path
	if j.PathFunc != nil {
		path = j.PathFunc(context)
	}
	httpAction := HttpRequestExecutor{
		Method:      j.Method,
		ContentType: "application/json",
		Path:        path,
		Body:        b,
		Headers:     j.Headers,
	}
	return httpAction.Execute(t, context, stepContext)

}
