// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"encoding/json"
)

// FormatPresenceData 格式化Presence成员列表为Pusher v8.3.0格式
// ⚠️ v8.3.0格式要求：hash字段只包含user_info，不含user_id
// ⚠️ Pusher 协议要求：必须包装在 presence 字段中
func FormatPresenceData(members map[string]map[string]interface{}, currentUserID string) PresenceData {
	ids := make([]string, 0, len(members))
	hash := make(map[string]interface{})

	// 当前用户信息
	var me *PresenceMember

	for userID, userInfo := range members {
		ids = append(ids, userID)
		hash[userID] = userInfo // ⚠️ 只存储user_info，不包含user_id

		// 设置当前用户信息
		if userID == currentUserID {
			me = &PresenceMember{
				UserID:   userID,
				UserInfo: userInfo,
			}
		}
	}

	return PresenceData{
		Presence: PresenceMemberList{
			Count: len(members),
			Ids:   ids,
			Hash:  hash,
			Me:    me,
		},
	}
}

// ParseChannelData 解析channel_data JSON字符串
func ParseChannelData(channelDataStr string) (*PresenceMember, error) {
	var member PresenceMember
	err := json.Unmarshal([]byte(channelDataStr), &member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

// EncodeChannelData 编码channel_data为JSON字符串（用于HTTP认证端点）
func EncodeChannelData(userID string, userInfo map[string]interface{}) (string, error) {
	member := PresenceMember{
		UserID:   userID,
		UserInfo: userInfo,
	}
	data, err := json.Marshal(member)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
