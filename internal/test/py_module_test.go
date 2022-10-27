package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v10/py"
	pydict "github.com/M-Quadra/go-python3-submodule/v10/py-dict"
	pyimport "github.com/M-Quadra/go-python3-submodule/v10/py-import"
	pymodule "github.com/M-Quadra/go-python3-submodule/v10/py-module"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v10/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyModuleCheck(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.False(t, pymodule.Check(nil))
	assert.False(t, pymodule.CheckExact(nil))

	module := pymodule.New("test_module")
	defer py.DecRef(module)

	assert.True(t, pymodule.Check(module))
	assert.True(t, pymodule.CheckExact(module))
}

func TestPyModuleNewObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.NewObject(nil))

	name := pyunicode.FromString("test_module")
	defer py.DecRef(name)
	nameRefCnt := py.RefCnt(name)
	defer func() { assert.Equal(t, nameRefCnt, py.RefCnt(name)) }()

	assert.NotNil(t, name)

	moduleA := pymodule.NewObject(name)
	defer py.DecRef(moduleA)
	defer func() { assert.Equal(t, 1, py.RefCnt(moduleA)) }()
	assert.NotNil(t, moduleA)

	moduleB := pymodule.NewObject(name)
	defer py.DecRef(moduleB)
	defer func() { assert.Equal(t, 1, py.RefCnt(moduleB)) }()
	assert.NotNil(t, moduleB)
}

func TestPyModuleNew(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	name := "test_module"

	moduleA := pymodule.New(name)
	defer py.DecRef(moduleA)
	defer func() { assert.Equal(t, 1, py.RefCnt(moduleA)) }()
	assert.NotNil(t, moduleA)

	moduleB := pymodule.New(name)
	defer py.DecRef(moduleB)
	defer func() { assert.Equal(t, 1, py.RefCnt(moduleB)) }()
	assert.NotNil(t, moduleB)
}

func TestPyModuleGetDict(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetDict(nil))

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	dic := pymodule.GetDict(sys)
	assert.True(t, pydict.Check(dic))
}

func TestPyModuleGetNameObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetNameObject(nil))

	name := "sys"

	sys := pyimport.ImportModule(name)
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	namePy := pymodule.GetNameObject(sys)
	namePyRefCnt := py.RefCnt(namePy)
	defer func() { assert.Equal(t, namePyRefCnt, py.RefCnt(namePy)) }()

	assert.Equal(t, name, pyunicode.AsString(namePy))
}

func TestPyModuleGetName(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, "", pymodule.GetName(nil))

	name := "sys"
	sys := pyimport.ImportModule(name)
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	assert.Equal(t, name, pymodule.GetName(sys))
}

func TestPyModuleGetState(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.True(t, pymodule.GetState(nil) == nil)

	sys := pyimport.ImportModule("sys")
	defer py.DecRef(sys)
	sysRefCnt := py.RefCnt(sys)
	defer func() { assert.Equal(t, sysRefCnt, py.RefCnt(sys)) }()
	assert.NotNil(t, sys)

	assert.True(t, pymodule.GetState(sys) == nil)
}

func TestPyModuleGetFilenameObject(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Nil(t, pymodule.GetFilenameObject(nil))

	queue := pyimport.ImportModule("queue")
	defer py.DecRef(queue)
	queueRefCnt := py.RefCnt(queue)
	defer func() { assert.Equal(t, queueRefCnt, py.RefCnt(queue)) }()

	name := pymodule.GetFilenameObject(queue)
	nameRefCnt := py.RefCnt(name)
	defer func() { assert.Equal(t, nameRefCnt, py.RefCnt(name)) }()
	assert.NotNil(t, name)

	assert.True(t, strings.HasSuffix(pyunicode.AsString(name), "/queue.py"))
}

func TestPyModuleGetFilename(t *testing.T) {
	fmt.Println(assert.CallerInfo()[0])

	assert.Equal(t, "", pymodule.GetFilename(nil))

	queue := pyimport.ImportModule("queue")
	defer py.DecRef(queue)
	queueRefCnt := py.RefCnt(queue)
	defer func() { assert.Equal(t, queueRefCnt, py.RefCnt(queue)) }()

	name := pymodule.GetFilename(queue)
	assert.True(t, strings.HasSuffix(name, "/queue.py"))
}
