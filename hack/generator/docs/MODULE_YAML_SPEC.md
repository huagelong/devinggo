# .module.yaml 配置文件规范

## 概述

`.module.yaml` 是 DevingGo 模块的标准化配置文件，提供比传统 `module.json` 更强大的功能。系统同时支持两种格式以保持向后兼容。

## 基本结构

```yaml
name: mymodule              # 模块名称（必填）
version: 1.0.0              # 版本号（必填） 
author: devinggo            # 作者（必填）
license: MIT                # 许可证
description: 模块描述       # 描述信息
homepage: https://...       # 主页URL
goVersion: 1.23+            # Go版本要求
tags: [api, web]            # 标签
keywords: [module, plugin]  # 关键词
```

## 依赖管理

```yaml
dependencies:                    # Go模块依赖
  github.com/gin-gonic/gin: v1.9.0
  
modules:                         # DevingGo模块依赖
  - system
  - api
```

## 文件管理

```yaml
files:
  go:                           # Go源码文件/目录
    - modules/mymodule
    - modules/bootstrap/worker/mymodule.go
  sql:                          # SQL迁移文件
    - resource/migrations/xxx.up.sql
    - resource/migrations/xxx.down.sql
  static:                       # 静态资源文件
    - resource/public/mymodule
  config:                       # 配置文件
    - config/mymodule.yaml
  other:                        # 其他文件
    - docs/README.md
  exclude:                      # 排除文件（支持通配符）
    - "**/*.test.go"
    - "**/temp/*"
```

## 配置文件合并

```yaml
configMerge:
  enabled: true
  files:
    - source: config/mymodule.yaml     # 模块内的配置文件
      target: manifest/config/config.yaml  # 项目的配置文件
      strategy: merge                  # 合并策略: merge, replace, skip
      keys:                            # 需要合并的键（仅merge策略）
        - database
        - redis
      variables:                       # 变量替换
        dbHost: "{{.DatabaseHost}}"
        dbPort: "{{.DatabasePort}}"
```

### 合并策略

- **merge**: 智能合并，保留目标文件的其他配置
- **replace**: 完全替换目标文件
- **skip**: 如果目标文件存在则跳过

## 静态资源部署

```yaml
staticDeploy:
  enabled: true
  rules:
    - source: resource/public/mymodule    # 模块内资源
      target: resource/public             # 项目资源目录
      method: copy                        # 部署方式: copy, symlink, merge
      overwrite: false                    # 是否覆盖已存在文件
```

### 部署方式

- **copy**: 复制文件
- **symlink**: 创建符号链接（节省空间）
- **merge**: 合并目录（智能处理冲突）

## 生命周期钩子

```yaml
hooks:
  preInstall:                           # 安装前执行
    - name: check-dependencies
      command: go mod download
      workDir: ./
      env:
        GO111MODULE: "on"
      ignoreError: false
      
  postInstall:                          # 安装后执行
    - name: run-migrations
      command: go run main.go migrate:up
      
  preUninstall:                         # 卸载前执行
    - name: backup-data
      command: ./scripts/backup.sh
      
  postUninstall:                        # 卸载后执行
    - name: cleanup
      command: rm -rf /tmp/mymodule
      
  preUpgrade:                           # 升级前执行
    - name: backup-config
      command: cp config.yaml config.yaml.bak
      
  postUpgrade:                          # 升级后执行
    - name: restart-service
      command: systemctl restart myapp
```

## 模板变量

```yaml
variables:
  moduleName: mymodule
  moduleNameCap: Mymodule
  dbPrefix: my_
  apiVersion: v1
```

### 使用变量

在配置文件、代码模板等文件中使用：

```go
// 使用 {{.moduleName}} 格式
package {{.moduleName}}

// 或使用 ${moduleName} 格式
const ModuleName = "${moduleName}"
```

## 安全配置

```yaml
security:
  signature:                            # 数字签名
    enabled: true
    publicKey: "-----BEGIN PUBLIC KEY-----..."
    algorithm: RSA                      # RSA, ECDSA
    
  permissions:                          # 权限要求
    fileSystem: true                    # 需要文件系统访问
    network: true                       # 需要网络访问
    database: true                      # 需要数据库访问
    admin: false                        # 需要管理员权限
    
  sensitiveFiles:                       # 敏感文件（导出时替换为变量）
    - config/database.yaml
    - .env
```

## 完整示例

```yaml
name: blog
version: 1.0.0
author: devinggo
license: MIT
description: 博客模块 - 提供文章管理功能
homepage: https://devinggo.com/modules/blog
goVersion: 1.23+
tags: [blog, cms, content]
keywords: [article, post, comment]

dependencies:
  github.com/russross/blackfriday/v2: v2.1.0

modules:
  - system
  - api

files:
  go:
    - modules/blog
    - modules/bootstrap/worker/blog.go
  sql:
    - resource/migrations/20260304_blog_module.up.sql
    - resource/migrations/20260304_blog_module.down.sql
  static:
    - resource/public/blog
  config:
    - config/blog.yaml
  exclude:
    - "**/*.test.go"
    - "**/.git"

configMerge:
  enabled: true
  files:
    - source: config/blog.yaml
      target: manifest/config/config.yaml
      strategy: merge
      keys:
        - blog
      variables:
        uploadPath: "{{.UploadPath}}"

staticDeploy:
  enabled: true
  rules:
    - source: resource/public/blog
      target: resource/public
      method: copy
      overwrite: false

hooks:
  postInstall:
    - name: init-database
      command: go run main.go migrate:up
    - name: seed-data
      command: go run main.go blog:seed

variables:
  moduleName: blog
  moduleNameCap: Blog
  tablePrefix: blog_

security:
  signature:
    enabled: false
  permissions:
    fileSystem: true
    database: true
    network: false
    admin: false
  sensitiveFiles:
    - config/blog.yaml
```

## 配置文件迁移

### 从 module.json 迁移到 .module.yaml

使用内置工具迁移：

```bash
go run hack/generator/main.go module:migrate -name mymodule
```

### 手动创建

使用 module:create 命令自动生成：

```bash
go run hack/generator/main.go module:create -name mymodule
```

会同时生成 `.module.yaml` 和 `module.json` 两种格式。

## 最佳实践

1. **版本管理**: 使用语义化版本号（major.minor.patch）
2. **依赖声明**: 明确声明所有外部依赖
3. **文件列表**: 尽可能详细列出所有文件，避免遗漏
4. **钩子使用**: 合理使用钩子处理复杂的安装/升级逻辑
5. **安全配置**: 敏感信息使用变量替换，启用签名验证
6. **向后兼容**: 暂时保留 `module.json` 以兼容旧系统

## 相关命令

```bash
# 创建模块（自动生成配置）
go run hack/generator/main.go module:create -name mymodule

# 验证配置
go run hack/generator/main.go module:validate -name mymodule

# 导出模块包
go run hack/generator/main.go module:export -name mymodule

# 导入模块包
go run hack/generator/main.go module:import -file mymodule.v1.0.0.zip

# 列出所有模块
go run hack/generator/main.go module:list

# 克隆模块
go run hack/generator/main.go module:clone -source blog -target news
```

## 技术参考

- 配置结构定义: `hack/generator/internal/config/module_config.go`
- 配置解析器: `hack/generator/internal/config/parser.go`
- 变量替换: `hack/generator/internal/config/variable_replacer.go`
- 模块生成器: `hack/generator/internal/generator/module_create.go`
