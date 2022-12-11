package main

import (
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pyeval "github.com/M-Quadra/go-python3-submodule/v10/py-eval"
	pyfloat "github.com/M-Quadra/go-python3-submodule/v10/py-float"
	pygilstate "github.com/M-Quadra/go-python3-submodule/v10/py-gil-state"
	pyimport "github.com/M-Quadra/go-python3-submodule/v10/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v10/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/v10/py-object"
	pysys "github.com/M-Quadra/go-python3-submodule/v10/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v10/py-unicode"
)

func init() {
	py.Finalize()
	py.Initialize()
	if !py.IsInitialized() {
		os.Exit(-1)
	}

	paths := pysys.GetObject("path")
	if paths == nil {
		os.Exit(-1)
	}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		os.Exit(-1)
	}

	wd := path.Dir(file)
	wdUnicode := pyunicode.FromString(wd)
	defer py.DecRef(wdUnicode)
	if !pylist.Append(paths, wdUnicode) {
		os.Exit(-1)
	}
}

var _m = sync.Mutex{}

// GetPopt python function
func GetPopt(trainX, trainY []int) []float64 {
	_m.Lock()
	defer _m.Unlock()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	trainXPy := pylist.FromInts(trainX)
	defer py.DecRef(trainXPy)
	trainYPy := pylist.FromInts(trainY)
	defer py.DecRef(trainYPy)

	curvefit := pyimport.ImportModule("curvefit")
	defer py.DecRef(curvefit)

	getPopt := pyobject.GetAttrString(curvefit, "getPopt")
	defer py.DecRef(getPopt)

	res := pyobject.CallFunctionObjArgs(getPopt, trainXPy, trainYPy)
	defer py.DecRef(res)

	opt := make([]float64, 0, pylist.Size(res))
	for i := 0; i < pylist.Size(res); i++ {
		item := pyfloat.AsFloat64(pylist.GetItem(res, i))
		opt = append(opt, item)
	}
	return opt
}
