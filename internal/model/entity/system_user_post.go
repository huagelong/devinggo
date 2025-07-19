// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemUserPost is the golang structure for table system_user_post.
type SystemUserPost struct {
	UserId int64 `json:"userId" orm:"user_id" description:"用户主键"` // 用户主键
	PostId int64 `json:"postId" orm:"post_id" description:"岗位主键"` // 岗位主键
}
