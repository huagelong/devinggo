<p align="center">
  <img src="docs/logo.svg" width="120" />
</p>
<p align="center">
 <img src="https://svg.hamm.cn/badge.svg?key=License&value=Apache-2.0%20%2F%20MIT" />
  <img src="https://svg.hamm.cn/badge.svg?key=Version&value=2.0.0%20LTS" />
	<img src="https://img.shields.io/badge/goframe-2.9-green" alt="goframe">
	<img src="https://img.shields.io/badge/golang-1.23.4-blue" alt="golang">
	<img src="https://img.shields.io/badge/postgresql-13%2B-blue" alt="postgresql">
</p>
<p  align="center">
<a href="https://zread.ai/huagelong/devinggo" target="_blank"><img src="https://img.shields.io/badge/Ask_Zread-_.svg?style=flat&color=00b0aa&labelColor=000000&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIHZpZXdCb3g9IjAgMCAxNiAxNiIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTQuOTYxNTYgMS42MDAxSDIuMjQxNTZDMS44ODgxIDEuNjAwMSAxLjYwMTU2IDEuODg2NjQgMS42MDE1NiAyLjI0MDFWNC45NjAxQzEuNjAxNTYgNS4zMTM1NiAxLjg4ODEgNS42MDAxIDIuMjQxNTYgNS42MDAxSDQuOTYxNTZDNS4zMTUwMiA1LjYwMDEgNS42MDE1NiA1LjMxMzU2IDUuNjAxNTYgNC45NjAxVjIuMjQwMUM1LjYwMTU2IDEuODg2NjQgNS4zMTUwMiAxLjYwMDEgNC45NjE1NiAxLjYwMDFaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00Ljk2MTU2IDEwLjM5OTlIMi4yNDE1NkMxLjg4ODEgMTAuMzk5OSAxLjYwMTU2IDEwLjY4NjQgMS42MDE1NiAxMS4wMzk5VjEzLjc1OTlDMS42MDE1NiAxNC4xMTM0IDEuODg4MSAxNC4zOTk5IDIuMjQxNTYgMTQuMzk5OUg0Ljk2MTU2QzUuMzE1MDIgMTQuMzk5OSA1LjYwMTU2IDE0LjExMzQgNS42MDE1NiAxMy43NTk5VjExLjAzOTlDNS42MDE1NiAxMC42ODY0IDUuMzE1MDIgMTAuMzk5OSA0Ljk2MTU2IDEwLjM5OTlaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik0xMy43NTg0IDEuNjAwMUgxMS4wMzg0QzEwLjY4NSAxLjYwMDEgMTAuMzk4NCAxLjg4NjY0IDEwLjM5ODQgMi4yNDAxVjQuOTYwMUMxMC4zOTg0IDUuMzEzNTYgMTAuNjg1IDUuNjAwMSAxMS4wMzg0IDUuNjAwMUgxMy43NTg0QzE0LjExMTkgNS42MDAxIDE0LjM5ODQgNS4zMTM1NiAxNC4zOTg0IDQuOTYwMVYyLjI0MDFDMTQuMzk4NCAxLjg4NjY0IDE0LjExMTkgMS42MDAxIDEzLjc1ODQgMS42MDAxWiIgZmlsbD0iI2ZmZiIvPgo8cGF0aCBkPSJNNCAxMkwxMiA0TDQgMTJaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00IDEyTDEyIDQiIHN0cm9rZT0iI2ZmZiIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIvPgo8L3N2Zz4K&logoColor=ffffff" alt="zread"/></a>
</p>

> 老版本请看v1分支

## 简介

`DevingGo` 是基于 **GoFrame v2.9** 的企业级前后端一体化开发框架，采用模块化架构，内置权限管理、代码生成、任务调度等功能，支持 **Vibe Coding**（AI 辅助编程）一键生成前后端代码。

**核心特性：**

- 🚀 **高性能** — 基于 GoFrame，发挥 Go 高并发优势
- 🤖 **Vibe Coding** — AI 辅助编程，一键生成前后端代码
- 📦 **模块化** — 插件化模块系统，支持独立开发和热插拔
- 🛠️ **代码生成** — 一键生成 CRUD / Worker / Module 代码
- 🔐 **安全可靠** — RBAC 权限模型，JWT 认证
- 🔗 **分布式 WebSocket** — 完整实现 Pusher Protocol v8.3.0
- ⚡ **异步任务队列** — 基于 Asynq，支持任务重试和监控
- 🔍 **链路追踪** — 集成 OpenTelemetry + Jaeger

