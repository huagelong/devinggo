# DevingGo 代码生成工具

DevingGo代码生成工具集，提供模块管理、Worker任务和CRUD代码生成功能。

## 功能特性

### 1. 模块管理 (module) ✅
- `module:create` - 创建新模块 ✅
- `module:clone` - 克隆现有模块 ✅
- `module:export` - 导出模块包 ✅
- `module:import` - 导入模块包 ✅
- `module:list` - 列出已安装模块 ✅
- `module:validate` - 验证模块完整性 ✅

### 2. Worker任务生成 (worker) ✅
- `worker:create` - 创建Worker任务（task/cron/both）✅
- 三种任务类型：异步任务/定时任务/混合任务 ✅
- 智能常量文件管理 ✅
- 自动注册到Worker管理器 ✅

### 3. CRUD代码生成 (crud) ✅
- `crud:generate` - 生成CRUD代码 ✅
- 智能字段解析和类型映射 ✅
- 自动生成后端API/请求/响应/Logic/Controller ✅
- **自动生成前端代码（Vue3 + TypeScript）** ✅
  - API 层 (`api/{module}/{resource}.ts`)
  - 页面组件 (`views/{module}/{resource}/index.vue`)
  - 类型定义 (`views/{module}/{resource}/model.ts`)
  - 列配置 (`views/{module}/{resource}/schemas.ts`)
  - CRUD 逻辑 (`views/{module}/{resource}/use-{resource}-crud.ts`)
  - 菜单SQL (`resource/migrations/menu_{module}_{resource}.sql`)
- 支持自定义模块和业务名称 ✅
- 支持同时生成前后端代码（`--frontend` 参数） ✅

## 目录结构

```
hack/generator/
├── main.go                      # CLI入口
├── README.md                    # 本文档
├── generator.yaml               # 批量生成配置
├── cmd/                         # 命令层
│   ├── module.go                # 模块管理命令
│   ├── worker.go                # Worker任务命令
│   └── crud.go                  # CRUD生成命令
├── internal/                    # 核心逻辑层
│   ├── generator/               # 生成器引擎
│   │   ├── crud.go              # CRUD生成器
│   │   ├── module.go            # 模块导入导出
│   │   ├── module_create.go     # 模块创建器
│   │   ├── module_test.go       # 模块测试
│   │   ├── template.go          # 模板渲染引擎
│   │   └── worker.go            # Worker生成器
│   ├── scanner/                 # 扫描器
│   │   └── module.go            # 模块扫描器
│   ├── utils/                   # 工具函数
│   │   ├── file.go              # 文件操作
│   │   ├── hooks.go             # 钩子执行器
│   │   ├── interactive.go       # 交互式提示
│   │   ├── naming.go            # 命名转换
│   │   ├── naming_test.go       # 命名转换测试
│   │   ├── output.go            # 命令结果输出
│   │   ├── sensitive.go         # 敏感信息清理
│   │   ├── zip.go               # 压缩/解压
│   │   └── zip_test.go          # 压缩测试
│   └── config/                  # 配置管理
│       └── config.go            # 配置解析
├── templates/                   # 模板文件
│   ├── crud/                    # CRUD模板
│   │   ├── api.go.tpl           # 后端API定义
│   │   ├── controller.go.tpl    # 后端控制器
│   │   ├── logic.go.tpl         # 后端逻辑层
│   │   ├── req.go.tpl           # 后端请求模型
│   │   ├── res.go.tpl           # 后端响应模型
│   │   └── frontend/            # 前端模板
│   │       ├── api.ts.tpl       # 前端API
│   │       ├── index.vue.tpl    # 前端页面
│   │       ├── menu.sql.tpl     # 菜单SQL
│   │       ├── model.ts.tpl     # 前端类型
│   │       ├── schemas.ts.tpl   # 前端列配置
│   │       └── use-crud.ts.tpl  # 前端CRUD逻辑
│   ├── module/                  # 模块模板
│   └── worker/                  # Worker模板
└── docs/                        # 文档
    └── MODULE_YAML_SPEC.md      # .module.yaml配置规范
```

## 使用方法

### 推荐：使用 Makefile 命令（简洁方式）

```bash
# 查看所有生成器命令
make gen-help

# 模块管理
make gen-module name=blog                    # 创建新模块
make clone-module name=news source=blog       # 克隆模块
make export-module name=blog                  # 导出模块
make import-module file=blog.v1.0.0.zip      # 导入模块
make list-modules                             # 列出所有模块
make validate-module name=blog                # 验证模块

# Worker任务生成
make gen-worker module=system worker=SendEmail

# CRUD生成（仅后端）
make gen-crud table=system_user module=system name=用户

# CRUD生成（前后端）
make gen-crud table=system_user module=system name=用户 frontend=1
```

### 完整命令示例

#### 模块管理

```bash
# 创建新模块
go run ./hack/generator/main.go module:create -name blog

# 克隆模块
go run ./hack/generator/main.go module:clone -source blog -target news

# 导出模块
go run ./hack/generator/main.go module:export -name blog

# 导入模块
go run ./hack/generator/main.go module:import -file blog.v1.0.0.zip

# 列出模块
go run ./hack/generator/main.go module:list

# 验证模块
go run ./hack/generator/main.go module:validate -name blog
```

#### Worker任务

```bash
# 创建异步任务
go run ./hack/generator/main.go worker:create -module system -name SendEmail -type task

# 创建定时任务
go run ./hack/generator/main.go worker:create -module system -name CleanCache -type cron

# 创建混合任务
go run ./hack/generator/main.go worker:create -module system -name DataSync -type both
```

#### CRUD生成

