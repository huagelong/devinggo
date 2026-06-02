# DevingGo 模块架构说明

本文档描述DevingGo的模块化架构设计、目录结构和开发规范。

## 模块化架构

DevingGo采用插件化的模块架构，每个模块都是一个独立的功能单元，包含完整的MVC结构和业务逻辑。

### 核心概念

- **模块独立性**: 每个模块都有独立的目录结构、配置和功能
- **热插拔**: 模块可以独立导入/导出，支持模块级别的部署
- **自动加载**: 通过bootstrap目录实现模块的自动注册和加载
- **标准化**: 统一的模块结构和配置格式(.module.yaml)

## 目录结构

```
modules/
├── bootstrap/                    # 模块自动加载器
│   ├── logic/                    # Logic层自动加载
│   │   ├── system.go             # 加载system模块logic
│   │   └── api.go                # 加载api模块logic
│   ├── modules/                  # 模块层自动加载
│   │   ├── system.go             # 加载system模块
│   │   └── api.go                # 加载api模块
│   └── worker/                   # Worker层自动加载
│       ├── system.go             # 加载system模块worker
│       └── api.go                # 加载api模块worker
│
├── system/                       # 系统核心模块
│   ├── module.go                 # 模块元信息
│   ├── .module.yaml              # 模块配置文件
│   ├── api/                      # API定义层
│   │   └── system/               # 按业务分组
│   │       ├── system_user.go    # 用户API定义
│   │       └── system_role.go    # 角色API定义
│   ├── controller/               # 控制器层
│   │   └── system/
│   │       ├── system_user.go    # 用户控制器
│   │       └── system_role.go    # 角色控制器
│   ├── logic/                    # 业务逻辑层
│   │   └── system/
│   │       ├── system_user.go    # 用户逻辑
│   │       └── system_role.go    # 角色逻辑
│   ├── service/                  # 服务接口层（自动生成）
│   │   ├── system_user.go
│   │   └── system_role.go
│   ├── model/                    # 数据模型层
│   │   ├── do/                   # 数据对象（自动生成）
│   │   ├── entity/               # 实体模型（自动生成）
│   │   ├── req/                  # 请求模型
│   │   └── res/                  # 响应模型
│   ├── router/                   # 路由注册
│   │   └── router.go
│   ├── worker/                   # 后台任务
│   │   ├── worker.go             # Worker管理器
│   │   ├── server/               # 异步任务
│   │   └── cron/                 # 定时任务
│   ├── cmd/                      # 命令行工具
│   │   ├── http.go               # HTTP服务启动
│   │   └── worker.go             # Worker服务启动
│   ├── consts/                   # 常量定义
│   │   └── worker.go             # Worker常量
│   ├── codes/                    # 错误码定义
│   └── pkg/                      # 内部工具包
│
└── api/                          # API模块（示例）
    ├── module.go
    ├── .module.yaml
    ├── api/
    ├── controller/
    ├── logic/
    └── service/
```

## 模块生命周期

### 1. 模块创建

```bash
# 创建新模块
make gen-module name=blog

# 生成的文件结构
modules/blog/
├── .module.yaml              # 模块配置
├── module.json               # 向后兼容配置
├── module.go                 # 模块元信息
├── api/blog/                 # API定义
├── controller/blog/          # 控制器
├── logic/blog/               # 业务逻辑
├── service/                  # 服务接口
├── model/                    # 数据模型
├── router/router.go          # 路由
└── cmd/init.go               # 初始化

# 同时创建bootstrap文件
modules/bootstrap/
├── logic/blog.go             # 自动加载logic
├── modules/blog.go           # 自动加载模块
└── worker/blog.go            # 自动加载worker
```

### 2. 模块加载

模块通过bootstrap目录自动加载：

```go
// main.go
import (
    _ "devinggo/modules/bootstrap/logic"   // 自动加载所有模块的logic
)

// modules/bootstrap/logic/system.go
package logic

import (
    _ "devinggo/modules/system/logic"      // 加载system模块
)
```

**加载顺序**：
1. main.go导入`modules/bootstrap/logic`
2. bootstrap包导入各模块的logic包
3. logic包的init()函数自动注册服务

### 3. 模块导出/导入

```bash
# 导出模块为zip包
make export-module name=blog
# 生成: blog.v1.0.0.zip

# 导入模块
make import-module file=blog.v1.0.0.zip
# 自动执行:
# 1. 解压文件
# 2. 复制Go代码到modules/
# 3. 复制静态资源到resource/
# 4. 执行数据库迁移
# 5. 合并配置文件
# 6. 更新bootstrap加载器
```

## 模块配置 (.module.yaml)

每个模块都有一个`.module.yaml`配置文件，描述模块的元信息和依赖关系。

### 基本配置

```yaml
name: blog                      # 模块名称
version: 1.0.0                  # 版本号
author: DevingGo Team           # 作者
license: MIT                    # 许可证
description: 博客管理模块       # 描述
homepage: https://example.com   # 主页
goVersion: 1.23+                # Go版本要求
```

