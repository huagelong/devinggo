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

## 完整 CLI 参考

### 1. 主程序 CLI (`go run main.go`)

入口：`main.go`，基于 GoFrame `gcmd` 框架。

| 命令 | 用法 | 说明 |
|------|------|------|
| 默认无参 | `go run main.go` | 启动所有服务（等同于 `all`） |
| `all` | `go run main.go all` | 启动 HTTP + Worker 服务 |
| `http` | `go run main.go http` | 仅启动 HTTP 服务 |
| `worker` | `go run main.go worker` | 仅启动 Worker 服务 |
| `version` | `go run main.go version` | 查看版本 |
| `unpack` | `go run main.go unpack` | 释放打包的资源文件（如配置文件） |
| `install` | `go run main.go install` | 系统初始化安装 |
| `migrate:up` | `go run main.go migrate:up [-n N]` | 应用 N 个 up 迁移，N 为空则应用全部 |
| `migrate:down` | `go run main.go migrate:down [-n N]` | 应用 N 个 down 迁移 |
| `migrate:goto` | `go run main.go migrate:goto -v 版本` | 迁移到指定版本 |
| `migrate:create` | `go run main.go migrate:create -name NAME` | 创建带时间戳的 up/down 迁移文件 |
| `migrate:force` | `go run main.go migrate:force -v 版本` | 强制设置版本（忽略脏状态） |
| `help` | `go run main.go help` | 查看帮助 |

**全局参数**：
- `-c, --config`：指定配置文件，默认 `config.yaml`
  ```bash
  go run main.go -c hack/config.yaml
  ```

### 2. GoFrame CLI (`gf`)

项目依赖 GoFrame CLI 进行代码生成，常用命令：

| 命令 | 说明 |
|------|------|
| `gf run main.go` | 热编译运行 |
| `gf gen dao` | 根据数据库表生成 Entity/DAO/DO |
| `gf gen service` | 根据 logic 生成 service 接口 |
| `gf gen ctrl` | 根据 api 生成 controller/sdk |
| `gf gen enums` | 扫描枚举生成 enums.go |
| `gf gen pb` | 解析 protobuf 并生成 go 文件 |
| `gf gen pbentity` | 根据数据库表生成 protobuf entity |
| `gf docker` | 构建 Docker 镜像 |
| `gf up -a` | 升级 GoFrame 到最新版 |

安装/更新 gf CLI：
```bash
make cli
```

### 3. Makefile 命令总览

文件：`Makefile`、`hack/hack.mk`、`hack/hack-cli.mk`、`hack/hack-cus.mk`

#### 开发与构建

| 命令 | 说明 |
|------|------|
| `make run` | 开发模式运行，先执行 `make dao service` 再 `go run main.go` |
| `make build` | 生产构建，自动构建前端并打包到 `resource/public/admin` |
| `make install` | 执行 `go run main.go install` |
| `make ui.install` | 安装前端依赖（`yarn install`） |
| `make ui.build` | 构建前端（`yarn build`） |

#### GoFrame 代码生成

| 命令 | 说明 |
|------|------|
| `make dao` | 生成 DAO/DO/Entity |
| `make service` | 生成 Service 接口 |
| `make ctrl` | 生成 Controller/SDK |
| `make enums` | 生成枚举文件 |
| `make pb` | 生成 protobuf 代码 |
| `make pbentity` | 生成 protobuf entity |

#### 代码质量

| 命令 | 说明 |
|------|------|
| `make lint` | 执行 golangci-lint |
| `make fmt` | 执行 goimports 和 gofmt 格式化 |

#### Docker / 部署

| 命令 | 说明 |
|------|------|
| `make image` | 构建 Docker 镜像 |
| `make image.push` | 构建并推送镜像 |
| `make deploy` | 使用 kustomize + kubectl 部署 |

#### 代码生成器封装（`make gen-help` 可查看）

