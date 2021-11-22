package main

import (
	"fmt"
	"math/rand"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
	pytuple "github.com/M-Quadra/go-python3-submodule/py-tuple"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
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
	assert.False(t, pydict.Check(list))
	assert.False(t, pydict.CheckExact(list))

	dic := pydict.New()
	defer py.DecRef(dic)

	assert.True(t, pydict.Check(dic))
	assert.True(t, pydict.CheckExact(dic))
}
func TestPyDictProxyNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.ProxyNew(nil))
	assert.Nil(t, pydict.ProxyNew(py.None))

	dicA := pydict.New()
	defer py.DecRef(dicA)

	dicB := pydict.ProxyNew(dicA)
	defer py.DecRef(dicB)

	str := pyunicode.FromString("1")
	defer py.DecRef(str)

	assert.True(t, pydict.SetItemString(dicA, "1", str))
	assert.False(t, pydict.SetItemString(dicB, "1", str))
}

func TestPyDictClear(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pydict.Clear(nil)

	v := pylong.FromInt(1)
	defer py.DecRef(v)

	dic := pydict.New()
	defer py.DecRef(dic)
	pydict.SetItemString(dic, "1", v)

	vA := pydict.GetItemString(dic, "1")
	defer py.DecRef(vA)
	assert.Equal(t, 1, pylong.AsInt(vA))

	pydict.Clear(dic)

	vB := pydict.GetItemString(dic, "1")
	defer py.DecRef(vB)
	assert.True(t, vB == nil)
}

func TestPyDictContains(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	dic := pydict.New()
	defer py.DecRef(dic)

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
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

	dicA := pydict.New()
	defer py.DecRef(dicA)

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	assert.True(t, pydict.SetItem(dicA, k, k))

	dicB := pydict.Copy(dicA)
	defer py.DecRef(dicB)
	assert.True(t, unsafe.Pointer(dicA) != unsafe.Pointer(dicB))

	pydict.Clear(dicA)
	assert.Equal(t, 1, pydict.Contains(dicB, k))
}

func TestPyDictSetItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	dic := pydict.New()
	defer py.DecRef(dic)

	assert.False(t, pydict.SetItemString(dic, "", nil))
	assert.False(t, pydict.SetItem(dic, nil, nil))

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	assert.False(t, pydict.SetItem(nil, k, k))
	assert.True(t, pydict.SetItem(dic, k, k))
	assert.True(t, pydict.SetItemString(dic, "", k))
}

func TestPyDictDelItem(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	dic := pydict.New()
	defer py.DecRef(dic)

	assert.False(t, pydict.DelItem(dic, nil))

	k := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(k)
	v := pylong.FromInt(rand.Intn(100))
	defer py.DecRef(v)

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

	dic := pydict.New()
	defer py.DecRef(dic)

	assert.Nil(t, pydict.GetItem(dic, nil))
	assert.Nil(t, pydict.GetItemWithError(dic, nil))

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)
	assert.True(t, pydict.SetItem(dic, k, k))
	assert.True(t, pydict.SetItemString(dic, "", k))

	assert.Nil(t, pydict.GetItem(nil, k))
	v := pydict.GetItem(dic, k)
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))

	v = pydict.GetItemString(dic, "")
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))
	v = pydict.GetItemWithError(dic, v)
	assert.True(t, unsafe.Pointer(k) == unsafe.Pointer(v))
}

func TestPyDictSetDefault(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	dic := pydict.New()
	defer py.DecRef(dic)

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)

	vRand := kRand + 1 + rand.Intn(100)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)

	pydict.SetItem(dic, k, v)
	res := pydict.SetDefault(dic, k, k)
	assert.Equal(t, vRand, pylong.AsInt(res))

	assert.Nil(t, pydict.SetDefault(dic, nil, nil))
	assert.Nil(t, pydict.SetDefault(nil, k, v))
}

func TestPyDictItems(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Items(nil))

	dic := pydict.New()
	defer py.DecRef(dic)

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)

	vRand := kRand + 1 + rand.Intn(100)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)

	pydict.SetItem(dic, k, v)

	list := pydict.Items(dic)
	defer py.DecRef(list)

	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	kv := pylist.GetItem(list, 0)
	assert.True(t, pytuple.Check(kv))

	assert.Equal(t, kRand, pylong.AsInt(pytuple.GetItem(kv, 0)))
	assert.Equal(t, vRand, pylong.AsInt(pytuple.GetItem(kv, 1)))
}

func TestPyDictKeys(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Keys(nil))

	dic := pydict.New()
	defer py.DecRef(dic)

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)

	pydict.SetItem(dic, k, k)

	list := pydict.Keys(dic)
	defer py.DecRef(list)
	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	assert.Equal(t, kRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyDictValues(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pydict.Values(nil))

	dic := pydict.New()
	defer py.DecRef(dic)

	kRand := rand.Intn(100)
	k := pylong.FromInt(kRand)
	defer py.DecRef(k)

	pydict.SetItem(dic, k, k)

	list := pydict.Values(dic)
	defer py.DecRef(list)
	assert.True(t, pylist.Check(list))
	assert.Equal(t, 1, pylist.Size(list))

	assert.Equal(t, kRand, pylong.AsInt(pylist.GetItem(list, 0)))
}

func TestPyDictSize(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1, pydict.Size(nil))

	dic := pydict.New()
	defer py.DecRef(dic)

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
