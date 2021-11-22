package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pylong "github.com/M-Quadra/go-python3-submodule/py-long"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyLongCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pylong.Check(nil))
	assert.False(t, pylong.CheckExact(nil))

	vPy := pylong.FromInt(rand.Intn(1000))
	defer py.DecRef(vPy)

	assert.True(t, pylong.Check(vPy))
	assert.True(t, pylong.CheckExact(vPy))
}

func TestPyLongFromAsInt(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1, pylong.AsInt(nil))

	v := rand.Intn(1000)
	vPy := pylong.FromInt(v)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsInt(vPy))
}

func TestPyLongFromAsUint(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, uint(math.MaxUint), pylong.AsUint(nil))

	v := uint(rand.Intn(1000))
	vPy := pylong.FromUint(v)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsUint(vPy))
}

func TestPyLongFromAsInt64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, int64(-1), pylong.AsInt64(nil))

	v := rand.Int63()
	vPy := pylong.FromInt64(v)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsInt64(vPy))
}

func TestPyLongFromAsUint64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, uint64(math.MaxUint64), pylong.AsUint64(nil))

	v := rand.Uint64()
	vPy := pylong.FromUint64(v)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsUint64(vPy))
}

func TestPyLongFromAsFloat64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1.0, pylong.AsFloat64(nil))

	v := 3.14
	vPy := pylong.FromFloat64(v)
	defer py.DecRef(vPy)
	assert.True(t, math.Abs(3.0-pylong.AsFloat64(vPy)) < 1e8)
}

func TestPyLongFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	v := rand.Intn(1000)
	vPy := pylong.FromString(strconv.Itoa(v), 10)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsInt(vPy))
}

func TestPyLongFromUnicode(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pylong.FromUnicodeObject(nil, 10))

	v := rand.Intn(1000)
	strPy := pyunicode.FromString(strconv.Itoa(v))
	defer py.DecRef(strPy)

	vPy := pylong.FromUnicodeObject(strPy, 10)
	defer py.DecRef(vPy)
	assert.Equal(t, v, pylong.AsInt(vPy))
}
