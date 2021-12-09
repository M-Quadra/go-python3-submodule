package main

import (
	"math"
	"runtime"
	"sync"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pyeval "github.com/M-Quadra/go-python3-submodule/v8/py-eval"
	pyfloat "github.com/M-Quadra/go-python3-submodule/v8/py-float"
	pygilstate "github.com/M-Quadra/go-python3-submodule/v8/py-gil-state"
	pyimport "github.com/M-Quadra/go-python3-submodule/v8/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v8/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/v8/py-object"
	"github.com/stretchr/testify/assert"
)

var _m = sync.Mutex{}

func curvefit(t assert.TestingT, isMultithreading bool) {
	if isMultithreading {
		_m.Lock()
		defer _m.Unlock()
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}
	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	curvefit := pyimport.ImportModule("curvefit")
	defer py.DecRef(curvefit)

	trainX := pylist.FromInts([]int{1, 2, 3, 4, 5})
	defer py.DecRef(trainX)
	trainY := pylist.FromInts([]int{1, 2, 3, 4, 5})
	defer py.DecRef(trainY)

	getPopt := pyobject.GetAttrString(curvefit, "getPopt")

	opt := pyobject.CallFunctionObjArgs(getPopt, trainX, trainY)
	defer py.DecRef(opt)
	assert.True(t, pylist.Check(opt))
	assert.Equal(t, 2, pylist.Size(opt))

	popts := []float64{1, 0}
	for i := 0; i < pylist.Size(opt) && i < len(popts); i++ {
		item := pyfloat.AsFloat64(pylist.GetItem(opt, i))
		assert.True(t, math.Abs(item-popts[i]) < 1e-8)
	}
}

func TestCurvefit(t *testing.T) {
	md := pyimport.ImportModule("curvefit")
	defer py.DecRef(md)

	trainX := pylist.FromInts([]int{1, 2, 3, 4, 5})
	defer py.DecRef(trainX)
	trainY := pylist.FromInts([]int{1, 2, 3, 4, 5})
	defer py.DecRef(trainY)

	getPopt := pyobject.GetAttrString(md, "getPopt")

	opt := pyobject.CallFunctionObjArgs(getPopt, trainX, trainY)
	defer py.DecRef(opt)
	assert.True(t, pylist.Check(opt))

	popts := []float64{1, 0}
	for i := 0; i < pylist.Size(opt) && i < len(popts); i++ {
		item := pyfloat.AsFloat64(pylist.GetItem(opt, i))
		assert.True(t, math.Abs(item-popts[i]) < 1e-8)
	}
}

func BenchmarkCurvefit(b *testing.B) {
	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gil := pygilstate.Ensure()
		defer pygilstate.Release(gil)
	}

	for i := 0; i < b.N; i++ {
		curvefit(b, false)
	}
}

func BenchmarkCurvefitMultithreading(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()

			curvefit(b, true)
		}()
	}
	wg.Wait()
}
