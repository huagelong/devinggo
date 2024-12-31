// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemAppApi is the golang structure for table system_app_api.
type SystemAppApi struct {
	AppId uint64 `json:"appId" orm:"app_id" description:"应用ID"`   // 应用ID
	ApiId uint64 `json:"apiId" orm:"api_id" description:"API—ID"` // API—ID
}