| 命令 | 用法示例 | 说明 |
|------|----------|------|
| `make gen-module` | `make gen-module name=blog` | 创建新模块 |
| `make clone-module` | `make clone-module name=news source=blog` | 克隆模块 |
| `make export-module` | `make export-module name=blog` | 导出模块 zip |
| `make import-module` | `make import-module file=./blog.zip` | 导入模块 zip |
| `make list-modules` | `make list-modules` | 列出所有模块 |
| `make validate-module` | `make validate-module name=blog` | 验证模块结构 |
| `make gen-worker` | `make gen-worker module=system worker=SendEmail type=task` | 创建 Worker |
| `make gen-crud` | `make gen-crud table=system_user module=system name=用户 frontend=1` | 生成 CRUD |

### 4. 代码生成器 CLI (`go run ./hack/generator/main.go`)

| 命令 | 用法 | 说明 |
|------|------|------|
| `module:create` | `-name <模块名>` | 创建新模块 |
| `module:clone` | `-source <源> -target <目标>` | 克隆现有模块 |
| `module:export` | `-name <模块名>` | 导出模块为 zip |
| `module:import` | `-file <zip路径>` | 导入模块 |
| `module:list` | 无参数 | 列出所有模块 |
| `module:validate` | `-name <模块名>` | 验证模块完整性 |
| `worker:create` | `-module <模块> -name <名称> [-type task/cron/both]` | 创建 Worker 任务 |
| `crud:generate` | `-m <模块> -t <表名> -n <中文名> [--frontend]` | 生成 CRUD 代码 |

### 5. 前端 admin-ui CLI

进入 `admin-ui/` 目录执行：

| 命令 | 说明 |
|------|------|
| `pnpm install` | 安装依赖 |
| `pnpm dev` | 启动所有前端应用开发服务器 |
| `pnpm dev:backend` | 仅启动 backend 应用 |
| `pnpm dev:docs` | 仅启动 docs 应用 |
| `pnpm build` | 构建所有应用 |
| `pnpm build:backend` | 仅构建 backend 应用 |
| `pnpm build:docs` | 仅构建 docs 应用 |
| `pnpm build:analyze` | 构建并分析包体积 |
| `pnpm lint` | 执行 ESLint 检查 |
| `pnpm format` | 自动格式化代码 |
| `pnpm check` | 执行所有检查（循环依赖、拼写、类型、依赖） |
| `pnpm check:type` | 执行 TypeScript 类型检查 |
| `pnpm check:circular` | 检查循环依赖 |
| `pnpm check:dep` | 检查依赖问题 |
| `pnpm check:cspell` | 拼写检查 |
| `pnpm test:unit` | 运行单元测试 |
| `pnpm test:e2e` | 运行 E2E 测试 |
| `pnpm clean` | 清理构建产物 |
| `pnpm reinstall` | 完全重新安装依赖 |
| `pnpm commit` | 使用 czg 提交（交互式） |

### 6. 数据库迁移

项目使用 GoFrame 的迁移组件，命令通过主程序 CLI 调用：

```bash
# 创建迁移
go run main.go migrate:create -name create_users_table

# 应用迁移
go run main.go migrate:up
go run main.go migrate:up -n 1

# 回滚迁移
go run main.go migrate:down
go run main.go migrate:down -n 1

# 跳到指定版本
go run main.go migrate:goto -v 20240101120000

# 强制设置版本
go run main.go migrate:force -v 20240101120000
```

### 7. 常用组合工作流

```bash
# 开发启动（生成 dao/service 后运行）
make run

# 新建表后更新实体
gf gen dao
# 或
make dao

# 修改 logic 后更新 service 接口
make service

# 修改 api 后生成 controller
make ctrl

# 为表生成前后端 CRUD
make gen-crud table=shop_product module=shop name=商品 frontend=1

# 生成后更新 service/controller
make service
make ctrl

# 构建生产包
make build

# 运行生产二进制
./main
```

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

