package goponent

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestContextGet(t *testing.T) {
	c := newContext()
	ContextSet(c, "key", "value")
	val, ok := ContextGet[string](c, "key")
	assert.Equal(t, true, ok)
	AssertEqual(t, "value", val)

	intVal, intOk := ContextGet[int](c, "key")
	assert.Equal(t, false, intOk)
	AssertEqual(t, 0, intVal)
}
