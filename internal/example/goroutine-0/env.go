package main

import (
	"os"
	"path"
	"runtime"
	"sync"

	python "github.com/M-Quadra/go-python3-submodule/v10"
	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pyeval "github.com/M-Quadra/go-python3-submodule/v10/py-eval"
	pygilstate "github.com/M-Quadra/go-python3-submodule/v10/py-gil-state"
	pylist "github.com/M-Quadra/go-python3-submodule/v10/py-list"
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

var (
	_m = sync.Mutex{}

	_ok     bool
	_save   *python.PyThreadState
	_gstate python.PyGILState
)

// Lock GIL
func Lock() {
	_m.Lock()
	runtime.LockOSThread()
	_ok = pygilstate.Check()
	if !_ok {
		_save = pyeval.SaveThread()
		_gstate = pygilstate.Ensure()
	}
}

// Unlock GIL
func Unlock() {
	if !_ok {
		pygilstate.Release(_gstate)
		pyeval.RestoreThread(_save)
	}
	runtime.UnlockOSThread()
	_m.Unlock()
}