# 4. 检查 i18n 翻译（前端页面显示需要）
# 翻译会自动生成到 admin-ui/apps/backend/src/locales/langs/{zh-CN,en-US}/system.json
# 如需调整，请手动修改对应语言文件
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

## 生成后检查与修复

CRUD 生成完成后，建议按以下步骤检查验证：

### 1. 后端编译检查
```bash
# 确保后端代码能正常编译
go build ./...
```

### 2. Service 接口检查
```bash
# 运行 gf gen service 重新生成 service 层
make service
```

### 3. 路由注册检查
```bash
# 确认控制器已注册到 router.go
grep -n "{Resource}Controller" modules/{module}/router/{module}/router.go
```

### 4. 前端组件检查

**常见问题**：浏览器控制台报错组件未注册
- `Failed to resolve component: Textarea`
- `Failed to resolve component: RadioGroup`
- `Failed to resolve component: InputNumber`

**修复方法**：检查 `views/{module}/{resource}/index.vue` 的 import 部分是否包含这些组件：
```typescript
import {
  Button, Input, InputNumber, RadioGroup, Textarea, ...
} from 'tdesign-vue-next';
```

### 5. i18n 翻译检查

**常见问题**：控制台警告 `Not found 'system.xxx.title' key in 'zh' locale messages`

**翻译数据来源**：
- generator 会自动读取数据库表字段的 `description`（注释）作为翻译值
- 如果字段没有注释，则回退到根据字段名生成的合理中文翻译（如 `title`→"标题"、`content`→"内容"、`status`→"状态"）
- 支持常见字段名智能翻译，也支持驼峰命名拆分（如 `userName`→"用户名称"）
- 因此，**建表时建议为字段添加中文注释**，这样前端显示最准确；即使无注释，也能获得合理的中文翻译

**示例**：
```sql
CREATE TABLE test_cms (
    id          bigserial PRIMARY KEY,
    title       varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    content     text COMMENT '内容',
    status      int NOT NULL DEFAULT 0 COMMENT '状态',
    remark      varchar(500) DEFAULT '' COMMENT '备注',
    sort        int NOT NULL DEFAULT 0 COMMENT '排序',
    created_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp DEFAULT CURRENT_TIMESTAMP
);
```

**检查位置**：
- `admin-ui/apps/backend/src/locales/langs/zh-CN/system.json`
- `admin-ui/apps/backend/src/locales/langs/en-US/system.json`

**自动生成的翻译示例**：
```json
{
  "testCms": {
    "title": "标题",
    "content": "内容",
    "status": "状态",
    "remark": "备注",
    "sort": "排序",
    "editTitle": "编辑内容管理",
    "createTitle": "新增内容管理"
  }
}
```

**如需要手动修复或调整翻译**：
1. 修改对应的数据库表字段注释
2. 重新运行 `make dao` 更新 entity 文件
3. 重新运行 generator 生成代码

> **最佳实践**：在建表阶段就为所有业务字段添加清晰的中文注释，这样生成的前端页面可以直接使用，无需二次修改翻译文件。

## 常见问题

### Q: 生成代码时报错 "读取 entity 文件失败"

**错误示例**：
```
❌ 生成CRUD代码失败：读取entity文件失败：open ...\internal\model\entity\xxx.go: The system cannot find the file specified.
```

**原因**：CRUD 生成器依赖 `internal/model/entity/` 目录下的 Entity 文件来解析表结构，如果表刚创建或 entity 文件未生成，就会出现此错误。

**解决步骤**：

1. **确认数据库表已存在**

2. **生成 Entity 文件**（以下方式任选其一）：
   ```bash
   # 方式1：使用 make
   make dao
   
   # 方式2：直接使用 gf 命令
   gf gen dao
   ```

3. **确认 entity 文件已生成**：
   ```bash
   ls internal/model/entity/xxx.go
   ```

4. **重新运行 CRUD 生成命令**：
   ```bash
   go run ./hack/generator/main.go crud:generate -m=<模块> -t=<表名> -n=<中文名> --frontend
   ```

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
