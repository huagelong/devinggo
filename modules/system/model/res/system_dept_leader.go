// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "devinggo/modules/system/model/req"

type SystemDeptLeaderInfo struct {
	req.SystemUser
	LeaderAddTime string `json:"leader_add_time"`
}
