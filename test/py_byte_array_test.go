package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/py"
	pybytearray "github.com/M-Quadra/go-python3-submodule/py-byte-array"
	"github.com/stretchr/testify/assert"
)

func TestPyByteArrayCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pybytearray.Check(nil))
	assert.False(t, pybytearray.CheckExact(nil))

	ary := pybytearray.FromString("身落红尘心已死")
	defer py.DecRef(ary)

	assert.True(t, pybytearray.Check(ary))
	assert.True(t, pybytearray.CheckExact(ary))
}

func TestPyByteArrayFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := "鹰语"

	ary := pybytearray.FromString(str)
	defer py.DecRef(ary)

	assert.Equal(t, str, pybytearray.AsString(ary))
}

func TestPyByteArrayConcat(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pybytearray.Concat(nil, nil))

	strA := "坎"
	aryA := pybytearray.FromString(strA)
	defer py.DecRef(aryA)

	aryA0 := pybytearray.Concat(aryA, nil)
	defer py.DecRef(aryA0)
	assert.NotEqual(t, unsafe.Pointer(aryA), unsafe.Pointer(aryA0))
	assert.Equal(t, strA, pybytearray.AsString(aryA0))

	aryA1 := pybytearray.Concat(nil, aryA)
	defer py.DecRef(aryA1)
	assert.NotEqual(t, unsafe.Pointer(aryA), unsafe.Pointer(aryA1))
	assert.Equal(t, strA, pybytearray.AsString(aryA1))

	strB := "离"
	aryB := pybytearray.FromString(strB)
	defer py.DecRef(aryB)

	aryAB := pybytearray.Concat(aryA, aryB)
	defer py.DecRef(aryAB)
	assert.NotEqual(t, unsafe.Pointer(aryA), unsafe.Pointer(aryAB))
	assert.NotEqual(t, unsafe.Pointer(aryB), unsafe.Pointer(aryAB))
	assert.Equal(t, strA+strB, pybytearray.AsString(aryAB))
}

func TestPyByteArrayResize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pybytearray.Resize(nil, 0)

	str := "对A"
	ary := pybytearray.FromString(str)
	defer py.DecRef(ary)
	assert.Zero(t, pybytearray.Resize(ary, len("对")))
	assert.Equal(t, "对", pybytearray.AsString(ary))

	// assert.Zero(t, pybytearray.Resize(ary, -1))
}
