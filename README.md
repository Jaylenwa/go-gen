## 介绍
一键搭建项目，自动化生成代码，大幅度提高开发效率，减少CRUD的重复工作

## 使用方法

前置条件

配置 `.gen-cli.yaml`

配置内容见文件注释

### 1、源码编译使用
```shell
go run main.go gen --config .gen-cli.yaml
```

### 2、二进制文件使用
```shell
# 生成可执行文件
go build .
# windows下执行
.\gogen.exe gen --config .gen-cli.yaml gen
# linux下执行
./gogen --config gen .gen-cli.yaml gen
```