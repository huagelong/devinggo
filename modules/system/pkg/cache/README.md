# Cache 缓存模块

## 简介

这是一个基于 Redis 的缓存管理模块，支持标签（Tag）功能，可以按标签批量清除缓存。该模块基于 GoFrame 框架开发，提供了丰富的缓存操作 API。

所有缓存操作方法都通过**可选参数 `tag ...interface{}`** 统一支持标签功能，无需单独的带标签方法。

## 文件结构

```
cache/
├── cache.go      # 核心缓存 API 和公开方法
├── tag_cache.go  # 标签缓存实现（Lua 脚本）
├── adapter.go    # gcache.Adapter 实现（GoFrame 框架集成）
└── README.md     # 使用文档
```

## 核心特性

- ✅ 基于 Redis 的分布式缓存
- ✅ 支持标签（Tag）功能，可按标签批量清除缓存
- ✅ 支持带过期时间的缓存
- ✅ 使用 Lua 脚本保证操作的原子性
- ✅ 实现 GoFrame gcache.Adapter 接口
- ✅ 支持缓存键模式匹配查询
- ✅ 统一的 API 接口，所有方法支持可选标签参数

## 初始化

在使用缓存之前，需要先初始化缓存适配器：

```go
import (
    "context"
    "github.com/huagelong/devinggo/modules/system/pkg/cache"
)

func init() {
    ctx := context.Background()
    cache.SetAdapter(ctx)
}
```

## 基本使用

### 1. 设置缓存

`Set` 方法通过可选参数支持标签，可以传入 0 个、1 个或多个标签。

```go
import (
    "context"
    "time"
    "github.com/huagelong/devinggo/modules/system/pkg/cache"
)

ctx := context.Background()

// 不带标签
err := cache.Set(ctx, "user:1", "张三", 5*time.Minute)

// 单个标签
err = cache.Set(ctx, "user:1", userData, 5*time.Minute, "user")

// 多个标签
err = cache.Set(ctx, "user:1", userData, 5*time.Minute, "user", "profile")

// 永久缓存（不过期）
err = cache.Set(ctx, "config:app", configData, 0)
```

### 2. 获取缓存

```go
// 获取缓存值
value, err := cache.Get(ctx, "user:1")
if err != nil {
    // 处理错误
}

// 转换为字符串
userName := value.String()

// 转换为其他类型
userID := value.Int()
data := value.Map()
```

### 3. 删除缓存

```go
// 删除单个缓存
lastValue, err := cache.Remove(ctx, "user:1")

// 删除多个缓存
keys := []interface{}{"user:1", "user:2", "user:3"}
lastValue, err = cache.Remove(ctx, keys)
```

### 4. 条件设置缓存

`SetIfNotExist` 支持可选标签参数，适用于分布式锁等场景。

```go
// 不带标签（分布式锁）
ok, err := cache.SetIfNotExist(ctx, "lock:order:123", "1", 10*time.Second)
if ok {
    // 成功获取锁
    defer cache.Remove(ctx, "lock:order:123")
    // 执行业务逻辑
}

// 带标签
ok, err = cache.SetIfNotExist(ctx, "user:1", userData, 5*time.Minute, "user")
```

## 标签（Tag）功能

标签功能允许您将多个缓存键关联到一个或多个标签，然后可以按标签批量清除缓存。

### 设置带标签的缓存

所有写入操作方法都支持可选标签参数：

```go
// Set - 设置缓存（单个标签）
err := cache.Set(ctx, "user:1001", userData, 5*time.Minute, "user")

// Set - 设置缓存（多个标签）
err = cache.Set(ctx, "product:2001", productData, 10*time.Minute, "product", "catalog")

// SetIfNotExist - 条件设置（带标签）
ok, err := cache.SetIfNotExist(ctx, "order:3001", orderData, 5*time.Minute, "order")

// GetOrSet - 获取或设置（带标签）
value, err := cache.GetOrSet(ctx, "user:1001", defaultUserData, 5*time.Minute, "user")

// GetOrSetFunc - 函数方式获取或设置（带标签）
value, err = cache.GetOrSetFunc(ctx, "user:1001", func(ctx context.Context) (interface{}, error) {
    return getUserFromDB(1001)
}, 5*time.Minute, "user")

// SetMap - 批量设置（所有键使用相同标签）
data := map[interface{}]interface{}{
    "user:1": userData1,
    "user:2": userData2,
}
err = cache.SetMap(ctx, data, 5*time.Minute, "user")
```

