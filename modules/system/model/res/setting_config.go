// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/encoding/gjson"

type SettingConfig struct {
	GroupId          int64       `json:"group_id"                     description:"组id"`      // 组id
	Key              string      `json:"key"                              description:"配置键名"` // 配置键名
	Value            string      `json:"value"                          description:"配置值"`    // 配置值
	Name             string      `json:"name"                            description:"配置名称"`  // 配置名称
	InputType        string      `json:"input_type"                 description:"数据输入类型"`     // 数据输入类型
	ConfigSelectData *gjson.Json `json:"config_select_data"  description:"配置选项数据"`            // 配置选项数据
	Sort             int         `json:"sort"                            description:"排序"`    // 排序
	Remark           string      `json:"remark"                        description:"备注"`      // 备注
}
