package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pybytes "github.com/M-Quadra/go-python3-submodule/v8/py-bytes"
	pyerr "github.com/M-Quadra/go-python3-submodule/v8/py-err"
	pytuple "github.com/M-Quadra/go-python3-submodule/v8/py-tuple"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v8/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyUnicodeCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pyunicode.Check(nil))
	assert.False(t, pyunicode.CheckExact(nil))

	u := pyunicode.FromString("")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.True(t, pyunicode.Check(u))
	assert.True(t, pyunicode.CheckExact(u))

	tuple := pytuple.New(1)
	defer py.DecRef(tuple)
	defer func() { assert.Equal(t, 1, py.RefCnt(tuple)) }()

	assert.False(t, pyunicode.Check(tuple))
	assert.False(t, pyunicode.CheckExact(tuple))
}

func TestPyUnicodeNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyunicode.New(-1, 'z'))
	pyerr.Clear()

	u := pyunicode.New(11, 'z')
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.NotNil(t, u)
}

func TestPyUnicodeFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	u := pyunicode.FromString("AA")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.Equal(t, "AA", pyunicode.AsString(u))
}

func TestPyUnicodeFromEncodedObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyunicode.FromEncodedObject(nil, "utf-8", "strict"))
	pyerr.Clear()

	str := "HiHi"
	b := pybytes.FromString(str)
	defer py.DecRef(b)
	bRefCnt := py.RefCnt(b)
	defer func() { assert.Equal(t, bRefCnt, py.RefCnt(b)) }()

	u := pyunicode.FromEncodedObject(b, "utf-8", "strict")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.Equal(t, str, pyunicode.AsString(u))
}

func TestPyUnicodeGetLength(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, 0, pyunicode.GetLength(nil))

	u := pyunicode.FromString("刃无锋")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.Equal(t, 3, pyunicode.GetLength(u))
}

func TestPyUnicodeCopyCharacters(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	uA := pyunicode.FromString("QB鸦")
	defer py.DecRef(uA)
	uARefCnt := py.RefCnt(uA)
	defer func() { assert.Equal(t, uARefCnt, py.RefCnt(uA)) }()

	uB := pyunicode.FromString("鬼鸟")
	defer py.DecRef(uB)
	uBRefCnt := py.RefCnt(uB)
	defer func() { assert.Equal(t, uBRefCnt, py.RefCnt(uB)) }()

	assert.Equal(t, 2, pyunicode.CopyCharacters(uA, 0, uB, 0, 3))
	assert.Equal(t, "鬼鸟鸦", pyunicode.AsString(uA))

	assert.Equal(t, -1, pyunicode.CopyCharacters(uA, 2, uB, 0, 3)) // SystemError: Cannot write 2 characters at 2 in a string of 3 characters
	pyerr.Clear()
	assert.Equal(t, -1, pyunicode.CopyCharacters(nil, 2, uB, 0, 3))
	assert.Equal(t, -1, pyunicode.CopyCharacters(uA, 2, nil, 0, 3))
}

func TestPyUnicodeFill(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	u := pyunicode.FromString("aaa")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	assert.Equal(t, 2, pyunicode.Fill(u, 1, 2, 'b'))
	assert.Equal(t, "abb", pyunicode.AsString(u))

	assert.Equal(t, 0, pyunicode.Fill(u, 4, 2, 'b'))
	assert.Equal(t, -1, pyunicode.Fill(u, -4, 2, 'b'))
	defer pyerr.Clear()

	assert.Equal(t, 0, pyunicode.Fill(nil, 1, 2, 'b'))
}

func TestPyUnicodeWriteReadChar(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	u := pyunicode.FromString("aaa")
	defer py.DecRef(u)
	uRefCnt := py.RefCnt(u)
	defer func() { assert.Equal(t, uRefCnt, py.RefCnt(u)) }()

	{
		assert.Equal(t, 0, pyunicode.WriteChar(u, 1, 'b'))
		assert.Equal(t, 'b', pyunicode.ReadChar(u, 1))

		assert.Equal(t, -1, pyunicode.WriteChar(u, -1, 'b'))
		pyerr.Clear()
		assert.Equal(t, 0, pyunicode.WriteChar(nil, 1, 'b'))
	}

	{
		assert.Equal(t, int32(-1), pyunicode.ReadChar(u, -1))
		pyerr.Clear()
		assert.Zero(t, pyunicode.ReadChar(nil, 1))

		assert.Equal(t, 'b', pyunicode.ReadChar(u, 1))
	}
}

func TestPyUnicodeSubstring(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	uA := pyunicode.FromString("SF")
	defer py.DecRef(uA)
	uARefCnt := py.RefCnt(uA)
	defer func() { assert.Equal(t, uARefCnt, py.RefCnt(uA)) }()

	uB := pyunicode.Substring(uA, 1, 2)
	defer py.DecRef(uB)
	uBRefCnt := py.RefCnt(uB)
	defer func() { assert.Equal(t, uBRefCnt, py.RefCnt(uB)) }()
	assert.Equal(t, "F", pyunicode.AsString(uB))

	uC := pyunicode.Substring(uA, 1, 3)
	defer py.DecRef(uC)
	uCRefCnt := py.RefCnt(uC)
	defer func() { assert.Equal(t, uCRefCnt, py.RefCnt(uC)) }()
	assert.Equal(t, "F", pyunicode.AsString(uC))

	assert.Nil(t, pyunicode.Substring(uA, 1, -1))
	pyerr.Clear()
	assert.Nil(t, pyunicode.Substring(nil, 1, 2))
}