### 按标签清除缓存

```go
// 清除单个标签的所有缓存
err := cache.RemoveByTag(ctx, "user")

// 清除多个标签的所有缓存
err = cache.RemoveByTag(ctx, "user", "product", "order")

// 按数据表清除缓存
err = cache.ClearByTable(ctx, "system_user")
```

## 高级功能

### 1. 获取或设置缓存

所有 GetOrSet 方法支持可选标签参数。

```go
// 不带标签
value, err := cache.GetOrSet(ctx, "user:1", "张三", 5*time.Minute)

// 带标签
value, err = cache.GetOrSet(ctx, "user:1", userData, 5*time.Minute, "user")

// 使用函数获取值（延迟执行，仅在缓存不存在时调用）
value, err = cache.GetOrSetFunc(ctx, "user:1", func(ctx context.Context) (interface{}, error) {
    // 从数据库查询用户
    user, err := db.GetUser(1)
    return user, err
}, 5*time.Minute, "user")
```

### 2. 更新缓存

更新操作也支持可选标签参数。

```go
// 更新缓存值（保持原有过期时间）
oldValue, exist, err := cache.Update(ctx, "user:1", "李四")

// 更新缓存值并更新标签
oldValue, exist, err = cache.Update(ctx, "user:1", userData, "user", "profile")

// 更新过期时间
oldDuration, err := cache.UpdateExpire(ctx, "user:1", 10*time.Minute)

// 更新过期时间并更新标签
oldDuration, err = cache.UpdateExpire(ctx, "user:1", 10*time.Minute, "user")

// 获取剩余过期时间
duration, err := cache.GetExpire(ctx, "user:1")
```

### 3. 检查缓存是否存在

```go
exists, err := cache.Contains(ctx, "user:1")
```

### 4. 批量操作

```go
// 批量设置缓存（不带标签）
data := map[interface{}]interface{}{
    "user:1": "张三",
    "user:2": "李四",
    "user:3": "王五",
}
err := cache.SetMap(ctx, data, 5*time.Minute)

// 批量设置缓存（带标签）
err = cache.SetMap(ctx, data, 5*time.Minute, "user")
```

## 管理功能

### 获取所有缓存键

```go
keys, err := cache.GetKeys(ctx)
for _, key := range keys {
    fmt.Println(key)
}
```

### 清空所有缓存

```go
// 清空 Redis 中的所有缓存（慎用）
err := cache.ClearCacheAll(ctx)
```

### 获取 Redis 信息

```go
// 获取 Redis 服务器信息
info, err := cache.GetInfo(ctx)

// info 是一个按节分组的 map
for section, data := range info {
    fmt.Printf("Section: %s\n", section)
    for key, value := range data {
        fmt.Printf("  %s: %v\n", key, value)
    }
}
```

## API 方法列表

> **说明**：所有方法的 `tag` 参数都是可选的，通过可变参数 `tag ...interface{}` 实现。

### 基础操作

| 方法 | 说明 |
|------|------|
| `Get(ctx, key)` | 获取缓存值 |
| `Set(ctx, key, value, duration, tag...)` | 设置缓存，可选标签 |
| `SetMap(ctx, data, duration, tag...)` | 批量设置缓存，可选标签 |
| `Remove(ctx, key)` | 删除一个或多个缓存键 |
| `Contains(ctx, key)` | 检查键是否存在 |

### 条件操作

| 方法 | 说明 |
|------|------|
| `SetIfNotExist(ctx, key, value, duration, tag...)` | 键不存在时设置，可选标签 |
| `GetOrSet(ctx, key, value, duration, tag...)` | 获取或设置缓存，可选标签 |
| `GetOrSetFunc(ctx, key, f, duration, tag...)` | 使用函数获取或设置，可选标签 |

### 更新操作

| 方法 | 说明 |
|------|------|
| `Update(ctx, key, value, tag...)` | 更新缓存值（保持过期时间），可选标签 |
| `UpdateExpire(ctx, key, duration, tag...)` | 更新过期时间，可选标签 |
| `GetExpire(ctx, key)` | 获取剩余过期时间 |

### 标签操作

| 方法 | 说明 |
|------|------|
| `RemoveByTag(ctx, tags...)` | 按标签批量删除缓存 |
| `ClearByTable(ctx, table)` | 清除指定表的缓存 |

### 管理操作

