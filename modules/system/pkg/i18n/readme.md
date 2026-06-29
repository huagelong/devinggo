# 多语言使用指南

## 概述

优化后的多语言方案提供了更优雅、更高效的 API 接口，支持多种使用场景。

## 核心特性

- ✅ **自动初始化**：资源路径只初始化一次，提高性能
- ✅ **优雅的 API**：提供简洁的 `T` 和 `Tf` 函数
- ✅ **灵活的参数**：支持任意类型的参数传递
- ✅ **智能语言检测**：自动解析 Accept-Language 头
- ✅ **上下文集成**：语言信息保存在上下文中

## 中间件配置

中间件已自动配置，会在每个请求中初始化多语言环境：

```go
func (s *sMiddleware) I18n(r *ghttp.Request) {
    ctx := i18n.InitI18n(r.GetCtx())
    r.SetCtx(ctx)
    r.Middleware.Next()
}
```

## 使用方式

### 1. 简单翻译（推荐）

```go
// 翻译单个键
msg := i18n.T(ctx, "user.login.success")
// 输出: "登录成功"
```

### 2. 带参数的翻译（推荐）

```go
// 翻译并格式化，支持任意类型参数
userName := "张三"
loginTime := time.Now()
count := 5

msg := i18n.Tf(ctx, "user.welcome", userName, loginTime)
// 输出: "欢迎 张三，登录时间: 2026-02-26 10:30:00"

msg := i18n.Tf(ctx, "user.unread_messages", count)
// 输出: "您有 5 条未读消息"
```

### 3. 获取当前语言

```go
lang := i18n.GetCurrentLanguage(ctx)
// 输出: "zh-CN" 或 "en"
```

## 语言检测优先级

1. **URL 参数**：`?lang=en`
2. **Accept-Language 头**：`Accept-Language: zh-CN,zh;q=0.9,en;q=0.8`
3. **默认语言**：`zh-CN`

## 在业务代码中使用

### Controller 层

```go
func (c *sController) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
    // 使用 T 函数翻译固定文本
    if req.Username == "" {
        return nil, myerror.NewErrorf(ctx, i18n.T(ctx, "error.username_required"))
    }
    
    // 使用 Tf 函数翻译带参数的文本
    user, err := service.User().GetByUsername(ctx, req.Username)
    if err != nil {
        return nil, myerror.NewErrorf(ctx, i18n.Tf(ctx, "error.user_not_found", req.Username))
    }
    
    // 返回成功消息
    return &v1.LoginRes{
        Message: i18n.T(ctx, "user.login.success"),
    }, nil
}
```

### Logic 层

```go
func (s *sLogic) ValidateUser(ctx context.Context, userId int64) error {
    user, err := s.dao.GetUser(ctx, userId)
    if err != nil {
        return myerror.NewErrorf(ctx, i18n.Tf(ctx, "error.query_failed", err.Error()))
    }
    
    if user.Status != 1 {
        return myerror.NewErrorf(ctx, i18n.T(ctx, "error.user_disabled"))
    }
    
    return nil
}
```

### 错误处理（已集成）

```go
// myerror 包已自动集成多语言
func CreateUser(ctx context.Context, req *v1.CreateUserReq) error {
    if req.Email == "" {
        // 自动翻译错误消息
        return myerror.MissingParameter(ctx, "user.email.required")
    }
    
    // 带参数的错误消息
    return myerror.InvalidParameter(ctx, "user.age.invalid", req.Age)
}
```

## 翻译文件结构

### 中文 (resource/i18n/zh-CN/common.toml)

```toml
[user]
login.success = "登录成功"
welcome = "欢迎 {0}，登录时间: {1}"
unread_messages = "您有 {0} 条未读消息"

[error]
username_required = "用户名不能为空"
user_not_found = "用户 {0} 不存在"
user_disabled = "用户已被禁用"
query_failed = "查询失败: {0}"
```

### 英文 (resource/i18n/en/common.toml)

```toml
[user]
login.success = "Login successful"
welcome = "Welcome {0}, login time: {1}"
unread_messages = "You have {0} unread messages"

[error]
username_required = "Username is required"
user_not_found = "User {0} not found"
user_disabled = "User has been disabled"
query_failed = "Query failed: {0}"
```

## API 参考

### InitI18n

初始化多语言环境，中间件中已自动调用。

```go
func InitI18n(ctx context.Context) context.Context
```

### T

简单翻译，不带参数。

```go
func T(ctx context.Context, key string) string

// 示例
msg := i18n.T(ctx, "user.login.success")
```

### Tf

翻译并格式化，支持任意类型参数。

```go
func Tf(ctx context.Context, key string, params ...interface{}) string

// 示例
msg := i18n.Tf(ctx, "user.welcome", userName, loginTime, count)
```

### GetCurrentLanguage

获取当前请求的语言设置。

```go
func GetCurrentLanguage(ctx context.Context) string

// 示例
lang := i18n.GetCurrentLanguage(ctx) // "zh-CN" or "en"
```

## 最佳实践

1. **优先使用 T 和 Tf**：新代码应使用 `T` 和 `Tf`，它们更简洁、更灵活
2. **翻译键命名**：使用点号分隔的层级结构，如 `module.action.message`
3. **参数顺序**：在翻译文件中使用 `{0}`, `{1}`, `{2}` 等占位符
4. **错误消息**：统一使用 `myerror` 包来处理错误翻译
5. **语言文件**：将相关的翻译键组织在同一个文件中

## 注意事项

- 中间件会自动处理语言初始化，无需手动调用 `InitI18n`
- 所有翻译函数都是线程安全的
- 语言信息存储在上下文中，跨服务调用时需要传递上下文
- 翻译文件支持 TOML、JSON、YAML 等多种格式
