package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v9/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/v9/py-err"
	pylist "github.com/M-Quadra/go-python3-submodule/v9/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/v9/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/v9/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyTupleCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pytuple.Check(nil)
	pytuple.CheckExact(nil)

	tuple := pytuple.New(0)
	defer py.DecRef(tuple)
	tupleRefCnt := py.RefCnt(tuple)
	defer func() { assert.Equal(t, tupleRefCnt, py.RefCnt(tuple)) }()
	assert.True(t, pytuple.Check(tuple))
	assert.True(t, pytuple.CheckExact(tuple))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()
	assert.False(t, pytuple.Check(list))
	assert.False(t, pytuple.CheckExact(list))
}

func TestPyTupleNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	tuple := pytuple.New(0)
	defer py.DecRef(tuple)
	tupleRefCnt := py.RefCnt(tuple)
	defer func() { assert.Equal(t, tupleRefCnt, py.RefCnt(tuple)) }()
	assert.NotNil(t, tuple)
}

func TestPyTupleSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Zero(t, pytuple.Size(nil))

	l := rand.Intn(1000)
	tuple := pytuple.New(l)
	defer py.DecRef(tuple)
	defer func() { assert.Equal(t, 1, py.RefCnt(tuple)) }()

	assert.Equal(t, l, pytuple.Size(tuple))
}

func TestPyTupleGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pytuple.GetItem(nil, 0))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	py.IncRef(v)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt-1, py.RefCnt(v)) }()

	tuple := pytuple.New(1)
	defer py.DecRef(tuple)
	defer func() { assert.Equal(t, 1, py.RefCnt(tuple)) }()
	assert.Nil(t, pytuple.GetItem(tuple, -1))
	pyerr.Clear()

	assert.True(t, pytuple.SetItem(tuple, 0, v))
	assert.Equal(t, vRand, pylong.AsInt(pytuple.GetItem(tuple, 0)))
}

func TestPyTupleGetSlice(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pytuple.GetSlice(nil, 0, 0)) // SystemError: Objects/tupleobject.c:455: bad argument to internal function
	pyerr.Clear()

	tupleA := pytuple.New(1)
	defer py.DecRef(tupleA)
	defer func() { assert.Equal(t, 1, py.RefCnt(tupleA)) }()

	tupleB := pytuple.GetSlice(tupleA, 0, -2)
	defer py.DecRef(tupleB)
	tupleBRefCnt := py.RefCnt(tupleB)
	defer func() { assert.Equal(t, tupleBRefCnt, py.RefCnt(tupleB)) }()
	assert.Zero(t, pytuple.Size(tupleB))
	assert.NotEqual(t, tupleA, tupleB)

	tupleC := pytuple.GetSlice(tupleA, 0, 2)
	defer py.DecRef(tupleC)
	assert.Equal(t, 1, pytuple.Size(tupleC))
	assert.Equal(t, tupleA, tupleC)
}

func TestPyTupleSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	py.IncRef(v)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt-1, py.RefCnt(v)) }()

	assert.False(t, pytuple.SetItem(nil, 0, v))

	tuple := pytuple.New(1)
	defer py.DecRef(tuple)
	defer func() { assert.Equal(t, 1, py.RefCnt(tuple)) }()

	assert.True(t, pytuple.SetItem(tuple, 0, nil))

	py.IncRef(v) // It' ll decref reference count of 'v' even SetItem was false.
	assert.False(t, pytuple.SetItem(tuple, -1, v))
	pyerr.Clear()

	assert.True(t, pytuple.SetItem(tuple, 0, v))
	assert.Equal(t, vRand, pylong.AsInt(pytuple.GetItem(tuple, 0)))
}
