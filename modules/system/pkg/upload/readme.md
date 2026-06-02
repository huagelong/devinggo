# Upload 包使用教程

## 简介

Upload 包是 DevingGo 框架中的文件上传工具包，提供了完整的文件上传解决方案，支持本地存储和云存储（S3兼容），包含普通文件上传、分片上传和网络图片保存等功能。

## 主要功能

- ✅ 普通文件上传
- ✅ 分片上传（支持大文件）
- ✅ 网络图片保存
- ✅ 本地存储和云存储
- ✅ 文件类型校验
- ✅ MD5 校验
- ✅ 自动资源类型识别
- ✅ 文件大小格式化
- ✅ 链式调用API
- ✅ 类型安全的常量定义

## 文件说明

| 文件名 | 说明 |
|--------|------|
| constants.go | 常量定义（存储模式、资源类型等） |
| validator.go | 统一的文件验证器 |
| uploader.go | 支持链式调用的上传器（核心） |
| upload.go | 工具函数（MD5计算、路径管理等） |
| chunk_upload.go | 分片合并辅助功能 |
| cloud.go | 云存储配置和操作 |
| save_network_image.go | 文件信息获取辅助功能 |

## 快速开始

### 1. 普通文件上传

```go
package main

import (
    "context"
    "devinggo/modules/system/pkg/upload"
    "github.com/gogf/gf/v2/net/ghttp"
)

// 最简单的方式（使用默认配置）
func uploadFileSimple(ctx context.Context, file *ghttp.UploadFile) {
    result, err := upload.NewUploader(ctx).UploadFile(file)
    if err != nil {
        return
    }
    println("文件URL:", result.Url)
}

// 使用链式调用的优雅方式
func uploadFile(ctx context.Context, file *ghttp.UploadFile) {
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().      // 使用本地存储
        SetRandomName(true).    // 使用随机文件名
        SetValidateType(true).  // 启用类型验证
        UploadFile(file)
    
    if err != nil {
        // 处理错误
        return
    }
    
    // 使用上传结果
    println("文件URL:", result.Url)
    println("文件MD5:", result.Hash)
    println("文件大小:", result.SizeInfo)
}
```

### 2. 分片上传（大文件）

```go
func chunkUpload(ctx context.Context, file *ghttp.UploadFile, index, total int64, hash, ext, fileType, fileName string) {
    // 使用新的分片上传API
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().
        SetRandomName(true).
        UploadChunk(file, index, total, hash, ext, fileType, fileName)
    
    if err != nil {
        // 处理错误
        return
    }
    
    // 注意：只有最后一个分片上传完成后才会返回结果
    if result != nil {
        println("文件上传完成:", result.Url)
    }
}
```

### 3. 保存网络图片

```go
func saveNetworkImage(ctx context.Context, imageUrl string) {
    // 从URL保存图片
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().      // 使用本地存储
        SetRandomName(true).    // 使用随机文件名
        SaveFromURL(imageUrl)
    
    if err != nil {
        // 处理错误
        return
    }
    
    println("图片已保存:", result.Url)
    println("本地路径:", result.LocalPath)
}
```

## 新特性说明

### 🎯 链式调用API

新版本提供了优雅的链式调用API，使代码更简洁易读：

```go
// 创建上传器并配置
result, err := upload.NewUploader(ctx).
    UseCloudStorage().              // 使用云存储
    SetRandomName(true).            // 随机文件名
    SetValidateType(true).          // 验证文件类型
    SetCustomName("myfile.jpg").    // 自定义文件名（可选）
    UploadFile(file)
```

### 📋 统一验证器

提供独立的验证器，支持多种验证场景：

```go
validator := upload.NewValidator(ctx)

// 验证文件（文件+图片）
err := validator.ValidateFile(file)

// 仅验证图片
err := validator.ValidateImage(file)

// 验证扩展名
err := validator.ValidateExtension("jpg", true)
```

