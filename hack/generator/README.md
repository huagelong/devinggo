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
- 自动生成API/请求/响应/Logic/Controller ✅
- 支持自定义模块和业务名称 ✅

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
│   │   ├── module.go            # 模块生成器
│   │   ├── worker.go            # Worker生成器
│   │   ├── crud.go              # CRUD生成器
│   │   └── template.go          # 模板渲染引擎
│   ├── scanner/                 # 扫描器
│   │   └── module.go            # 模块扫描器
│   ├── utils/                   # 工具函数
│   │   ├── zip.go               # 压缩/解压
│   │   ├── file.go              # 文件操作
│   │   └── naming.go            # 命名转换
│   └── config/                  # 配置管理
│       └── config.go            # 配置解析
└── templates/                   # 模板文件
    ├── module/                  # 模块模板
    ├── worker/                  # Worker模板
    └── crud/                    # CRUD模板
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
# 1. 创建新模块
make gen-module name=blog
# 或完整命令
go run ./hack/generator/main.go module:create -name blog

# 2. 克隆现有模块（快速创建）
make clone-module name=news source=blog
# 或完整命令
go run ./hack/generator/main.go module:clone -source blog -target news

# 3. 导出模块为zip包
make export-module name=blog
# 生成文件: blog.v1.0.0.zip

# 4. 导入模块包
make import-module file=blog.v1.0.0.zip
# 自动部署文件、执行迁移、合并配置

# 5. 列出已安装模块
make list-modules
# 显示所有模块的名称、版本、作者信息

# 6. 验证模块完整性
make validate-module name=blog
# 检查配置文件、文件路径、依赖关系
```

### Worker任务示例

```bash
# 1. 创建异步任务（Task）
make gen-worker module=system worker=SendEmail
# 或完整命令
go run ./hack/generator/main.go worker:create -module system -name SendEmail -type task

# 生成文件:
# - modules/system/worker/server/send_email_worker.go
# - modules/system/consts/worker.go (更新常量)

# 2. 创建定时任务（Cron）
make gen-worker module=system worker=CleanCache
go run ./hack/generator/main.go worker:create -module system -name CleanCache -type cron

# 生成文件:
# - modules/system/worker/cron/clean_cache_cron.go
# - modules/system/consts/worker.go (更新常量)

# 3. 创建混合任务（Both）
make gen-worker module=system worker=DataSync
go run ./hack/generator/main.go worker:create -module system -name DataSync -type both

# 生成文件:
# - modules/system/worker/server/data_sync_worker.go
# - modules/system/worker/cron/data_sync_cron.go (复用task数据结构)
# - modules/system/consts/worker.go (更新常量)

# 4. 生成后需要运行
# 1. 使用默认模块（system）生成CRUD
make gen-crud table=system_user
# 或完整命令
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户

# 生成文件:
# - modules/system/api/system/system_user.go        # API定义
# - modules/system/model/req/system_user.go         # 请求模型
# - modules/system/model/res/system_user.go         # 响应模型
# - modules/system/logic/system/system_user.go      # 业务逻辑
# - modules/system/controller/system/system_user.go # 控制器

# 2. 指定模块生成CRUD
make gen-crud table=system_post module=system
go run ./hack/generator/main.go crud:generate -m=system -t=system_post -n=帖子

# 3. 自定义业务名称
go run ./hack/generator/main.go crud:generate -m=system -t=system_user -n=用户

# 4. 生成后需要运行
make service  # 更新service接口
make ctrl     # 生成controller方法

# 5. 完整工作流
make gen-crud table=system_dept
make service
make ctrl
# 现在可以启动服务测试API了
```

## 完整工作流示例

### 场景1：创建新业务模块

```bash
# Step 1: 创建模块
make gen-module name=blog

# Step 2: 生成CRUD代码（假设已有blog_post表）
make gen-crud table=blog_post module=blog

# Step 3: 创建Worker任务
make gen-worker module=blog worker=PublishPost

# Step 4: 更新代码
make dao      # 更新数据访问层
make service  # 更新服务层
make ctrl     # 更新控制器

# Step 5: 启动服务
go run main.go
```

### 场景2：从现有模块快速开发

```bash
# Step 1: 克隆现有模块
make clone-module name=product source=blog

# Step 2: 生成CRUD
make gen-crud table=product_item module=product
make gen-crud table=product_category module=product

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
# 自动执行迁移、部署文件

# 3. 更新代码并启动
make service && make dao
go run main.go

# 批量生成（使用配置文件）
go run hack/generator/main.go crud:generate -config=generator.yaml
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
6) - 首个完整版本
  - ✅ 完成基础架构搭建
  - ✅ 完成模块管理功能（创建/克隆/导入/导出/列表/验证）
  - ✅ 实现.module.yaml配置系统
  - ✅ 实现变量替换引擎
  - ✅ 创建17个模块模板文件
  - ✅ 完成Worker任务生成器（task/cron/both三种类型）
  - ✅ 完成CRUD代码生成器
  - ✅ 集成到Makefile工作流
  - ✅ 清理旧代码，重命名模块加载器

## 相关文档

- [技术方案](../../docs/PLAN-CodeGeneratorUnification.md)
- [任务清单](../../docs/TODO-CodeGeneratorUnification.md)
- [阶段一完成报告](../../docs/STAGE1_COMPLETION_REPORT.md)
- [阶段二完成报告](../../docs/STAGE2_COMPLETION_REPORT.md)
- [阶段三完成报告](../../docs/STAGE3_COMPLETION_REPORT.md)
- [阶段四完成报告](../../docs/STAGE4_COMPLETION_REPORT.md)
- [阶段五完成报告](../../docs/STAGE5_COMPLETION_REPORT.md)
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

### file.go - 文件操作

```go
// 检查路径
PathExists(path) bool
IsDir(path) bool
IsFile(path) bool

// 目录操作
EnsureDir(dir) error
GetProjectRoot() (string, error)
GetModuleName() (string, error)

// 文件操作
CopyFile(src, dst) error
WriteFile(path, content, overwrite) error

// 代码格式化
FormatGoCode(filePath) error
FormatGoCodeInDir(dir) error
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
4. **模板变量**: 使用 `{{.Variable}}` 格式，遵循Go template语法

## 版本历史

- **v1.0.0** (2026-03-04)
  - ✅ 完成基础架构搭建
  - ✅ 完成模块管理功能（创建/克隆/导入/导出/列表/验证）
  - ✅ 实现.module.yaml配置系统
  - ✅ 实现变量替换引擎
  - ✅ 创建17个模块模板文件
  - ⏳ Worker任务生成（待开发）
  -.module.yaml 配置规范](docs/MODULE_YAML_SPEC.md) ✅
- [ ⏳ CRUD代码生成（待开发）

## 相关文档

- [技术方案](../../docs/PLAN-CodeGeneratorUnification.md)
- [任务清单](../../docs/TODO-CodeGeneratorUnification.md)
- [模块开发指南](../../docs/MODULE.md) - 待创建
- [Worker开发指南](../../docs/WORKER.md) - 待创建

## 许可证

MIT License
