---
name: devinggo-project-init
description: DevingGo 项目初始化与本地开发环境配置流程。用于用户询问项目初始化、开发环境搭建、依赖检查、配置文件生成、数据库迁移、本地启动后端/前端/Worker 服务、或希望 AI 大模型一步步把 DevingGo 跑起来时。适用于基于本项目 CLI、Makefile、manifest/config、hack/config、admin-ui 前端工作区的初始化排障与启动验证。
---

# DevingGo 项目初始化

## 核心原则

使用这个 skill 时，先检查再行动，优先复用用户本机已经安装好的工具。只有在缺少依赖、版本不满足、配置缺失或服务不可达时，才提出安装、升级或修改建议。

所有输出使用中文。不要创建临时文件或总结文档。不要覆盖用户已有配置，除非用户明确确认。不要执行破坏性数据库操作，例如 `migrate:down`、`migrate:force`、删除数据库、清空表或重建卷，除非用户明确要求并理解后果。

## 项目事实

- 后端入口：`go run main.go`，默认启动 `all`，即 HTTP + Worker。
- 可单独启动：`go run main.go http`、`go run main.go worker`。
- 数据库迁移：`go run main.go migrate:up`，迁移文件在 `resource/migrations/`。
- 当前代码注册的命令不包含 `install`。如果文档提到 `go run main.go install` 或 `make install`，先用 `go run main.go help` 验证，不要盲目执行。
- 后端默认端口来自 `manifest/config/config.yaml` 的 `server.address`，常见值为 `:8070`。
- 前端在 `admin-ui/`，开发命令是 `pnpm dev:backend`，默认端口来自 `admin-ui/apps/backend/.env.development`，常见值为 `5999`。
- 前端开发代理在 `admin-ui/apps/backend/vite.config.ts`，`/api` 转发到 `http://localhost:8070`。
- 前端要求来自 `admin-ui/package.json`：Node `>=20.19.0`，pnpm `>=10.0.0`，声明包管理器 `pnpm@10.32.1`。
- 后端要求来自 `go.mod`：Go `1.23.4`。GoFrame 依赖为 `github.com/gogf/gf/v2 v2.9.0`。
- 必需外部服务：PostgreSQL `>=13`、Redis `>=5`。
- 配置样例：`manifest/config/config.example.yaml`、`hack/config.example.yaml`。真实配置通常是 `manifest/config/config.yaml`、`hack/config.yaml`。

## 工作流

1. 明确目标和边界
   - 先复述本次初始化目标：后端、前端、Worker、数据库迁移、还是全量本地开发环境。
   - 如果用户只问“初始化项目”，默认目标是：检查工具版本，准备配置，确认 PostgreSQL/Redis 可连接，执行迁移，启动后端和前端，验证页面和接口可访问。
   - 如果需要安装软件、修改 PATH、写入配置、执行迁移，先说明影响并取得确认。

2. 检查本机工具
   - Windows PowerShell 常用检查：
     ```powershell
     go version
     gf -v
     node -v
     pnpm -v
     make --version
     psql --version
     redis-cli --version
     ```
   - macOS/Linux 常用检查：
     ```bash
     go version
     gf -v
     node -v
     pnpm -v
     make --version
     psql --version
     redis-cli --version
     ```
   - 版本判断：Go 必须满足 `1.23.4` 或兼容的更高补丁版本；Node 必须 `>=20.19.0`；pnpm 必须 `>=10.0.0`，优先使用 `corepack prepare pnpm@10.32.1 --activate` 对齐项目声明；PostgreSQL 必须 `>=13`；Redis 必须 `>=5`。
   - 如果工具已安装且版本满足要求，不要重复安装。

3. 检查 GoFrame CLI
   - 先执行 `gf -v`。
   - 如果缺失，再说明 `make cli` 会从 GitHub 下载 GoFrame CLI，需要网络和用户确认。
   - Windows 上 `hack/hack-cli.mk` 使用 `wget/chmod/rm`，可能依赖类 Unix 工具；若失败，按 GoFrame 官方方式或手动下载 `gf`，不要强行改 Makefile。

4. 检查外部服务
   - 读取 `manifest/config/config.yaml`，确认 PostgreSQL 和 Redis 地址、端口、用户、密码、库名。
   - 如果配置文件不存在，可以建议从样例复制：
     ```powershell
     Copy-Item manifest/config/config.example.yaml manifest/config/config.yaml
     Copy-Item hack/config.example.yaml hack/config.yaml
     ```
     只在目标文件不存在时直接复制；目标文件已存在时先询问。
   - Windows 端口探测：
     ```powershell
     Test-NetConnection 127.0.0.1 -Port 5432
     Test-NetConnection 127.0.0.1 -Port 6379
     ```
   - macOS/Linux 端口探测：
     ```bash
     nc -vz 127.0.0.1 5432
     nc -vz 127.0.0.1 6379
     ```
   - 如果用户使用 Docker 启动 PostgreSQL/Redis，先检查是否已有容器和端口占用；创建新容器前必须确认容器名、端口、密码和数据卷策略。

5. 准备依赖
   - 后端依赖：
     ```bash
     go mod download
     ```
   - 前端依赖：
     ```bash
     cd admin-ui
     pnpm install
     ```
   - 如果网络失败，不要反复重试；报告失败的源、包名或错误信息，并建议配置代理或镜像。

6. 初始化数据库
   - 执行前确认当前连接的是目标开发库，不是生产库或共享库。
   - 首选命令：
     ```bash
     go run main.go migrate:up
     ```
   - 如果迁移报 dirty/version 错误，先展示当前错误和版本，不要自动 `migrate:force`。
   - 迁移完成后，可用 `go run main.go migrate:up -n 1` 之类命令仅在用户明确要求时处理增量场景。

7. 启动并验证服务
   - 后端 HTTP：
     ```bash
     go run main.go http
     ```
   - Worker：
     ```bash
     go run main.go worker
     ```
   - 或全量：
     ```bash
     go run main.go
     ```
   - 前端：
     ```bash
     cd admin-ui
     pnpm dev:backend
     ```
   - 验证：
     ```bash
     curl http://localhost:8070/swagger
     curl http://localhost:8070/api.json
     ```
     浏览器访问 `http://localhost:5999`，登录账号参考 README：`superAdmin / admin123`。

## 排障优先级

- 服务起不来：先看终端错误、端口占用、配置文件路径、数据库/Redis 连通性。
- 前端请求失败：先看 `VITE_GLOB_API_URL=/api`、Vite proxy 是否指向 `localhost:8070`、后端是否已启动。
- 登录失败：先确认迁移是否完成、数据库是否是当前配置库、Redis 是否连通、AES key 是否与 `settings.aesKey` 和 `VITE_APP_AES_KEY` 一致。
- 生成 DAO 或 CRUD 失败：使用 `$devinggo-generator`，不要把代码生成问题混入初始化流程。
- README 或旧说明与代码冲突：以当前代码、`go run main.go help`、`Makefile`、`admin-ui/package.json` 为准，并在回复中指出冲突。

## 完成标准

初始化完成时，必须给用户一个简短状态清单：

- 工具版本是否满足要求。
- `manifest/config/config.yaml` 和 `hack/config.yaml` 是否已准备好。
- PostgreSQL 和 Redis 是否可连接。
- 数据库迁移是否成功。
- 后端、Worker、前端是否启动，以及访问地址。
- 仍需用户处理的事项，例如安装缺失软件、修改密码、处理端口占用或网络下载失败。
