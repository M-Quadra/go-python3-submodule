package main

import (
	"testing"

	"github.com/M-Quadra/go-python3-submodule/py"
	pycallable "github.com/M-Quadra/go-python3-submodule/py-callable"
	pydict "github.com/M-Quadra/go-python3-submodule/py-dict"
	pyerr "github.com/M-Quadra/go-python3-submodule/py-err"
	pyeval "github.com/M-Quadra/go-python3-submodule/py-eval"
	pyimport "github.com/M-Quadra/go-python3-submodule/py-import"
	pylist "github.com/M-Quadra/go-python3-submodule/py-list"
	pyobject "github.com/M-Quadra/go-python3-submodule/py-object"
	pysys "github.com/M-Quadra/go-python3-submodule/py-sys"
	pyunicode "github.com/M-Quadra/go-python3-submodule/py-unicode"
	"github.com/stretchr/testify/assert"
)

func TestPyImportImportModule(t *testing.T) {
	os := pyimport.ImportModule("os")
	defer py.DecRef(os)
	assert.NotNil(t, os)
}

func TestPyImportImportModuleEx(t *testing.T) {
	os := pyimport.ImportModuleEx("os", nil, nil, nil)
	defer py.DecRef(os)
	assert.NotNil(t, os)
}

func TestPyImportImportModuleLevelObject(t *testing.T) {
	assert.Nil(t, pyimport.ImportModuleLevelObject(nil, nil, nil, nil, 0))
	pyerr.Clear()

	name := pyunicode.FromString("math")
	defer py.DecRef(name)

	math := pyimport.ImportModuleLevelObject(name, nil, nil, nil, 0)
	defer py.DecRef(math)
	assert.NotNil(t, math)
}

func TestPyImportImportModuleLevel(t *testing.T) {
	math := pyimport.ImportModuleLevel("math", nil, nil, nil, 0)
	defer py.DecRef(math)
	assert.NotNil(t, math)
}

func TestPyImportImport(t *testing.T) {
	assert.Nil(t, pyimport.Import(nil))

	name := pyunicode.FromString("platform")
	defer py.DecRef(name)

	platform := pyimport.Import(name)
	defer py.DecRef(platform)
	assert.NotNil(t, platform)
}

func TestPyImportReloadModule(t *testing.T) {
	assert.Nil(t, pyimport.ReloadModule(nil))

	osA := pyimport.ImportModule("os")
	defer py.DecRef(osA)
	assert.NotNil(t, osA)

	osB := pyimport.ReloadModule(osA)
	defer py.DecRef(osB)
	assert.NotNil(t, osB)

	// [DataDog/go-python3]: PyImport_ReloadModule return a new reference, pointer should be the same
	assert.Equal(t, osA, osB)
}

func TestPyImportAddModule(t *testing.T) {
	assert.Nil(t, pyimport.AddModuleObject(nil))

	os := pyimport.ImportModule("os")
	defer py.DecRef(os)
	assert.NotNil(t, os)

	name := "os.new"
	newA := pyimport.AddModule(name)
	defer py.DecRef(newA)
	assert.NotNil(t, newA)

	nameUnicode := pyunicode.FromString(name)
	defer py.DecRef(nameUnicode)
	assert.NotNil(t, nameUnicode)

	newB := pyimport.AddModuleObject(nameUnicode)
	defer py.DecRef(newB)
	assert.NotNil(t, newB)
}

func TestPyImportExecCodeModule(t *testing.T) {
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
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModule("test_module", code)
	defer py.DecRef(module)
	assert.NotNil(t, module)
}

func TestPyImportExecCodeModuleEx(t *testing.T) {
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
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModuleEx("test_module", code, "test_module.py")
	defer py.DecRef(module)
	assert.NotNil(t, module)
}

func TestPyImportExecCodeModuleObject(t *testing.T) {
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
	assert.NotNil(t, code)

	modulename := pyunicode.FromString("test_module")
	defer py.DecRef(modulename)

	module := pyimport.ExecCodeModuleObject(modulename, code, filename, filename)
	defer py.DecRef(module)
	assert.NotNil(t, module)

	{ // nil
		assert.Nil(t, pyimport.ExecCodeModuleObject(nil, code, filename, filename))
		assert.Nil(t, pyimport.ExecCodeModuleObject(modulename, nil, filename, filename))

		moduleA := pyimport.ExecCodeModuleObject(modulename, code, nil, filename)
		defer py.DecRef(moduleA)
		assert.NotNil(t, moduleA)

		moduleB := pyimport.ExecCodeModuleObject(modulename, code, filename, nil)
		defer py.DecRef(moduleB)
		assert.NotNil(t, moduleB)
	}
}

func TestPyImportExecCodeModuleWithPathnames(t *testing.T) {
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
	assert.NotNil(t, code)

	module := pyimport.ExecCodeModuleWithPathnames("test_module", code, "test_module.py", "test_module.py")
	defer py.DecRef(module)
	assert.NotNil(t, module)

	assert.Nil(t, pyimport.ExecCodeModuleWithPathnames("test_module", nil, "test_module.py", "test_module.py"))
}

func TestPyImportGetMagicNumber(t *testing.T) {
	assert.NotNil(t, pyimport.GetMagicNumber())
}

func TestPyImportGetMagicTag(t *testing.T) {
	assert.NotNil(t, pyimport.GetMagicTag())
}

func TestPyImportGetModuleDict(t *testing.T) {
	dic := pyimport.GetModuleDict()
	defer py.DecRef(dic)
	assert.True(t, pydict.Check(dic))
}

func TestPyImportGetModule(t *testing.T) {
	assert.Nil(t, pyimport.GetModule(nil))

	name := "os"
	osA := pyimport.ImportModule(name)
	defer py.DecRef(osA)

	nameUnicode := pyunicode.FromString(name)
	defer py.DecRef(nameUnicode)
	osB := pyimport.GetModule(nameUnicode)
	defer py.DecRef(osB)

	assert.Equal(t, osA, osB)
}

func TestPyImportGetImporter(t *testing.T) {
	pathList := pysys.GetObject("path")
	path := pylist.GetItem(pathList, 0)
	assert.NotNil(t, path)

	importer := pyimport.GetImporter(path)
	defer py.DecRef(importer)
	assert.NotNil(t, importer)

	pyimport.GetImporter(nil)
}
