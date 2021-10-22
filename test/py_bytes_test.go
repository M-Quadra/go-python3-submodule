package main

import (
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/py"
	pybytes "github.com/M-Quadra/go-python3-submodule/py-bytes"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	"github.com/stretchr/testify/assert"
)

func TestPyBytesCheck(t *testing.T) {
	assert.False(t, pybytes.Check(nil))
	assert.False(t, pybytes.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.False(t, pybytes.Check(list))
	assert.False(t, pybytes.CheckExact(list))

	bytes := pybytes.FromString("aria")
	defer py.DecRef(bytes)
	assert.True(t, pybytes.Check(bytes))
	assert.True(t, pybytes.CheckExact(bytes))
}

func TestPyBytesFromString(t *testing.T) {
	str := "hidan"
	bytes := pybytes.FromString(str)
	defer py.DecRef(bytes)
	assert.Equal(t, str, pybytes.AsString(bytes))
}

func TestPyBytesSize(t *testing.T) {
	strA := "no"
	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	assert.Equal(t, len(strA), pybytes.Size(bytesA))

	strB := "の"
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)
	assert.Equal(t, len(strB), pybytes.Size(bytesB))
}

func TestPyBytesConcatAndDel(t *testing.T) {
	strA := "a"
	strB := "b"

	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)

	bytesC := pybytes.ConcatAndDel(bytesA, bytesB)
	defer py.DecRef(bytesC)

	assert.NotEqual(t, unsafe.Pointer(bytesA), unsafe.Pointer(bytesC))
	assert.NotEqual(t, unsafe.Pointer(bytesB), unsafe.Pointer(bytesC))
	assert.Equal(t, strA+strB, pybytes.AsString(bytesC))
}
