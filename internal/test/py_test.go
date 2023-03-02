package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPyIsNone(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	require.True(t, py.IsNone(py.None))
	require.False(t, py.IsNone(nil))
	require.False(t, py.IsNone(py.True))
}
