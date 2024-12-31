// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

type CacheInfo struct {
	Keys   []string   `json:"keys"`
	Server ServerInfo `json:"server"`
}

type ServerInfo struct {
	Version      string `json:"version"`
	RedisMode    string `json:"redis_mode"`
	RunDays      string `json:"run_days"`
	AofEnabled   string `json:"aof_enabled"`
	UseMemory    string `json:"use_memory"`
	Port         string `json:"port"`
	Clients      string `json:"clients"`
	ExpiredKeys  string `json:"expired_keys"`
	SysTotalKeys int    `json:"sys_total_keys"`
}
