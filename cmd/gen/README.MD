# 项目架构

## 六边形架构
服务采用了六边形架构

1、领域层[domain]

领域模型包含了所有的应用逻辑与规则。

领域层中不会直接引用技术实现，例如 HTTP 上下文或数据库调用，这样就能够确保在技术方面的改动不会影响到领域层面。

2、端口层[port]

负责接收与用例相关的所有请求，这些请求负责在领域层中协调工作。

端口层在端口内部作为领域层的边界，在端口外部则扮演了外部实体的角色。

3、适配器层[adapter]

这一层的技术实现负责以某种格式接收输入、及产生输出。

在适配器层不存在领域逻辑，它的唯一职责就是在外部世界与领域层之间进行技术性的转换。

## 项目结构
```
├─adapter                       # 适配器层
│  ├─driven                     # 从动端口适配器
│  └─driver                     # 驱动端口适配器
│      ├─dto                    # 数据传输对象
│      │  ├─request             # 请求
│      │  └─response            # 请求响应
│      ├─handler                # 处理适配器
│      └─router                 # 路由适配器
├─domain                        # 领域层
│  ├─enums                      # 领域常量
│  ├─entity                     # 领域实体
│  └─service                    # 领域服务
├─global                        # 全局变量
├─infra                         # 基础设施
│  ├─middleware                 # 中间件
│  ├─model                      # 模型
│  └─utils                      # 通用工具
├─init                          # 初始化
│  ├─config                     # 配置
│  ├─db                         # 数据库
│  └─log                        # 日志
├─manifest                      # 清单
│  ├─config                     # 配置
│  └─i18n                       # 国际化
└─port                          # 端口层
    ├─driven                    # 从动端口
    └─driver                    # 驱动端口
```
### 调用链路
http请求 -> adapter/driver/handler -> port/driver -> domain -> port/driven/ -> adapter/driven