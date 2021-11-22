package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pyimport "github.com/M-Quadra/go-python3-submodule/py-import"
	pymodule "github.com/M-Quadra/go-python3-submodule/py-module"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyModuleCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pymodule.Check(nil))
	assert.False(t, pymodule.CheckExact(nil))

	modelPy := pymodule.New("test_module")
	defer py.DecRef(modelPy)
	assert.True(t, pymodule.Check(modelPy))
	assert.True(t, pymodule.CheckExact(modelPy))
}

func TestPyModuleNewObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.NewObject(nil))

	namePy := pyunicode.FromString("test_module")
	defer py.DecRef(namePy)
	assert.NotNil(t, namePy)

	modelPy := pymodule.NewObject(namePy)
	defer py.DecRef(modelPy)
	assert.NotNil(t, modelPy)
}

func TestPyModuleNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	modelPy := pymodule.New("test_module")
	defer py.DecRef(modelPy)
	assert.NotNil(t, modelPy)
}

func TestPyModuleGetDict(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetDict(nil))

	name := "sys"
	namePy := pyunicode.FromString(name)
	defer py.DecRef(namePy)
	assert.NotNil(t, namePy)

	sysPy := pyimport.ImportModule(name)
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	dicPy := pymodule.GetDict(sysPy)
	assert.True(t, pydict.Check(dicPy))
}

func TestPyModuleGetNameObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetNameObject(nil))

	name := "sys"
	sysPy := pyimport.ImportModule(name)
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	namePy := pymodule.GetNameObject(sysPy)
	assert.Equal(t, name, pyunicode.AsString(namePy))
}

func TestPyModuleGetName(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, "", pymodule.GetName(nil))

	name := "sys"
	sysPy := pyimport.ImportModule(name)
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	assert.Equal(t, name, pymodule.GetName(sysPy))
}

func TestPyModuleGetState(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.True(t, pymodule.GetState(nil) == nil)

	sysPy := pyimport.ImportModule("sys")
	defer py.DecRef(sysPy)
	assert.NotNil(t, sysPy)

	assert.True(t, pymodule.GetState(sysPy) == nil)
}

func TestPyModuleGetFilenameObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetFilenameObject(nil))

	testPy := pyimport.ImportModule("test")
	defer py.DecRef(testPy)

	namePy := pymodule.GetFilenameObject(testPy)
	defer py.DecRef(namePy)
	assert.NotNil(t, namePy)

	assert.True(t, strings.HasSuffix(pyunicode.AsString(namePy), "test/__init__.py"))
}

func TestPyModuleGetFilename(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	pymodule.GetFilename(nil)

	testPy := pyimport.ImportModule("test")
	defer py.DecRef(testPy)

	name := pymodule.GetFilename(testPy)
	assert.True(t, strings.HasSuffix(name, "test/__init__.py"))
}
