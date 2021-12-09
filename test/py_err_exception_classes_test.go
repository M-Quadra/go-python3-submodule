package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pydict "github.com/M-Quadra/go-python3-submodule/v8/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/v8/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/v8/py-exc"
	pylong "github.com/M-Quadra/go-python3-submodule/v8/py-long"
	"github.com/stretchr/testify/assert"
)

func TestPyErrNewException(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.NewException("", nil, nil))
	assert.Nil(t, pyerr.NewException("module.class", nil, nil))
	pyerr.Clear()

	{
		assert.Nil(t, pyerr.Occurred())

		exc := pyerr.NewException("module.class_0", nil, nil)
		defer py.DecRef(exc)
		assert.NotNil(t, exc)

		pyerr.SetNone(exc)
		assert.NotNil(t, pyerr.Occurred())
		pyerr.Print()
		assert.Nil(t, pyerr.Occurred())
	}

	{
		assert.Nil(t, pyerr.Occurred())

		exc := pyerr.NewException("module.class_1", pyexc.BaseException, nil)
		defer py.DecRef(exc)
		assert.NotNil(t, exc)

		pyerr.SetNone(exc)
		assert.NotNil(t, pyerr.Occurred())
		pyerr.Print()
		assert.Nil(t, pyerr.Occurred())
	}

	{
		assert.Nil(t, pyerr.Occurred())

		dic := pydict.New()
		defer py.DecRef(dic)
		defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()
		pydict.SetItemString(dic, strconv.Itoa(rand.Intn(100)), pylong.FromInt(rand.Intn(100)))

		exc := pyerr.NewException("module.class_2", nil, dic)
		defer py.DecRef(exc)
		pyerr.Print()
		assert.NotNil(t, exc)

		pyerr.SetNone(exc)
		assert.NotNil(t, pyerr.Occurred())
		pyerr.Print()
		assert.Nil(t, pyerr.Occurred())
	}
}

func TestPyErrNewExceptionWithDoc(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyerr.NewExceptionWithDoc("", "", nil, nil))
	pyerr.Clear()

	exc := pyerr.NewExceptionWithDoc("module.class", "doc", nil, nil)
	defer py.DecRef(exc)
	assert.NotNil(t, exc)
}
