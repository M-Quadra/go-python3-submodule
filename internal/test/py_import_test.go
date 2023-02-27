package main

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/go-python3-submodule/v11/py"
	pycallable "github.com/M-Quadra/go-python3-submodule/v11/py-callable"
	pydict "github.com/M-Quadra/go-python3-submodule/v11/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/v11/py-err"
	pyeval "github.com/M-Quadra/go-python3-submodule/v11/py-eval"
	pyimport "github.com/M-Quadra/go-python3-submodule/v11/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/v11/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/v11/py-object"
	pysys "github.com/M-Quadra/go-python3-submodule/v11/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/v11/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyImportImportModule(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	os := pyimport.ImportModule("os")
	defer py.DecRef(os)
	osRefCnt := py.RefCnt(os)
	defer func() { assert.Equal(t, osRefCnt, py.RefCnt(os)) }()

	assert.NotNil(t, os)
}

func TestPyImportImportModuleEx(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	os := pyimport.ImportModuleEx("os", nil, nil, nil)
	defer py.DecRef(os)
	osRefCnt := py.RefCnt(os)
	defer func() { assert.Equal(t, osRefCnt, py.RefCnt(os)) }()

	assert.NotNil(t, os)
}

func TestPyImportImportModuleLevelObject(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.ImportModuleLevelObject(nil, nil, nil, nil, 0))
	pyerr.Clear()

	name := pyunicode.FromString("math")
	defer py.DecRef(name)

	math := pyimport.ImportModuleLevelObject(name, nil, nil, nil, 0)
	defer py.DecRef(math)
	mathRefCnt := py.RefCnt(math)
	defer func() { assert.Equal(t, mathRefCnt, py.RefCnt(math)) }()

	assert.NotNil(t, math)
}

func TestPyImportImportModuleLevel(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	math := pyimport.ImportModuleLevel("math", nil, nil, nil, 0)
	defer py.DecRef(math)
	mathRefCnt := py.RefCnt(math)
	defer func() { assert.Equal(t, mathRefCnt, py.RefCnt(math)) }()

	assert.NotNil(t, math)
}

func TestPyImportImport(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.Import(nil))

	name := pyunicode.FromString("platform")
	defer py.DecRef(name)

	platform := pyimport.Import(name)
	defer py.DecRef(platform)
	platformRefCnt := py.RefCnt(platform)
	defer func() { assert.Equal(t, platformRefCnt, py.RefCnt(platform)) }()

	assert.NotNil(t, platform)
}

func TestPyImportReloadModule(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.ReloadModule(nil))

	osA := pyimport.ImportModule("os")
	defer py.DecRef(osA)
	osARefCnt := py.RefCnt(osA)
	defer func() { assert.Equal(t, osARefCnt, py.RefCnt(osA)) }()
	assert.NotNil(t, osA)

	osB := pyimport.ReloadModule(osA)
	defer py.DecRef(osB)
	osBRefCnt := py.RefCnt(osB)
	defer func() { assert.Equal(t, osBRefCnt, py.RefCnt(osB)) }()
	assert.NotNil(t, osB)

	// [DataDog/go-python3]: PyImport_ReloadModule return a new reference, pointer should be the same
	assert.Equal(t, osA, osB)
}

func TestPyImportAddModule(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.AddModuleObject(nil))

	os := pyimport.ImportModule("os")
	defer py.DecRef(os)
	osRefCnt := py.RefCnt(os)
	defer func() { assert.Equal(t, osRefCnt, py.RefCnt(os)) }()
	assert.NotNil(t, os)

	name := "os.new"
	newA := pyimport.AddModule(name)
	newARefCnt := py.RefCnt(newA)
	defer func() { assert.Equal(t, newARefCnt, py.RefCnt(newA)) }()
	assert.NotNil(t, newA)

	nameUnicode := pyunicode.FromString(name)
	defer py.DecRef(nameUnicode)
	nameUnicodeRefCnt := py.RefCnt(nameUnicode)
	defer func() { assert.Equal(t, nameUnicodeRefCnt, py.RefCnt(nameUnicode)) }()
	assert.NotNil(t, nameUnicode)

	newB := pyimport.AddModuleObject(nameUnicode)
	newBRefCnt := py.RefCnt(newB)
	defer func() { assert.Equal(t, newBRefCnt, py.RefCnt(newB)) }()

	assert.NotNil(t, newB)
	assert.Equal(t, newA, newB)
}

