// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

// 存储模式常量
const (
	StorageModeLocal = 1 // 本地存储
	StorageModeCloud = 2 // 云存储
)

// 资源类型常量
const (
	ResourceTypeImage = "image"
	ResourceTypeText  = "text"
	ResourceTypeAudio = "audio"
	ResourceTypeVideo = "video"
	ResourceTypeZip   = "zip"
	ResourceTypeOther = "other"
)

// 配置键常量
const (
	ConfigKeyUploadMode       = "upload_mode"
	ConfigKeyUploadAllowFile  = "upload_allow_file"
	ConfigKeyUploadAllowImage = "upload_allow_image"
	ConfigKeyUploadDomain     = "upload_domain"
	ConfigKeyEndpoint         = "endpoint"
	ConfigKeyAccessKey        = "access_key"
	ConfigKeySecretKey        = "secret_key"
	ConfigKeyRegion           = "region"
	ConfigKeyBucket           = "bucket"
	ConfigKeyUseSsl           = "use_ssl"
	ConfigKeyHostnameImmu     = "hostname_immutable"
	ConfigKeyDelLocal         = "del_local"
	ConfigGroupUpload         = "upload_config"
)

// 路径常量
const (
	PathResource   = "resource"
	PathPublic     = "public"
	PathRuntime    = "runtime"
	PathChunk      = "chunk"
	PathNetwork    = "network"
	DefaultUpload  = "uploads"
	DateFormat     = "Ymd"
	ChunkExtension = ".chunk"
)

// 文件大小单位
const (
	ByteSize = 1024
)

// 布尔值字符串
const (
	BoolTrue = "true"
)