**演示地址：** http://122.152.211.127:8070 （账号 `superAdmin` / `admin123`） ｜ **文档：** [zread.ai/huagelong/devinggo](https://zread.ai/huagelong/devinggo)

## 功能特性

- **系统管理** — 用户 / 部门 / 岗位 / 角色 / 菜单 / 字典 / 参数 / 应用管理（RBAC + 数据权限）
- **开发工具** — 代码生成器（CRUD / Worker / Module）、API 管理与授权、OpenAPI 文档自动生成
- **任务调度** — Cron 定时任务、任务日志、Asynq 异步消息队列（重试与失败处理）
- **监控与日志** — 操作日志、登录日志、API 日志、系统监控（Redis / 在线用户 / 服务器状态）
- **通用功能** — 通知公告、消息管理、附件管理（本地 / 对象存储）、模块导入导出
- **基础能力** — 分布式 WebSocket、多语言、多主题、Docker 部署、链路追踪、Redis 缓存、数据库版本化迁移

## 技术栈

| 类别 | 后端 | 前端 (admin-ui) |
|------|------|------|
| 核心 | GoFrame v2.9 / Go 1.23.4 | Vue 3 + TypeScript |
| 数据库 | PostgreSQL >= 13 | — |
| 缓存/队列 | Redis >= 5 / Asynq v0.24.1 | — |
| 构建 | — | Vite |
| UI / 状态 | — | Vue Vben Admin / Tdesign Vue / Pinia |
| 其他 | JWT、golang-migrate、excelize、OpenTelemetry、雪花算法 ID | Vue Router / Axios |

## 项目结构

```
devinggo/
├── api/                    # 全局 API 定义
├── internal/               # 内部核心（cmd/controller/logic/service/dao/model/router）
├── modules/                # 业务模块（插件化，含 system 核心模块）
├── admin-ui/               # 前端管理界面（Vue3 + Vite）
├── hack/generator/         # 代码生成器
├── .codex/skills/          # Codex Skill 配置（AI 代码生成）
├── manifest/config/        # 配置文件
├── resource/               # 静态资源（i18n / migrations / public）
└── main.go                 # 程序入口
```

> **分层架构**：API → Controller → Logic → Service → DAO（其中 DAO / Entity / DO / Service 由工具自动生成）。详见 [模块架构说明](modules/README.md)。

## AI 大模型一键初始化

> **推荐使用 Codex。** 本项目内置两个 Skill：`devinggo-project-init`（环境检查与项目初始化）和 `devinggo-generator`（代码生成）。
> 在 Codex 中直接使用下列提示词即可；使用其他 AI 工具时，请将 `.codex/skills/` 下的 Skill 文件配置到对应工具，并保留 Skill 名称。

**使用codex初始化环境与服务：**

```text
使用 $devinggo-project-init 帮我检查本机环境，配置 DevingGo，初始化数据库，并启动后端、Worker 和前端服务。
```

该 Skill 会检查 Go、GoFrame CLI、Node、pnpm、PostgreSQL、Redis、Make 等依赖，缺失时提示安装，并在迁移、安装依赖、启动服务前说明影响范围。

## 快速开始

### 环境要求

Go >= 1.23.4、PostgreSQL >= 13、Redis >= 5、Make（推荐）

> Windows 需安装 make，可通过 [Chocolatey](https://chocolatey.org/)、[Scoop](https://scoop.sh/) 或下载 [MinGW](https://www.mingw-w64.org/) make 并加入 PATH。

**make 编译环境安装（Chocolatey 方式）：**

打开 cmd（管理员权限）执行：

```bat
@powershell -NoProfile -ExecutionPolicy Bypass -Command "iex ((new-object net.webclient).DownloadString('https://chocolatey.org/install.ps1'))" && SET PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin
```

或打开 PowerShell（管理员权限）执行：

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))
```

安装 make、sed、golang、nodejs：

```bash
choco install make
choco install sed
choco install golang
choco install nodejs
```

### 手动步骤

```bash
git clone https://github.com/huagelong/devinggo.git && cd devinggo

# 1. 配置文件准备
# 后端配置（生产部署）
cp manifest/config/config.example.yaml manifest/config/config.yaml   # 复制配置模板
# 修改 manifest/config/config.yaml 中的 PostgreSQL、Redis 连接信息

