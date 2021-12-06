# go-python3-submodule

Go -> C API -> Python3, submodule style

此物源自另一个<del>有生之年</del>工程的拆分, 尝试将[DataDog/go-python3](https://github.com/DataDog/go-python3)改造更加面向对象与强类型未果, 经历合久必分, 分久必合, 反复折腾后只余构建速度催人泪下。

> 阿库娅大人说: "阿库西斯教徒都是努力的人, 就算失败了也是世界的错！遇到困难先跑路, 大不了弃坑。"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

历史的倒车再次启动, 分离子模块, 先完成API绑定。

# 莽 / Installation

```
go get github.com/M-Quadra/go-python3-submodule/v9
```

# 润 / Usage

调用方式同`Python/C API`方法, 如`PyBool_Check(x)`调用方法为`pybool.Check`, 包名由编辑器自动补全。

用例同测试, 见[test](./test)文件夹。

方法名针对数据类型做了部分转换, 具体如下:

Python/C API | Go
:---:|:---:
PyFloat_AsDouble | pyfloat.AsFloat64
PyLong_AsLong | pylong.AsInt
PyLong_AsLongLong | pylong.AsInt64
... | ...

如遇卡死, 多半是GIL, 单线程可尝试添加以下代码:

```
if !pygilstate.Check() {
	save := pyeval.SaveThread()
	defer pyeval.RestoreThread(save)

	gstate := pygilstate.Ensure()
	defer pygilstate.Release(gstate)
}

// do something...
```

多线程自觉上锁:

```
var _m = sync.Mutex{}

func xx() {
	_m.Lock()
	defer _m.Unlock()
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if !pygilstate.Check() {
		save := pyeval.SaveThread()
		defer pyeval.RestoreThread(save)

		gstate := pygilstate.Ensure()
		defer pygilstate.Release(gstate)
	}

	// do something...
}
```

具体用例可参考[此处](./test/benchmark/curvefit_test.go)

# 肝 / Progress

开发环境: macOS 12.0, python 3.9。

已迁移[DataDog/go-python3](https://github.com/DataDog/go-python3)的大部分方法, 当然还有些小问题龟速修复中...

测试用例大部分添加了引用计数检查。

已知不同环境会有所差异, 大体完善后预计进Docker修缮。

已完成`python:3.9`容器中的测试。测试过更多的环境后也许会打上版本标签。

# 坑 / Todo

- `PyModule_GetDef`涉及到结构体转换, 目前搁置。

- `Py_Main`调用卡死, 先注释掉, 随缘解决。

- `PyObject.ob_refcnt`应该弄成方法还是走`Py_REFCNT`? 同理`PyObject`是否也应该开辟方法?

- `Exception`的引用计数没搞懂, 摸了。

- 有余力会加上多版本支持, 预计为`go get .../go-python3-submodule/v9`, `go get .../go-python3-submodule/v8`形式。

# 杂 / Other

[容器测试记录](./test/dockerfiles/test_in_container_v9.md)