package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pybytes "github.com/M-Quadra/go-python3-submodule/v11/py-bytes"
	pylist "github.com/M-Quadra/go-python3-submodule/v11/py-list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPyBytesCheck(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	require.False(t, pybytes.Check(nil))
	require.False(t, pybytes.CheckExact(nil))

	list := pylist.New(0)
	defer py.DecRef(list)
	defer func() { require.Equal(t, 1, py.RefCnt(list)) }()

	require.False(t, pybytes.Check(list))
	require.False(t, pybytes.CheckExact(list))

	bytes := pybytes.FromString("aria")
	defer py.DecRef(bytes)
	defer func() { require.Equal(t, 1, py.RefCnt(bytes)) }()

	require.True(t, pybytes.Check(bytes))
	require.True(t, pybytes.CheckExact(bytes))
}

func TestPyBytesFromString(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	str := "hidan"
	bytes := pybytes.FromString(str)
	defer py.DecRef(bytes)
	defer func() { assert.Equal(t, 1, py.RefCnt(bytes)) }()

	require.Equal(t, str, pybytes.AsString(bytes))
}

func TestPyBytesSize(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	strA := "no"
	bytesA := pybytes.FromString(strA)
	defer py.DecRef(bytesA)
	defer func() { require.Equal(t, 1, py.RefCnt(bytesA)) }()

	require.Equal(t, len(strA), pybytes.Size(bytesA))

	strB := "の"
	bytesB := pybytes.FromString(strB)
	defer py.DecRef(bytesB)
	defer func() { require.Equal(t, 1, py.RefCnt(bytesB)) }()

	require.Equal(t, len(strB), pybytes.Size(bytesB))
}

func subTestPyBytesConcatAndDel(
	t *testing.T,
	strA string,
	strB string,
) (initRefCntA, initRefCntB int) {
	bytesA := pybytes.FromString(strA)
	refCntA := py.RefCnt(bytesA)
	defer func() {
		assert.True(t, py.RefCnt(bytesA) != refCntA)
		assert.True(t, py.RefCnt(bytesA) < math.MinInt32 ||
			py.RefCnt(bytesA) > math.MaxInt32 ||
			py.RefCnt(bytesA) == refCntA-1)
	}()

	bytesB := pybytes.FromString(strB)
	refCntB := py.RefCnt(bytesB)
	defer func() {
		assert.True(t, py.RefCnt(bytesB) != refCntB)
		assert.True(t, py.RefCnt(bytesB) < math.MinInt32 ||
			py.RefCnt(bytesB) > math.MaxInt32 ||
			py.RefCnt(bytesB) == refCntB-1)
	}()

	bytesC := pybytes.ConcatAndDel(bytesA, bytesB)
	defer py.DecRef(bytesC)
	defer require.Equal(t, 1, py.RefCnt(bytesC))

	require.NotEqual(t, unsafe.Pointer(bytesA), unsafe.Pointer(bytesC))
	require.NotEqual(t, unsafe.Pointer(bytesB), unsafe.Pointer(bytesC))
	require.Equal(t, strA+strB, pybytes.AsString(bytesC))
	return refCntA, refCntB
}

func TestPyBytesConcatAndDel(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	const maxCnt = 100
	lastRefCntA, lastRefCntB := -1, -1

	// refCnt != 1
	for i := 0; i < maxCnt; i++ {
		refCntA, refCntB := subTestPyBytesConcatAndDel(t, "a", "b")
		if i == 0 {
			lastRefCntA, lastRefCntB = refCntA, refCntB
		} else {
			require.Equal(t, lastRefCntA, refCntA)
			require.Equal(t, lastRefCntB, refCntB)
		}
	}

	// refCnt == 1
	for i := 0; i < maxCnt; i++ {
		a, b := make([]byte, 64), make([]byte, 64)
		_, err := _rand.Read(a)
		require.NoError(t, err)
		_, err = _rand.Read(b)
		require.NoError(t, err)

		strA := hex.EncodeToString(a)
		strB := hex.EncodeToString(b)

		refCntA, refCntB := subTestPyBytesConcatAndDel(t, strA, strB)
		if i == 0 {
			lastRefCntA, lastRefCntB = refCntA, refCntB
		} else {
			require.Equal(t, lastRefCntA, refCntA)
			require.Equal(t, lastRefCntB, refCntB)
		}
	}
}