```bash
# 仅生成后端代码
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户

# 同时生成前后端代码
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户 --frontend

# 预览模式（不写入文件）
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户 --dry-run

# 强制覆盖
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户 --force

# 批量生成（配置文件）
go run ./hack/generator/main.go crud:generate -c=generator.yaml

# JSON格式输出
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户 -o=json
```

### 生成文件说明

#### 后端文件
```
modules/{module}/
├── api/{module}/{resource}.go        # API定义
├── model/req/{table}.go              # 请求模型
├── model/res/{table}.go              # 响应模型
├── logic/{module}/{table}.go         # 业务逻辑
└── controller/{module}/{resource}.go # 控制器
```

#### 前端文件（使用 --frontend 时生成）
```
admin-ui/apps/backend/src/
├── api/{module}/{resource}.ts                           # 前端API
├── views/{module}/{resource}/
│   ├── index.vue                                        # 页面组件
│   ├── model.ts                                         # 类型定义
│   ├── schemas.ts                                       # 列配置
│   └── use-{resource}-crud.ts                          # CRUD逻辑
└── router/routes/modules/{module}.ts                    # 路由配置（数据库驱动，需执行SQL）

resource/migrations/
└── menu_{module}_{resource}.sql                         # 菜单SQL
```

## 完整工作流示例

### 场景1：创建新业务模块

```bash
# Step 1: 创建模块
make gen-module name=blog

# Step 2: 生成CRUD代码（假设已有blog_post表）
make gen-crud table=blog_post module=blog name=帖子 frontend=1

# Step 3: 创建Worker任务
make gen-worker module=blog worker=PublishPost

# Step 4: 更新代码
make dao      # 更新数据访问层
make service  # 更新服务层
make ctrl     # 更新控制器

# Step 5: 执行菜单SQL
# 将生成的 resource/migrations/menu_blog_post.sql 导入数据库

# Step 6: 启动服务
go run main.go
```

### 场景2：从现有模块快速开发

```bash
# Step 1: 克隆现有模块
make clone-module name=product source=blog

# Step 2: 生成CRUD
make gen-crud table=product_item module=product name=商品 frontend=1
make gen-crud table=product_category module=product name=分类 frontend=1

# Step 3: 更新代码并启动
make service && make ctrl
go run main.go
```

### 场景3：模块分发和部署

```bash
# 在开发机器上：
# 1. 导出模块
make export-module name=blog
# 生成: blog.v1.0.0.zip

# 在生产机器上：
# 2. 导入模块
make import-module file=blog.v1.0.0.zip

# 3. 更新代码并启动
make service && make dao
go run main.go
```

## 配置文件

### generator.yaml

批量生成CRUD代码的配置文件：

```yaml
# 目标模块
module: system

# 生成配置列表
tables:
  - table: users
    business: User
    description: 用户管理
    
  - table: roles
    business: Role
    description: 角色管理
    
  - table: permissions
    business: Permission
    description: 权限管理
```

## 工具函数

### naming.go - 命名转换

```go
// snake_case
ToSnakeCase("CategoryName") // "category_name"

// camelCase
ToCamelCase("category_name") // "categoryName"

// PascalCase
ToPascalCase("category_name") // "CategoryName"

// CONST_CASE
ToConstCase("CategoryName") // "CATEGORY_NAME"

// kebab-case
ToKebabCase("CategoryName") // "category-name"
```

### output.go - 命令结果输出

```go
// 创建结果
result := utils.NewCommandResult(true, "操作成功")

// 添加文件
result.AddFile("path/to/file.go")

// 添加警告
result.AddWarning("提示信息")

// 添加错误
result.AddError("错误信息")

// 输出结果
result.Print(utils.OutputFormatText)  // 文本格式
result.Print(utils.OutputFormatJSON)  // JSON格式
```

### file.go - 文件操作

```go
// 检查路径
PathExists(path) bool
IsDir(path) bool
IsFile(path) bool

// 目录操作
EnsureDir(dir) error
GetProjectRoot() (string, error)

// 文件操作
CopyFile(src, dst) error
```

### zip.go - 压缩/解压

```go
// 压缩目录
ZipDirectory(ctx, srcDir, dstZip) error

// 解压文件
UnzipFile(srcZip, dstDir) error
```

## 注意事项

1. **独立性**: 工具不依赖项目业务代码，可独立编译和分发
2. **安全性**: 解压文件时会检查路径穿越攻击
3. **代码格式化**: 生成的Go代码会自动格式化（gofmt/goimports）
4. **模板变量**: 
   - 后端模板使用 `{{.Variable}}` 格式（Go template语法）
   - 前端模板使用 `<%.Variable%>` 格式（避免与Vue语法冲突）
5. **菜单SQL**: 生成后需要手动执行SQL或放入migration文件

## 版本历史

- **v1.0.0** (2026-06-10)
  - ✅ 完成基础架构搭建
  - ✅ 完成模块管理功能（创建/克隆/导入/导出/列表/验证）
  - ✅ 实现.module.yaml配置系统
  - ✅ 实现变量替换引擎
  - ✅ 完成Worker任务生成器（task/cron/both三种类型）
  - ✅ 完成CRUD代码生成器（后端）
  - ✅ 完成CRUD前端代码生成（Vue3 + TypeScript）
  - ✅ 集成到Makefile工作流
  - ✅ 添加菜单SQL生成功能

## 相关文档

- [.module.yaml 配置规范](docs/MODULE_YAML_SPEC.md)

## 贡献指南

欢迎提交Issue和Pull Request来改进这个工具。

### 开发环境要求

- Go 1.23+
- GoFrame v2

### 本地开发

```bash
# 克隆项目
git clone https://github.com/huagelong/devinggo.git
cd devinggo/hack/generator

# 运行工具
go run main.go -h

# 运行测试
go test ./...
```

## 许可证

MIT License
