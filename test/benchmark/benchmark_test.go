package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v9/py"
	pyeval "github.com/M-Quadra/go-python3-submodule/v9/py-eval"
	pygilstate "github.com/M-Quadra/go-python3-submodule/v9/py-gil-state"
	pylist "github.com/M-Quadra/go-python3-submodule/v9/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/v9/py-sys"
	pythreadstate "github.com/M-Quadra/go-python3-submodule/v9/py-thread-state"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v9/py-unicode"
	"github.com/stretchr/testify/assert"
)

func init() {
	py.Finalize()
	py.Initialize()
	// defer py.Finalize()
	if !py.IsInitialized() {
		os.Exit(-1)
	}

	paths := pysys.GetObject("path")
	if paths == nil {
		os.Exit(-1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	wdUnicode := pyunicode.FromString(wd)
	defer py.DecRef(wdUnicode)
	if !pylist.Append(paths, wdUnicode) {
		os.Exit(-1)
	}
}

func BenchmarkPyGILState(b *testing.B) {
	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	for i := 0; i < b.N; i++ {
		assert.True(b, pygilstate.Check())
	}
}

func BenchmarkPyEvalSaveRestoreThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			threadState := pyeval.SaveThread()
			defer pyeval.RestoreThread(threadState)

			assert.False(b, pygilstate.Check())
		}()
	}
}

func BenchmarkPyThreadState(b *testing.B) {
	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gil := pygilstate.Ensure()
		defer pygilstate.Release(gil)
	}

	for i := 0; i < b.N; i++ {
		threadStateA := pygilstate.GetThisThreadState()
		threadStateB := pythreadstate.Get()
		assert.Equal(b, threadStateA, threadStateB)

		threadStateC := pythreadstate.Swap(threadStateA)
		assert.Equal(b, threadStateA, threadStateC)
	}
}
