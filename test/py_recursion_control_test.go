package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyRecursiveCall(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Zero(t, py.EnterRecursiveCall("in test function"))
	py.LeaveRecursiveCall()
}

func TestPyRepr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := pyunicode.FromString("hello world")
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	assert.Zero(t, py.ReprEnter(str))
	assert.True(t, py.ReprEnter(str) > 0)

	py.ReprLeave(str)
	py.ReprLeave(str)
}
