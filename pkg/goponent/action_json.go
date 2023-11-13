package goponent

import (
	"encoding/json"
	"testing"
)

var _ Actor = JsonRequestAction[string]{}

type JsonRequestAction[T any] struct {
	Method   string
	Path     string
	PathFunc func(ctx *Context) string
	Body     T
	BodyFunc func(ctx *Context) T
}

func (j JsonRequestAction[T]) Act(t *testing.T, context *Context, stepContext *Context) error {
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
	httpAction := HttpRequestAction{
		Method:      j.Method,
		ContentType: "application/json",
		Path:        path,
		Body:        b,
	}
	return httpAction.Act(t, context, stepContext)

}
