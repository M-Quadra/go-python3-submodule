# go-python3-submodule

Go -> C API -> Python3, submodule style

这是另一个有生之年项目的拆分,

尝试将[DataDog/go-python3](https://github.com/DataDog/go-python3)改造得更加OO未果, 为绕过循环引用整合模块后遭遇十分感人的编译速度, 令我抓狂。

历史的倒车再次上演, 索性分离子模块, 先完成API绑定

# 开发环境

macOS 12.0, python 3.9

已完成[DataDog/go-python3](https://github.com/DataDog/go-python3)的大部分方法, 可以通过单独的方法测试, 但是依然跑不动`go test`。龟速修复中...

大体完善估计会进Docker修缮

# 坑 / Todo

- `PyModule_GetDef`涉及到结构体转换, 目前搁置。

- `pyimport.ImportModuleLevelObject`与某些方法存在冲突, 无法保证整体通过`go test`。目前优先保证方法可以独立调用, 后续再处理。

- `Py_Main`调用卡死, 先注释掉, 随缘解决。