func TestPyImportExecCodeModule(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.ExecCodeModule("", nil))

	// [DataDog/go-python3]: fake module
	source := pyunicode.FromString("__version__ = '2.0'")
	defer py.DecRef(source)
	filename := pyunicode.FromString("test_module.py")
	defer py.DecRef(filename)
	mode := pyunicode.FromString("exec")
	defer py.DecRef(mode)

	// [DataDog/go-python3]: perform module load
	builtins := pyeval.GetBuiltins()
	assert.True(t, pydict.Check(builtins))

	compile := pydict.GetItemString(builtins, "compile")
	assert.True(t, pycallable.Check(compile))

	code := pyobject.CallFunctionObjArgs(compile, source, filename, mode)
	defer py.DecRef(code)
	defer func() { assert.Equal(t, 1, py.RefCnt(code)) }()
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModule("test_module", code)
	defer py.DecRef(module)
	assert.NotNil(t, module)
}

func TestPyImportExecCodeModuleEx(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.ExecCodeModuleEx("", nil, ""))

	// [DataDog/go-python3]: fake module
	source := pyunicode.FromString("__version__ = '2.0'")
	defer py.DecRef(source)
	filename := pyunicode.FromString("test_module.py")
	defer py.DecRef(filename)
	mode := pyunicode.FromString("exec")
	defer py.DecRef(mode)

	// [DataDog/go-python3]: perform module load
	builtins := pyeval.GetBuiltins()
	assert.True(t, pydict.Check(builtins))

	compile := pydict.GetItemString(builtins, "compile")
	assert.True(t, pycallable.Check(compile))

	code := pyobject.CallFunctionObjArgs(compile, source, filename, mode)
	defer py.DecRef(code)
	defer func() { assert.Equal(t, 1, py.RefCnt(code)) }()
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModuleEx("test_module", code, "test_module.py")
	defer py.DecRef(module)
	assert.NotNil(t, module)
}

func TestPyImportExecCodeModuleObject(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	// [DataDog/go-python3]: fake module
	source := pyunicode.FromString("__version__ = '2.0'")
	defer py.DecRef(source)
	filename := pyunicode.FromString("test_module.py")
	defer py.DecRef(filename)
	mode := pyunicode.FromString("exec")
	defer py.DecRef(mode)

	// [DataDog/go-python3]: perform module load
	builtins := pyeval.GetBuiltins()
	assert.True(t, pydict.Check(builtins))

	compile := pydict.GetItemString(builtins, "compile")
	assert.True(t, pycallable.Check(compile))

	code := pyobject.CallFunctionObjArgs(compile, source, filename, mode)
	defer py.DecRef(code)
	defer func() { assert.Equal(t, 1, py.RefCnt(code)) }()
	assert.NotNil(t, code)

	modulename := pyunicode.FromString("test_module")
	defer py.DecRef(modulename)

	module := pyimport.ExecCodeModuleObject(modulename, code, filename, filename)
	defer py.DecRef(module)
	moduleRefCnt := py.RefCnt(module)
	defer func() { assert.Equal(t, moduleRefCnt, py.RefCnt(module)) }()
	assert.NotNil(t, module)

	{ // nil
		assert.Nil(t, pyimport.ExecCodeModuleObject(nil, code, filename, filename))
		assert.Nil(t, pyimport.ExecCodeModuleObject(modulename, nil, filename, filename))

		moduleA := pyimport.ExecCodeModuleObject(modulename, code, nil, filename)
		defer py.DecRef(moduleA)
		moduleARefCnt := py.RefCnt(moduleA)
		assert.Equal(t, moduleRefCnt+1, moduleARefCnt)

		moduleB := pyimport.ExecCodeModuleObject(modulename, code, filename, nil)
		defer py.DecRef(moduleB)
		moduleBRefCnt := py.RefCnt(moduleB)
		assert.Equal(t, moduleARefCnt+1, moduleBRefCnt)
	}
}