### 🎨 常量定义

所有魔法数字和字符串都定义为常量：

```go
// 存储模式
upload.StorageModeLocal  // 本地存储 = 1
upload.StorageModeCloud  // 云存储 = 2

// 资源类型
upload.ResourceTypeImage  // 图片
upload.ResourceTypeVideo  // 视频
upload.ResourceTypeAudio  // 音频
upload.ResourceTypeText   // 文本/文档
upload.ResourceTypeZip    // 压缩包
upload.ResourceTypeOther  // 其他
```

## 核心API说明

### NewUploader - 创建上传器

创建一个新的文件上传器实例，支持链式配置。

```go
func NewUploader(ctx context.Context) *Uploader
```

**链式配置方法：**
- `SetStorageMode(mode int)` - 设置存储模式
- `SetRandomName(random bool)` - 设置是否使用随机文件名
- `SetCustomName(name string)` - 设置自定义文件名
- `SetValidateType(validate bool)` - 设置是否验证文件类型
- `UseCloudStorage()` - 快捷设置云存储
- `UseLocalStorage()` - 快捷设置本地存储

**上传方法：**
- `UploadFile(file)` - 上传文件
- `UploadImage(file)` - 上传图片（自动验证图片类型）
- `SaveFromURL(url)` - 从URL保存文件
- `UploadChunk(...)` - 上传文件分片

**使用示例:**
```go
// 基础上传
result, err := upload.NewUploader(ctx).UploadFile(file)

// 图片上传
result, err := upload.NewUploader(ctx).
    SetRandomName(true).
    UploadImage(file)

// 云存储上传
result, err := upload.NewUploader(ctx).
    UseCloudStorage().
    SetValidateType(true).
    UploadFile(file)

// 分片上传
result, err := upload.NewUploader(ctx).
    SetRandomName(true).
    UploadChunk(file, index, total, hash, ext, fileType, fileName)
```

### NewValidator - 创建验证器

创建一个文件验证器实例，支持多种验证场景。

```go
func NewValidator(ctx context.Context) *Validator
```

**验证器方法：**
- `ValidateFile(file)` - 验证文件类型（支持文件和图片）
- `ValidateImage(file)` - 验证图片类型
- `ValidateExtension(ext, allowImage)` - 验证文件扩展名

**使用示例:**
```go
validator := upload.NewValidator(ctx)

// 验证文件类型
if err := validator.ValidateFile(file); err != nil {
    // 文件类型不允许上传
    return err
}

// 验证图片类型
if err := validator.ValidateImage(file); err != nil {
    // 图片类型不允许上传
    return err
}

// 验证扩展名
if err := validator.ValidateExtension("jpg", true); err != nil {
    // 扩展名不允许
    return err
}
```

## 工具函数

### GetResourceType - 获取资源类型

根据文件的MIME类型自动识别资源类型。

```go
func GetResourceType(mimeType string) string
```

**支持的资源类型:**
- `image`: 图片类型
- `text`: 文本/文档类型（包括PDF、Word等）
- `audio`: 音频类型
- `video`: 视频类型
- `zip`: 压缩包类型
- `other`: 其他类型

**示例:**
```go
resourceType := upload.GetResourceType("image/jpeg")  // 返回 "image"
resourceType := upload.GetResourceType("video/mp4")   // 返回 "video"
```

### CalcFileMd5 - 计算文件MD5

计算上传文件的MD5哈希值，用于文件唯一性校验。

```go
func CalcFileMd5(file *ghttp.UploadFile) (string, error)
```

**使用示例:**
```go
md5Hash, err := upload.CalcFileMd5(file)
if err != nil {
    return err
}
println("文件MD5:", md5Hash)
```

### FormatSize - 格式化文件大小

将字节数格式化为人类可读的文件大小。

```go
func FormatSize(size int64) string
```

