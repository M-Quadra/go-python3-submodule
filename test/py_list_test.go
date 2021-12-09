package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/v8/py-err"
	pylist "github.com/M-Quadra/go-python3-submodule/v8/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/v8/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/v8/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyListCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.Check(nil))
	assert.False(t, pylist.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.Check(list))
	assert.True(t, pylist.CheckExact(list))

	assert.False(t, pylist.Check(py.True))
	assert.False(t, pylist.CheckExact(py.True))
}

func TestPyListSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, 0, pylist.Size(nil))

	l := rand.Intn(1000)
	list := pylist.New(l)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, l, pylist.Size(list))
}

func TestPyListGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pylist.GetItem(nil, 0))

	vRand := rand.Intn(1000)

	list := pylist.New(1)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.SetItem(list, 0, pylong.FromInt(vRand)))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.SetItem(nil, 0, nil))

	list := pylist.New(1)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.SetItem(list, 0, nil))
	assert.Equal(t, 1, pylist.Size(list))
	assert.Nil(t, pylist.GetItem(list, 0))
}

func TestPyListInsert(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.Insert(nil, 0, nil))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.False(t, pylist.Insert(list, 0, nil))
	assert.Equal(t, 0, pylist.Size(list))

	assert.Equal(t, vRefCnt, py.RefCnt(v))
	assert.True(t, pylist.Insert(list, 0, v)) // reference count + 1
	assert.Equal(t, vRefCnt+1, py.RefCnt(v))

	assert.Equal(t, 1, pylist.Size(list))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListAppend(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.Append(nil, nil))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.False(t, pylist.Append(list, nil))
	assert.Equal(t, 0, pylist.Size(list))

	assert.Equal(t, vRefCnt, py.RefCnt(v))
	assert.True(t, pylist.Append(list, v)) // reference count + 1
	assert.Equal(t, vRefCnt+1, py.RefCnt(v))

	assert.Equal(t, 1, pylist.Size(list))
	assert.Equal(t, vRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyListGetSlice(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pylist.GetSlice(nil, 1, 2))

	vRand := rand.Intn(1000)
	vA := pylong.FromInt(vRand)
	defer py.DecRef(vA)
	vARefCnt := py.RefCnt(vA)
	defer func() { assert.Equal(t, vARefCnt, py.RefCnt(vA)) }()

	vB := pylong.FromInt(vRand + 1)
	defer py.DecRef(vB)
	vBRefCnt := py.RefCnt(vB)
	defer func() { assert.Equal(t, vBRefCnt, py.RefCnt(vB)) }()

	listA := pylist.New(0)
	defer py.DecRef(listA)
	defer func() { assert.Equal(t, 1, py.RefCnt(listA)) }()

	assert.True(t, pylist.Append(listA, vA)) // reference count + 1
	assert.True(t, pylist.Append(listA, vB)) // reference count + 1

	listB := pylist.GetSlice(listA, 1, 2)
	defer py.DecRef(listB)
	defer func() { assert.Equal(t, 1, py.RefCnt(listB)) }()

	assert.Equal(t, listA, listB) // !!!
	assert.NotEqual(t, unsafe.Pointer(listA), unsafe.Pointer(listB))

	vAInt := pylong.AsInt(pylist.GetItem(listA, 0))
	assert.Equal(t, vRand, vAInt)
	vBInt := pylong.AsInt(pylist.GetItem(listB, 0))
	assert.Equal(t, vRand+1, vBInt)
	assert.NotEqual(t, vAInt, vBInt)
}

func TestPyListSort(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.Sort(nil))
	pyerr.Clear()

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.Sort(list))
	sort.Ints(ary)
	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pylist.GetItem(list, i)))
	}
}

func TestPyListReverse(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylist.Reverse(nil))

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.Reverse(list))
	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pylist.GetItem(list, cnt-1-i)))
	}
}

func TestPyListAsTuple(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pylist.AsTuple(nil))

	cnt := rand.Intn(1000) + 10
	ary := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		ary = append(ary, rand.Intn(1000))
	}

	list := pylist.FromInts(ary)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	tuple := pylist.AsTuple(list)
	defer py.DecRef(tuple)
	defer func() { assert.Equal(t, 1, py.RefCnt(tuple)) }()

	for i := 0; i < cnt; i++ {
		assert.Equal(t, ary[i], pylong.AsInt(pytuple.GetItem(tuple, i)))
	}
}
