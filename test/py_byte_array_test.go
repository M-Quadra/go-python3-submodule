package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pybytearray "github.com/M-Quadra/go-python3-submodule/v8/py-byte-array"
	"github.com/stretchr/testify/assert"
)

func TestPyByteArrayCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pybytearray.Check(nil))
	assert.False(t, pybytearray.CheckExact(nil))

	arr := pybytearray.FromString("身落红尘心已死")
	defer py.DecRef(arr)
	defer func() { assert.Equal(t, 1, py.RefCnt(arr)) }()

	assert.True(t, pybytearray.Check(arr))
	assert.True(t, pybytearray.CheckExact(arr))
}

func TestPyByteArrayFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := "鹰语"

	arr := pybytearray.FromString(str)
	defer py.DecRef(arr)
	defer func() { assert.Equal(t, 1, py.RefCnt(arr)) }()

	assert.Equal(t, str, pybytearray.AsString(arr))
}

func TestPyByteArrayConcat(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pybytearray.Concat(nil, nil))

	strA := "坎"
	arrA := pybytearray.FromString(strA)
	defer py.DecRef(arrA)
	defer func() { assert.Equal(t, 1, py.RefCnt(arrA)) }()

	arrA0 := pybytearray.Concat(arrA, nil)
	defer py.DecRef(arrA0)
	defer func() { assert.Equal(t, 1, py.RefCnt(arrA0)) }()

	assert.NotEqual(t, unsafe.Pointer(arrA), unsafe.Pointer(arrA0))
	assert.Equal(t, strA, pybytearray.AsString(arrA0))

	arrA1 := pybytearray.Concat(nil, arrA)
	defer py.DecRef(arrA1)
	defer func() { assert.Equal(t, 1, py.RefCnt(arrA1)) }()

	assert.NotEqual(t, unsafe.Pointer(arrA), unsafe.Pointer(arrA1))
	assert.Equal(t, strA, pybytearray.AsString(arrA1))

	strB := "离"
	arrB := pybytearray.FromString(strB)
	defer py.DecRef(arrB)
	defer func() { assert.Equal(t, 1, py.RefCnt(arrB)) }()

	arrAB := pybytearray.Concat(arrA, arrB)
	defer py.DecRef(arrAB)
	defer func() { assert.Equal(t, 1, py.RefCnt(arrAB)) }()

	assert.NotEqual(t, unsafe.Pointer(arrA), unsafe.Pointer(arrAB))
	assert.NotEqual(t, unsafe.Pointer(arrB), unsafe.Pointer(arrAB))
	assert.Equal(t, strA+strB, pybytearray.AsString(arrAB))
}

func TestPyByteArrayResize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pybytearray.Resize(nil, 0)

	str := "对A"
	arr := pybytearray.FromString(str)
	defer py.DecRef(arr)
	defer func() { assert.Equal(t, 1, py.RefCnt(arr)) }()

	assert.Zero(t, pybytearray.Resize(arr, len("对")))
	assert.Equal(t, "对", pybytearray.AsString(arr))

	// assert.Zero(t, pybytearray.Resize(arr, -1))
	// assert.Equal(t, -1, pybytearray.Size(arr))
}
