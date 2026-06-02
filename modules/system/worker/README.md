# Worker 使用指南

本模块提供了基于新版Worker API的完整实现，支持自动加载和注册机制。

## 架构说明

### 自动加载机制

系统使用了分层的自动加载架构：

1. **模块加载层** (`modules/bootstrap/worker/system.go`)
   - 通过 `import _ "devinggo/modules/system/worker"` 自动加载system模块的worker
   - 支持多模块扩展，新模块只需添加对应的加载文件

2. **自动注册层** (`modules/system/worker/worker.go`)
   - 在包的 `init()` 函数中创建全局Manager
   - 自动注册所有Worker和Cron任务到全局Manager
   - 无需手动调用注册函数

3. **使用层** (`cmd/worker.go`)
   - 通过 `worker.GetManager()` 获取全局Manager
   - 直接使用已注册的Worker启动服务

### 优势

✅ **自动发现** - 新增Worker只需创建文件，init自动注册  
✅ **模块化** - 支持多模块各自定义Worker  
✅ **简单易用** - 启动代码只需获取Manager即可  
✅ **向后兼容** - 保持原有的自动加载机制  

## 快速开始

### 1. Worker自动启动

Worker服务会通过以下链路自动启动：

```
cmd/worker.go 
  ↓ import
modules/bootstrap/worker/system.go 
  ↓ import  
modules/system/worker (package)
  ↓ init()
自动注册所有Worker和Cron任务
```

在应用启动代码中：

```go
package main

import (
    "context"
    "devinggo/modules/system/worker"
    "log"
)

func main() {
    ctx := context.Background()
    
    // 获取全局Manager（已在init中自动注册所有Worker）
    mgr := worker.GetManager()
    
    // 启动Worker服务器（处理任务）
    go func() {
        log.Println("启动Worker服务器...")
        if err := mgr.RunServer(); err != nil {
            log.Fatalf("Worker服务器错误: %v", err)
        }
    }()
    
    // 启动Cron调度器（定时任务）
    go func() {
        log.Println("启动Cron调度器...")
        if err := mgr.RunCron(); err != nil {
            log.Fatalf("Cron调度器错误: %v", err)
        }
    }()
    
    // 应用主逻辑
    // ...
}
```

### 2. 发送任务

#### 方式一：使用便捷函数（test_task）

```go
import (
    "devinggo/modules/system/worker/task/test_task"
)

// 立即发送任务
err := test_task.Send(ctx, test_task.TestTaskData{
    Name: "测试任务",
})

// 延迟5秒发送
err := test_task.SendWithDelay(ctx, test_task.TestTaskData{
    Name: "延迟任务",
}, 5)
```

#### 方式二：使用TaskBuilder（推荐）

```go
import (
    "devinggo/modules/system/pkg/worker"
    "devinggo/modules/system/worker/consts"
    "time"
)

// 立即执行
err := worker.NewTaskBuilder(ctx, consts.TEST_TASK).
    WithData(map[string]interface{}{
        "name": "测试",
    }).
    Send()

// 延迟5分钟执行
err := worker.NewTaskBuilder(ctx, consts.TEST_TASK).
    WithData(map[string]interface{}{
        "name": "测试",
    }).
    WithDelay(5 * time.Minute).
    WithQueue("critical").
    Send()

// 在指定时间执行
tomorrow := time.Now().Add(24 * time.Hour)
err := worker.NewTaskBuilder(ctx, consts.URL_CRON).
    WithData(map[string]interface{}{
        "url": "https://api.example.com/webhook",
        "method": "POST",
    }).
    WithProcessAt(tomorrow).
    Send()
```

## 已实现的Worker和任务

### Worker处理器

1. **TestWorker** (`consts.TEST_TASK`) - 测试任务处理器
2. **TestCronWorker** (`consts.TEST_CRON`) - 测试定时任务处理器
3. **CmdCronWorker** (`consts.CMD_CRON`) - 命令执行处理器
4. **UrlCronWorker** (`consts.URL_CRON`) - HTTP请求处理器

### 定时任务

1. **TestCron** - 测试定时任务
2. **CmdCron** - 命令执行定时任务
3. **UrlCron** - HTTP请求定时任务

## 任务类型说明

### 1. 测试任务（TEST_TASK）

```go
type TestTaskData struct {
    Name string `json:"name"`
}

// 发送方式
test_task.Send(ctx, test_task.TestTaskData{
    Name: "测试名称",
})
```

### 2. URL请求任务（URL_CRON）

```go
type UrlCronData struct {
    Url         string                 `json:"url"`
    Method      string                 `json:"method"`
    Headers     map[string]string      `json:"headers"`
    Params      map[string]interface{} `json:"params"`
    Timeout     int64                  `json:"timeout"`
    Retry       int                    `json:"retry"`
    Cookies     map[string]string      `json:"cookie"`
    ContentType string                 `json:"content_type"`
    Proxy       string                 `json:"proxy"`
}

// 发送方式
worker.NewTaskBuilder(ctx, consts.URL_CRON).
    WithData(map[string]interface{}{
        "url": "https://api.example.com/webhook",
        "method": "POST",
        "headers": map[string]string{
            "Authorization": "Bearer token",
        },
        "params": map[string]interface{}{
            "key": "value",
        },
    }).
    Send()
```

### 3. 命令执行任务（CMD_CRON）

```go
type CmdCronData struct {
    Cmd string `json:"cmd"`
}

// 发送方式
worker.NewTaskBuilder(ctx, consts.CMD_CRON).
    WithData(map[string]interface{}{
        "cmd": "echo 'Hello World'",
    }).
    Send()
```

