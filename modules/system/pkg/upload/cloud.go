// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"

	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/huagelong/goss"
)

// CloudConfig 云存储配置
type CloudConfig struct {
	Endpoint          string
	AccessKey         string
	SecretKey         string
	Region            string
	Bucket            string
	UseSsl            bool
	HostnameImmutable bool
	DelLocal          bool
}

// GetCloudConfig 获取云存储配置
func GetCloudConfig(ctx context.Context) (*CloudConfig, error) {
	getConfigValue := func(key string) (string, error) {
		return service.SettingConfig().GetConfigByKey(ctx, key, ConfigGroupUpload)
	}

	endpoint, err := getConfigValue(ConfigKeyEndpoint)
	if err != nil {
		return nil, gerror.Wrap(err, "获取endpoint配置失败")
	}

	accessKey, err := getConfigValue(ConfigKeyAccessKey)
	if err != nil {
		return nil, gerror.Wrap(err, "获取access_key配置失败")
	}

	secretKey, err := getConfigValue(ConfigKeySecretKey)
	if err != nil {
		return nil, gerror.Wrap(err, "获取secret_key配置失败")
	}

	region, err := getConfigValue(ConfigKeyRegion)
	if err != nil {
		return nil, gerror.Wrap(err, "获取region配置失败")
	}

	bucket, err := getConfigValue(ConfigKeyBucket)
	if err != nil {
		return nil, gerror.Wrap(err, "获取bucket配置失败")
	}

	useSsl, _ := getConfigValue(ConfigKeyUseSsl)
	hostnameImmutable, _ := getConfigValue(ConfigKeyHostnameImmu)
	delLocal, _ := getConfigValue(ConfigKeyDelLocal)

	return &CloudConfig{
		Endpoint:          endpoint,
		AccessKey:         accessKey,
		SecretKey:         secretKey,
		Region:            region,
		Bucket:            bucket,
		UseSsl:            useSsl == BoolTrue,
		HostnameImmutable: hostnameImmutable == BoolTrue,
		DelLocal:          delLocal == BoolTrue,
	}, nil
}

// GetCloudUpload 获取云存储客户端
func GetCloudUpload(ctx context.Context) (*goss.Goss, error) {
	cloudConfig, err := GetCloudConfig(ctx)
	if err != nil {
		return nil, err
	}

	boolTrue := true
	boolFalse := false
	config := &goss.Config{
		Endpoint:  cloudConfig.Endpoint,
		AccessKey: cloudConfig.AccessKey,
		SecretKey: cloudConfig.SecretKey,
		Region:    cloudConfig.Region,
		Bucket:    cloudConfig.Bucket,
	}

	if cloudConfig.UseSsl {
		config.UseSsl = &boolTrue
	} else {
		config.UseSsl = &boolFalse
	}

	if cloudConfig.HostnameImmutable {
		config.HostnameImmutable = &boolTrue
	}

	g.Log().Debug(ctx, "upload_config:", config)

	client, err := goss.New(goss.WithConfig(config))
	if err != nil {
		return nil, gerror.Wrap(err, "创建云存储客户端失败")
	}

	return client, nil
}

// PutFromFile 上传文件到云存储
func PutFromFile(ctx context.Context, filePath string, remotePath string) error {
	client, err := GetCloudUpload(ctx)
	if err != nil {
		return err
	}

	if err = client.PutFromFile(ctx, remotePath, filePath); err != nil {
		return gerror.Wrap(err, "上传文件到云存储失败")
	}

	// 检查是否需要删除本地文件
	cloudConfig, err := GetCloudConfig(ctx)
	if err != nil {
		return err
	}

	if cloudConfig.DelLocal {
		if err := gfile.RemoveFile(filePath); err != nil {
			return gerror.Wrap(err, "删除本地文件失败")
		}
	}

	return nil
}

// IsLocalUpload 判断是否为本地存储模式
func IsLocalUpload(ctx context.Context) bool {
	uploadMode, _ := service.SettingConfig().GetConfigByKey(ctx, ConfigKeyUploadMode, ConfigGroupUpload)
	return g.IsEmpty(uploadMode) || uploadMode == "1"
}
