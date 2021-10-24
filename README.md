# go-python3-submodule

Go -> C API -> Python3, submodule style

尝试将[DataDog/go-python3]改造得更加OO未果, 为绕过循环引用整合模块后遭遇十分感人的编译速度, 令我抓狂。

历史的倒车再次上演, 索性分离子模块, 先完成绑定...

# 坑 / Todo

- `PyModule_GetDef`涉及到结构体转换, 目前搁置
