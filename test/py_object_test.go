package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pydict "github.com/M-Quadra/go-python3-submodule/v8/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/v8/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/v8/py-exc"
	pyimport "github.com/M-Quadra/go-python3-submodule/v8/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v8/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/v8/py-long"
	pyobject "github.com/M-Quadra/go-python3-submodule/v8/py-object"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v8/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyObjectHasAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	name := pyunicode.FromString("stdout")
	defer py.DecRef(name)
	nameRefCnt := py.RefCnt(name)

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()

	assert.Equal(t, nameRefCnt, py.RefCnt(name))
	assert.True(t, pyobject.HasAttr(sys, name)) // reference count + 1
	assert.Equal(t, nameRefCnt+1, py.RefCnt(name))

	assert.False(t, pyobject.HasAttr(sys, nil))
	assert.False(t, pyobject.HasAttr(nil, name))
}

func TestPyObjectHasAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()

	assert.True(t, pyobject.HasAttrString(sys, "stdout"))

	assert.False(t, pyobject.HasAttrString(sys, ""))
	assert.False(t, pyobject.HasAttrString(nil, "stdout"))
}

func TestPyObjectGetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	name := pyunicode.FromString("stdout")
	defer py.DecRef(name)
	nameRefCnt := py.RefCnt(name)

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()

	assert.Equal(t, nameRefCnt, py.RefCnt(name))
	assert.NotNil(t, pyobject.GetAttr(sys, name)) // reference count + 1
	assert.Equal(t, nameRefCnt+1, py.RefCnt(name))

	assert.Nil(t, pyobject.GetAttr(nil, name))
	assert.Nil(t, pyobject.GetAttr(sys, nil))
}

func TestPyObjectGetAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	name := pyimport.ImportModule("sys")
	defer py.DecRef(name)
	nameRefCnt := py.RefCnt(name)
	defer func() { assert.Equal(t, nameRefCnt, py.RefCnt(name)) }()

	assert.NotNil(t, pyobject.GetAttrString(name, "stdout"))

	assert.Nil(t, pyobject.GetAttrString(nil, "stdout"))
	assert.Nil(t, pyobject.GetAttrString(name, ""))
}

func TestPyObjectGenericGetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectSetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	name := pyunicode.FromString("stdout")
	defer py.DecRef(name)
	assert.NotNil(t, name)

	stdout := pyobject.GetAttr(sys, name)
	assert.NotNil(t, stdout)

	assert.True(t, pyobject.HasAttr(sys, name))
	assert.False(t, pyobject.SetAttr(nil, name, stdout))
	assert.False(t, pyobject.SetAttr(sys, nil, stdout))
	assert.True(t, pyobject.SetAttr(sys, name, nil))
	assert.False(t, pyobject.HasAttr(sys, name))

	assert.True(t, pyobject.SetAttr(sys, name, stdout))
	assert.True(t, pyobject.HasAttr(sys, name))
}

func TestPyObjectSetAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	name := "stdout"
	stdout := pyobject.GetAttrString(sys, name)
	stdoutRefCnt := py.RefCnt(stdout)
	defer func() { assert.Equal(t, stdoutRefCnt, py.RefCnt(stdout)) }()
	assert.NotNil(t, stdout)

	assert.True(t, pyobject.HasAttrString(sys, name))
	assert.False(t, pyobject.SetAttrString(nil, name, stdout))
	assert.True(t, pyobject.SetAttrString(sys, name, nil))
	assert.False(t, pyobject.HasAttrString(sys, name))

	assert.False(t, pyobject.HasAttrString(sys, ""))
	assert.True(t, pyobject.SetAttrString(sys, "", stdout))
	assert.True(t, pyobject.HasAttrString(sys, ""))
	assert.True(t, pyobject.SetAttrString(sys, "", nil))
	assert.False(t, pyobject.HasAttrString(sys, ""))

	assert.True(t, pyobject.SetAttrString(sys, name, stdout))
	assert.True(t, pyobject.HasAttrString(sys, name))
}

func TestPyObjectGenericSetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectDelAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()

	name := pyunicode.FromString("stdout")
	defer py.DecRef(name)

	stdout := pyobject.GetAttr(sys, name)
	stdoutRefCnt := py.RefCnt(stdout)
	defer func() { assert.Equal(t, stdoutRefCnt, py.RefCnt(stdout)) }()

	assert.True(t, pyobject.HasAttr(sys, name))
	assert.True(t, pyobject.DelAttr(sys, name))
	assert.False(t, pyobject.HasAttr(sys, name))

	assert.False(t, pyobject.DelAttr(nil, name))
	assert.False(t, pyobject.DelAttr(sys, nil))

	assert.True(t, pyobject.SetAttr(sys, name, stdout))
	assert.True(t, pyobject.HasAttr(sys, name))
}

func TestPyObjectDelAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()

	name := "stdout"
	stdout := pyobject.GetAttrString(sys, name)

	assert.True(t, pyobject.HasAttrString(sys, name))
	assert.True(t, pyobject.DelAttrString(sys, name))
	assert.False(t, pyobject.HasAttrString(sys, name))

	assert.False(t, pyobject.DelAttrString(nil, name))
	assert.False(t, pyobject.DelAttrString(sys, ""))

	assert.True(t, pyobject.SetAttrString(sys, name, stdout))
	assert.True(t, pyobject.HasAttrString(sys, name))
}

func TestPyObjectRichCompare(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strA := pyunicode.FromString("strA")
	defer py.DecRef(strA)
	strARefCnt := py.RefCnt(strA)
	defer func() { assert.Equal(t, strARefCnt, py.RefCnt(strA)) }()

	strB := pyunicode.FromString("strB")
	defer py.DecRef(strB)
	strBRefCnt := py.RefCnt(strB)
	defer func() { assert.Equal(t, strBRefCnt, py.RefCnt(strB)) }()

	assert.Equal(t, py.False, pyobject.RichCompare(strA, strB, py.EQ))
	assert.Equal(t, py.True, pyobject.RichCompare(strA, strB, py.LE))

	assert.Nil(t, pyobject.RichCompare(strA, nil, py.GT))
	assert.Nil(t, pyobject.RichCompare(nil, strB, py.LE))
	assert.Nil(t, pyobject.RichCompare(nil, nil, py.EQ))
	assert.Nil(t, pyobject.RichCompare(nil, nil, -1))
	assert.Nil(t, pyobject.RichCompare(strA, strB, -1))
}

func TestPyObjectRichCompareBool(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strA := pyunicode.FromString("strA")
	defer py.DecRef(strA)
	strARefCnt := py.RefCnt(strA)
	defer func() { assert.Equal(t, strARefCnt, py.RefCnt(strA)) }()

	strB := pyunicode.FromString("strB")
	defer py.DecRef(strB)
	strBRefCnt := py.RefCnt(strB)
	defer func() { assert.Equal(t, strBRefCnt, py.RefCnt(strB)) }()

	assert.Equal(t, 0, pyobject.RichCompareBool(strA, strB, py.EQ))
	assert.Equal(t, 1, pyobject.RichCompareBool(strA, strB, py.LE))

	assert.Equal(t, -1, pyobject.RichCompareBool(strA, nil, py.GT))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, strB, py.LE))
	assert.Equal(t, 1, pyobject.RichCompareBool(nil, nil, py.EQ))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, nil, py.LE))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, nil, -1))
	assert.Equal(t, -1, pyobject.RichCompareBool(strA, strB, -1))
}

func TestPyObjectRepr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	repr := pyobject.Repr(list)
	defer py.DecRef(repr)
	defer func() { assert.Equal(t, 1, py.RefCnt(repr)) }()
	assert.Equal(t, "[]", pyunicode.AsString(repr))

	reprNil := pyobject.Repr(nil)
	defer py.DecRef(reprNil)
	defer func() { assert.Equal(t, 1, py.RefCnt(reprNil)) }()
	assert.Equal(t, "<NULL>", pyunicode.AsString(reprNil))
}

