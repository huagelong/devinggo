# Contexts 上下文管理包

## 概述

contexts 包提供了统一的上下文管理功能，用于在请求处理过程中存储和获取各类运行时信息，包括用户身份、租户信息、权限配置、自定义数据等。

## 核心结构

### Context 结构体

```go
type Context struct {
    Module          string      // 应用模块 system|api|home|websocket
    User            *Identity   // 上下文用户信息
    Data            g.Map       // 自定义kv变量
    AppId           string      // 应用ID
    Permission      string      // 权限标识
    ExceptAuth      bool        // 是否排除权限验证
    ExceptLogin     bool        // 是否排除登录验证
    ExceptAccessLog bool        // 是否排除记录访问日志
    TenantId        int64       // 租户ID
    RequestBody     string      // 请求体内容
}
```

### Identity 结构体

```go
type Identity struct {
    Id       int64      // 用户ID
    Username string     // 用户名
    AppId    string     // 应用ID
    RoleIds  []int64    // 角色ID列表
    DeptIds  []int64    // 部门ID列表
}
```

## 功能特性

### 1. 上下文初始化与获取

#### 初始化上下文

```go
func Init(r *ghttp.Request, customCtx *model.Context)
```

在中间件或请求处理开始时初始化上下文：

```go
customCtx := &model.Context{
    Module: "system",
    User: &model.Identity{
        Id: 1001,
        Username: "admin",
    },
    Data: g.Map{},
}
contexts.Init(r, customCtx)
```

#### 获取完整上下文

```go
ctx := contexts.Get(ctx)
if ctx != nil {
    // 使用上下文信息
}
```

#### 获取模块名称

```go
module := contexts.GetModule(ctx)  // 返回 "system" 等模块名
```

### 2. 用户信息管理

#### 设置用户信息

```go
user := &model.Identity{
    Id:       1001,
    Username: "zhangsan",
    AppId:    "app001",
    RoleIds:  []int64{1, 2, 3},
    DeptIds:  []int64{10, 20},
}
contexts.SetUser(ctx, user)
```

#### 获取用户信息

```go
// 获取完整用户信息
user := contexts.GetUser(ctx)
if user != nil {
    fmt.Printf("用户ID: %d, 用户名: %s\n", user.Id, user.Username)
}

// 仅获取用户ID
userId := contexts.GetUserId(ctx)  // 未登录时返回 0
```

#### 清除用户信息

```go
contexts.DelUser(ctx)  // 清除当前用户，实际上设置 Id 为 0
```

### 3. 应用ID管理

```go
// 设置应用ID
contexts.SetAppId(ctx, "app123")

// 获取应用ID
appId := contexts.GetAppId(ctx)
```

### 4. 自定义数据存储

contexts 包支持在请求生命周期内存储自定义键值对数据：

```go
// 设置单个键值对
contexts.SetData(ctx, "order_id", "20240101001")
contexts.SetData(ctx, "trace_id", "xyz-abc-123")

// 批量设置
contexts.SetDataMap(ctx, g.Map{
    "request_source": "mobile",
    "api_version":    "v2.0",
    "extra_info":     map[string]interface{}{"key": "value"},
})

// 获取所有自定义数据
data := contexts.GetData(ctx)
orderId := data["order_id"]
```

### 5. 权限管理

```go
// 设置权限标识
contexts.SetPermission(ctx, "system:user:create")

// 获取权限标识
permission := contexts.GetPermission(ctx)

// 设置是否排除权限验证
contexts.SetExceptAuth(ctx, true)

// 检查是否排除权限验证
if contexts.GetExceptAuth(ctx) {
    // 跳过权限检查
}
```

### 6. 登录控制

```go
// 设置是否排除登录验证（如公开API）
contexts.SetExceptLogin(ctx, true)

// 检查是否需要登录
if !contexts.GetExceptLogin(ctx) {
    // 需要验证登录状态
}
```

### 7. 访问日志控制

```go
// 设置是否排除访问日志记录（如健康检查接口）
contexts.SetExceptAccessLog(ctx, true)

// 检查是否需要记录访问日志
if !contexts.GetExceptAccessLog(ctx) {
    // 记录访问日志
}
```

### 8. 多租户支持

```go
// 设置租户ID
contexts.SetTenantId(ctx, 10001)

// 获取租户ID
tenantId := contexts.GetTenantId(ctx)
```

**注意事项：**
- 当配置文件中 `tenant.enable` 为 `false` 时，`GetTenantId` 返回默认租户ID（0）
- 当多租户开启且 TenantId 为空时，`GetTenantId` 会触发 panic

### 9. 请求信息获取

#### 获取请求体

```go
// 设置请求体内容（通常在中间件中自动设置）
contexts.SetRequestBody(ctx, `{"name":"test","value":123}`)

// 获取请求体内容，未设置时返回 "{}"
body := contexts.GetRequestBody(ctx)
```

#### 获取请求耗时

```go
// 获取从请求进入到当前时刻的耗时（毫秒）
takeUpTime := contexts.GetTakeUpTime(ctx)
fmt.Printf("请求耗时: %d ms\n", takeUpTime)
```

