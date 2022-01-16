# go-python3-submodule

#### English (poorly_(ˊཀˋ」∠)_) | [简体中文](./README_zh-cn.md)

This project is an incorporated submodule of Go binds to Python3/C APIs. 

This project is inspired from [DataDog/go-python3](https://github.com/DataDog/go-python3) that has Faster compilation with multiple submodules.

# Installation

Go modules are fully preferable.

```
go get github.com/M-Quadra/go-python3-submodule/v9
```

Python version | Package URL
:---:|:---:
3.9 | github.com/M-Quadra/go-python3-submodule/v9
3.8 | github.com/M-Quadra/go-python3-submodule/v8

# Usage

Call as `Python/C API` functions. e.g. `PyBool_Check(x)` call with `pybool.Check`. The gopls will import the rest automately.

A test example, watch [test](./test) folder here.

Function names are coming from types. As follow:

Python/C API | Go
:---:|:---:
PyFloat_AsDouble | pyfloat.AsFloat64
PyLong_AsLong | pylong.AsInt
PyLong_AsLongLong | pylong.AsInt64
... | ...

If you are confusing while running the code, to check the GIL. For using under a single thread, try following code:

```
if !pygilstate.Check() {
	save := pyeval.SaveThread()
	defer pyeval.RestoreThread(save)

	gstate := pygilstate.Ensure()
	defer pygilstate.Release(gstate)
}

// do something...
```

In goroutines, just add a Lock.

```
var _m = sync.Mutex{}

func xx() {
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

	// do something...
}
```

This is an [example](./test/benchmark/curvefit_test.go).

# Progress

Development environment: macOS 12.1, python 3.9.

Incorporate the most of the original functions in [DataDog/go-python3](https://github.com/DataDog/go-python3).

Add the counting check in references that cover the most of test cases.

# Todo

- `PyModule_GetDef`, need convert to structure.

- Support `Py_Main`.

- Add `PyObject.ob_refcnt` or keep `Py_REFCNT`?

- Check reference counts for `Exception`.

- CI.

# Other

[Compile Test](https://github.com/M-Quadra/go-python3-submodule/wiki/Compile-Test)
