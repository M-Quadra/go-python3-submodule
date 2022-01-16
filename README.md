# go-python3-submodule

English (poorly_(ˊཀˋ」∠)_) | [简体中文](./README_zh-cn.md)

This is Go bindings for Python3/C API with submodule style. 

This project was inspired by [DataDog/go-python3](https://github.com/DataDog/go-python3). Faster compilation with multiple submodules.

# Installation

Recommend Go modules.

```
go get github.com/M-Quadra/go-python3-submodule/v9
```

Python version | Package URL
:---:|:---:
3.9 | github.com/M-Quadra/go-python3-submodule/v9
3.8 | github.com/M-Quadra/go-python3-submodule/v8

# Usage

Use as `Python/C API` functions. Eg: `PyBool_Check(x)` call with `pybool.Check`. Let gopls auto autocomplete package import.

Test as example, watch [test](./test)folder.

The function names converted by type. As follows:

Python/C API | Go
:---:|:---:
PyFloat_AsDouble | pyfloat.AsFloat64
PyLong_AsLong | pylong.AsInt
PyLong_AsLongLong | pylong.AsInt64
... | ...

If you get stuck when running, check the GIL. For single threads try following code:

```
if !pygilstate.Check() {
	save := pyeval.SaveThread()
	defer pyeval.RestoreThread(save)

	gstate := pygilstate.Ensure()
	defer pygilstate.Release(gstate)
}

// do something...
```

In goroutines, just add Lock.

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

Add most of functions in [DataDog/go-python3](https://github.com/DataDog/go-python3).

Add reference count check for most of test cases.

# Todo

- `PyModule_GetDef`, need convert to structure.

- Support `Py_Main`.

- Add `PyObject.ob_refcnt` or keep `Py_REFCNT`?

- Reference count check for `Exception`.

- CI.

# Other

[Compile Test](https://github.com/M-Quadra/go-python3-submodule/wiki/Compile-Test)