# 开发配置（本地开发时使用）
cp hack/config.example.yaml hack/config.yaml   # 复制开发配置模板
# 修改 hack/config.yaml 中的 PostgreSQL、Redis 连接信息

# 2. 构建与启动
make build                                                           # 构建
go run main.go migrate:up                                                # 迁移数据库
go run main.go                                                          # 启动（HTTP + Worker）
```

**访问地址：** API http://localhost:8070 ｜ 文档 http://localhost:8070/swagger ｜ 账号 `superAdmin` / `admin123`

## 部署

### Docker

```text
# 使用codex

请帮我基于当前项目构建 DevingGo Docker 镜像，并根据我的数据库和 Redis 地址生成部署步骤。

```

### 二进制部署

```bash
make build              # 产物输出到 bin/<version>/<os>_<arch>/
./devinggo unpack       # 释放配置与资源（首次执行，生成 .init.lock）
./devinggo migrate:up   # 迁移
./devinggo              # 启动
```

> `unpack` 只需首次执行一次，重新初始化请先删除 `.init.lock`。生产环境推荐使用 systemd 或 supervisor 托管进程（参考 `docs/devinggo.service`），并务必修改 `token.secretKey`、`settings.aesKey` 及数据库、Redis 密码等默认配置。

## 开发指南

### 分层架构职责

| 层级 | 职责 |
|------|------|
| API | 定义请求/响应结构，参数验证 |
| Controller | 接收请求，参数绑定，调用 Logic |
| Logic | 业务逻辑实现，事务控制 |
| Service | 服务接口（由 Logic 实现） |
| DAO | 数据访问封装（自动生成） |

### API 规范

启动后访问 http://localhost:8070/swagger

- **认证**：JWT，Header `Authorization: Bearer <token>`
- **响应**：`{ code, message, data, timestamp }`
- **分页**：`page`（页码）、`pageSize`（每页数量）

### 代码生成（Codex Skill）

创建模块、生成 CRUD、创建 Worker 任务均可通过 Skill 完成：

```text
使用 $devinggo-generator 创建名为 blog 的新模块。
使用 $devinggo-generator 为 system 模块的 system_user 表生成用户 CRUD 代码。
使用 $devinggo-generator 在 system 模块创建名为 SendEmail 的异步任务。
使用 $devinggo-generator 导出 / 导入模块。
使用 $devinggo-generator 查看支持的代码生成能力。
```

> 也可使用 `$devinggo-project-init` 执行数据库迁移、查看迁移状态、启动服务、更新 DAO/Service 接口等操作。详见 [代码生成工具](hack/generator/README.md) 和 [Skill 配置](.codex/skills/devinggo-generator/SKILL.md)。

## 配置说明

主配置文件：`manifest/config/config.yaml`

```yaml
server:
  address: ":8070"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"
database:
  default:
    link: "pgsql:host=127.0.0.1 port=5432 dbname=devinggo user=postgres password=123456 sslmode=disable"
redis:
  default:
    address: "127.0.0.1:6379"
jwt:
  secret: "your-secret-key"
  expire: 7200
asynq:
  addr: "127.0.0.1:6379"
  db: 1
  concurrency: 10
```

## 常见问题

| 问题 | 解决方案 |
|------|----------|
| 修改端口 | `manifest/config/config.yaml` → `server.address` |
| 启用调试 | `database.default.debug: true` |
| 迁移失败 | 检查 PostgreSQL 配置，查看 `logs/` 日志 |
| JWT 过期时间 | 修改 `jwt.expire`（秒） |
| Worker 不执行 | 用 `$devinggo-project-init` 检查 Worker 与 Redis 连接 |
| API 返回 401 | Token 过期，重新登录 |

## 交流与协议

- **QQ群**：483122520 ｜ 感谢 [GoFrame](https://github.com/gogf/gf)、[Asynq](https://github.com/hibiken/asynq)、[PostgreSQL](https://www.postgresql.org/)、[Redis](https://redis.io/)、[OpenTelemetry](https://opentelemetry.io/) 等开源项目
- **后端协议**：[Apache License 2.0](LICENSE)
- **前端协议**：[MIT License](admin-ui/LICENSE)（基于 Vue Vben Admin）

如果项目对您有帮助，欢迎点右上角 Star 收藏支持！

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=huagelong/devinggo&type=Date)](https://www.star-history.com/#huagelong/devinggo&Date)