### 依赖管理

```yaml
dependencies:                   # Go依赖包
  github.com/gin-gonic/gin: v1.9.0
  
modules:                        # DevingGo模块依赖
  - system                      # 依赖system模块
```

### 文件清单

```yaml
files:
  go:                           # Go源码文件
    - modules/blog
    - modules/bootstrap/logic/blog.go
    - modules/bootstrap/modules/blog.go
    - modules/bootstrap/worker/blog.go
  sql:                          # SQL迁移文件
    - resource/migrations/20260306_blog_module.up.sql
    - resource/migrations/20260306_blog_module.down.sql
  static:                       # 静态资源
    - resource/public/blog
  config:                       # 配置文件
    - config/blog.yaml
```

详细配置规范请参考：[MODULE_YAML_SPEC.md](../hack/generator/docs/MODULE_YAML_SPEC.md)

## 模块开发流程

### 1. 创建模块

```bash
# 创建新模块
make gen-module name=blog

# 或从现有模块克隆
make clone-module name=blog source=system
```

### 2. 开发CRUD功能

```bash
# 生成CRUD代码（假设已有blog_post表）
make gen-crud table=blog_post module=blog

# 更新service接口
make service

# 更新controller
make ctrl
```

### 3. 添加Worker任务

```bash
# 创建后台任务
make gen-worker module=blog worker=PublishPost

# 更新service接口
make service
```

### 4. 自定义开发

编辑生成的文件，添加自定义业务逻辑：

- `logic/` - 业务逻辑实现
- `controller/` - HTTP处理器
- `api/` - API定义和文档
- `model/req/` - 请求参数验证
- `model/res/` - 响应数据结构

### 5. 测试和部署

```bash
# 运行数据库迁移
go run main.go migrate:up

# 启动服务
go run main.go

# 或使用热重载
gf run main.go

# 导出模块（用于分发）
make export-module name=blog
```

## 最佳实践

### 1. 模块命名规范

- 模块名使用小写字母和下划线：`user_center`、`blog`、`product`
- 避免使用系统保留名：`system`、`internal`、`bootstrap`
- 包名与目录名保持一致

### 2. 目录组织

- **按业务分组**：`api/user/`、`logic/user/`、`controller/user/`
- **统一命名**：文件名与数据表名对应，如`system_user.go`
- **避免循环依赖**：logic可以调用service，但service不能调用logic

### 3. 代码生成

- 优先使用代码生成器创建CRUD代码
- 生成后的代码可以自定义修改，不会被覆盖
- 定期运行`make service`更新服务接口

### 4. 数据库迁移

- 每个模块的SQL迁移文件存放在`resource/migrations/`
- 文件名格式：`{时间戳}_{模块名}_module.up.sql` 和 `.down.sql`
- 使用`go run main.go migrate:up`执行迁移

### 5. 模块依赖

- 在`.module.yaml`中声明模块依赖
- 避免模块之间的强耦合
- 通过service接口进行模块间通信

## 常见问题

### Q: 如何添加新的模块到现有项目？

```bash
# 方法1: 创建新模块
make gen-module name=newmodule

# 方法2: 导入现有模块包
make import-module file=newmodule.v1.0.0.zip

# 方法3: 克隆现有模块
make clone-module name=newmodule source=system
```

### Q: 模块之间如何通信？

通过service接口调用：

```go
// 在blog模块中调用system模块的用户服务
import "devinggo/modules/system/service"

func (s *sBlogPost) GetAuthor(ctx context.Context, userId int64) (*model.User, error) {
    // 调用system模块的服务
    user, err := service.SystemUser().GetById(ctx, userId)
    return user, err
}
```

### Q: 如何更新已有模块？

```bash
# 1. 修改代码后，更新service接口
make service

# 2. 如果添加了新表，重新生成DAO
make dao

# 3. 导出新版本
make export-module name=blog
# 版本号会自动递增
```

### Q: bootstrap目录的作用是什么？

bootstrap目录负责自动加载所有模块：

- `logic/` - 加载各模块的logic包（注册service）
- `modules/` - 加载各模块的module包（模块初始化）
- `worker/` - 加载各模块的worker包（注册后台任务）

新建模块时，代码生成器会自动创建对应的bootstrap文件。

### Q: 如何删除模块？

```bash
# 1. 手动删除模块目录
rm -rf modules/mymodule

# 2. 删除bootstrap加载器
rm modules/bootstrap/logic/mymodule.go
rm modules/bootstrap/modules/mymodule.go
rm modules/bootstrap/worker/mymodule.go

# 3. 回滚数据库迁移（可选）
go run main.go migrate:down

# 4. 更新代码
make service
```

## 相关文档

- [代码生成器使用说明](../hack/generator/README.md)
- [.module.yaml配置规范](../hack/generator/docs/MODULE_YAML_SPEC.md)
- [技术方案文档](../docs/PLAN-CodeGeneratorUnification.md)
- [Worker开发指南](system/worker/README.md)

## 许可证

MIT License