| 方法 | 说明 |
|------|------|
| `GetKeys(ctx)` | 获取所有缓存键 |
| `GetInfo(ctx)` | 获取 Redis 服务器信息 |
| `ClearCacheAll(ctx)` | 清空所有缓存 |
| `GetCache()` | 获取缓存实例 |
| `SetAdapter(ctx)` | 初始化缓存适配器 |

## 配置说明

缓存模块需要在 GoFrame 配置文件中配置 Redis 连接：

```yaml
redis:
  # 默认 Redis 配置
  default:
    address: "127.0.0.1:6379"
    db: 0
    pass: ""
    
  # 缓存专用 Redis 配置（groupKey: cache）
  cache:
    address: "127.0.0.1:6379"
    db: 1
    pass: ""
```

## 实现原理

### 标签管理

该模块使用 Redis 的 SET 数据结构来管理标签关系：

- `tags:{tag}`: 存储某个标签关联的所有缓存键
- `item_tags:{key}`: 存储某个缓存键关联的所有标签

### Lua 脚本

为了保证操作的原子性，模块使用了三个 Lua 脚本：

1. **setScript**: 设置缓存并关联标签
2. **deleteScript**: 删除缓存并清理标签关系
3. **invalidateTagsScript**: 按标签批量删除缓存

## 注意事项

1. 确保 Redis 服务正常运行
2. 合理设置缓存过期时间，避免内存溢出
3. 使用 `ClearCacheAll()` 时要特别小心，它会清空整个 Redis 数据库
4. 标签功能会占用额外的 Redis 存储空间，请根据实际情况权衡使用
5. 高并发场景下建议使用 `SetIfNotExist` 来避免缓存击穿
6. 所有方法的 `tag` 参数都是可选的，根据需要传入

## 性能优化建议

1. 使用合适的缓存键命名规范，便于管理和查询
2. 为不同类型的数据设置不同的过期时间
3. 使用标签功能批量清除相关缓存，提高效率
4. 避免缓存大对象，考虑序列化优化
5. 定期清理过期的标签关系数据

## 示例代码

完整的使用示例：

```go
package main

import (
    "context"
    "fmt"
    "time"
    
    "github.com/huagelong/devinggo/modules/system/pkg/cache"
)

func main() {
    ctx := context.Background()
    
    // 初始化缓存
    cache.SetAdapter(ctx)
    
    // 设置带标签的用户缓存
    userData := map[string]interface{}{
        "id":   1001,
        "name": "张三",
        "age":  25,
    }
    err := cache.Set(ctx, "user:1001", userData, 5*time.Minute, "user", "profile")
    if err != nil {
        panic(err)
    }
    
    // 获取用户缓存
    user, err := cache.Get(ctx, "user:1001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("用户信息: %v\n", user.Map())
    
    // 使用 GetOrSet 获取或设置缓存
    product, err := cache.GetOrSet(ctx, "product:2001", map[string]interface{}{
        "id":   2001,
        "name": "商品A",
    }, 10*time.Minute, "product")
    if err != nil {
        panic(err)
    }
    fmt.Printf("商品信息: %v\n", product.Map())
    
    // 按标签批量清除缓存
    err = cache.RemoveByTag(ctx, "user", "product")
    if err != nil {
        panic(err)
    }
    
    fmt.Println("缓存已清除")
}
```

## 许可证