## 添加新的Worker

### 方式一：函数式注册（推荐）

在 `modules/system/worker/server/` 目录下创建新文件：

```go
package server

import (
    "context"
    "devinggo/modules/system/pkg/worker"
    "devinggo/modules/system/worker/consts"
    "github.com/hibiken/asynq"
)

// 定义数据结构
type MyTaskData struct {
    Field1 string `json:"field1"`
    Field2 int    `json:"field2"`
}

// 定义执行函数
func executeMyTask(ctx context.Context, t *asynq.Task) error {
    // 解析参数
    data, err := worker.GetParameters[MyTaskData](ctx, t)
    if err != nil {
        return err
    }
    
    // 执行业务逻辑
    worker.GetLogger().Infof(ctx, "处理任务: %+v", data)
    
    return nil
}

// 自动注册
func init() {
    worker.RegisterWorkerFunc(consts.MY_TASK, executeMyTask)
}
```

### 方式二：使用Builder（进阶）

如果需要更多控制，可以使用WorkerBuilder：

```go
func init() {
    worker.NewWorkerBuilder(consts.MY_TASK).
        WithExecute(executeMyTask).
        Register()
}
```

### 发送任务

```go
worker.NewTaskBuilder(ctx, consts.MY_TASK).
    WithData(MyTaskData{
        Field1: "value1",
        Field2: 123,
    }).
    Send()
```

## 迁移说明

### 从旧的结构体方式迁移到函数式注册

#### Worker迁移

**旧代码（结构体+接口）：**
```go
type TestWorker struct {
    Type string
}

func NewTestWorker() *TestWorker {
    return &TestWorker{
        Type: consts.TEST_TASK,
    }
}

func (w *TestWorker) GetType() string {
    return w.Type
}

func (w *TestWorker) Execute(ctx context.Context, t *asynq.Task) error {
    data, err := worker.GetParameters[TestTaskData](ctx, t)
    if err != nil {
        return err
    }
    // 处理逻辑...
    return nil
}

// 在worker.go中注册
func init() {
    globalManager.RegisterWorker(server.NewTestWorker())
}
```

**新代码（函数式注册）：**
```go
// 定义执行函数
func executeTestTask(ctx context.Context, t *asynq.Task) error {
    data, err := worker.GetParameters[TestTaskData](ctx, t)
    if err != nil {
        return err
    }
    // 处理逻辑...
    return nil
}

// 在本文件的init中直接注册
func init() {
    worker.RegisterWorkerFunc(consts.TEST_TASK, executeTestTask)
}
```

#### Cron迁移

**旧代码（结构体+5个接口方法）：**
```go
type TestCron struct {
    Type string
}

func (c *TestCron) GetType() string { return consts.TEST_CRON }
func (c *TestCron) GetDescription() string { return "测试定时任务" }
func (c *TestCron) GetPayload(params string) (interface{}, error) {
    return map[string]string{"data": params}, nil
}

func init() {
    globalManager.RegisterCronTask(&TestCron{})
}
```

**新代码（函数式注册）：**
```go
func init() {
    worker.RegisterCronFunc(
        consts.TEST_CRON,
        "测试定时任务",
        func(params string) (interface{}, error) {
            return map[string]string{"data": params}, nil
        },
    )
}
```

#### 任务发送方式

**旧的任务发送方式（已简化）：**
```go
taskItem := test_task.New()
taskItem.Send(ctx, data)
```

**新的任务发送方式（推荐）：**
```go
// 方式1：使用便捷函数
test_task.Send(ctx, data)

// 方式2：使用TaskBuilder（更强大）
worker.NewTaskBuilder(ctx, consts.TEST_TASK).
    WithData(data).
    WithQueue("critical").
    WithDelay(5 * time.Minute).
    Send()
```

## 优势

✅ **函数式注册** - 使用RegisterWorkerFunc/RegisterCronFunc，代码减少90%  
✅ **集中管理** - 所有Worker和Cron任务自动注册，清晰明了  
✅ **按需加载** - 支持init函数自动注册，无需手动调用  
✅ **灵活配置** - 可以根据环境选择性注册Worker  
✅ **易于测试** - 可以为测试创建独立的Worker集合  
✅ **代码简洁** - 使用TaskBuilder大幅简化任务发送代码  
✅ **无需结构体** - 不再需要定义Worker结构体和接口方法  

## 配置

在 `config.yaml` 中配置Worker参数：

```yaml
worker:
  redis:
    address: "localhost:6379"
    pass: ""
    db: 3
  queues:
    critical: 6  # 高优先级
    default: 3   # 默认优先级
    low: 1       # 低优先级
  concurrency: 10  # 并发worker数量
  shutdownTimeout: "10s"
  location: "Asia/Shanghai"  # 时区
```

## 故障排查

### Worker未执行任务

1. 检查Worker服务器是否启动：`mgr.RunServer()`
2. 检查Redis连接配置
3. 查看日志是否有错误信息

### Cron任务未执行

1. 检查Cron调度器是否启动：`mgr.RunCron()`
2. 检查数据库中的定时任务配置
3. 确认定时任务状态为启用（status=1）

### 任务发送失败

1. 检查任务类型是否已注册
2. 检查数据结构是否匹配
3. 检查Redis连接是否正常

## 更多信息

- [Worker包文档](../pkg/worker/README.md)
- [快速入门](../pkg/worker/QUICKSTART.md)
- [API对比](../pkg/worker/COMPARISON.md)
- [完整示例](../pkg/worker/example/example.go)
