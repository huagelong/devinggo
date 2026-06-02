// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenButtons is the golang structure of table code_gen_buttons for DAO operations like Where/Data.
type CodeGenButtons struct {
	g.Meta        `orm:"table:code_gen_buttons, do:true"`
	Id            any         //
	TableId       any         // 所属表ID
	ButtonCode    any         // 按钮编码
	ButtonName    any         // 按钮名称
	ButtonComment any         // 按钮描述
	IsShow        any         // 是否显示: 1=显示, 2=隐藏
	Sort          any         // 排序
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