本模块遵循 [LICENSE](https://github.com/huagelong/devinggo/blob/master/LICENSE) 开源协议。


### 2. 获取缓存

```go
// 获取缓存值
value, err := cache.Get(ctx, "user:1")
if err != nil {
    // 处理错误
}

// 转换为字符串
userName := value.String()

// 转换为其他类型
userID := value.Int()
data := value.Map()
```

### 3. 删除缓存

```go
// 删除单个缓存
lastValue, err := cache.Remove(ctx, "user:1")

// 删除多个缓存
keys := []interface{}{"user:1", "user:2", "user:3"}
lastValue, err := cache.Remove(ctx, keys)
```

### 4. 条件设置缓存

```go
// 仅当 key 不存在时设置（不带标签）
ok, err := cache.SetIfNotExist(ctx, "lock:order:123", "1", 10*time.Second)
if ok {
    // 成功获取锁
    defer cache.Remove(ctx, "lock:order:123")
    // 执行业务逻辑
}

// 条件设置缓存（带标签）
ok, err = cache.SetIfNotExist(ctx, "user:1", userData, 5*time.Minute, "user")
```

## 标签（Tag）功能

标签功能允许您将多个缓存键关联到一个或多个标签，然后可以按标签批量清除缓存。

###直接使用公开 API 设置带标签的缓存
err := cache.Set(ctx, "user:1001", userData, 5*time.Minute, "user")

// 设置多个标签
err = cache.Set(ctx, "product:2001", productData, 10*time.Minute, "product", "catalog")

// 条件设置带标签的缓存
ok, err := cache.SetIfNotExistWithTag(ctx, "order:3001", orderData, 5*time.Minute, "order")

// GetOrSet 带标签
value, err := cache.GetOrSet(ctx, "user:1001", defaultUserData, 5*time.Minute, "user")

// GetOrSetFunc 带标签
value, err = cache.GetOrSetFunc(ctx, "user:1001", func(ctx context.Context) (interface{}, error) {
    return getUserFromDB(1001)
}, 5*time.Minute, "user")名解析）

// 手动指定标签
// 注意：需要使用内部方法，建议通过表名自动关联
```

### 按标签清除缓存

```go
// 清除所有与 "user" 标签相关的缓存
err := cache.RemoveByTag(ctx, "user")

// 清除多个标签
err = cache.RemoveByTag(ctx, "user", "product", "order")

// 按数据表清除缓存
err = cache.ClearByTable(ctx, "system_user")
```
ache.GetOrSet(ctx, "user:1", "张三", 5*time.Minute)

// 带标签的获取或设置
value, err = cache.GetOrSet(ctx, "user:1", userData, 5*time.Minute, "user")

// 使用函数获取值（延迟执行）
value, err = cache.GetOrSetFunc(ctx, "user:1", func(ctx context.Context) (interface{}, error) {
    // 从数据库查询用户
    user, err := db.GetUser(1)
    return user, err
}, 5*time.Minute, "user"GetOrSet(ctx, "user:1", "张三", 5*time.Minute)

// 使用函数获取值（延迟执行）
value, err = c.GetOrSetFunc(ctx, "user:1", func(ctx context.Context) (interface{}, error) {
    // 从数据库查询用户ache.Update(ctx, "user:1", "李四")

// 更新带标签的缓存
oldValue, exist, err = cache.Update(ctx, "user:1", userData, "user", "profile")

// 更新过期时间
oldDuration, err := cache.UpdateExpire(ctx, "user:1", 10*time.Minute)

// 更新过期时间（带标签）
oldDuration, err = cache.UpdateExpire(ctx, "user:1", 10*time.Minute, "user")

// 获取剩余过期时间
duration, err := cache

```go
// 更新缓存值（保持原有过期时间）
oldValue, exist, err := c.Update(ctx, "user:1", "李四")

// 更新过期时间
oldDuration, errache := c.UpdateExpire(ctx, "user:1", 10*time.Minute)

// 获取剩余过期时间
duration, err := c.GetExpire(ctx, "user:1")
```

### 3. 检查缓存是否存在

```go
exists, err := c.Contains(ctx, "user:1")
```

### 4. 批量操作

```go
// 批量设置缓存
data := map[interface{}]interface{}{
    "user:1": "张三",
    "user:2": "李四",
    "user:3": "王五",
}
err := cache.SetMap(ctx, data, 5*time.Minute)

// 批量设置带标签的缓存
err = cache.SetMap(ctx, data, 5*time.Minute, "user")
```

## 管理功能

### 获取所有缓存键

```go
keys, err := cache.GetKeys(ctx)
for _, key := range keys {
    fmt.Println(key)
}
```

### 清空所有缓存

```go
// 清空 Redis 中的所有缓存（慎用）
err := cache.ClearCacheAll(ctx)
```

### 获取 Redis 信息

```go
// 获取 Redis 服务器信息
info, err := cache.GetInfo(ctx)

// info 是一个按节分组的 map
for section, data := range info {
    fmt.Printf("Section: %s\n", section)
    for key, value := range data {
        fmt.Printf("  %s: %v\n", key, value)
    }
}
```

## 配置说明

缓存模块需要在 GoFrame 配置文件中配置 Redis 连接：

```yaml
redis:
  # 默认 Redis 配置
  default:
    address: "127.0.0.1:6379"
    db: 0
    pass: ""
    
  # 缓存专用 Redis 配置（groupKey: cache）
  cache:
    address: "127.0.0.1:6379"
    db: 1
    pass: ""
```

## 实现原理

### 标签管理

该模块使用 Redis 的 SET 数据结构来管理标签关系：

- `tags:{tag}`: 存储某个标签关联的所有缓存键
- `item_tags:{key}`: 存储某个缓存键关联的所有标签

### Lua 脚本

为了保证操作的原子性，模块使用了三个 Lua 脚本：

1. **setScript**: 设置缓存并关联标签
2. **deleteScript**: 删除缓存并清理标签关系
3. **invalidateTagsScript**: 按标签批量删除缓存

## 注意事项

1. 确保 Redis 服务正常运行
2. 合理设置缓存过期时间，避免内存溢出
3. 使用 `ClearCacheAll()` 时要特别小心，它会清空整个 Redis 数据库
4. 标签功能会占用额外的 Redis 存储空间，请根据实际情况权衡使用
5. 高并发场景下建议使用 `SetIfNotExist` 来避免缓存击穿

## API 方法列表

### 基础操作

| 方法 | 说明 | 支持标签 |
|------|------|---------|
| `Get(ctx, key)` | 获取缓存值 | - |
| `Set(ctx, key, value, duration, tag...)` | 设置缓存 | ✅ |
| `SetMap(ctx, data, duration, tag...)` | 批量设置缓存 | ✅ |
| `Remove(ctx, key)` | 删除缓存 | - |
| `Contains(ctx, key)` | 检查键是否存在 | - |

### 条件操作

| 方法 | 说明 | 支持标签 |
|------|------|---------|
| `SetIfNotExist(ctx, key, value, duration)` | 键不存在时设置 | ❌ |
| `SetIfNotExistWithTag(ctx, key, value, duration, tag...)` | 键不存在时设置（带标签） | ✅ |
| `GetOrSet(ctx, key, value, duration, tag...)` | 获取或设置缓存 | ✅ |
| `GetOrSetFunc(ctx, key, f, duration, tag...)` | 使用函数获取或设置 | ✅ |

### 更新操作

| 方法 | 说明 | 支持标签 |
|------|------|---------|
| `Update(ctx, key, value, tag...)` | 更新缓存值 | ✅ |
| `UpdateExpire(ctx, key, duration, tag...)` | 更新过期时间 | ✅ |
| `GetExpire(ctx, key)` | 获取剩余过期时间 | - |

### 标签操作

| 方法 | 说明 |
|------|------|
| `RemoveByTag(ctx, tags...)` | 按标签批量删除缓存 |
| `ClearByTable(ctx, table)` | 清除指定表的缓存 |

### 管理操作

| 方法 | 说明 |
|------|------|
| `GetKeys(ctx)` | 获取所有缓存键 |
| `GetInfo(ctx)` | 获取 Redis 服务器信息 |
| `ClearCacheAll(ctx)` | 清空所有缓存 |
| `GetCache()` | 获取缓存实例 |
| `SetAdapter(ctx)` | 初始化缓存适配器 |

## 性能优化建议

1. 使用合适的缓存键命名规范，便于管理和查询
2. 为不同类型的数据设置不同的过期时间
3. 使用标签功能批量清除相关缓存，提高效率
4. 避免缓存大对象，考虑序列化优化
5. 定期清理过期的标签关系数据

## 示例代码

完整的使用示例：

```go
package main

import (
    "context"
    "fmt"
    "time"
    
    "gi设置带标签的用户缓存
    userData := map[string]interface{}{
        "id":   1001,
        "name": "张三",
        "age":  25,
    }
    err := cache.Set(ctx, "user:1001", userData, 5*time.Minute, "user", "profile")
    if err != nil {
        panic(err)
    }
    
    // 获取用户缓存
    user, err := cache.Get(ctx, "user:1001")
    if err != nil {
        panic(err)
    }
    fmt.Printf("用户信息: %v\n", user.Map())
    
    // 使用 GetOrSet 获取或设置缓存
    product, err := cache.GetOrSet(ctx, "product:2001", map[string]interface{}{
        "id":   2001,
        "name": "商品A",
    }, 10*time.Minute, "product")
    if err != nil {
        panic(err)
    }
    fmt.Printf("商品信息: %v\n", product.Map())
    
    // 按标签批量清除缓存
    err = cache.RemoveByTag(ctx, "user", "product
    user, err := cache.Get(ctx, "user:1001")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("用户信息: %v\n", user.Map())
    
    // 按标签清除缓存
    err = cache.RemoveByTag(ctx, "user")
    if err != nil {
        panic(err)
    }
    
    fmt.Println("缓存已清除")
}
```

## 许可证

本模块遵循 [LICENSE](https://github.com/huagelong/devinggo/blob/master/LICENSE) 开源协议。
