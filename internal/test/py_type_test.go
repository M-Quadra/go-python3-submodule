package main

import (
	"fmt"
	"testing"

	pytype "github.com/M-Quadra/go-python3-submodule/v11/py-type"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPyTypeCheck(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	require.False(t, pytype.Check(nil))
	require.False(t, pytype.CheckExact(nil))

	require.True(t, pytype.Check(pytype.Type))
	require.True(t, pytype.CheckExact(pytype.Type))
}

func TestPyTypeClearCache(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.NotEqual(t, uint(0xffffffff), pytype.ClearCache())
	assert.Equal(t, uint(0xffffffff), pytype.ClearCache())
	assert.Equal(t, uint(0xffffffff), pytype.ClearCache())
}
