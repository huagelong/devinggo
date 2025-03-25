// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SettingConfigGroupSave struct {
	Name   string `json:"name"    v:"required|length:1,32#请输入配置组名称|配置组名称长度必须在1到32之间"         description:"配置组名称"` // 配置组名称
	Code   string `json:"code"   v:"required|length:1,64#请输入配置组标识|配置组标识长度必须在1到64之间"          description:"配置组标识"` // 配置组标识
	Remark string `json:"remark"         description:"备注"`                                                        // 备注
}

type SettingConfigGroupUpdate struct {
	Id     int64  `json:"id" v:"required#请输入配置组ID" description:"配置组ID"`                                             // 配置组ID
	Name   string `json:"name"    v:"required|length:1,32#请输入配置组名称|配置组名称长度必须在1到32之间"           description:"配置组名称"` // 配置组名称
	Code   string `json:"code"    v:"required|length:1,64#请输入配置组标识|配置组标识长度必须在1到64之间"         description:"配置组标识"`   // 配置组标识
	Remark string `json:"remark"         description:"备注"`                                                          // 备注
}
