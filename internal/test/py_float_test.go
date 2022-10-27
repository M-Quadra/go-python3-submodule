package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pyerr "github.com/M-Quadra/go-python3-submodule/v10/py-err"
	pyfloat "github.com/M-Quadra/go-python3-submodule/v10/py-float"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v10/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyFloatCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pyfloat.Check(nil))

	f := pyfloat.FromFloat64(rand.Float64())
	defer py.DecRef(f)
	defer func() { assert.Equal(t, 1, py.RefCnt(f)) }()
	assert.True(t, pyfloat.Check(f))
	assert.True(t, pyfloat.CheckExact(f))

	str := pyunicode.FromString("123")
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()
	assert.False(t, pyfloat.Check(str))
	assert.False(t, pyfloat.CheckExact(str))
}

func TestPyFloatFromString(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pyfloat.FromStringPy(nil))

	v := rand.Float64()
	strF := strconv.FormatFloat(v, 'f', -1, 64)
	str := pyunicode.FromString(strF)
	defer py.DecRef(str)
	strRefCnt := py.RefCnt(str)
	defer func() { assert.Equal(t, strRefCnt, py.RefCnt(str)) }()

	vA := pyfloat.FromString(strF)
	defer py.DecRef(vA)
	vARefCnt := py.RefCnt(vA)
	defer func() { assert.Equal(t, vARefCnt, py.RefCnt(vA)) }()

	vB := pyfloat.FromStringPy(str)
	defer py.DecRef(vB)
	vBRefCnt := py.RefCnt(vB)
	defer func() { assert.Equal(t, vBRefCnt, py.RefCnt(vB)) }()

	assert.True(t, math.Abs(v-pyfloat.AsFloat64(vA)) < 1e8)
	assert.True(t, math.Abs(v-pyfloat.AsFloat64(vB)) < 1e8)
}

func TestPyFloatAsFloat64(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, -1.0, pyfloat.AsFloat64(nil))
	defer pyerr.Clear()

	vRand := rand.Float64()
	f := pyfloat.FromFloat64(vRand)
	defer py.DecRef(f)
	fRefCnt := py.RefCnt(f)
	defer func() { assert.Equal(t, fRefCnt, py.RefCnt(f)) }()

	assert.True(t, math.Abs(pyfloat.AsFloat64(f)-vRand) < 1e8)
}

func TestPyFloatGetInfo(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.NotNil(t, pyfloat.GetInfo())
}

func TestPyFloatGetMinMax(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.True(t, math.Abs(2.2250738585072014e-308-pyfloat.GetMin()) < 1e8)
	assert.True(t, math.Abs(math.MaxFloat64-pyfloat.GetMax()) < 1e8)
}
