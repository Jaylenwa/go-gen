## 介绍
一键搭建项目，自动化生成代码，大幅度提高开发效率，减少CRUD的重复工作

## 使用方法

**前置条件说明**

方法1，方法2；需要结合配置文件使用

方法3（命令行模式）无需配置

配置文件生成方式：
1. 源代码中的.gen-cli.yaml
2. 可执行文件自动生成，执行后生成.gen-cli.yaml
```shell
# 方式1
go run main.go config
# 方式2
# go build后 执行
[可执行文件] config
```

配置内容见文件注释

### 1、源码编译使用
```shell
go run main.go gen --config .gen-cli.yaml
```

### 2、二进制文件使用
生成可执行文件`go build .`

需要结合配置文件一起使用，配置文件变量值及配置文件名称可自定义
```shell
# linux下执行
[可执行文件] gen --config .gen-cli.yaml
```

### 3、使用命令行（推荐）
**说明**：[可执行文件名]为 go build 后的可执行文件

使用 `[可执行文件名] gen --help` 查看命令行参数帮助列表

使用案例：
```shell
[可执行文件名] gen -a 127.0.0.1 -u root --pwd root -p 3306 -d dbName -t tableName -m go_module
```
还有额外参数可以配置，使用`[可执行文件名] gen --help`命令查看 按需填写
