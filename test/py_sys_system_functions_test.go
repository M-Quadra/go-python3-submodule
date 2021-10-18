package main

import (
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPySysGetSetObject(t *testing.T) {
	name := "platform"
	platformPy := pysys.GetObject(name)
	assert.True(t, pyunicode.Check(platformPy))
	py.IncRef(platformPy)

	nPlatformPy := pyunicode.FromString("test")
	defer py.DecRef(nPlatformPy)

	assert.True(t, pysys.SetObject(name, nil))
	assert.Nil(t, pysys.GetObject(name))
	assert.True(t, pysys.SetObject(name, nil))

	assert.True(t, pysys.SetObject(name, nPlatformPy))
	assert.Equal(t, nPlatformPy, pysys.GetObject(name))
	assert.True(t, pysys.SetObject(name, platformPy))
}

func TestPySysWarnOptions(t *testing.T) {
	warnoptionsPy := pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptionsPy))

	s := "ignore"
	sPy := pyunicode.FromString(s)
	defer py.DecRef(sPy)

	pysys.AddWarnOption(s)
	warnoptionsPy = pysys.GetObject("warnoptions")
	assert.Equal(t, s, pyunicode.AsString(pylist.GetItem(warnoptionsPy, 0)))

	pysys.ResetWarnOptions()
	warnoptionsPy = pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptionsPy))

	pysys.AddWarnOptionUnicode(nil)
	pysys.AddWarnOptionUnicode(sPy)
	warnoptionsPy = pysys.GetObject("warnoptions")
	assert.Equal(t, s, pyunicode.AsString(pylist.GetItem(warnoptionsPy, 0)))

	pysys.ResetWarnOptions()
	warnoptionsPy = pysys.GetObject("warnoptions")
	assert.Zero(t, pylist.Size(warnoptionsPy))
}

func TestPySysSetPath(t *testing.T) {
	pathPy := pysys.GetObject("path")
	py.IncRef(pathPy)

	pysys.SetPath("test")
	nPathPy := pysys.GetObject("path")
	assert.Equal(t, "test", pyunicode.AsString(pylist.GetItem(nPathPy, 0)))

	pysys.SetObject("path", pathPy)
}

func TestPySysXOption(t *testing.T) {
	pysys.AddXOption("faulthandler")

	xOptionsPy := pysys.GetXOptions()
	faulthandler := pydict.GetItemString(xOptionsPy, "faulthandler")

	assert.Equal(t, py.True, faulthandler)
}
