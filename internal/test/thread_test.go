package main

import (
	"fmt"
	"testing"

	pyeval "github.com/M-Quadra/go-python3-submodule/v9/py-eval"
	pygilstate "github.com/M-Quadra/go-python3-submodule/v9/py-gil-state"
	pythreadstate "github.com/M-Quadra/go-python3-submodule/v9/py-thread-state"
	"github.com/stretchr/testify/assert"
)

func TestPyEvalInitThreads(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyeval.InitThreads()
	assert.True(t, pyeval.ThreadsInitialized())
	pyeval.InitThreads()
}

func TestPyGILState(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	assert.True(t, pygilstate.Check())
}

func TestPyThreadState(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	threadStateA := pygilstate.GetThisThreadState()
	threadStateB := pythreadstate.Get()
	assert.Equal(t, threadStateA, threadStateB)

	threadStateC := pythreadstate.Swap(threadStateA)
	assert.Equal(t, threadStateA, threadStateC)
}

func TestPyEvalSaveRestoreThread(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	threadState := pyeval.SaveThread()
	defer pyeval.RestoreThread(threadState)

	assert.False(t, pygilstate.Check())
}
