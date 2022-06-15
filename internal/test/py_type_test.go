package main

import (
	"fmt"
	"testing"

	pytype "github.com/M-Quadra/go-python3-submodule/v9/py-type"
	"github.com/stretchr/testify/assert"
)

func TestPyTypeCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pytype.Check(nil))
	assert.False(t, pytype.CheckExact(nil))

	assert.True(t, pytype.Check(pytype.Type))
	assert.True(t, pytype.CheckExact(pytype.Type))
}

func TestPyTypeClearCache(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.NotEqual(t, uint(0xffffffff), pytype.ClearCache())
	assert.Equal(t, uint(0xffffffff), pytype.ClearCache())
	assert.Equal(t, uint(0xffffffff), pytype.ClearCache())
}
