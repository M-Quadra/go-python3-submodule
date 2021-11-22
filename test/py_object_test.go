package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyexc "github.com/M-Quadra/go-python3-submodule/py-exc"
	pyimport "github.com/M-Quadra/go-python3-submodule/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
	pyobject "github.com/M-Quadra/go-python3-submodule/py-object"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyObjectHasAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)

	namePy := pyunicode.FromString("stdout")
	defer py.DecRef(namePy)
	assert.True(t, pyobject.HasAttr(sysPy, namePy))

	assert.False(t, pyobject.HasAttr(sysPy, nil))
	assert.False(t, pyobject.HasAttr(nil, namePy))
}

func TestPyObjectHasAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)
	assert.True(t, pyobject.HasAttrString(sysPy, "stdout"))

	assert.False(t, pyobject.HasAttrString(sysPy, ""))
	assert.False(t, pyobject.HasAttrString(nil, "stdout"))
}

func TestPyObjectGetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)

	namePy := pyunicode.FromString("stdout")
	defer py.DecRef(namePy)
	assert.NotNil(t, pyobject.GetAttr(sysPy, namePy))

	assert.Nil(t, pyobject.GetAttr(nil, namePy))
	assert.Nil(t, pyobject.GetAttr(sysPy, nil))
}

func TestPyObjectGetAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)
	assert.NotNil(t, pyobject.GetAttrString(sysPy, "stdout"))

	assert.Nil(t, pyobject.GetAttrString(nil, "stdout"))
	assert.Nil(t, pyobject.GetAttrString(sysPy, ""))
}

func TestPyObjectGenericGetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectSetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	name := "stdout"
	namePy := pyunicode.FromString(name)
	defer py.DecRef(namePy)
	assert.NotNil(t, namePy)

	stdoutPy := pyobject.GetAttr(sysPy, namePy)
	assert.NotNil(t, stdoutPy)

	assert.True(t, pyobject.HasAttr(sysPy, namePy))
	assert.False(t, pyobject.SetAttr(nil, namePy, stdoutPy))
	assert.False(t, pyobject.SetAttr(sysPy, nil, stdoutPy))
	assert.True(t, pyobject.SetAttr(sysPy, namePy, nil))
	assert.False(t, pyobject.HasAttr(sysPy, namePy))

	assert.True(t, pyobject.SetAttr(sysPy, namePy, stdoutPy))
	assert.True(t, pyobject.HasAttr(sysPy, namePy))
}

func TestPyObjectSetAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	name := "stdout"
	stdoutPy := pyobject.GetAttrString(sysPy, name)
	assert.NotNil(t, stdoutPy)

	assert.True(t, pyobject.HasAttrString(sysPy, name))
	assert.False(t, pyobject.SetAttrString(nil, name, stdoutPy))
	assert.True(t, pyobject.SetAttrString(sysPy, name, nil))
	assert.False(t, pyobject.HasAttrString(sysPy, name))

	assert.False(t, pyobject.HasAttrString(sysPy, ""))
	assert.True(t, pyobject.SetAttrString(sysPy, "", stdoutPy))
	assert.True(t, pyobject.HasAttrString(sysPy, ""))
	assert.True(t, pyobject.SetAttrString(sysPy, "", nil))
	assert.False(t, pyobject.HasAttrString(sysPy, ""))

	assert.True(t, pyobject.SetAttrString(sysPy, name, stdoutPy))
	assert.True(t, pyobject.HasAttrString(sysPy, name))
}

func TestPyObjectGenericSetAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectDelAttr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)

	name := "stdout"
	namePy := pyunicode.FromString(name)
	defer py.DecRef(namePy)

	stdoutPy := pyobject.GetAttr(sysPy, namePy)

	assert.True(t, pyobject.HasAttr(sysPy, namePy))
	assert.True(t, pyobject.DelAttr(sysPy, namePy))
	assert.False(t, pyobject.HasAttr(sysPy, namePy))

	assert.False(t, pyobject.DelAttr(nil, namePy))
	assert.False(t, pyobject.DelAttr(sysPy, nil))

	assert.True(t, pyobject.SetAttr(sysPy, namePy, stdoutPy))
	assert.True(t, pyobject.HasAttr(sysPy, namePy))
}

func TestPyObjectDelAttrString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)

	name := "stdout"
	stdoutPy := pyobject.GetAttrString(sysPy, name)

	assert.True(t, pyobject.HasAttrString(sysPy, name))
	assert.True(t, pyobject.DelAttrString(sysPy, name))
	assert.False(t, pyobject.HasAttrString(sysPy, name))

	assert.False(t, pyobject.DelAttrString(nil, name))
	assert.False(t, pyobject.DelAttrString(sysPy, ""))

	assert.True(t, pyobject.SetAttrString(sysPy, name, stdoutPy))
	assert.True(t, pyobject.HasAttrString(sysPy, name))
}

