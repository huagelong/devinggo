# 自动生成目录说明

⚠️ **警告：本目录下的所有文件都是由GoFrame自动生成的，请勿手动修改！**

## 目录作用

`internal/dao/internal/` 目录包含由 `gf gen dao` 命令自动生成的数据访问对象（DAO）实现代码。

## 文件说明

### 自动生成的文件

每个数据表对应一个DAO实现文件，例如：

- `system_user.go` - system_user表的DAO实现
- `system_role.go` - system_role表的DAO实现
- `system_menu.go` - system_menu表的DAO实现

### 文件内容

每个DAO文件包含：

```go
// 表字段定义
type SystemUserColumns struct {
    Id        string // 主键ID
    Username  string // 用户名
    Password  string // 密码
    CreatedAt string // 创建时间
    UpdatedAt string // 更新时间
    // ... 其他字段
}

// DAO操作对象
var SystemUser = sSystemUser{
    Table:   "system_user",
    Group:   "default",
    Columns: SystemUserColumns{
        Id:        "id",
        Username:  "username",
        Password:  "password",
        // ... 字段映射
    },
}

// DAO结构
type sSystemUser struct {
    table   string
    group   string
    columns SystemUserColumns
}

// 数据库操作方法
func (dao *sSystemUser) DB() gdb.DB
func (dao *sSystemUser) TX(tx gdb.TX) *sSystemUser
func (dao *sSystemUser) Ctx(ctx context.Context) *sSystemUser
func (dao *sSystemUser) Table() string
func (dao *sSystemUser) Group() string
func (dao *sSystemUser) Columns() SystemUserColumns
// ...更多方法
```

## 使用方式

### 在业务代码中使用DAO

```go
import "devinggo/internal/dao"

// 查询单条记录
user, err := dao.SystemUser.Ctx(ctx).Where("id", userId).One()

// 查询列表
users, err := dao.SystemUser.Ctx(ctx).
    Where("status", 1).
    OrderDesc("created_at").
    All()

// 插入记录
_, err = dao.SystemUser.Ctx(ctx).Data(data).Insert()

// 更新记录
_, err = dao.SystemUser.Ctx(ctx).
    Where("id", userId).
    Data(data).
    Update()

// 删除记录
_, err = dao.SystemUser.Ctx(ctx).
    Where("id", userId).
    Delete()

// 使用字段常量（类型安全）
dao.SystemUser.Ctx(ctx).
    Where(dao.SystemUser.Columns().Username, "admin").
    One()
```

### 在Logic层使用

```go
// modules/system/logic/system/system_user.go
import "devinggo/internal/dao"

func (s *sSystemUser) GetById(ctx context.Context, id int64) (*entity.SystemUser, error) {
    var user *entity.SystemUser
    err := dao.SystemUser.Ctx(ctx).
        Where(dao.SystemUser.Columns().Id, id).
        Scan(&user)
    return user, err
}
```

## 代码生成

### 何时需要重新生成？

当数据库表结构发生变化时，需要重新生成DAO代码：

- 添加了新表
- 修改了表字段
- 删除了表或字段
- 修改了字段类型或属性

### 如何重新生成？

```bash
# 方法1: 使用Makefile命令
make dao

# 方法2: 使用gf命令
gf gen dao

# 方法3: 完整命令
cd /path/to/project
gf gen dao
```

### 生成配置

DAO生成器配置在项目根目录的 `hack/config.yaml` 文件中：

```yaml
gfcli:
  gen:
    dao:
    - link: "pgsql:user:password@tcp(127.0.0.1:5432)/database"
      group: "default"
      tables: ""                    # 空表示所有表
      tablesEx: ""                  # 排除的表
      removePrefix: ""              # 移除表名前缀
      descriptionTag: true          # 生成字段描述
      noModelComment: false         # 是否生成模型注释
      path: "./internal"            # 生成路径
      modelPath: "./internal/model" # 模型生成路径
      daoPath: "./internal/dao"     # DAO生成路径
```

## 目录结构

```
internal/
├── dao/                          # DAO接口目录（手动维护）
│   ├── setting_config.go         # 配置表DAO接口
│   ├── system_user.go            # 用户表DAO接口
│   └── ...
│
└── dao/internal/                 # DAO实现目录（自动生成）⚠️
    ├── README.md                 # 本说明文件
    ├── setting_config.go         # 配置表DAO实现
    ├── system_user.go            # 用户表DAO实现
    └── ...
```

## 注意事项

### ⚠️ 不要手动修改

1. **本目录所有文件都会被覆盖**
   - 每次运行 `gf gen dao` 都会重新生成所有文件
   - 手动修改的内容会丢失

2. **如需自定义DAO方法**
   - 在 `internal/dao/` 目录（上级目录）中扩展
   - 不要修改 `internal/dao/internal/` 中的文件

3. **版本控制**
   - 通常将此目录加入 `.gitignore`
   - 或者提交到代码库以便其他开发者直接使用

### ✅ 正确的扩展方式

如果需要自定义DAO方法，在上级目录扩展：

```go
// internal/dao/system_user.go（上级目录，手动维护）
package dao

import (
    "devinggo/internal/dao/internal"
)

// SystemUser是对internal.SystemUser的扩展
var SystemUser = &sSystemUser{sSystemUser: internal.SystemUser}

type sSystemUser struct {
    *internal.sSystemUser
}

// 自定义方法：根据用户名查询
func (dao *sSystemUser) GetByUsername(ctx context.Context, username string) (*entity.SystemUser, error) {
    var user *entity.SystemUser
    err := dao.Ctx(ctx).
        Where(dao.Columns().Username, username).
        Scan(&user)
    return user, err
}

// 自定义方法：查询活跃用户
func (dao *sSystemUser) GetActiveUsers(ctx context.Context) ([]*entity.SystemUser, error) {
    var users []*entity.SystemUser
    err := dao.Ctx(ctx).
        Where(dao.Columns().Status, 1).
        Where(dao.Columns().DeletedAt, nil).
        OrderDesc(dao.Columns().LastLoginAt).
        Scan(&users)
    return users, err
}
```

## 相关文档

- [GoFrame官方文档 - DAO生成](https://goframe.org/pages/viewpage.action?pageId=1114367)
- [GoFrame官方文档 - Model关联](https://goframe.org/pages/viewpage.action?pageId=1114692)
- [项目模块架构说明](../../modules/README.md)
- [代码生成器使用说明](../../hack/generator/README.md)

## 常见问题

### Q: 为什么会自动生成这些文件？

GoFrame采用代码生成的方式来提供类型安全的数据库访问。通过分析数据库表结构，自动生成：
- 表字段常量（避免硬编码字符串）
- 标准CRUD方法
- 链式查询构造器

### Q: 生成失败怎么办？

```bash
# 1. 检查数据库连接配置
cat hack/config.yaml

# 2. 测试数据库连接
go run main.go -h

# 3. 清理后重新生成
rm -rf internal/dao/internal/*
make dao

# 4. 查看详细错误信息
gf gen dao -v
```

### Q: 是否需要提交到Git？

建议提交到Git，原因：
1. 其他开发者可以直接使用，无需配置数据库
2. 保持团队代码一致性
3. 便于Code Review

但也可以加入 `.gitignore`，让每个开发者自己生成。

### Q: 如何只生成特定表的DAO？

修改 `hack/config.yaml` 中的 `tables` 配置：

```yaml
gfcli:
  gen:
    dao:
    - tables: "system_user,system_role,system_menu"  # 只生成这些表
      # 或
      tablesEx: "temp_*,test_*"                      # 排除这些表
```

## 许可证

MIT License
