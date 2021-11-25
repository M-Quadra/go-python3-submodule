# go-python3-submodule

Go -> C API -> Python3, submodule style

这是另一个有生之年 repository 的拆分, 尝试将[DataDog/go-python3](https://github.com/DataDog/go-python3)改造更加OO未果, 为绕过循环引用限制整合单一 package 后编译速度催人泪下。

历史的倒车再次启动, 索性分离子模块, 先完成API绑定。

# 开发环境

macOS 12.0, python 3.9

已完成[DataDog/go-python3](https://github.com/DataDog/go-python3)的大部分方法, 可以通过单独的方法测试, 但是依然跑不动`go test`。龟速修复中...

大体完善估计会进Docker修缮

# 坑 / Todo

- `PyModule_GetDef`涉及到结构体转换, 目前搁置

- `Py_Main`调用卡死, 先注释掉, 随缘解决

- `PyObject.ob_refcnt`应该弄成方法还是走`Py_REFCNT`? 同理`PyObject`是否也应该开辟方法

- test大一统已完成, 开始为单元测试增加引用计数