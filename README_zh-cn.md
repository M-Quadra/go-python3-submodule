# go-python3-submodule

[English (稀烂_(ˊཀˋ」∠)_)](./README.md) | 简体中文

子模块风格的 Go Python3/C API 绑定。

此物源自另一个<del>有生之年</del>工程的拆分, 尝试将[DataDog/go-python3](https://github.com/DataDog/go-python3)改造更加面向对象与强类型未果, 经历合久必分, 分久必合, 反复折腾后只余构建速度催人泪下。

> 阿库娅大人说: "阿库西斯教徒都是努力的人, 就算失败了也是世界的错！遇到困难先跑路, 大不了弃坑。"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

历史的倒车再次启动, 分离子模块, 先完成API绑定。

## 莽

推荐使用 Go Modules

```
go get github.com/M-Quadra/go-python3-submodule/v10
```

Python version | Package URL
:---:|:---:
3.10 | github.com/M-Quadra/go-python3-submodule/v10 
3.9 | github.com/M-Quadra/go-python3-submodule/v9 
3.8 | github.com/M-Quadra/go-python3-submodule/v8 

## 润

调用方式同`Python/C API`方法, 如`PyBool_Check(x)`使用`pybool.Check(x)`调用, 包名补全交给gopls。

方法名针对数据类型做了部分转换, 具体如下:

Python/C API | Go
:---:|:---:
PyFloat_AsDouble | pyfloat.AsFloat64
PyLong_AsLong | pylong.AsInt
PyLong_AsLongLong | pylong.AsInt64
... | ...

如遇卡死, 建议排查GIL或引用计数



协程调用有2种方式: [example-0](./internal/example/goroutine-0) , [example-1](/internal/example/goroutine-1)

个人偏好example-0



## 坑

- [ ] `Exception`的单元测试添加引用计数检查。不过目前没搞懂, 摸了。
- [ ] GitHub Action



## 肝

开发环境: macOS (x86-64)

运行目标: Linux (Docker x86-64)


已迁移[DataDog/go-python3](https://github.com/DataDog/go-python3)的大部分方法, 当然还有些小问题龟速修复中...

测试用例大部分添加了引用计数检查。

已知不同环境会有所差异, 适配目标为通过[单元测试](./internal/test)。

- `PyModule_GetDef`涉及到结构体转换, 目前搁置。
- `Py_Main`调用卡死, 先注释掉, 随缘解决。