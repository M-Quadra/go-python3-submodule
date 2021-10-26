package main

import (
	"math/rand"
	"sort"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyListCheck(t *testing.T) {
	assert.False(t, pylist.Check(nil))
	assert.False(t, pylist.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.True(t, pylist.Check(list))
	assert.True(t, pylist.CheckExact(list))

	assert.False(t, pylist.Check(py.True))
	assert.False(t, pylist.CheckExact(py.True))
}

func TestPyListSize(t *testing.T) {
	assert.Equal(t, 0, pylist.Size(nil))

	l := rand.Intn(1000)
	list := pylist.New(l)
	defer py.DecRef(list)
	assert.Equal(t, l, pylist.Size(list))
}

func TestPyListGetItem(t *testing.T) {
	assert.Nil(t, pylist.GetItem(nil, 0))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)

	list := pylist.New(1)
	defer py.DecRef(list)
	assert.True(t, pylist.SetItem(list, 0, v))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListSetItem(t *testing.T) {
	assert.False(t, pylist.SetItem(nil, 0, nil))

	list := pylist.New(1)
	defer py.DecRef(list)
	assert.True(t, pylist.SetItem(list, 0, nil))
	assert.Equal(t, 1, pylist.Size(list))
	assert.Nil(t, pylist.GetItem(list, 0))
}

func TestPyListInsert(t *testing.T) {
	assert.False(t, pylist.Insert(nil, 0, nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.False(t, pylist.Insert(list, 0, nil))
	assert.Equal(t, 0, pylist.Size(list))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	assert.True(t, pylist.Insert(list, 0, v))
	assert.Equal(t, 1, pylist.Size(list))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListAppend(t *testing.T) {
	assert.False(t, pylist.Append(nil, nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.False(t, pylist.Append(list, nil))
	assert.Equal(t, 0, pylist.Size(list))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	assert.True(t, pylist.Append(list, v))
	assert.Equal(t, 1, pylist.Size(list))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListGetSlice(t *testing.T) {
	assert.Nil(t, pylist.GetSlice(nil, 1, 2))

	listA := pylist.New(0)
	defer py.DecRef(listA)

	vRand := rand.Intn(1000)
	vA := pylong.FromInt(vRand)
	defer py.DecRef(vA)
	vB := pylong.FromInt(vRand + 1)
	defer py.DecRef(vB)
	assert.True(t, pylist.Append(listA, vA))
	assert.True(t, pylist.Append(listA, vB))

	listB := pylist.GetSlice(listA, 1, 2)
	defer py.DecRef(listB)
	assert.Equal(t, listA, listB) // !!!
	assert.NotEqual(t, unsafe.Pointer(listA), unsafe.Pointer(listB))

	vAInt := pylong.AsInt(pylist.GetItem(listA, 0))
	assert.Equal(t, vRand, vAInt)
	vBInt := pylong.AsInt(pylist.GetItem(listB, 0))
	assert.Equal(t, vRand+1, vBInt)
	assert.NotEqual(t, vAInt, vBInt)
}

func TestPyListSort(t *testing.T) {
	assert.False(t, pylist.Sort(nil))
	pyerr.Clear()

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)

	assert.True(t, pylist.Sort(list))
	sort.Ints(ary)
	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pylist.GetItem(list, i)))
	}
}

func TestPyListReverse(t *testing.T) {
	assert.False(t, pylist.Reverse(nil))

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)

	assert.True(t, pylist.Reverse(list))
	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pylist.GetItem(list, cnt-1-i)))
	}
}

func TestPyListAsTuple(t *testing.T) {
	assert.Nil(t, pylist.AsTuple(nil))

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)

	tuple := pylist.AsTuple(list)
	defer py.DecRef(tuple)

	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pytuple.GetItem(tuple, i)))
	}
}
