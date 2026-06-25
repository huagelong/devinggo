<p align="center">
  <img src="docs/logo.svg" width="120" />
</p>
<p align="center">
 <img src="https://svg.hamm.cn/badge.svg?key=License&value=Apache-2.0%20%2F%20MIT" />
  <img src="https://svg.hamm.cn/badge.svg?key=Version&value=1.0.0%20LTS" />
	<img src="https://img.shields.io/badge/goframe-2.9-green" alt="goframe">
	<img src="https://img.shields.io/badge/golang-1.23.4-blue" alt="golang">
	<img src="https://img.shields.io/badge/postgresql-13%2B-blue" alt="postgresql">
</p>
<p  align="center">
<a href="https://zread.ai/huagelong/devinggo" target="_blank"><img src="https://img.shields.io/badge/Ask_Zread-_.svg?style=flat&color=00b0aa&labelColor=000000&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIHZpZXdCb3g9IjAgMCAxNiAxNiIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTQuOTYxNTYgMS42MDAxSDIuMjQxNTZDMS44ODgxIDEuNjAwMSAxLjYwMTU2IDEuODg2NjQgMS42MDE1NiAyLjI0MDFWNC45NjAxQzEuNjAxNTYgNS4zMTM1NiAxLjg4ODEgNS42MDAxIDIuMjQxNTYgNS42MDAxSDQuOTYxNTZDNS4zMTUwMiA1LjYwMDEgNS42MDE1NiA1LjMxMzU2IDUuNjAxNTYgNC45NjAxVjIuMjQwMUM1LjYwMTU2IDEuODg2NjQgNS4zMTUwMiAxLjYwMDEgNC45NjE1NiAxLjYwMDFaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00Ljk2MTU2IDEwLjM5OTlIMi4yNDE1NkMxLjg4ODEgMTAuMzk5OSAxLjYwMTU2IDEwLjY4NjQgMS42MDE1NiAxMS4wMzk5VjEzLjc1OTlDMS42MDE1NiAxNC4xMTM0IDEuODg4MSAxNC4zOTk5IDIuMjQxNTYgMTQuMzk5OUg0Ljk2MTU2QzUuMzE1MDIgMTQuMzk5OSA1LjYwMTU2IDE0LjExMzQgNS42MDE1NiAxMy43NTk5VjExLjAzOTlDNS42MDE1NiAxMC42ODY0IDUuMzE1MDIgMTAuMzk5OSA0Ljk2MTU2IDEwLjM5OTlaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik0xMy43NTg0IDEuNjAwMUgxMS4wMzg0QzEwLjY4NSAxLjYwMDEgMTAuMzk4NCAxLjg4NjY0IDEwLjM5ODQgMi4yNDAxVjQuOTYwMUMxMC4zOTg0IDUuMzEzNTYgMTAuNjg1IDUuNjAwMSAxMS4wMzg0IDUuNjAwMUgxMy43NTg0QzE0LjExMTkgNS42MDAxIDE0LjM5ODQgNS4zMTM1NiAxNC4zOTg0IDQuOTYwMVYyLjI0MDFDMTQuMzk4NCAxLjg4NjY0IDE0LjExMTkgMS42MDAxIDEzLjc1ODQgMS42MDAxWiIgZmlsbD0iI2ZmZiIvPgo8cGF0aCBkPSJNNCAxMkwxMiA0TDQgMTJaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00IDEyTDEyIDQiIHN0cm9rZT0iI2ZmZiIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIvPgo8L3N2Zz4K&logoColor=ffffff" alt="zread"/></a>
</p>

## AI 大模型一键初始化

本项目内置初始化 Skill：`.codex/skills/devinggo-project-init/SKILL.md`。在支持 Skill 的 AI 大模型工具中输入：

```text
使用 $devinggo-project-init 帮我检查本机环境，配置 DevingGo，初始化数据库，并启动后端、Worker 和前端服务。
```