## 使用示例

### 示例1：在中间件中初始化上下文

```go
func AuthMiddleware(r *ghttp.Request) {
    // 创建上下文
    customCtx := &model.Context{
        Module: "system",
        User:   &model.Identity{},
        Data:   g.Map{},
    }
    
    // 从token解析用户信息
    if token := r.GetHeader("Authorization"); token != "" {
        user := parseToken(token)
        customCtx.User = user
        customCtx.TenantId = getTenantIdFromUser(user)
    }
    
    // 初始化上下文
    contexts.Init(r, customCtx)
    
    r.Middleware.Next()
}
```

### 示例2：在业务逻辑中使用上下文

```go
func (s *sUserService) CreateUser(ctx context.Context, req *model.CreateUserReq) error {
    // 获取当前操作用户ID
    operatorId := contexts.GetUserId(ctx)
    
    // 获取租户ID
    tenantId := contexts.GetTenantId(ctx)
    
    // 存储自定义数据
    contexts.SetData(ctx, "operation", "create_user")
    
    // 执行业务逻辑
    err := dao.User.Insert(ctx, &entity.User{
        Username:  req.Username,
        TenantId:  tenantId,
        CreatedBy: operatorId,
    })
    
    return err
}
```

### 示例3：权限验证

```go
func (s *sAuthService) CheckPermission(ctx context.Context, permission string) bool {
    // 如果排除权限验证，直接返回true
    if contexts.GetExceptAuth(ctx) {
        return true
    }
    
    // 获取用户信息
    user := contexts.GetUser(ctx)
    if user == nil {
        return false
    }
    
    // 检查用户角色权限
    return s.hasPermission(ctx, user.RoleIds, permission)
}
```

### 示例4：日志记录

```go
func AccessLogMiddleware(r *ghttp.Request) {
    r.Middleware.Next()
    
    ctx := r.GetCtx()
    
    // 如果排除访问日志，则不记录
    if contexts.GetExceptAccessLog(ctx) {
        return
    }
    
    // 记录访问日志
    log := &entity.ApiLog{
        UserId:      contexts.GetUserId(ctx),
        Module:      contexts.GetModule(ctx),
        Permission:  contexts.GetPermission(ctx),
        RequestBody: contexts.GetRequestBody(ctx),
        TakeUpTime:  contexts.GetTakeUpTime(ctx),
        TenantId:    contexts.GetTenantId(ctx),
    }
    
    dao.ApiLog.Insert(ctx, log)
}
```

### 示例5：公共接口配置

```go
func (c *cUserController) PublicInfo(ctx context.Context, req *v1.PublicInfoReq) (res *v1.PublicInfoRes, err error) {
    // 设置为公共接口，无需登录
    contexts.SetExceptLogin(ctx, true)
    contexts.SetExceptAuth(ctx, true)
    
    // 处理业务逻辑
    info := service.User().GetPublicInfo(ctx, req.UserId)
    res = &v1.PublicInfoRes{Info: info}
    return
}
```

## 设计模式

### 单例模式

contexts 包使用包级单例 `std` 避免重复创建对象，提高性能：

```go
var std = &sContexts{}

func New() *sContexts { return std }
```

### 包级函数

提供了包级函数，可直接调用，无需显式创建实例：

```go
// 推荐方式
userId := contexts.GetUserId(ctx)

// 也可以使用（向后兼容）
userId := contexts.New().GetUserId(ctx)
```

### 泛型优化

使用泛型函数 `getField` 统一字段读取逻辑，减少代码重复：

```go
func getField[T any](s *sContexts, ctx context.Context, getter func(*model.Context) T) T {
    c := s.Get(ctx)
    if c == nil {
        var zero T
        return zero
    }
    return getter(c)
}
```

## 注意事项

1. **上下文必须初始化**：在使用其他函数之前，必须先调用 `Init` 初始化上下文
2. **线程安全**：contexts 基于 Go 的 context.Context，在同一个请求生命周期内是安全的
3. **空值处理**：当上下文未初始化时，Get 方法会打印警告日志并返回零值
4. **租户ID处理**：开启多租户后，必须确保 TenantId 被正确设置，否则会触发 panic
5. **请求体内容**：`GetRequestBody` 在未设置时返回 `"{}"`，确保返回有效JSON

## 依赖包

- `github.com/gogf/gf/v2/frame/g` - GoFrame 基础框架
- `github.com/gogf/gf/v2/net/ghttp` - HTTP 服务
- `github.com/gogf/gf/v2/os/gtime` - 时间处理
- `github.com/gogf/gf/v2/util/gconv` - 类型转换

## 相关链接

- [GoFrame 官方文档](https://goframe.org)
- [项目地址](https://github.com/huagelong/devinggo)

## 版权信息

```
@Link  https://github.com/huagelong/devinggo
@Copyright  Copyright (c) 2024 devinggo
@Author  Kai <hpuwang@gmail.com>
@License  https://github.com/huagelong/devinggo/blob/master/LICENSE
```
