# go-python3-submodule

#### English (poorly_(ˊཀˋ」∠)_) | [简体中文](./README_zh-cn.md)

This project is an incorporated submodule of Go binds to Python3/C APIs. 

This project is inspired from [DataDog/go-python3](https://github.com/DataDog/go-python3) that has Faster compilation with multiple submodules.

## Installation

### Go Modules

```
go get github.com/M-Quadra/go-python3-submodule/v10
```

Python version | Package URL
:---:|:---:
3.10 | github.com/M-Quadra/go-python3-submodule/v10 
3.9 | github.com/M-Quadra/go-python3-submodule/v9 
3.8 | github.com/M-Quadra/go-python3-submodule/v8 

## Usage

Call as `Python/C API` functions. e.g. `PyBool_Check(x)` call with `pybool.Check`. The gopls will import the rest automately.

Function names are coming from types. As follow:

Python/C API | Go
:---:|:---:
PyFloat_AsDouble | pyfloat.AsFloat64
PyLong_AsLong | pylong.AsInt
PyLong_AsLongLong | pylong.AsInt64
... | ...

If you are confusing while running the code, to check the GIL or reference count.

There are two example for goroutine: [example-0](./internal/example/goroutine-0) , [example-1](/internal/example/goroutine-1).

I'd like to use example-0.

## Progress

Development environment: macOS (x86-64).

Test environment: Linux (Docker x86-64)

Incorporate the most of the original functions in [DataDog/go-python3](https://github.com/DataDog/go-python3).

Add the counting check in references that cover the most of test cases.

## Todo

- [ ] Unit test of reference counts in `Exception`.
- [ ] GitHub Action.