该 Skill 会先检查本机已安装的 Go、GoFrame CLI、Node、pnpm、PostgreSQL、Redis、Make 等工具及版本，满足要求时直接复用现有环境；缺失或版本不满足时再提示安装或升级。初始化过程中会优先使用本项目的 CLI、Makefile、配置样例和前端工作区，避免覆盖已有配置，并在执行数据库迁移、安装依赖、启动服务前说明影响范围。

## 简介

`DevingGo` 是基于 **GoFrame v2.9** 的企业级前后端一体化开发框架，采用模块化架构，内置权限管理、代码生成、任务调度等企业级功能，支持 **Vibe Coding**（AI辅助编程）一键生成前后端代码。

### 特点

- 🚀 **高性能**: 基于GoFrame，发挥Go高并发优势
- 🤖 **Vibe Coding**: AI辅助编程，通过 `skills/generator/` 一键生成前后端代码
- 📦 **模块化**: 插件化模块系统，支持独立开发和热插拔
- 🛠️ **代码生成**: 一键生成CRUD/Worker/Module前后端代码
- 🔐 **安全可靠**: RBAC权限模型，JWT认证
- 🌍 **国际化**: 内置多语言支持
- 🐳 **容器化**: 支持Docker一键部署

### 技术亮点

- **分布式WebSocket**: 完整实现Pusher Protocol v8.3.0
- **异步任务队列**: 基于Asynq，支持任务重试和监控
- **数据库迁移**: 版本化管理，支持迁移回滚
- **链路追踪**: 集成OpenTelemetry + Jaeger
- **API文档**: 自动生成OpenAPI文档，支持在线调试

## 演示地址

- **访问地址**: http://122.152.211.127:4000/admin/#/
- **测试账号**: superAdmin / admin123

## 文档

