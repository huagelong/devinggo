# Hook 使用指南

Hook 包提供了优雅的数据库钩子功能，支持自动填充字段、缓存清理和用户关联查询。

## 功能特性

- **自动填充字段**: 自动填充 `created_by` 和 `updated_by` 字段
- **缓存清理**: 在数据变更时自动清理相关缓存
- **用户关联**: 自动关联查询用户信息
- **函数选项模式**: 优雅的配置方式
- **类型安全**: 强类型参数传递

## 快速开始

### 使用默认配置

```go
import "devinggo/modules/system/pkg/hook"

// 使用默认配置（启用自动填充和缓存清理）
dao.YourModel.Ctx(ctx).Hook(hook.Default()).Insert(data)
```

### 自定义配置

```go
// 使用函数选项模式自定义配置
dao.YourModel.Ctx(ctx).Hook(hook.New(
    hook.WithAutoCreatedUpdatedBy(),
    hook.WithCacheEvict(),
    hook.WithUserRelate("created_by", "updated_by"),
)).All()
```

## 配置选项

### WithAutoCreatedUpdatedBy()

启用自动填充 `created_by` 和 `updated_by` 字段。

```go
// 插入时自动填充 created_by 和 updated_by
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithAutoCreatedUpdatedBy(),
)).Insert(g.Map{
    "username": "admin",
    // created_by 和 updated_by 会自动填充
})

// 更新时自动填充 updated_by
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithAutoCreatedUpdatedBy(),
)).Where("id", 1).Update(g.Map{
    "username": "new_name",
    // updated_by 会自动填充
})
```

### WithoutAutoCreatedUpdatedBy()

禁用自动填充功能。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithoutAutoCreatedUpdatedBy(),
)).Insert(data)
```

### WithCacheEvict()

启用缓存清理功能，在数据变更（Insert/Update/Delete）时自动清理该表的缓存。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithCacheEvict(),
)).Update(data)
```

### WithoutCacheEvict()

禁用缓存清理功能。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithoutCacheEvict(),
)).Update(data)
```

### WithUserRelate(fieldNames ...string)

启用用户关联查询，自动查询指定字段对应的用户信息。

```go
// 查询结果会自动关联 created_by 和 updated_by 对应的用户信息
result, err := dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithUserRelate("created_by", "updated_by"),
)).All()

// 查询结果中会自动添加 created_by_relate 和 updated_by_relate 字段
// created_by_relate 包含 created_by 对应的用户完整信息
// updated_by_relate 包含 updated_by 对应的用户完整信息
```

### WithoutUserRelate()

禁用用户关联查询。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithoutUserRelate(),
)).All()
```

## 预设配置

### Default()

默认配置，启用自动填充和缓存清理。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.Default()).Insert(data)

// 等同于
dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithAutoCreatedUpdatedBy(),
    hook.WithCacheEvict(),
)).Insert(data)
```

### ReadOnly(fieldNames ...string)

只读模式，只启用用户关联查询，适用于查询操作。

```go
result, err := dao.SystemLog.Ctx(ctx).Hook(
    hook.ReadOnly("created_by", "operator_id"),
).All()
```

### Minimal()

最小化配置，禁用所有功能。

```go
dao.SystemUser.Ctx(ctx).Hook(hook.Minimal()).All()
```

## 使用场景

### 场景 1：标准增删改操作

```go
// 插入数据，自动填充 created_by 和 updated_by，并清理缓存
_, err := dao.SystemUser.Ctx(ctx).Hook(hook.Default()).Insert(g.Map{
    "username": "admin",
    "nickname": "管理员",
})

// 更新数据，自动填充 updated_by，并清理缓存
_, err := dao.SystemUser.Ctx(ctx).Hook(hook.Default()).
    Where("id", 1).
    Update(g.Map{"nickname": "新昵称"})

// 删除数据，自动清理缓存
_, err := dao.SystemUser.Ctx(ctx).Hook(hook.Default()).
    Where("id", 1).
    Delete()
```

### 场景 2：查询操作并关联用户信息

```go
// 查询操作日志，并关联操作人信息
logs, err := dao.SystemOperLog.Ctx(ctx).Hook(
    hook.ReadOnly("created_by"),
).All()

// 查询结果中会包含 created_by_relate 字段
for _, log := range logs {
    fmt.Println(log["created_by_relate"]) // 用户完整信息
}
```

### 场景 3：禁用某些功能

```go
// 批量插入数据，只启用缓存清理，不自动填充字段
_, err := dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithoutAutoCreatedUpdatedBy(),
    hook.WithCacheEvict(),
)).Insert(batchData)
```

### 场景 4：完全自定义

```go
// 自定义配置
_, err := dao.SystemUser.Ctx(ctx).Hook(hook.New(
    hook.WithAutoCreatedUpdatedBy(),      // 启用自动填充
    hook.WithCacheEvict(),                 // 启用缓存清理
    hook.WithUserRelate("supervisor_id"),  // 关联主管信息
)).All()
```

## 组合使用

```go
// 在一个复杂的查询中组合使用多个配置
result, err := dao.SystemDept.Ctx(ctx).
    Hook(hook.New(
        hook.WithUserRelate("created_by", "updated_by", "leader_id"),
    )).
    Where("status", 1).
    Order("sort asc").
    All()

// 结果中会包含：
// - created_by_relate: 创建人信息
// - updated_by_relate: 更新人信息
// - leader_id_relate: 部门负责人信息
```

## 最佳实践

1. **读操作**: 使用 `ReadOnly()` 配置用户关联
   ```go
   hook.ReadOnly("created_by", "updated_by")
   ```

2. **写操作**: 使用 `Default()` 启用自动填充和缓存清理
   ```go
   hook.Default()
   ```

3. **批量操作**: 根据需要禁用某些功能以提高性能
   ```go
   hook.New(
       hook.WithoutAutoCreatedUpdatedBy(),
       hook.WithCacheEvict(),
   )
   ```

4. **特殊场景**: 使用 `Minimal()` 禁用所有功能
   ```go
   hook.Minimal()
   ```

## 注意事项

1. `created_by` 和 `updated_by` 字段必须存在于表结构中才会自动填充
2. 用户关联查询需要在 `Select` 操作中使用
3. 缓存清理功能依赖于 `cache.ClearByTable()` 方法
4. 用户ID从上下文中获取，需要确保上下文中包含用户信息

## API 参考

### 函数

- `New(opts ...Option) gdb.HookHandler` - 创建自定义配置的 Hook
- `Default() gdb.HookHandler` - 返回默认配置
- `ReadOnly(fieldNames ...string) gdb.HookHandler` - 返回只读配置
- `Minimal() gdb.HookHandler` - 返回最小配置

### 选项函数

- `WithAutoCreatedUpdatedBy() Option` - 启用自动填充
- `WithoutAutoCreatedUpdatedBy() Option` - 禁用自动填充
- `WithCacheEvict() Option` - 启用缓存清理
- `WithoutCacheEvict() Option` - 禁用缓存清理
- `WithUserRelate(fieldNames ...string) Option` - 启用用户关联
- `WithoutUserRelate() Option` - 禁用用户关联
