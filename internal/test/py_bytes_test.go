package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pybytes "github.com/M-Quadra/go-python3-submodule/v11/py-bytes"
	pylist "github.com/M-Quadra/go-python3-submodule/v11/py-list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPyBytesCheck(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	require.False(t, pybytes.Check(nil))
	require.False(t, pybytes.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { require.Equal(t, 1, py.RefCnt(list)) }()

	require.False(t, pybytes.Check(list))
	require.False(t, pybytes.CheckExact(list))

	bytes := pybytes.FromString("aria")
	defer py.DecRef(bytes)
	defer func() { require.Equal(t, 1, py.RefCnt(bytes)) }()

	require.True(t, pybytes.Check(bytes))
	require.True(t, pybytes.CheckExact(bytes))
}

func TestPyBytesFromString(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	str := "hidan"
	bytes := pybytes.FromString(str)
	defer py.DecRef(bytes)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytes)) }()

	require.Equal(t, str, pybytes.AsString(bytes))
}

func TestPyBytesSize(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	strA := "no"
	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	defer func() { require.Equal(t, 1, py.RefCnt(bytesA)) }()

	require.Equal(t, len(strA), pybytes.Size(bytesA))

	strB := "の"
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)
	defer func() { require.Equal(t, 1, py.RefCnt(bytesB)) }()

	require.Equal(t, len(strB), pybytes.Size(bytesB))
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
