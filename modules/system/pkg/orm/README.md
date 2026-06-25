# ORM 包使用指南

## 概述

`orm` 包提供了一套优雅的数据库查询构建器和执行器，支持链式调用、泛型和多种查询模式。

## 核心组件

### 1. QueryBuilder - 查询构建器

链式构建数据库查询，提供流畅的 API 接口。

#### 基础用法

```go
import "devinggo/modules/system/pkg/orm"

// 创建查询构建器
query := orm.NewQuery(dao.SystemUser.Ctx(ctx))

// 链式调用构建查询
results := []entity.SystemUser{}
err := query.
    WithFields("id, username, nickname").
    WithWhere(g.Map{"status": 1}).
    WithOrder("created_at", "desc").
    WithPage(1, 20).
    ScanAll(&results)
```

#### 使用 PageListReq

```go
// 使用 PageListReq 快速构建分页查询
query := orm.NewQuery(m).WithPageListReq(req, g.Map{
    "dept_id": deptId,
})

var results []res.SystemUser
var total int
err := query.ScanAndCount(&results, &total)
```

#### 使用 ListReq

```go
// 使用 ListReq 构建列表查询（不分页）
query := orm.NewQuery(m).WithListReq(req, g.Map{
    "type": "admin",
})

var results []res.SystemUser
err := query.ScanAll(&results)
```

### 2. Executor - 查询执行器

提供泛型支持的查询执行器，自动处理类型转换。

#### 基础用法

```go
// 创建执行器
executor := orm.NewExecutor[res.SystemUser](m)

// 查询单条数据
user, err := executor.One(ctx)

// 查询所有数据
users, err := executor.All(ctx)

// 分页查询
users, total, err := executor.PagedList(ctx, 1, 20)

// 统计数量
count, err := executor.Count(ctx)

// 判断是否存在
exists, err := executor.Exists(ctx)
```

#### 结合 QueryBuilder 使用

```go
// 构建复杂查询并执行
users, total, err := orm.NewQuery(m).
    WithWhere(g.Map{"status": 1}).
    WithOrder("created_at", "desc").
    Executor[res.SystemUser]().
    PagedList(ctx, 1, 20)
```

## 迁移指南

### ~~旧代码（已删除）~~

```go
// 以下方法已被删除，不再支持
// err = orm.GetPageList(s.Model(ctx), req).ScanAndCount(&res, &total, false)
// err = orm.GetList(m, req).Scan(&res)
```

### 新代码（唯一方式）

```go
// 使用链式调用
err = orm.NewQuery(s.Model(ctx)).
    WithPageListReq(req).
    ScanAndCount(&res, &total)

// 使用泛型执行器
res, total, err := orm.NewQuery(m).
    WithPageListReq(req).
    Executor[res.SystemUser]().
    AllWithCount(ctx)
```

## 高级用法

### 组合查询条件

```go
query := orm.NewQuery(m)

// 动态添加条件
if !g.IsEmpty(username) {
    query.WithWhere(g.Map{"username": username})
}

if isRecycle {
    query.WithRecycle(true)
}

if needAuth {
    query.WithFilterAuth(true)
}

// 执行查询
results, total, err := query.
    Executor[res.SystemUser]().
    PagedList(ctx, page, pageSize)
```

### 自定义查询

```go
// 获取底层 Model 进行自定义操作
model := orm.NewQuery(m).
    WithWhere(g.Map{"status": 1}).
    Build()

// 继续使用 gdb.Model 的原生方法
model = model.LeftJoin("system_dept d", "d.id=system_user.dept_id")
model = model.Where("d.status", 1)

// 执行查询
var results []res.SystemUser
err := model.Scan(&results)
```

## API 参考

### QueryBuilder 方法

| 方法 | 说明 |
|------|------|
| `NewQuery(m *gdb.Model)` | 创建查询构建器 |
| `WithRecycle(bool)` | 设置回收站查询 |
| `WithFilterAuth(bool)` | 设置权限过滤 |
| `WithWhere(...g.Map)` | 设置查询条件 |
| `WithFields(interface{})` | 设置查询字段 |
| `WithOrder(string, ...string)` | 设置排序 |
| `WithPage(int, int)` | 设置分页 |
| `WithPageListReq(*model.PageListReq, ...g.Map)` | 使用 PageListReq 配置 |
| `WithListReq(*model.ListReq, ...g.Map)` | 使用 ListReq 配置 |
| `Build()` | 构建 Model |
| `Executor[T]()` | 返回泛型执行器 |
| `ScanOne(interface{})` | 扫描单条 |
| `ScanAll(interface{})` | 扫描所有 |
| `ScanAndCount(interface{}, *int)` | 扫描并统计 |

### Executor 方法

| 方法 | 说明 |
|------|------|
| `NewExecutor[T](m *gdb.Model)` | 创建执行器 |
| `One(ctx)` | 查询单条 |
| `All(ctx)` | 查询所有 |
| `AllWithCount(ctx)` | 查询所有并统计 |
| `PagedList(ctx, page, pageSize)` | 分页查询 |
| `Count(ctx)` | 统计数量 |
| `Exists(ctx)` | 判断存在 |

## 最佳实践

1. **优先使用泛型执行器**：类型安全，减少手动类型转换
2. **链式调用构建查询**：代码更简洁，逻辑更清晰
3. **利用 WithPageListReq/WithListReq**：快速处理标准查询请求
4. **不再支持旧方法**：GetPageList/GetList 已删除，必须使用新的链式调用

## 注意事项

- 泛型执行器需要 Go 1.18+
- 权限过滤和回收站查询需要在模型层正确配置
