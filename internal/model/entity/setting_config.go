// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SettingConfig is the golang structure for table setting_config.
type SettingConfig struct {
	GroupId          int64  `json:"groupId"          orm:"group_id"           description:"组id"`  // 组id
	Key              string `json:"key"              orm:"key"                description:"配置键名"` // 配置键名
	Value            string `json:"value"            orm:"value"              description:"配置值"`  // 配置值
	Name             string `json:"name"             orm:"name"               description:""`     //
	InputType        string `json:"inputType"        orm:"input_type"         description:""`     //
	ConfigSelectData string `json:"configSelectData" orm:"config_select_data" description:""`     //
	Sort             int    `json:"sort"             orm:"sort"               description:""`     //
	Remark           string `json:"remark"           orm:"remark"             description:""`     //
}
