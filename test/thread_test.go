package main

import (
	"fmt"
	"testing"

	pyeval "github.com/M-Quadra/go-python3-submodule/py-eval"
	pygilstate "github.com/M-Quadra/go-python3-submodule/py-gil-state"
	pythreadstate "github.com/M-Quadra/go-python3-submodule/py-thread-state"
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

	pyeval.InitThreads()
	gil := pygilstate.Ensure()
	assert.True(t, pygilstate.Check())
	pygilstate.Release(gil)
}

func TestPyThreadState(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyeval.InitThreads()

	threadStateA := pygilstate.GetThisThreadState()
	threadStateB := pythreadstate.Get()
	assert.Equal(t, threadStateA, threadStateB)

	threadStateC := pythreadstate.Swap(threadStateA)
	assert.Equal(t, threadStateA, threadStateC)
}

func TestPyEvalSaveRestoreThread(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pyeval.InitThreads()

	threadState := pyeval.SaveThread()
	assert.False(t, pygilstate.Check())
	pyeval.RestoreThread(threadState)
}