**使用示例:**
```go
sizeStr := upload.FormatSize(1024)           // "1.00 KB"
sizeStr := upload.FormatSize(1024 * 1024)    // "1.00 MB"
sizeStr := upload.FormatSize(1024 * 1024 * 1024) // "1.00 GB"
```

## 云存储配置

### 配置说明

云存储功能支持所有S3兼容的对象存储服务（如阿里云OSS、腾讯云COS、MinIO等）。

**配置项:**
- `upload_mode`: 上传模式（1: 本地, 2: 云存储）
- `endpoint`: 服务端点
- `access_key`: 访问密钥
- `secret_key`: 密钥
- `region`: 区域
- `bucket`: 存储桶名称
- `use_ssl`: 是否使用SSL
- `hostname_immutable`: 主机名是否不可变
- `del_local`: 上传到云端后是否删除本地文件

### GetCloudUpload - 获取云存储客户端

获取配置好的云存储客户端实例。

```go
func GetCloudUpload(ctx context.Context) (*goss.Goss, error)
```

### PutFromFile - 上传文件到云存储

将本地文件上传到云存储服务。

```go
func PutFromFile(ctx context.Context, filePath string, remotePath string) error
```

### IsLocalUpload - 判断是否本地存储

检查当前配置的存储模式是否为本地存储。

```go
func IsLocalUpload(ctx context.Context) bool
```

## 完整使用示例

### 示例1: 带类型校验的文件上传

```go
package controller

import (
    "context"
    "devinggo/modules/system/pkg/upload"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/net/ghttp"
)

type FileController struct{}

func (c *FileController) Upload(ctx context.Context, r *ghttp.Request) {
    // 获取上传文件
    file := r.GetUploadFile("file")
    if file == nil {
        r.Response.WriteJson(g.Map{
            "code": 400,
            "msg":  "请选择文件",
        })
        return
    }
    
    // 使用新的链式API上传（自动验证类型）
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().      // 本地存储
        SetRandomName(true).    // 随机文件名
        SetValidateType(true).  // 启用验证
        UploadFile(file)
    
    if err != nil {
        r.Response.WriteJson(g.Map{
            "code": 500,
            "msg":  "上传失败",
            "data": err.Error(),
        })
        return
    }
    
    // 返回结果
    r.Response.WriteJson(g.Map{
        "code": 0,
        "msg":  "上传成功",
        "data": result,
    })
}
```

### 示例2: 分片上传完整流程

```go
func (c *FileController) ChunkUpload(ctx context.Context, r *ghttp.Request) {
    // 获取分片信息
    file := r.GetUploadFile("file")
    index := r.GetInt64("index")      // 当前分片序号
    total := r.GetInt64("total")      // 总分片数
    hash := r.GetString("hash")       // 文件唯一标识
    ext := r.GetString("ext")         // 文件扩展名
    fileName := r.GetString("name")   // 原始文件名
    fileType := r.GetString("type")   // MIME类型
    
    // 执行分片上传
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().
        SetRandomName(true).
        SetValidateType(true).
        UploadChunk(file, index, total, hash, ext, fileType, fileName)
    
    if err != nil {
        r.Response.WriteJson(g.Map{
            "code": 500,
            "msg":  "上传失败",
        })
        return
    }
    
    // 判断是否为最后一个分片
    if result != nil {
        // 所有分片上传完成
        r.Response.WriteJson(g.Map{
            "code": 0,
            "msg":  "上传完成",
            "data": result,
        })
    } else {
        // 分片上传成功，等待下一个分片
        r.Response.WriteJson(g.Map{
            "code": 0,
            "msg":  "分片上传成功",
        })
    }
}
```

### 示例3: 保存网络图片