func TestPyObjectRichCompare(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strAPy := pyunicode.FromString("strA")
	defer py.DecRef(strAPy)
	strBPy := pyunicode.FromString("strB")
	defer py.DecRef(strBPy)

	assert.Equal(t, py.False, pyobject.RichCompare(strAPy, strBPy, py.EQ))
	assert.Equal(t, py.True, pyobject.RichCompare(strAPy, strBPy, py.LE))

	assert.Nil(t, pyobject.RichCompare(strAPy, nil, py.GT))
	assert.Nil(t, pyobject.RichCompare(nil, strBPy, py.LE))
	assert.Nil(t, pyobject.RichCompare(nil, nil, py.EQ))
	assert.Nil(t, pyobject.RichCompare(nil, nil, -1))
	assert.Nil(t, pyobject.RichCompare(strAPy, strBPy, -1))
}

func TestPyObjectRichCompareBool(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strAPy := pyunicode.FromString("strA")
	defer py.DecRef(strAPy)
	strBPy := pyunicode.FromString("strB")
	defer py.DecRef(strBPy)

	assert.Equal(t, 0, pyobject.RichCompareBool(strAPy, strBPy, py.EQ))
	assert.Equal(t, 1, pyobject.RichCompareBool(strAPy, strBPy, py.LE))

	assert.Equal(t, -1, pyobject.RichCompareBool(strAPy, nil, py.GT))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, strBPy, py.LE))
	assert.Equal(t, 1, pyobject.RichCompareBool(nil, nil, py.EQ))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, nil, py.LE))
	assert.Equal(t, -1, pyobject.RichCompareBool(nil, nil, -1))
	assert.Equal(t, -1, pyobject.RichCompareBool(strAPy, strBPy, -1))
}

func TestPyObjectRepr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	reprPy := pyobject.Repr(listPy)
	defer py.DecRef(reprPy)
	assert.Equal(t, "[]", pyunicode.AsString(reprPy))

	reprNilPy := pyobject.Repr(nil)
	defer py.DecRef(reprNilPy)
	assert.Equal(t, "<NULL>", pyunicode.AsString(reprNilPy))
}

func TestPyObjectASCII(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	asciiPy := pyobject.ASCII(listPy)
	defer py.DecRef(asciiPy)
	assert.Equal(t, "[]", pyunicode.AsString(asciiPy))

	asciiNilPy := pyobject.ASCII(nil)
	defer py.DecRef(asciiNilPy)
	assert.Equal(t, "<NULL>", pyunicode.AsString(asciiNilPy))
}

func TestPyObjectStr(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	strPy := pyobject.Str(listPy)
	defer py.DecRef(strPy)
	assert.Equal(t, "[]", pyunicode.AsString(strPy))

	strNilPy := pyobject.Str(nil)
	defer py.DecRef(strNilPy)
	assert.Equal(t, "<NULL>", pyunicode.AsString(strNilPy))
}

func TestPyObjectBytes(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

}

func TestPyObjectIsSubclass(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	assert.Equal(t, 1, pyobject.IsSubclass(pyexc.Warning, pyexc.Exception))
	assert.Equal(t, 0, pyobject.IsSubclass(pyexc.Exception, pyexc.Warning))
	assert.Equal(t, -1, pyobject.IsSubclass(py.False, listPy))

	assert.Equal(t, -1, pyobject.IsSubclass(py.False, nil))
	assert.Equal(t, -1, pyobject.IsSubclass(nil, listPy))
	assert.Equal(t, -1, pyobject.IsSubclass(nil, nil))
}

func TestPyObjectIsInstance(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	assert.Equal(t, -1, pyobject.IsInstance(py.False, nil))
	assert.Equal(t, -1, pyobject.IsInstance(nil, listPy))
	assert.Equal(t, -1, pyobject.IsInstance(nil, nil))
}

func TestPyObjectHash(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strPy := pyunicode.FromString("test string")
	defer py.DecRef(strPy)

	assert.NotEqual(t, -1, pyobject.Hash(strPy))
	assert.Equal(t, -1, pyobject.Hash(nil))
}

func TestPyObjectHashNotImplemented(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	strPy := pyunicode.FromString("test string")
	defer py.DecRef(strPy)

	assert.Equal(t, -1, pyobject.HashNotImplemented(strPy))
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

	vPy := pylong.FromInt(rand.Int())
	defer py.DecRef(vPy)
	assert.Equal(t, pylong.Type, pyobject.Type(vPy))

	assert.Nil(t, pyobject.Type(nil))
}

func TestPyObjectSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defer pyerr.Clear()

	vPy := pylong.FromInt(rand.Int())
	defer py.DecRef(vPy)
	assert.Equal(t, -1, pyobject.Size(vPy))

	l := rand.Intn(100)
	listPy := pylist.New(l)
	defer py.DecRef(listPy)
	assert.Equal(t, l, pyobject.Size(listPy))

	assert.Equal(t, -1, pyobject.Size(nil))
}

func TestPyObjectLength(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defer pyerr.Clear()

	vPy := pylong.FromInt(rand.Int())
	defer py.DecRef(vPy)
	assert.Equal(t, -1, pyobject.Length(vPy))

	l := rand.Intn(100)
	listPy := pylist.New(l)
	defer py.DecRef(listPy)
	assert.Equal(t, l, pyobject.Length(listPy))

	assert.Equal(t, -1, pyobject.Length(nil))
}

func TestPyObjectLengthHint(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	l := rand.Intn(100)
	listPy := pylist.New(l)
	defer py.DecRef(listPy)
	assert.Equal(t, l, pyobject.LengthHint(listPy, 0))

	vPy := pylong.FromInt(l)
	defer py.DecRef(vPy)
	assert.Equal(t, l, pyobject.LengthHint(vPy, l))

	assert.Equal(t, l, pyobject.LengthHint(nil, l))
}

func TestPyObjectGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defer pyerr.Clear()

	dicPy := pydict.New()
	defer py.DecRef(dicPy)

	keyPy := pyunicode.FromString("k")
	defer py.DecRef(keyPy)
	valPy := pyunicode.FromString("v")
	defer py.DecRef(valPy)
	assert.True(t, pydict.SetItem(dicPy, keyPy, valPy))
	assert.Equal(t, valPy, pyobject.GetItem(dicPy, keyPy))

	assert.Nil(t, pyobject.GetItem(dicPy, nil))
	assert.Nil(t, pyobject.GetItem(nil, keyPy))
	assert.Nil(t, pyobject.GetItem(nil, nil))
}

func TestPyObjectSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defer pyerr.Clear()

	dicPy := pydict.New()
	defer py.DecRef(dicPy)

	keyPy := pyunicode.FromString("k")
	defer py.DecRef(keyPy)
	valPy := pyunicode.FromString("v")
	defer py.DecRef(valPy)
	assert.True(t, pyobject.SetItem(dicPy, keyPy, valPy))
	assert.Equal(t, valPy, pyobject.GetItem(dicPy, keyPy))

	assert.False(t, pyobject.SetItem(dicPy, keyPy, nil))
	assert.False(t, pyobject.SetItem(dicPy, nil, valPy))
	assert.False(t, pyobject.SetItem(dicPy, nil, nil))
}

func TestPyObjectDelItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	defer pyerr.Clear()

	dicPy := pydict.New()
	defer py.DecRef(dicPy)

	keyPy := pyunicode.FromString("k")
	defer py.DecRef(keyPy)
	valPy := pyunicode.FromString("v")
	defer py.DecRef(valPy)
	assert.False(t, pyobject.DelItem(dicPy, keyPy))
	assert.True(t, pyobject.SetItem(dicPy, keyPy, valPy))
	assert.Equal(t, valPy, pyobject.GetItem(dicPy, keyPy))

	assert.True(t, pyobject.DelItem(dicPy, keyPy))
	assert.Nil(t, pyobject.GetItem(dicPy, keyPy))

	assert.False(t, pyobject.DelItem(dicPy, nil))
	assert.False(t, pyobject.DelItem(nil, keyPy))
	assert.False(t, pyobject.DelItem(nil, nil))
}

func TestPyObjectDir(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	listPy := pylist.New(0)
	defer py.DecRef(listPy)

	dirPy := pyobject.Dir(listPy)
	defer py.DecRef(dirPy)
	reprPy := pyobject.Repr(dirPy)
	defer py.DecRef(reprPy)

	str := "['__add__', '__class__', '__class_getitem__', '__contains__', '__delattr__', '__delitem__', '__dir__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__gt__', '__hash__', '__iadd__', '__imul__', '__init__', '__init_subclass__', '__iter__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__reversed__', '__rmul__', '__setattr__', '__setitem__', '__sizeof__', '__str__', '__subclasshook__', 'append', 'clear', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort']"
	assert.Equal(t, str, pyunicode.AsString(reprPy))
}

func TestPyObjectGetIter(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	vPy := pylong.FromInt(rand.Int())
	defer py.DecRef(vPy)
	assert.Nil(t, pyobject.GetIter(vPy))
	assert.True(t, pyerr.ExceptionMatches(pyexc.TypeError))
	defer pyerr.Clear()

	listPy := pylist.New(rand.Intn(100))
	defer py.DecRef(listPy)

	iterPy := pyobject.GetIter(listPy)
	defer py.DecRef(iterPy)
	assert.NotNil(t, iterPy)

	assert.Nil(t, pyobject.GetIter(nil))
}
