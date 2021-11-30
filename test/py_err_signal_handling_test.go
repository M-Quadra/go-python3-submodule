package main

import (
	"fmt"
	"runtime"
	"testing"

	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	"github.com/stretchr/testify/assert"
)

func TestPyErrCheckSignals(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyerr.SetInterrupt()

	switch runtime.GOOS {
	case "linux":
		assert.Nil(t, pyerr.Occurred())
		assert.Equal(t, 0, pyerr.CheckSignals())
		assert.Nil(t, pyerr.Occurred())
	default:
		assert.Nil(t, pyerr.Occurred())
		assert.Equal(t, -1, pyerr.CheckSignals())
		assert.NotNil(t, pyerr.Occurred())

		// TypeError: 'NoneType' object is not callable
		assert.True(t, pyerr.GivenExceptionMatches(pyerr.Occurred(), pyexc.TypeError))
		pyerr.Print()
		assert.Nil(t, pyerr.Occurred())
	}
}
