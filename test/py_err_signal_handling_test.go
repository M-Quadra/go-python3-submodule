package main

import (
	"fmt"
	"testing"

	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	"github.com/stretchr/testify/assert"
)

func TestPyErrCheckSignals(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.Clear()

	pyerr.SetInterrupt()
	assert.Equal(t, -1, pyerr.CheckSignals())

	assert.True(t, pyerr.GivenExceptionMatches(pyerr.Occurred(), pyexc.TypeError))
	pyerr.Print()
	assert.Nil(t, pyerr.Occurred())
}