- [zread](https://zread.ai/huagelong/devinggo)

## 项目结构

```
devinggo/
├── api/                    # 全局API定义
├── internal/               # 内部核心代码
│   ├── cmd/                # 命令入口
│   ├── controller/         # 控制器层
│   ├── dao/                # 数据访问层（自动生成）
│   ├── logic/              # 业务逻辑层
│   ├── model/              # 数据模型（do/entity自动生成）
│   ├── service/            # 服务接口层（自动生成）
│   └── router/             # 路由注册
├── modules/                # 业务模块（插件化）
│   ├── system/             # 系统核心模块
│   │   ├── api/            # API定义
│   │   ├── controller/     # 控制器
│   │   ├── logic/          # 业务逻辑
│   │   ├── service/        # 服务接口
│   │   ├── model/          # 数据模型
│   │   ├── router/         # 路由
│   │   ├── worker/         # 异步任务
│   │   ├── cmd/            # 命令
│   │   └── pkg/            # 工具包
│   └── api/                # API模块（示例）
├── admin-ui/               # 前端管理界面（Vue3 + Vite）
│   ├── apps/               # 应用目录
│   ├── packages/           # 共享包
│   └── scripts/            # 构建脚本
├── hack/                   # 开发工具
│   └── generator/          # 代码生成器
├── .opencode/skills/       # Vibe Coding Skill配置
│   └── generator/          # AI代码生成Skill
├── manifest/               # 配置清单
│   ├── config/             # 配置文件
│   ├── docker/             # Docker配置
│   └── deploy/             # 部署配置
├── resource/               # 静态资源
│   ├── i18n/               # 国际化
│   ├── migrations/         # 数据库迁移
│   └── public/             # 公共资源
├── docs/                   # 文档
├── logs/                   # 日志
└── main.go                 # 程序入口
```

### 架构说明

- **分层架构**: API -> Controller -> Logic -> Service -> DAO
- **模块化**: 支持模块独立开发、热插拔，详见 [模块架构说明](modules/README.md)
- **代码生成**: DAO/Entity/DO/Service 由工具自动生成
- **配置分离**: 支持多环境部署

## 技术栈

### 后端
| 类别 | 技术 |
|------|------|
| 核心框架 | GoFrame v2.9.0 |
| 数据库 | PostgreSQL >= 13.0 |
| 缓存/消息 | Redis >= 5.0 |
| 任务队列 | Asynq v0.24.1 |
| WebSocket | Gorilla WebSocket (Pusher Protocol v8.3.0) |
| 认证 | JWT v5.2.1 |
| 数据库迁移 | golang-migrate v4.17.1 |
| Excel处理 | excelize v2.9.0 |
| 链路追踪 | OpenTelemetry + Jaeger |
| ID生成 | Yitter IdGenerator (雪花算法) |
| API文档 | OpenAPI/Swagger 自动生成 |

### 前端 (admin-ui)
| 类别 | 技术 |
|------|------|
| 框架 | Vue 3 + TypeScript |
| 构建工具 | Vite |
| UI库 | 基于 Vue Vben Admin |
| 状态管理 | Pinia |
| 路由 | Vue Router |
| HTTP请求 | Axios |

## 功能特性

### 系统管理
- [x] **用户管理**: 用户增删改查、密码重置、状态管理
- [x] **部门管理**: 树形组织架构，支持数据权限按部门划分
- [x] **岗位管理**: 用户岗位配置，支持一人多岗
- [x] **角色管理**: RBAC权限分配，支持菜单/数据/接口权限
- [x] **菜单管理**: 菜单、按钮、接口权限配置
- [x] **字典管理**: 字典类型与数据分级维护
- [x] **参数管理**: 系统参数动态配置
- [x] **应用管理**: 多应用管理，支持应用分组和API授权

### Vibe Coding (AI辅助编程)
- [x] **智能代码生成**: 基于 `.opencode/skills/generator/` 配置，通过自然语言描述一键生成前后端代码
- [x] **前后端一体化**: `--frontend` 参数同时生成后端API和前端Vue3页面
- [x] **模块快速搭建**: 自动创建模块结构、数据库迁移、菜单配置
- [x] **代码模板定制**: 支持自定义 `hack/generator/templates/` 下的代码模板

### 开发工具
- [x] **代码生成器**: 基于数据表一键生成CRUD/Worker/Module代码，详见 [代码生成工具](hack/generator/README.md) 和 [Skill配置](.opencode/skills/generator/SKILL.md)
- [x] **API管理**: API接口与分组管理，应用授权
- [x] **接口文档**: 自动生成OpenAPI/Swagger文档

### 监控与日志
- [x] **操作日志**: 用户操作行为记录与导出
- [x] **登录日志**: 登录行为记录，含IP、设备信息
- [x] **API日志**: API请求参数与响应记录
- [x] **系统监控**: Redis、在线用户、服务器状态监控
- [x] **数据维护**: 数据表信息查看与维护

### 任务调度
- [x] **定时任务**: Cron任务在线管理，支持启停和手动触发
- [x] **任务日志**: 执行结果与错误追踪
- [x] **消息队列**: 基于Asynq的异步任务，支持重试和失败处理

### 通用功能
- [x] **通知公告**: 系统公告发布，支持定时发布和撤回
- [x] **消息管理**: 系统消息通知
- [x] **附件管理**: 文件上传管理，支持本地/对象存储
- [x] **模块管理**: 模块启用/禁用/导出/导入，详见 [模块架构](modules/README.md)

### 基础特性
- [x] **WebSocket**: 完整实现Pusher Protocol v8.3.0，支持Public/Private/Presence频道，详见 [WebSocket教程](modules/system/pkg/websocket/README.md)
- [x] **多语言**: 内置国际化，支持中文/英文切换
- [x] **多主题**: 亮色/暗色主题切换
- [x] **Docker部署**: 完整Docker镜像和Nginx配置
- [x] **链路追踪**: OpenTelemetry + Jaeger 分布式追踪
- [x] **缓存系统**: 基于Redis的全局缓存，支持自动刷新
- [x] **数据库迁移**: 版本化管理，支持迁移回滚

## 快速开始

### 环境要求
- **Go**: >= 1.23.4
- **PostgreSQL**: >= 13.0
- **Redis**: >= 5.0
- **Make**: 推荐安装

### Windows 安装 Make

项目使用 `Makefile` 管理构建命令，Windows 系统需要手动安装 make 工具：

**方式1: 通过 Chocolatey 安装（推荐）**

```powershell
# 以管理员身份运行 PowerShell，安装 Chocolatey（如未安装）
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# 安装 make
choco install make
```

**方式2: 通过 Scoop 安装**

```powershell
# 安装 Scoop（如未安装）
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression

# 安装 make
scoop install make
```

**方式3: 直接下载**

下载 [MinGW](https://www.mingw-w64.org/) 或 [GnuWin32 make](https://gnuwin32.sourceforge.net/packages/make.htm)，将 `make.exe` 所在目录添加到系统 PATH。

### 1. 克隆项目

```bash
git clone https://github.com/huagelong/devinggo.git
cd devinggo
```

### 2. 配置环境

```bash
# 复制配置文件
cp manifest/config/config.example.yaml manifest/config/config.yaml
cp hack/config.example.yaml hack/config.yaml

# 编辑配置文件，配置PostgreSQL和Redis连接
vim manifest/config/config.yaml
```

### 3. 编译运行

```bash
# 安装GoFrame CLI
make cli

# 编译项目
make build

# 初始化数据库
go run main.go migrate:up

# 启动服务
go run main.go
```

**访问地址**:
- API: http://localhost:8070
- 文档: http://localhost:8070/swagger
- 测试账号: superAdmin / admin123

### Docker部署

```bash
docker build -t devinggo:latest .
docker run -d -p 8070:8070 \
  -e DATABASE_HOST=your_db_host \
  -e DATABASE_PASSWORD=your_password \
  -e REDIS_HOST=your_redis_host \
  --name devinggo devinggo:latest
```

### 常用命令

```bash
# 代码生成相关
go run hack/generator/main.go crud:generate -m=system -t=system_user -n=用户                    # 生成CRUD代码
go run hack/generator/main.go worker:create -module=system -name=SendEmail -type=task            # 创建Worker任务
go run hack/generator/main.go module:create -name=blog                                            # 创建新模块
go run hack/generator/main.go module:export -name=blog                                            # 导出模块
go run hack/generator/main.go module:import -file=blog.v1.0.0.zip                                 # 导入模块
go run hack/generator/main.go -h                                                                # 查看所有生成器命令

# 数据库迁移相关
go run main.go migrate:up        # 执行迁移（更新数据库）
go run main.go migrate:down      # 回滚迁移
go run main.go migrate:status    # 查看迁移状态
go run main.go migrate:version   # 查看当前版本

# 服务启动相关
go run main.go                   # 启动HTTP服务（默认）
go run main.go worker            # 启动Worker服务（处理异步任务和定时任务）

# GoFrame CLI工具
make cli                         # 安装GoFrame CLI工具
gf gen dao                       # 生成DAO代码
gf gen service                   # 生成Service接口
gf build                         # 编译项目
```

## 开发指南

### 创建新模块

使用代码生成器快速创建标准化的模块结构：

```bash
# 创建名为blog的新模块
go run hack/generator/main.go module:create --name=blog
```

这将创建以下结构：
```
modules/blog/
├── module.go                 # 模块元信息
├── api/blog/                 # API定义
├── controller/blog/          # 控制器
├── logic/blog/               # 业务逻辑
├── service/                  # 服务接口
├── model/                    # 数据模型
├── router/router.go          # 路由配置
└── cmd/init.go               # 初始化命令
```

### 生成CRUD代码

基于数据表快速生成完整的后端CRUD代码：

```bash
# 生成system_user表的CRUD代码
go run hack/generator/main.go crud:generate \
  -m=system \
  -t=system_user \
  -n=用户
```

将自动生成后端代码：
- API定义 (`api/system/system_user.go`)
- 请求/响应模型 (`model/req/system_user.go`, `model/res/system_user.go`)
- 控制器 (`controller/system/system_user.go`)
- 业务逻辑 (`logic/system/system_user.go`)

### 创建Worker任务

创建异步任务或定时任务：

```bash
# 创建异步任务
go run hack/generator/main.go worker:create \
  -module=system \
  -name=SendEmail \
  -type=task \
  -desc="发送邮件任务"

# 创建定时任务
go run hack/generator/main.go worker:create \
  -module=system \
  -name=CleanLog \
  -type=cron \
  -desc="清理日志任务"

# 创建混合任务（既可异步调用，也可定时执行）
go run hack/generator/main.go worker:create \
  -module=system \
  -name=DataSync \
  -type=both \
  -desc="数据同步任务"
```

### 配置说明

主配置文件: `manifest/config/config.yaml`

```yaml
# 服务器配置
server:
  address: ":8070"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"

# 数据库配置
database:
  default:
    link: "pgsql:host=127.0.0.1 port=5432 dbname=devinggo user=postgres password=123456 sslmode=disable"

# Redis配置
redis:
  default:
    address: "127.0.0.1:6379"
    db: 0
    pass: ""

# JWT配置
jwt:
  secret: "your-secret-key"
  expire: 7200

# 任务队列配置
asynq:
  addr: "127.0.0.1:6379"
  db: 1
  concurrency: 10
```

### 开发规范

**分层架构**: API -> Controller -> Logic -> Service -> DAO

| 层级 | 职责 |
|------|------|
| API | 定义请求/响应结构，参数验证 |
| Controller | 接收HTTP请求，参数绑定，调用Logic |
| Logic | 业务逻辑实现，事务控制 |
| Service | 服务接口定义（由Logic实现） |
| DAO | 数据访问封装（自动生成） |

### API文档

启动后访问: http://localhost:8070/swagger

**统一响应格式**:
```json
{
  "code": 0,
  "message": "成功",
  "data": {},
  "timestamp": 1234567890
}
```

**认证**: JWT Token，Header: `Authorization: Bearer <token>`

**分页**: `page` (页码), `pageSize` (每页数量)

### 常见问题

| 问题 | 解决方案 |
|------|----------|
| 修改服务端口 | `manifest/config/config.yaml` 中修改 `server.address` |
| 启用调试模式 | 设置 `database.default.debug: true` |
| 数据库迁移失败 | 检查PostgreSQL连接配置，查看 `logs/` 目录日志 |
| JWT过期时间 | `manifest/config/config.yaml` 中修改 `jwt.expire`（秒） |
| Worker任务不执行 | 确保已启动Worker服务 (`go run main.go worker`)，检查Redis连接 |
| 跨域访问 | 已内置CORS中间件，可在配置中调整 |
| API返回401 | 检查Token是否过期，重新登录获取Token |

## 交流社区

- QQ群：483122520

## 🎉 关注项目

- 如果项目对您有帮助，请点右上角 "Star" 收藏项目，您的支持是我创作的动力！

## 开源协议

- **后端代码**: [Apache License 2.0](LICENSE)
- **前端 admin-ui**: [MIT License](admin-ui/LICENSE)（基于 Vue Vben Admin）

## 鸣谢

感谢以下开源项目：

1. [GoFrame](https://github.com/gogf/gf) - 企业级Go应用开发框架
2. [Asynq](https://github.com/hibiken/asynq) - Go异步任务队列
3. [PostgreSQL](https://www.postgresql.org/) - 开源关系型数据库
4. [Redis](https://redis.io/) - 高性能缓存数据库
5. [OpenTelemetry](https://opentelemetry.io/) - 可观测性框架

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=huagelong/devinggo&type=Date)](https://www.star-history.com/#huagelong/devinggo&Date)
