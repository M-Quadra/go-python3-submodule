package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/py-tuple"
	"github.com/stretchr/testify/assert"
)

func TestPyTupleCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pytuple.Check(nil)
	pytuple.CheckExact(nil)

	tuple := pytuple.New(0)
	defer py.DecRef(tuple)
	assert.True(t, pytuple.Check(tuple))
	assert.True(t, pytuple.CheckExact(tuple))

	list := pylist.New(0)
	defer py.DecRef(list)
	assert.False(t, pytuple.Check(list))
	assert.False(t, pytuple.CheckExact(list))
}

func TestPyTupleNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	tuple := pytuple.New(0)
	defer py.DecRef(tuple)
	assert.NotNil(t, tuple)
}

func TestPyTupleSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Zero(t, pytuple.Size(nil))

	l := rand.Intn(1000)
	tuple := pytuple.New(l)
	defer py.DecRef(tuple)

	assert.Equal(t, l, pytuple.Size(tuple))
}

func TestPyTupleGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pytuple.GetItem(nil, 0))

	tuple := pytuple.New(1)
	defer py.DecRef(tuple)
	assert.Nil(t, pytuple.GetItem(tuple, -1))
	pyerr.Clear()

	v := rand.Intn(1000)
	assert.True(t, pytuple.SetItem(tuple, 0, pylong.FromInt(v)))

	assert.Equal(t, v, pylong.AsInt(pytuple.GetItem(tuple, 0)))
}

func TestPyTupleGetSlice(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pytuple.GetSlice(nil, 0, 0))
	pyerr.Clear()

	tupleA := pytuple.New(1)
	defer py.DecRef(tupleA)

	tupleB := pytuple.GetSlice(tupleA, 0, -2)
	defer py.DecRef(tupleB)
	assert.Zero(t, pytuple.Size(tupleB))

	tupleC := pytuple.GetSlice(tupleA, 0, 2)
	defer py.DecRef(tupleC)
	assert.Equal(t, 1, pytuple.Size(tupleC))
}

func TestPyTupleSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := rand.Intn(1000)
	long := pylong.FromInt(v)
	assert.False(t, pytuple.SetItem(nil, 0, long))

	tuple := pytuple.New(1)
	defer py.DecRef(tuple)

	assert.True(t, pytuple.SetItem(tuple, 0, nil))

	py.IncRef(long) // It' ll decref reference count of 'long' even SetItem was false.
	assert.False(t, pytuple.SetItem(tuple, -1, long))
	pyerr.Clear()

	assert.True(t, pytuple.SetItem(tuple, 0, long))
	assert.Equal(t, v, pylong.AsInt(pytuple.GetItem(tuple, 0)))
}
