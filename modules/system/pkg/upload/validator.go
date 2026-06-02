// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"

	"devinggo/modules/system/myerror"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

// Validator 文件验证器
type Validator struct {
	ctx context.Context
}

// NewValidator 创建文件验证器
func NewValidator(ctx context.Context) *Validator {
	return &Validator{ctx: ctx}
}

// ValidateFile 验证文件类型（文件和图片）
func (v *Validator) ValidateFile(file *ghttp.UploadFile) error {
	ext := Ext(file.Filename)
	return v.ValidateExtension(ext, true)
}

// ValidateImage 验证图片类型
func (v *Validator) ValidateImage(file *ghttp.UploadFile) error {
	ext := Ext(file.Filename)
	return v.ValidateImageExtension(ext)
}

// ValidateExtension 验证文件扩展名（支持文件和图片）
func (v *Validator) ValidateExtension(ext string, allowImage bool) error {
	// 检查文件类型
	allowedFiles, err := v.getAllowedFiles()
	if err != nil {
		return err
	}

	if gstr.InArray(allowedFiles, ext) {
		return nil
	}

	// 如果允许图片，则继续检查图片类型
	if allowImage {
		return v.ValidateImageExtension(ext)
	}

	return myerror.ValidationFailed(v.ctx, "不允许上传此类型文件")
}

// ValidateImageExtension 验证图片扩展名
func (v *Validator) ValidateImageExtension(ext string) error {
	allowedImages, err := v.getAllowedImages()
	if err != nil {
		return err
	}
	g.Log().Debugf(v.ctx, "Allowed image extensions: %v", allowedImages)
	if !gstr.InArray(allowedImages, ext) {
		return myerror.ValidationFailed(v.ctx, "不允许上传此类型图片")
	}

	return nil
}

// getAllowedFiles 获取允许的文件类型列表
func (v *Validator) getAllowedFiles() ([]string, error) {
	configValue, err := service.SettingConfig().GetConfigByKey(v.ctx, ConfigKeyUploadAllowFile, ConfigGroupUpload)
	if err != nil {
		return nil, err
	}
	return gstr.Split(configValue, ","), nil
}

// getAllowedImages 获取允许的图片类型列表
func (v *Validator) getAllowedImages() ([]string, error) {
	configValue, err := service.SettingConfig().GetConfigByKey(v.ctx, ConfigKeyUploadAllowImage, ConfigGroupUpload)
	if err != nil {
		return nil, err
	}
	return gstr.Split(configValue, ","), nil
}