```go
func (c *FileController) SaveNetworkImage(ctx context.Context, r *ghttp.Request) {
    // 获取图片URL
    imageUrl := r.GetString("url")
    if imageUrl == "" {
        r.Response.WriteJson(g.Map{
            "code": 400,
            "msg":  "请提供图片URL",
        })
        return
    }
    
    // 保存网络图片
    result, err := upload.NewUploader(ctx).
        UseLocalStorage().
        SetRandomName(true).
        SetValidateType(false).
        SaveFromURL(imageUrl)
    
    if err != nil {
        r.Response.WriteJson(g.Map{
            "code": 500,
            "msg":  "保存失败",
            "data": err.Error(),
        })
        return
    }
    
    // 返回结果
    r.Response.WriteJson(g.Map{
        "code": 0,
        "msg":  "保存成功",
        "data": result,
    })
}
```

### 示例4: 使用云存储

```go
func (c *FileController) UploadToCloud(ctx context.Context, r *ghttp.Request) {
    file := r.GetUploadFile("file")
    
    // 使用新API上传到云存储
    result, err := upload.NewUploader(ctx).
        UseCloudStorage().      // 使用云存储
        SetRandomName(true).    // 随机文件名
        UploadFile(file)
    
    if err != nil {
        r.Response.WriteJson(g.Map{
            "code": 500,
            "msg":  "上传失败",
        })
        return
    }
    
    r.Response.WriteJson(g.Map{
        "code": 0,
        "msg":  "上传成功",
        "data": result,
    })
}
```

## 配置文件示例

在系统配置表中需要配置以下上传相关参数：

```yaml
upload_config:
  # 上传模式: 1-本地存储, 2-云存储
  upload_mode: "1"
  
  # 允许上传的文件类型
  upload_allow_file: "doc,docx,xls,xlsx,ppt,pptx,pdf,zip,rar,txt"
  
  # 允许上传的图片类型
  upload_allow_image: "jpg,jpeg,png,gif,bmp,webp"
  
  # 云存储配置（当upload_mode为2时生效）
  endpoint: "oss-cn-hangzhou.aliyuncs.com"
  access_key: "your_access_key"
  secret_key: "your_secret_key"
  region: "cn-hangzhou"
  bucket: "your-bucket-name"
  use_ssl: "true"
  hostname_immutable: "false"
  
  # 上传到云端后是否删除本地文件
  del_local: "false"
```

## 注意事项

1. **文件大小限制**: 需要在服务器和框架配置中设置合适的上传大小限制
2. **分片大小**: 建议分片大小为 5MB - 10MB，根据网络环境调整
3. **文件类型校验**: 使用 `SetValidateType(true)` 启用文件类型校验
4. **MD5校验**: 可以使用MD5值进行文件去重和完整性校验
5. **云存储**: 使用云存储时，需要确保配置正确且有相应的权限
6. **路径安全**: 所有文件路径都经过处理，防止路径遍历攻击
7. **并发上传**: 分片上传支持并发，但需要注意服务器负载
8. **常量使用**: 使用预定义常量（如 `StorageModeLocal`）代替魔法数字

## 常见问题

### Q: 如何支持更大的文件？
A: 使用分片上传功能 `UploadChunk()`，将大文件分割成多个小分片上传。

### Q: 如何防止重复上传？
A: 可以通过文件的MD5值进行判断，上传前先检查系统中是否已存在相同MD5的文件。

### Q: 云存储上传失败如何处理？
A: 文件会先保存到本地，如果云存储上传失败，可以根据 `del_local` 配置决定是否保留本地文件。

### Q: 如何自定义文件存储路径？
A: 可以修改配置文件中的 `upload.dir` 配置项，或直接调用相关路径函数。

### Q: 支持哪些云存储服务？
A: 支持所有S3兼容的对象存储服务，包括阿里云OSS、腾讯云COS、AWS S3、MinIO等。

## 技术支持

如有问题，请访问：
- GitHub: https://github.com/huagelong/devinggo
- 文档: 查看项目文档获取更多信息

## 许可证

本项目遵循 MIT 许可证。详见 LICENSE 文件。
