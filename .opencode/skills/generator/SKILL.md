---
name: devinggo-generator
description: |
  DevingGo 代码生成器 CLI 工具集。当用户需要生成 CRUD 代码、创建模块、管理 Worker 任务、
  导入导出模块、或需要快速搭建后端/前端代码时使用。适用于任何涉及代码生成、模块管理、
  数据库表对应的前后端代码生成的场景。
  
  触发关键词：生成代码、CRUD、模块管理、Worker任务、生成器、代码生成器、
  创建模块、克隆模块、导入模块、导出模块、定时任务、异步任务
---

# DevingGo 代码生成器 Skill

## 概述

DevingGo Generator 是一个统一的代码生成 CLI 工具，支持：
- **模块管理**：创建、克隆、导入、导出、验证模块
- **Worker 任务生成**：创建异步任务、定时任务、混合任务
- **CRUD 代码生成**：根据数据库表自动生成完整的前后端代码

## 工具位置

```
./hack/generator/main.go
```

## 命令速查表

### 1. 模块管理 (module)

| 命令 | 功能 | 用法 |
|------|------|------|
| `module:create` | 创建新模块 | `-name <模块名>` |
| `module:clone` | 克隆现有模块 | `-source <源> -target <目标>` |
| `module:export` | 导出模块为 zip | `-name <模块名>` |
| `module:import` | 从 zip 导入模块 | `-file <zip路径>` |
| `module:list` | 列出所有模块 | 无参数 |
| `module:validate` | 验证模块完整性 | `-name <模块名>` |

### 2. Worker 任务 (worker)

| 命令 | 功能 | 用法 |
|------|------|------|
| `worker:create` | 创建 Worker 任务 | `-module <模块> -name <名称> -type <task/cron/both>` |

### 3. CRUD 生成 (crud)

| 命令 | 功能 | 用法 |
|------|------|------|
| `crud:generate` | 生成 CRUD 代码 | `-m <模块> -t <表名> -n <中文名> [--frontend]` |

## 使用指南

### 场景1：创建新模块

当用户需要创建一个新的业务模块时使用。

```bash
# 创建模块
go run ./hack/generator/main.go module:create -name blog

# 然后需要运行
go run main.go service  # 更新服务层
go run main.go dao      # 更新数据访问层
```

**生成文件**：
- `modules/blog/` 目录及基础文件结构
- `.module.yaml` 配置文件

### 场景2：为已有表生成 CRUD

当用户已有数据库表，需要生成对应的前后端代码时使用。

```bash
# 基础用法（仅后端）
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户

# 同时生成前端代码
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户 --frontend

# 使用 Makefile
make gen-crud table=system_user module=system name=用户
make gen-crud table=system_user module=system name=用户 frontend=1
```

**参数说明**：
- `-m, --module`: 模块名（如 system, blog）
- `-t, --table`: 数据库表名（如 system_user）
- `-n, --name`: 资源中文名（如 用户）
- `--frontend`: 同时生成前端代码
- `--force`: 覆盖已存在的文件
- `--dry-run`: 仅预览，不实际生成

**生成的后端文件**：
```
modules/{module}/
├── api/{module}/{resource}.go        # API定义
├── model/req/{table}.go              # 请求模型
├── model/res/{table}.go              # 响应模型
├── logic/{module}/{table}.go         # 业务逻辑
└── controller/{module}/{resource}.go # 控制器
```

**生成的前端文件**（使用 `--frontend`）：
```
admin-ui/apps/backend/src/
├── api/{module}/{resource}.ts                           # 前端API
├── views/{module}/{resource}/
│   ├── index.vue                                        # 页面组件
│   ├── model.ts                                         # 类型定义
│   ├── schemas.ts                                       # 列配置
│   └── use-{resource}-crud.ts                          # CRUD逻辑

resource/migrations/
└── menu_{module}_{resource}.sql                         # 菜单SQL
```

