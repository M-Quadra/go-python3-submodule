package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pybytes "github.com/M-Quadra/go-python3-submodule/v10/py-bytes"
	pylist "github.com/M-Quadra/go-python3-submodule/v10/py-list"
	"github.com/stretchr/testify/assert"
)

func TestPyBytesCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pybytes.Check(nil))
	assert.False(t, pybytes.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.False(t, pybytes.Check(list))
	assert.False(t, pybytes.CheckExact(list))

	bytes := pybytes.FromString("aria")
	defer py.DecRef(bytes)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytes)) }()

	assert.True(t, pybytes.Check(bytes))
	assert.True(t, pybytes.CheckExact(bytes))
}

func TestPyBytesFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := "hidan"
	bytes := pybytes.FromString(str)
	defer py.DecRef(bytes)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytes)) }()

	assert.Equal(t, str, pybytes.AsString(bytes))
}

func TestPyBytesSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strA := "no"
	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytesA)) }()

	assert.Equal(t, len(strA), pybytes.Size(bytesA))

	strB := "の"
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytesB)) }()

	assert.Equal(t, len(strB), pybytes.Size(bytesB))
}

func TestPyBytesConcatAndDel(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strA := "a"
	strB := "b"

	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytesA)) }()
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytesB)) }()

	bytesC := pybytes.ConcatAndDel(bytesA, bytesB)
	defer py.DecRef(bytesC)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytesC)) }()

	assert.NotEqual(t, unsafe.Pointer(bytesA), unsafe.Pointer(bytesC))
	assert.NotEqual(t, unsafe.Pointer(bytesB), unsafe.Pointer(bytesC))
	assert.Equal(t, strA+strB, pybytes.AsString(bytesC))
}
