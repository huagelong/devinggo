// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package upload

import (
	"context"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/huagelong/goss"
)

func GetCloudUpload(ctx context.Context) (*goss.Goss, error) {

	endpoint, err := service.SettingConfig().GetConfigByKey(ctx, "endpoint", "upload_config")
	if err != nil {
		return nil, err
	}

	accessKey, err := service.SettingConfig().GetConfigByKey(ctx, "access_key", "upload_config")
	if err != nil {
		return nil, err
	}

	secretKey, err := service.SettingConfig().GetConfigByKey(ctx, "secret_key", "upload_config")
	if err != nil {
		return nil, err
	}

	region, err := service.SettingConfig().GetConfigByKey(ctx, "region", "upload_config")
	if err != nil {
		return nil, err
	}

	bucket, err := service.SettingConfig().GetConfigByKey(ctx, "bucket", "upload_config")
	if err != nil {
		return nil, err
	}

	useSsl, err := service.SettingConfig().GetConfigByKey(ctx, "use_ssl", "upload_config")
	if err != nil {
		return nil, err
	}

	hostnameImmutable, err := service.SettingConfig().GetConfigByKey(ctx, "hostname_immutable", "upload_config")
	if err != nil {
		return nil, err
	}
	config := &goss.Config{
		Endpoint:  endpoint,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
		Bucket:    bucket,
	}
	boolTrue := true
	boolFalse := false
	if !g.IsEmpty(useSsl) && (useSsl == "true") {
		config.UseSsl = &boolTrue
	} else {
		config.UseSsl = &boolFalse
	}
	if !g.IsEmpty(hostnameImmutable) && (hostnameImmutable == "true") {
		config.HostnameImmutable = &boolTrue
	}
	g.Log().Debug(ctx, "upload_config:", config)
	goss, err := goss.New(goss.WithConfig(config))
	if err != nil {
		return nil, err
	}
	return goss, nil
}

func PutFromFile(ctx context.Context, filePath string, remotePath string) error {
	goss, err := GetCloudUpload(ctx)
	if err != nil {
		return err
	}
	err = goss.PutFromFile(ctx, remotePath, filePath)
	if err != nil {
		return err
	}
	delLocal, err := service.SettingConfig().GetConfigByKey(ctx, "del_local", "upload_config")
	if err != nil {
		return err
	}
	if !g.IsEmpty(delLocal) && (delLocal == "true") {
		err := gfile.RemoveFile(filePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsLocalUpload(ctx context.Context) bool {
	uploadMode, _ := service.SettingConfig().GetConfigByKey(ctx, "upload_mode", "upload_config")
	//g.Log().Debug(ctx, "upload_mode:", uploadMode)
	if g.IsEmpty(uploadMode) {
		return true
	}
	if uploadMode == "1" {
		return true
	}
	return false
}