**后续步骤**：
```bash
# 1. 更新服务层
make service

# 2. 生成控制器方法
make ctrl

# 3. 执行菜单SQL（如果使用 --frontend）
# 将 resource/migrations/menu_*.sql 导入数据库
```

### 场景3：克隆现有模块

当用户想基于已有模块快速创建新模块时使用。

```bash
# 克隆模块
go run ./hack/generator/main.go module:clone -source blog -target news

# 或使用 Makefile
make clone-module name=news source=blog
```

### 场景4：创建 Worker 任务

当用户需要创建后台任务时使用。

```bash
# 异步任务
go run ./hack/generator/main.go worker:create -module system -name SendEmail -type task

# 定时任务
go run ./hack/generator/main.go worker:create -module system -name CleanCache -type cron

# 混合任务
go run ./hack/generator/main.go worker:create -module system -name DataSync -type both
```

**生成文件**：
```
modules/{module}/
├── worker/server/{name}_worker.go   # 异步任务
├── worker/cron/{name}_cron.go       # 定时任务
└── consts/worker.go                 # 常量定义（更新）
```

### 场景5：导出/导入模块

当用户需要在不同环境之间迁移模块时使用。

```bash
# 导出模块
go run ./hack/generator/main.go module:export -name blog
# 生成: blog.v1.0.0.zip

# 导入模块
go run ./hack/generator/main.go module:import -file blog.v1.0.0.zip
```

### 场景6：批量生成 CRUD

当用户需要为多个表同时生成代码时使用。

```bash
# 创建 generator.yaml
cat > generator.yaml << 'EOF'
module: system
tables:
  - table: users
    business: User
    description: 用户管理
  - table: roles
    business: Role
    description: 角色管理
EOF

# 批量生成
go run ./hack/generator/main.go crud:generate -c=generator.yaml --frontend
```

## 完整工作流示例

### 从零开始创建业务模块

```bash
# Step 1: 创建模块
go run ./hack/generator/main.go module:create -name shop

# Step 2: 创建数据库表 shop_product
# ... 使用数据库迁移工具 ...

# Step 3: 生成 CRUD（前后端）
go run ./hack/generator/main.go crud:generate -m=shop -t=shop_product -n=商品 --frontend

# Step 4: 更新代码
make service
make ctrl

# Step 5: 导入菜单SQL
# psql -d devinggo -f resource/migrations/menu_shop_product.sql

# Step 6: 启动服务
go run main.go
```

## 注意事项

1. **数据库表必须先存在**：CRUD 生成器通过解析 `internal/model/entity/` 下的 Entity 文件来获取字段信息
2. **生成后必须运行 `make service`**：更新 service 接口
3. **前端代码生成后需要执行菜单 SQL**：菜单是通过数据库驱动的，需要导入生成的 SQL
4. **Worker 任务创建后需要重启服务**：Worker 服务需要重新加载任务
5. **系统模块（system）不能被克隆或导出**
6. **模板分隔符**：
   - 后端模板使用 `{{.Variable}}`
   - 前端模板使用 `<%.Variable%>`（避免与 Vue 语法冲突）

## 常见问题

### Q: 生成代码时报错 "读取 entity 文件失败"
**A**: 需要先通过 `make dao` 或 `gf gen dao` 生成 Entity 文件

### Q: 前端页面不显示菜单
**A**: 需要执行生成的菜单 SQL 文件（`resource/migrations/menu_*.sql`）

### Q: 如何只生成前端代码？
**A**: 无法单独生成前端代码，必须使用 `--frontend` 参数同时生成前后端

### Q: 生成的代码可以自定义吗？
**A**: 可以修改 `hack/generator/templates/crud/` 下的模板文件

## 相关文件

- `hack/generator/main.go` - CLI 入口
- `hack/generator/templates/crud/` - CRUD 模板
- `hack/generator/templates/module/` - 模块模板
- `hack/generator/templates/worker/` - Worker 模板
- `hack/hack.mk` - Makefile 命令定义
