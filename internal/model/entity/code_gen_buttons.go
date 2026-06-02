// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenButtons is the golang structure for table code_gen_buttons.
type CodeGenButtons struct {
	Id            int64       `json:"id"            orm:"id"             description:""`                 //
	TableId       int64       `json:"tableId"       orm:"table_id"       description:"所属表ID"`            // 所属表ID
	ButtonCode    string      `json:"buttonCode"    orm:"button_code"    description:"按钮编码"`             // 按钮编码
	ButtonName    string      `json:"buttonName"    orm:"button_name"    description:"按钮名称"`             // 按钮名称
	ButtonComment string      `json:"buttonComment" orm:"button_comment" description:"按钮描述"`             // 按钮描述
	IsShow        int         `json:"isShow"        orm:"is_show"        description:"是否显示: 1=显示, 2=隐藏"` // 是否显示: 1=显示, 2=隐藏
	Sort          int         `json:"sort"          orm:"sort"           description:"排序"`               // 排序
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`             // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`             // 更新时间
}