func TestPyImportExecCodeModuleWithPathnames(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	// [DataDog/go-python3]: fake module
	source := pyunicode.FromString("__version__ = '2.0'")
	defer py.DecRef(source)
	filename := pyunicode.FromString("test_module.py")
	defer py.DecRef(filename)
	mode := pyunicode.FromString("exec")
	defer py.DecRef(mode)

	// [DataDog/go-python3]: perform module load
	builtins := pyeval.GetBuiltins()
	assert.True(t, pydict.Check(builtins))

	compile := pydict.GetItemString(builtins, "compile")
	assert.True(t, pycallable.Check(compile))

	code := pyobject.CallFunctionObjArgs(compile, source, filename, mode)
	defer py.DecRef(code)
	defer func() { assert.Equal(t, 1, py.RefCnt(code)) }()
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModuleWithPathnames("test_module", code, "test_module.py", "test_module.py")
	defer py.DecRef(module)
	assert.NotNil(t, module)

	assert.Nil(t, pyimport.ExecCodeModuleWithPathnames("test_module", nil, "test_module.py", "test_module.py"))
}

func TestPyImportGetMagicNumber(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.NotNil(t, pyimport.GetMagicNumber())
}

func TestPyImportGetMagicTag(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.NotNil(t, pyimport.GetMagicTag())
}

func TestPyImportGetModuleDict(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	dicA := pyimport.GetModuleDict()
	assert.True(t, pydict.Check(dicA))
	dicARefCnt := py.RefCnt(dicA)

	dicB := pyimport.GetModuleDict()
	assert.True(t, pydict.Check(dicB))
	dicBRefCnt := py.RefCnt(dicB)

	assert.Equal(t, dicA, dicB)
	assert.Equal(t, dicARefCnt, dicBRefCnt)
}

func TestPyImportGetModule(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	assert.Nil(t, pyimport.GetModule(nil))

	name := "os"
	osA := pyimport.ImportModule(name)
	defer py.DecRef(osA)
	osARefCnt := py.RefCnt(osA)
	defer func() { assert.Equal(t, osARefCnt, py.RefCnt(osA)) }()

	nameUnicode := pyunicode.FromString(name)
	defer py.DecRef(nameUnicode)
	nameUnicodeRefCnt := py.RefCnt(nameUnicode)
	defer func() { assert.Equal(t, nameUnicodeRefCnt, py.RefCnt(nameUnicode)) }()

	osB := pyimport.GetModule(nameUnicode)
	defer py.DecRef(osB)
	osBRefCnt := py.RefCnt(osB)
	defer func() { assert.Equal(t, osBRefCnt, py.RefCnt(osB)) }()

	assert.Equal(t, osA, osB)
	assert.Equal(t, osARefCnt+1, osBRefCnt)
}

func TestPyImportGetImporter(t *testing.T) {
	fmt.Println("current:", assert.CallerInfo()[0])

	pathList := pysys.GetObject("path")
	pathListRefCnt := py.RefCnt(pathList)
	defer func() { assert.Equal(t, pathListRefCnt, py.RefCnt(pathList)) }()

	path := pylist.GetItem(pathList, 0)
	assert.NotNil(t, path)

	importer := pyimport.GetImporter(path)
	defer py.DecRef(importer)
	assert.NotNil(t, importer)

	pyimport.GetImporter(nil)
}
