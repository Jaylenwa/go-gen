## 介绍
一键搭建项目，自动化生成代码，大幅度提高开发效率，减少CRUD的重复工作

## 使用方法

**前置条件说明**

方法1，方法2；需要配置`.gen-cli.yaml`方可使用

方法3（命令行模式）无需配置

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
.\gogen.exe gen --config .gen-cli.yaml
# linux下执行
./gogen --config gen .gen-cli.yaml
```

### 3、使用命令行（推荐）
**说明**：[可执行文件名]为 go build 后的可执行文件

使用 `[可执行文件名] gen --help` 查看命令行参数帮助列表

使用案例：
```shell
[可执行文件名] gen -a 127.0.0.1 -u root --pwd root -p 3306 -d dbName -t tableName -s serverName
```
还有额外参数可以配置，使用`[可执行文件名] gen --help`命令查看 按需填写