func TestPyObjectASCII(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	ascii := pyobject.ASCII(list)
	defer py.DecRef(ascii)
	defer func() { assert.Equal(t, 1, py.RefCnt(ascii)) }()
	assert.Equal(t, "[]", pyunicode.AsString(ascii))

	asciiNil := pyobject.ASCII(nil)
	defer py.DecRef(asciiNil)
	defer func() { assert.Equal(t, 1, py.RefCnt(asciiNil)) }()
	assert.Equal(t, "<NULL>", pyunicode.AsString(asciiNil))
}

func TestPyObjectStr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	str := pyobject.Str(list)
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()
	assert.Equal(t, "[]", pyunicode.AsString(str))

	strNil := pyobject.Str(nil)
	defer py.DecRef(strNil)
	defer func() { assert.Equal(t, 1, py.RefCnt(strNil)) }()
	assert.Equal(t, "<NULL>", pyunicode.AsString(strNil))
}

func TestPyObjectBytes(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectIsSubclass(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, 1, pyobject.IsSubclass(pyexc.Warning, pyexc.Exception))
	assert.Equal(t, 0, pyobject.IsSubclass(pyexc.Exception, pyexc.Warning))
	assert.Equal(t, -1, pyobject.IsSubclass(py.False, list))

	assert.Equal(t, -1, pyobject.IsSubclass(py.False, nil))
	assert.Equal(t, -1, pyobject.IsSubclass(nil, list))
	assert.Equal(t, -1, pyobject.IsSubclass(nil, nil))
}

func TestPyObjectIsInstance(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, -1, pyobject.IsInstance(py.False, nil))
	assert.Equal(t, -1, pyobject.IsInstance(nil, list))
	assert.Equal(t, -1, pyobject.IsInstance(nil, nil))
}

func TestPyObjectHash(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := pyunicode.FromString("test string")
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	assert.NotEqual(t, -1, pyobject.Hash(str))
	assert.Equal(t, -1, pyobject.Hash(nil))
}

func TestPyObjectHashNotImplemented(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	str := pyunicode.FromString("test string")
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	assert.Equal(t, -1, pyobject.HashNotImplemented(str))
	assert.True(t, pyerr.ExceptionMatches(pyexc.TypeError))
	defer pyerr.Clear()

	assert.Equal(t, -1, pyobject.HashNotImplemented(nil))
}

func TestPyObjectIsTrue(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, 1, pyobject.IsTrue(py.True))
	assert.Equal(t, 0, pyobject.IsTrue(py.False))

	assert.Equal(t, -1, pyobject.IsTrue(nil))
}

func TestPyObjectNot(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, 0, pyobject.Not(py.True))
	assert.Equal(t, 1, pyobject.Not(py.False))

	assert.Equal(t, -1, pyobject.Not(nil))
}

func TestPyObjectType(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := pylong.FromInt(rand.Int())
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, pylong.Type, pyobject.Type(v))

	assert.Nil(t, pyobject.Type(nil))
}

func TestPyObjectSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := pylong.FromInt(rand.Int())
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, -1, pyobject.Size(v))
	pyerr.Clear()

	l := rand.Intn(100)
	list := pylist.New(l)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, l, pyobject.Size(list))

	assert.Equal(t, -1, pyobject.Size(nil))
	pyerr.Clear()
}

func TestPyObjectLength(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := pylong.FromInt(rand.Int())
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, -1, pyobject.Length(v))
	pyerr.Clear()

	l := rand.Intn(100)
	list := pylist.New(l)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, l, pyobject.Length(list))

	assert.Equal(t, -1, pyobject.Length(nil))
	pyerr.Clear()
}

func TestPyObjectLengthHint(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	l := rand.Intn(100)
	list := pylist.New(l)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.Equal(t, l, pyobject.LengthHint(list, 0))

	v := pylong.FromInt(l)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, l, pyobject.LengthHint(v, l))
	assert.Equal(t, l, pyobject.LengthHint(nil, l))
}

func TestPyObjectGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	key := pyunicode.FromString("k")
	defer py.DecRef(key)
	keyRefCnt := py.RefCnt(key)
	defer func() { assert.Equal(t, keyRefCnt, py.RefCnt(key)) }()

	val := pyunicode.FromString("v")
	defer py.DecRef(val)
	valRefCnt := py.RefCnt(val)
	defer func() { assert.Equal(t, valRefCnt, py.RefCnt(val)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.Equal(t, valRefCnt, py.RefCnt(val))
	assert.True(t, pydict.SetItem(dic, key, val))
	assert.Equal(t, valRefCnt+1, py.RefCnt(val))

	assert.Equal(t, val, pyobject.GetItem(dic, key))
	defer py.DecRef(val)
	assert.Equal(t, valRefCnt+2, py.RefCnt(val))

	assert.Equal(t, val, pyobject.GetItem(dic, key))
	defer py.DecRef(val)
	assert.Equal(t, valRefCnt+3, py.RefCnt(val))

	assert.Nil(t, pyobject.GetItem(dic, nil))
	pyerr.Clear()
	assert.Nil(t, pyobject.GetItem(nil, key))
	pyerr.Clear()
	assert.Nil(t, pyobject.GetItem(nil, nil))
	pyerr.Clear()
}

func TestPyObjectSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	key := pyunicode.FromString("k")
	defer py.DecRef(key)
	keyRefCnt := py.RefCnt(key)
	defer func() { assert.Equal(t, keyRefCnt, py.RefCnt(key)) }()

	val := pyunicode.FromString("v")
	defer py.DecRef(val)
	valRefCnt := py.RefCnt(val)
	defer func() { assert.Equal(t, valRefCnt, py.RefCnt(val)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.True(t, pyobject.SetItem(dic, key, val))
	assert.Equal(t, val, pydict.GetItem(dic, key))

	assert.False(t, pyobject.SetItem(dic, key, nil))
	pyerr.Clear()
	assert.False(t, pyobject.SetItem(dic, nil, val))
	pyerr.Clear()
	assert.False(t, pyobject.SetItem(dic, nil, nil))
	pyerr.Clear()
}

func TestPyObjectDelItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	key := pyunicode.FromString("k")
	defer py.DecRef(key)
	keyRefCnt := py.RefCnt(key)
	defer func() { assert.Equal(t, keyRefCnt, py.RefCnt(key)) }()

	val := pyunicode.FromString("v")
	defer py.DecRef(val)
	valRefCnt := py.RefCnt(val)
	defer func() { assert.Equal(t, valRefCnt, py.RefCnt(val)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.False(t, pyobject.DelItem(dic, key))
	pyerr.Clear()
	assert.True(t, pyobject.SetItem(dic, key, val))
	assert.Equal(t, val, pydict.GetItem(dic, key))

	assert.True(t, pyobject.DelItem(dic, key))
	assert.Nil(t, pyobject.GetItem(dic, key))
	pyerr.Clear()

	assert.False(t, pyobject.DelItem(dic, nil))
	pyerr.Clear()
	assert.False(t, pyobject.DelItem(nil, key))
	pyerr.Clear()
	assert.False(t, pyobject.DelItem(nil, nil))
	pyerr.Clear()
}

func TestPyObjectDir(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	dirPy := pyobject.Dir(list)
	defer py.DecRef(dirPy)
	defer func() { assert.Equal(t, 1, py.RefCnt(dirPy)) }()

	reprPy := pyobject.Repr(dirPy)
	defer py.DecRef(reprPy)
	defer func() { assert.Equal(t, 1, py.RefCnt(reprPy)) }()

	str := "['__add__', '__class__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']"
	assert.Equal(t, str, pyunicode.AsString(reprPy))
}

func TestPyObjectGetIter(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := pylong.FromInt(rand.Int())
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Nil(t, pyobject.GetIter(v)) // TypeError: 'int' object is not iterable
	assert.True(t, pyerr.ExceptionMatches(pyexc.TypeError))
	pyerr.Clear()

	list := pylist.New(rand.Intn(100))
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	iterPy := pyobject.GetIter(list)
	defer py.DecRef(iterPy)
	defer func() { assert.Equal(t, 1, py.RefCnt(iterPy)) }()
	assert.NotNil(t, iterPy)

	assert.Nil(t, pyobject.GetIter(nil))
}
