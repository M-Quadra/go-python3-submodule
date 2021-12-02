# 容器测试记录

这是一份包含不同环境下测试结果的记录, 测试镜像包含`scipy`, 运行环境均为容器。

(这玩意似乎放到wiki中更合适, 或许后续该迁移)

基础镜像 | 构建参考 | 测试镜像大小 | 测试日期
:---:|:---:|:---:|:---:
python:3.9 | [Dockerfile](./Dockerfile.python3.9) | 1.53 GB | 2021-12-02
python:3.9-slim | [Dockerfile](./Dockerfile.python3.9-slim) | 1.02 GB | 2021-12-02
