# FilterAuth 数据权限过滤器使用教程

## 概述

FilterAuth 是一个基于 GoFrame 框架的数据权限过滤处理器，用于在数据库查询时自动应用用户的数据访问权限规则。它支持多种数据权限范围，确保用户只能访问其有权限查看的数据。

## 功能特性

- 🔐 **自动权限过滤**：根据用户角色自动过滤数据查询结果
- 👥 **多种权限范围**：支持 6 种不同的数据权限范围
- 🚀 **无侵入式设计**：通过 GoFrame 的 Model Handler 机制实现
- 💾 **缓存优化**：内置缓存策略，提高查询性能
- 🏢 **部门层级支持**：支持部门及子部门的权限控制

## 数据权限范围说明

| 权限范围 | 说明 | 适用场景 |
|---------|------|---------|
| 1 - 全部数据权限 | 可以查看所有数据 | 超级管理员或特定高权限角色 |
| 2 - 自定义数据权限 | 根据角色关联的部门查看数据 | 跨部门协作场景 |
| 3 - 本部门数据权限 | 只能查看所属部门的数据 | 部门内部管理 |
| 4 - 本部门及以下数据权限 | 可以查看本部门及下级部门的数据 | 上级部门管理下级部门 |
| 5 - 本人数据权限 | 只能查看自己创建的数据 | 个人数据隔离场景 |
| 6 - 部门及子部门权限 | 针对部门表的特殊权限控制 | 部门树形结构查询 |

## 核心方法

### 1. FilterAuth

自动检测数据表是否包含 `created_by` 字段，如果包含则自动应用权限过滤。

```go
func FilterAuth(m *gdb.Model) *gdb.Model
```

**使用示例：**

```go
// 自动检测并应用权限过滤
users, err := dao.SystemUser.Ctx(ctx).
    Handler(handler.FilterAuth).
    All()
```

### 2. FilterAuthWithField

指定用于权限过滤的字段名，更加灵活。

```go
func FilterAuthWithField(filterField string) func(m *gdb.Model) *gdb.Model
```

**使用示例：**

```go
// 使用自定义字段进行权限过滤
logs, err := dao.SystemOperLog.Ctx(ctx).
    Handler(handler.FilterAuthWithField("user_id")).
    All()
```

## 快速开始

### 基本使用

```go
package example

import (
    "context"
    "devinggo/internal/dao"
    "devinggo/modules/system/pkg/handler"
)

func GetUsers(ctx context.Context) (users []*entity.SystemUser, err error) {
    // 方式1：使用 FilterAuth 自动检测
    err = dao.SystemUser.Ctx(ctx).
        Handler(handler.FilterAuth).
        Scan(&users)
    return
}

func GetUserLogs(ctx context.Context, userId int64) (logs []*entity.SystemOperLog, err error) {
    // 方式2：使用 FilterAuthWithField 指定字段
    err = dao.SystemOperLog.Ctx(ctx).
        Where(dao.SystemOperLog.Columns().UserId, userId).
        Handler(handler.FilterAuthWithField("created_by")).
        Scan(&logs)
    return
}
```

### 配合分页查询使用

```go
func GetUsersWithPage(ctx context.Context, page, pageSize int) (list []*entity.SystemUser, total int, err error) {
    model := dao.SystemUser.Ctx(ctx).
        Handler(handler.FilterAuth)
    
    // 获取总数
    total, err = model.Count()
    if err != nil {
        return
    }
    
    // 获取分页数据
    err = model.Page(page, pageSize).Scan(&list)
    return
}
```

### 复杂查询场景

```go
func SearchOrders(ctx context.Context, keyword string) (orders []*entity.Order, err error) {
    err = dao.Order.Ctx(ctx).
        Where("status", 1).
        WhereLike("order_no", "%"+keyword+"%").
        Handler(handler.FilterAuth). // 权限过滤
        OrderDesc("created_at").
        Scan(&orders)
    return
}
```

## 工作原理

### 执行流程

```
1. 从上下文获取当前登录用户信息
   ↓
2. 查询用户所属的角色和部门
   ↓
3. 检查是否为超级管理员（跳过过滤）
   ↓
4. 根据角色的 data_scope 确定权限范围
   ↓
5. 构建对应的 WHERE 条件附加到查询
   ↓
6. 返回过滤后的查询结果
```

### 数据表要求

要使用 FilterAuth 自动检测功能，数据表需要包含以下字段之一：

- `created_by` - 创建者 ID（推荐）
- 或使用 `FilterAuthWithField` 指定其他字段

### 依赖的数据表

过滤器需要以下系统表支持：

| 表名 | 说明 |
|-----|------|
| system_user_role | 用户角色关联表 |
| system_user_dept | 用户部门关联表 |
| system_role | 角色表（包含 data_scope 字段） |
| system_role_dept | 角色部门关联表 |
| system_dept | 部门表（包含层级关系） |

## 配置说明

### 缓存配置

过滤器支持全局数据库缓存配置，在配置文件中设置：

```yaml
settings:
  enableGlobalDbCache: true  # 启用全局缓存
```

**缓存策略：**
- 启用缓存时：默认缓存 24 小时
- 禁用缓存时：duration 设置为 -1（不缓存）

### 超级管理员配置

在 `consts` 包中定义超级管理员角色代码：

```go
const (
    SuperRoleKey = "SuperAdmin"  // 超级管理员角色标识
)
```

超级管理员角色拥有全部数据权限，不受任何过滤限制。

## 高级用法

### 自定义部门权限处理

```go
// 权限范围 6：针对部门表的特殊处理
// 在角色的 data_scope 设置为 6 时：
// - 如果查询的是 system_dept 表，直接过滤 id 字段
// - 如果表包含 dept_id 字段，则过滤 dept_id 字段
```

