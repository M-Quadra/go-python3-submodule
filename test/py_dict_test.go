package main

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pydict "github.com/M-Quadra/go-python3-submodule/v8/py-dict"
	pylist "github.com/M-Quadra/go-python3-submodule/v8/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/v8/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/v8/py-tuple"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v8/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyDictCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pydict.Check(nil))
	assert.False(t, pydict.Check(py.None))
	assert.False(t, pydict.CheckExact(nil))
	assert.False(t, pydict.CheckExact(py.None))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()
	assert.False(t, pydict.Check(list))
	assert.False(t, pydict.CheckExact(list))

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.True(t, pydict.Check(dic))
	assert.True(t, pydict.CheckExact(dic))
}

func TestPyDictProxyNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.ProxyNew(nil))
	assert.Nil(t, pydict.ProxyNew(py.None))

	str := pyunicode.FromString("1")
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	dicA := pydict.New()
	defer py.DecRef(dicA)
	defer func() { assert.Equal(t, 1, py.RefCnt(dicA)) }()

	dicB := pydict.ProxyNew(dicA)
	defer py.DecRef(dicB)
	defer func() { assert.Equal(t, 1, py.RefCnt(dicB)) }()

	assert.True(t, pydict.SetItemString(dicA, "1", str))
	assert.False(t, pydict.SetItemString(dicB, "1", str))
}

func TestPyDictClear(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pydict.Clear(nil)

	v := pylong.FromInt(1)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	pydict.SetItemString(dic, "1", v)

	vA := pydict.GetItemString(dic, "1")
	vARedCnt := py.RefCnt(vA)
	defer func() { assert.Equal(t, vARedCnt-1, py.RefCnt(vA)) }()
	assert.Equal(t, 1, pylong.AsInt(vA))

	pydict.Clear(dic)

	vB := pydict.GetItemString(dic, "1")
	assert.True(t, vB == nil)
}

func TestPyDictContains(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()
	assert.True(t, pydict.SetItem(dic, k, k))
	assert.Equal(t, 1, pydict.Contains(dic, k))

	assert.Equal(t, 0, pydict.Contains(dic, nil))
	assert.Equal(t, -1, pydict.Contains(nil, dic))
	assert.Equal(t, 0, pydict.Contains(dic, py.None))
	assert.Equal(t, -1, pydict.Contains(py.None, dic))
}

func TestPyDictCopy(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Copy(nil))

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dicA := pydict.New()
	defer py.DecRef(dicA)
	defer func() { assert.Equal(t, 1, py.RefCnt(dicA)) }()

	assert.True(t, pydict.SetItem(dicA, k, k))

	dicB := pydict.Copy(dicA)
	defer py.DecRef(dicB)
	defer func() { assert.Equal(t, 1, py.RefCnt(dicB)) }()

	assert.True(t, unsafe.Pointer(dicA) != unsafe.Pointer(dicB))

	pydict.Clear(dicA)
	assert.Equal(t, 1, pydict.Contains(dicB, k))
}

func TestPyDictSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.False(t, pydict.SetItemString(dic, "", nil))
	assert.False(t, pydict.SetItem(dic, nil, nil))

	assert.False(t, pydict.SetItem(nil, k, k))
	assert.True(t, pydict.SetItem(dic, k, k))
	assert.True(t, pydict.SetItemString(dic, "", k))
}

func TestPyDictDelItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	v := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.False(t, pydict.DelItem(dic, nil))

	assert.False(t, pydict.DelItem(nil, k))
	assert.False(t, pydict.DelItem(dic, k))
	assert.False(t, pydict.DelItemString(dic, ""))

	pydict.SetItem(dic, k, v)
	pydict.SetItemString(dic, "", v)

	assert.True(t, pydict.DelItem(dic, k))
	assert.True(t, pydict.DelItemString(dic, ""))

	assert.False(t, pydict.DelItem(dic, k))
	assert.False(t, pydict.DelItemString(dic, ""))
}

func TestPyDictGetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	assert.Nil(t, pydict.GetItem(dic, nil))
	assert.Nil(t, pydict.GetItemWithError(dic, nil))

	assert.True(t, pydict.SetItem(dic, k, k))
	assert.True(t, pydict.SetItemString(dic, "", k))

	assert.Nil(t, pydict.GetItem(nil, k))
	v := pydict.GetItem(dic, k)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))

	v = pydict.GetItemString(dic, "")
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))
	v = pydict.GetItemWithError(dic, v)
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))
}

func TestPyDictSetDefault(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	vRand := kRand + 1 + rand.Intn(100)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	pydict.SetItem(dic, k, v)
	res := pydict.SetDefault(dic, k, k)
	resRefCnt := py.RefCnt(res)
	defer func() { assert.Equal(t, resRefCnt, py.RefCnt(res)) }()
	assert.Equal(t, vRand, pylong.AsInt(res))

	assert.Nil(t, pydict.SetDefault(dic, nil, nil))
	assert.Nil(t, pydict.SetDefault(nil, k, v))
}

func TestPyDictItems(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Items(nil))

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	vRand := kRand + 1 + rand.Intn(100)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	pydict.SetItem(dic, k, v)

	list := pydict.Items(dic)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()

	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	kv := pylist.GetItem(list, 0)
	defer py.DecRef(kv)
	defer func() { assert.Equal(t, 1, py.RefCnt(kv)) }()
	assert.True(t, pytuple.Check(kv))

	assert.Equal(t, kRand, pylong.AsInt(pytuple.GetItem(kv, 0)))
	assert.Equal(t, vRand, pylong.AsInt(pytuple.GetItem(kv, 1)))
}

func TestPyDictKeys(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Keys(nil))

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	pydict.SetItem(dic, k, k)

	list := pydict.Keys(dic)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()
	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	assert.Equal(t, kRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyDictValues(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Values(nil))

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)
	kRefCnt := py.RefCnt(k)
	defer func() { assert.Equal(t, kRefCnt, py.RefCnt(k)) }()

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	pydict.SetItem(dic, k, k)

	list := pydict.Values(dic)
	defer py.DecRef(list)
	defer func() { assert.Equal(t, 1, py.RefCnt(list)) }()
	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	assert.Equal(t, kRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyDictSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1, pydict.Size(nil))

	dic := pydict.New()
	defer py.DecRef(dic)
	defer func() { assert.Equal(t, 1, py.RefCnt(dic)) }()

	keyMap := map[int]struct{}{}
	cnt := rand.Intn(100)
	for i := 0; i < cnt; i++ {
		keyMap[rand.Intn(1e5)] = struct{}{}
	}

	for k := range keyMap {
		pydict.SetItem(dic, pylong.FromInt(k), pylong.FromInt(k))
	}

	assert.Equal(t, len(keyMap), pydict.Size(dic))
}
