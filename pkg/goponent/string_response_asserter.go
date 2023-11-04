package goponent

import (
	"io"
	"net/http"
	"testing"
)

var _ ResponseAsserter = StringResponseAsserter{}

type StringResponseAsserter struct {
	ExpectedBody       string
	ExpectedStatusCode int
}

func (s StringResponseAsserter) Assert(t *testing.T, res *http.Response) error {
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	AssertEqual(t, s.ExpectedBody, string(b))
	AssertEqual(t, s.ExpectedStatusCode, res.StatusCode)

	return nil
}
