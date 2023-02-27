package main

import (
	"fmt"
	"testing"

	pyerr "github.com/M-Quadra/go-python3-submodule/v11/py-err"
	"github.com/stretchr/testify/assert"
)

func TestPyErrCheckSignals(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	pyerr.SetInterrupt()

	assert.Nil(t, pyerr.Occurred())
	assert.Equal(t, 0, pyerr.CheckSignals())
	assert.Nil(t, pyerr.Occurred())
}
