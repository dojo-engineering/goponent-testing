package goponent

import (
	"testing"

	"gotest.tools/v3/assert"
)

var AssertEqualFunction func(t assert.TestingT, x, y interface{}, msgAndArgs ...interface{}) = assert.Equal

func AssertEqual[T any](t *testing.T, expected T, actual T, messages ...string) {
	AssertEqualFunction(t, expected, actual, messages)
}
