// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SettingConfigSave struct {
	GroupId          int64  `json:"group_id" v:"required|integer|min:1"`
	Key              string `json:"key" v:"required|max-length:32"`
	Name             string `json:"name" v:"required|max-length:255"`
	Value            string `json:"value"`
	InputType        string `json:"input_type"`
	ConfigSelectData string `json:"config_select_data"`
	Sort             int    `json:"sort"`
	Remark           string `json:"remark"`
}

type SettingConfigUpdate struct {
	GroupId          int64  `json:"group_id" v:"required|integer|min:1"`
	Key              string `json:"key" v:"required|max-length:32"`
	Name             string `json:"name" v:"required|max-length:255"`
	Value            string `json:"value"`
	InputType        string `json:"input_type"`
	ConfigSelectData string `json:"config_select_data"`
	Sort             int    `json:"sort"`
	Remark           string `json:"remark"`
}

type SettingConfigSearch struct {
	GroupId int64  `json:"group_id" v:"required|integer|min:1"`
	Key     string `json:"key"`
	Name    string `json:"name"`
}