### 多角色权限合并

当用户拥有多个角色时，过滤器会：
1. 遍历所有角色的 data_scope
2. 收集每个角色对应的用户 ID 列表
3. 合并所有用户 ID（取并集）
4. 最终使用 `WHERE IN` 过滤

### 获取表信息的辅助函数

```go
// 获取表名
func getTableName(m *gdb.Model) string

// 获取表字段列表
func getTableFieds(m *gdb.Model) []string

// 设置缓存选项
func setCacheOption(ctx context.Context, duration ...time.Duration) gdb.CacheOption
```

## 性能优化建议

### 1. 启用缓存

```yaml
settings:
  enableGlobalDbCache: true
```

缓存以下查询结果：
- 用户角色关联
- 用户部门关联
- 角色信息
- 部门信息

### 2. 添加索引

为相关字段添加数据库索引：

```sql
-- 用户角色表
CREATE INDEX idx_user_id ON system_user_role(user_id);
CREATE INDEX idx_role_id ON system_user_role(role_id);

-- 用户部门表
CREATE INDEX idx_user_id ON system_user_dept(user_id);
CREATE INDEX idx_dept_id ON system_user_dept(dept_id);

-- 业务表的权限字段
CREATE INDEX idx_created_by ON your_table(created_by);
```

### 3. 避免过度嵌套查询

对于复杂的部门层级结构，建议：
- 使用 `level` 字段存储部门路径
- 使用 `LIKE` 查询快速定位子部门
- 定期清理无用的部门关联数据

## 常见问题

### Q1: 为什么我的数据没有被过滤？

**可能原因：**
1. 数据表没有 `created_by` 字段
2. 当前用户是超级管理员
3. 用户角色的 data_scope 设置为 1（全部数据权限）

**解决方案：**
- 使用 `FilterAuthWithField` 指定过滤字段
- 检查角色的 data_scope 配置

### Q2: 如何对部门表进行权限过滤？

**解决方案：**
将角色的 `data_scope` 设置为 6，系统会自动处理部门表的特殊过滤逻辑。

```go
depts, err := dao.SystemDept.Ctx(ctx).
    Handler(handler.FilterAuth).
    All()
```

### Q3: 性能是否会受到影响？

**优化建议：**
- 启用全局缓存减少数据库查询
- 为关联表添加适当的索引
- 避免在循环中调用带权限过滤的查询

### Q4: 如何调试权限过滤问题？

```go
// 添加调试日志
model := dao.SystemUser.Ctx(ctx).Handler(handler.FilterAuth)
g.Log().Debug(ctx, "SQL:", model.GetSql())
// 执行查询
err := model.Scan(&users)
```

## 注意事项

⚠️ **重要提醒：**

1. **上下文传递**：必须确保请求上下文中包含用户信息（通过 `contexts.GetUser(ctx)` 获取）
2. **字段命名**：默认使用 `created_by` 作为权限过滤字段，请保持命名一致性
3. **超级管理员**：超级管理员不受任何权限限制，请谨慎分配
4. **缓存更新**：修改用户角色或部门后，缓存可能不会立即生效
5. **性能考虑**：在大数据量场景下，建议优化部门层级查询

## 最佳实践

### 1. 统一权限字段命名

```go
// 推荐：所有业务表使用统一的字段名
type BaseModel struct {
    CreatedBy int64  `json:"createdBy" dc:"创建者"`
    // ... 其他公共字段
}
```

### 2. 中间件集成

```go
// 在路由中间件中注入用户信息到上下文
func AuthMiddleware(r *ghttp.Request) {
    user := getCurrentUser(r)
    ctx := contexts.SetUser(r.Context(), user)
    r.SetCtx(ctx)
    r.Middleware.Next()
}
```

### 3. Service 层封装

```go
type IUserService interface {
    GetList(ctx context.Context, req *v1.UserListReq) ([]*entity.SystemUser, error)
}

func (s *sUser) GetList(ctx context.Context, req *v1.UserListReq) ([]*entity.SystemUser, error) {
    var users []*entity.SystemUser
    err := dao.SystemUser.Ctx(ctx).
        Handler(handler.FilterAuth). // 统一应用权限过滤
        Page(req.Page, req.PageSize).
        Scan(&users)
    return users, err
}
```

### 4. 单元测试

```go
func TestFilterAuth(t *testing.T) {
    ctx := context.Background()
    // 模拟用户信息
    ctx = contexts.SetUser(ctx, &contexts.UserInfo{
        Id:      1,
        RoleIds: []int64{2},
        DeptIds: []int64{1},
    })
    
    var users []*entity.SystemUser
    err := dao.SystemUser.Ctx(ctx).
        Handler(handler.FilterAuth).
        Scan(&users)
    
    assert.Nil(t, err)
    // 验证返回的数据是否符合权限要求
}
```

## 扩展阅读

- [GoFrame 数据库模型](https://goframe.org/pages/viewpage.action?pageId=1114686)
- [GoFrame Model Handler](https://goframe.org/pages/viewpage.action?pageId=1114695)
- [RBAC 权限模型](https://en.wikipedia.org/wiki/Role-based_access_control)

## 版本历史

- v1.0.0 - 初始版本，支持基础的数据权限过滤
- 支持 6 种数据权限范围
- 内置缓存优化机制
- 支持部门层级结构

## 贡献者

- Kai <hpuwang@gmail.com>

## 许可证

遵循项目主许可证：[LICENSE](https://github.com/huagelong/devinggo/blob/master/LICENSE)

---

📝 **反馈与建议**：如有问题或建议，欢迎提交 Issue 或 Pull Request。
