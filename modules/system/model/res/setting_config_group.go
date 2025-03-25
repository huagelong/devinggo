// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SettingConfigGroup struct {
	Id        int64       `json:"id"                 description:"主键"`  // 主键
	Name      string      `json:"name"             description:"配置组名称"` // 配置组名称
	Code      string      `json:"code"             description:"配置组标识"` // 配置组标识
	CreatedBy int64       `json:"createdBy"  description:"创建者"`         // 创建者
	UpdatedBy int64       `json:"updatedBy"  description:"更新者"`         // 更新者
	CreatedAt *gtime.Time `json:"createdAt"  description:""`            //
	UpdatedAt *gtime.Time `json:"updatedAt"  description:""`            //
	Remark    string      `json:"remark"         description:"备注"`      // 备注
}
