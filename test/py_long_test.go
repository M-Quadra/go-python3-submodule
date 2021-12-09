package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v8/py"
	pylong "github.com/M-Quadra/go-python3-submodule/v8/py-long"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v8/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyLongCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylong.Check(nil))
	assert.False(t, pylong.CheckExact(nil))

	v := pylong.FromInt(rand.Intn(1000))
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.True(t, pylong.Check(v))
	assert.True(t, pylong.CheckExact(v))
}

func TestPyLongFromAsInt(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1, pylong.AsInt(nil))

	vRand := rand.Intn(1000)
	v := pylong.FromInt(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsInt(v))
}

func TestPyLongFromAsUint(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, uint(math.MaxUint), pylong.AsUint(nil))

	vRand := uint(rand.Intn(1000))
	v := pylong.FromUint(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsUint(v))
}

func TestPyLongFromAsInt64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, int64(-1), pylong.AsInt64(nil))

	vRand := rand.Int63()
	v := pylong.FromInt64(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsInt64(v))
}

func TestPyLongFromAsUint64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, uint64(math.MaxUint64), pylong.AsUint64(nil))

	vRand := rand.Uint64()
	v := pylong.FromUint64(vRand)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsUint64(v))
}

func TestPyLongFromAsFloat64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1.0, pylong.AsFloat64(nil))

	vFloat64 := 3.14
	v := pylong.FromFloat64(vFloat64)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.True(t, math.Abs(3.0-pylong.AsFloat64(v)) < 1e8)
}

func TestPyLongFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	vRand := rand.Intn(1000)
	v := pylong.FromString(strconv.Itoa(vRand), 10)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsInt(v))
}

func TestPyLongFromUnicode(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pylong.FromUnicodeObject(nil, 10))

	vRand := rand.Intn(1000)
	str := pyunicode.FromString(strconv.Itoa(vRand))
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	v := pylong.FromUnicodeObject(str, 10)
	defer py.DecRef(v)
	vRefCnt := py.RefCnt(v)
	defer func() { assert.Equal(t, vRefCnt, py.RefCnt(v)) }()

	assert.Equal(t, vRand, pylong.AsInt(v))
}
