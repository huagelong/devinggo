// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Migrations is the golang structure for table migrations.
type Migrations struct {
	Id        uint   `json:"id"        orm:"id"        description:""` //
	Migration string `json:"migration" orm:"migration" description:""` //
	Batch     int    `json:"batch"     orm:"batch"     description:""` //
}
