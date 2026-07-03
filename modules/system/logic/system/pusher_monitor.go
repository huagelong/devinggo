package system

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"

	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sPusherMonitor struct {
	base.BaseService
}

func init() {
	service.RegisterPusherMonitor(NewPusherMonitor())
}

func NewPusherMonitor() *sPusherMonitor {
	return &sPusherMonitor{}
}

func (s *sPusherMonitor) ListApps(ctx context.Context) ([]res.PusherMonitorAppItem, error) {
	var apps []*entity.SystemApp
	if err := service.SystemApp().Model(ctx).Where("status", 1).Order("id asc").Scan(&apps); err != nil {
		return nil, err
	}

	groupMap, err := s.getGroupMap(ctx, apps)
	if err != nil {
		return nil, err
	}

	items := make([]res.PusherMonitorAppItem, 0, len(apps))
	for _, app := range apps {
		items = append(items, res.PusherMonitorAppItem{
			Id:        app.Id,
			GroupId:   app.GroupId,
			GroupName: groupMap[app.GroupId],
			AppName:   app.AppName,
			AppId:     app.AppId,
			AppKey:    app.AppKey,
			Status:    app.Status,
		})
	}
	return items, nil
}

func (s *sPusherMonitor) GetOverview(ctx context.Context, appID string) (*res.PusherMonitorOverview, error) {
	app, err := s.getApp(ctx, appID)
	if err != nil {
		return nil, err
	}

	channelItems, err := s.listChannelItems(ctx, appID)
	if err != nil {
		return nil, err
	}
	socketIDs := websocket.GetAllSocketIDsByApp(ctx, appID)
	subscriptionCount := 0
	presenceChannelCount := 0
	presenceUsers := make(map[string]struct{})
	for _, item := range channelItems {
		subscriptionCount += item.SubscriptionCount
		if item.ChannelType == "presence" {
			presenceChannelCount++
			detailUserIDs := websocket.GetUserIDsBySocketIDs(ctx, websocket.GetAllSocketIDByChannelForApp(ctx, appID, item.Name))
			for _, userID := range detailUserIDs {
				presenceUsers[userID] = struct{}{}
			}
		}
	}

	eventCount, errorCount := websocket.GetRecentAppMetrics(appID)
	return &res.PusherMonitorOverview{
		App: res.PusherMonitorAppItem{
			Id:      app.Id,
			GroupId: app.GroupId,
			AppName: app.AppName,
			AppId:   app.AppId,
			AppKey:  app.AppKey,
			Status:  app.Status,
		},
		ConnectionCount:      len(socketIDs),
		ActiveChannelCount:   len(channelItems),
		SubscriptionCount:    subscriptionCount,
		PresenceChannelCount: presenceChannelCount,
		PresenceUserCount:    len(presenceUsers),
		RecentEventCount:     eventCount,
		RecentErrorCount:     errorCount,
	}, nil
}

func (s *sPusherMonitor) ListChannels(ctx context.Context, appID, keyword, channelType string, reqPage *page.PageReq) ([]res.PusherMonitorChannelItem, int, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, 0, err
	}

	items, err := s.listChannelItems(ctx, appID)
	if err != nil {
		return nil, 0, err
	}
	filtered := make([]res.PusherMonitorChannelItem, 0, len(items))
	for _, item := range items {
		if keyword != "" && !strings.Contains(strings.ToLower(item.Name), strings.ToLower(keyword)) {
			continue
		}
		if channelType != "" && item.ChannelType != channelType {
			continue
		}
		filtered = append(filtered, item)
	}

	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i].SubscriptionCount != filtered[j].SubscriptionCount {
			return filtered[i].SubscriptionCount > filtered[j].SubscriptionCount
		}
		return filtered[i].Name < filtered[j].Name
	})

	total := len(filtered)
	pageNum, pageSize, offset := page.CalPage(reqPage.Page, reqPage.PageSize)
	_ = pageNum
	if offset >= total {
		return []res.PusherMonitorChannelItem{}, total, nil
	}
	end := offset + pageSize
	if end > total {
		end = total
	}
	return filtered[offset:end], total, nil
}

func (s *sPusherMonitor) GetChannelDetail(ctx context.Context, appID, channelName string) (*res.PusherMonitorChannelDetail, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, err
	}

	socketIDs := websocket.GetAllSocketIDByChannelForApp(ctx, appID, channelName)
	serverNames := make([]string, 0)
	serverSeen := make(map[string]struct{})
	for _, socketID := range socketIDs {
		serverName := websocket.GetServerNameBySocketId4Redis(ctx, socketID)
		if serverName == "" {
			continue
		}
		if _, ok := serverSeen[serverName]; ok {
			continue
		}
		serverSeen[serverName] = struct{}{}
		serverNames = append(serverNames, serverName)
	}
	sort.Strings(serverNames)

	userIDs := websocket.GetUserIDsBySocketIDs(ctx, socketIDs)
	sort.Strings(socketIDs)
	sort.Strings(userIDs)

	return &res.PusherMonitorChannelDetail{
		Channel: res.PusherMonitorChannelItem{
			Name:              channelName,
			ChannelType:       s.channelTypeName(channelName),
			SubscriptionCount: len(socketIDs),
			UserCount:         len(userIDs),
			ServerCount:       len(serverNames),
			ServerNames:       serverNames,
		},
		SocketIds: socketIDs,
		UserIds:   userIDs,
	}, nil
}

func (s *sPusherMonitor) ListConnections(ctx context.Context, appID, socketID, userID, channel string, reqPage *page.PageReq) ([]res.PusherMonitorConnectionItem, int, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, 0, err
	}

	allSocketIDs := websocket.GetAllSocketIDsByApp(ctx, appID)
	items := make([]res.PusherMonitorConnectionItem, 0, len(allSocketIDs))
	for _, currentSocketID := range allSocketIDs {
		item := s.buildConnectionItem(ctx, appID, currentSocketID)
		if socketID != "" && !strings.Contains(item.SocketId, socketID) {
			continue
		}
		if userID != "" && item.UserId != userID {
			continue
		}
		if channel != "" && !slices.Contains(item.Channels, channel) {
			continue
		}
		items = append(items, item)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].SocketId < items[j].SocketId
	})

	total := len(items)
	pageNum, pageSize, offset := page.CalPage(reqPage.Page, reqPage.PageSize)
	_ = pageNum
	if offset >= total {
		return []res.PusherMonitorConnectionItem{}, total, nil
	}
	end := offset + pageSize
	if end > total {
		end = total
	}
	return items[offset:end], total, nil
}

func (s *sPusherMonitor) TerminateConnections(ctx context.Context, appID string, socketIDs []string) (*res.PusherMonitorTerminateResult, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, err
	}

	terminated := make([]string, 0, len(socketIDs))
	for _, socketID := range socketIDs {
		if websocket.GetAppIDBySocketId4Redis(ctx, socketID) != appID {
			continue
		}
		serverName := websocket.GetServerNameBySocketId4Redis(ctx, socketID)
		if serverName == "" {
			continue
		}
		if serverName == websocket.GetServerName() {
			if websocket.TerminateLocalClient(socketID) {
				terminated = append(terminated, socketID)
			}
		} else {
			if err := websocket.PublishTerminateConnectionMessage(ctx, socketID, &websocket.TopicTerminateConnection{
				SocketID: socketID,
			}); err == nil {
				terminated = append(terminated, socketID)
			}
		}
	}

	sort.Strings(terminated)
	return &res.PusherMonitorTerminateResult{
		RequestedCount:  len(socketIDs),
		TerminatedCount: len(terminated),
		SocketIds:       terminated,
	}, nil
}

func (s *sPusherMonitor) TerminateUsers(ctx context.Context, appID string, userIDs []string) (*res.PusherMonitorTerminateResult, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, err
	}

	socketIDs := make([]string, 0)
	seen := make(map[string]struct{})
	for _, userID := range userIDs {
		for _, socketID := range websocket.FilterSocketIDsByApp(ctx, appID, websocket.GetAllSocketIdsByUserId(ctx, userID)) {
			if _, ok := seen[socketID]; ok {
				continue
			}
			seen[socketID] = struct{}{}
			socketIDs = append(socketIDs, socketID)
		}
	}
	return s.TerminateConnections(ctx, appID, socketIDs)
}

func (s *sPusherMonitor) DebugEvent(ctx context.Context, appID, name string, channels []string, data string, socketID string) (*res.PusherMonitorDebugEventResult, error) {
	if _, err := s.getApp(ctx, appID); err != nil {
		return nil, err
	}

	if nameErr := websocket.ValidateEventName(name); nameErr != nil {
		return nil, myerror.ValidationFailed(ctx, nameErr.Error())
	}
	if dataErr := websocket.ValidateEventData(data); dataErr != nil {
		return nil, myerror.ValidationFailed(ctx, dataErr.Error())
	}
	if chErr := websocket.ValidateChannels(channels); chErr != nil {
		return nil, myerror.ValidationFailed(ctx, chErr.Error())
	}

	triggeredCount := 0
	for _, channel := range channels {
		payload := data
		if websocket.IsEncryptedChannel(channel) {
			sharedSecret, err := websocket.GetSharedSecret(ctx, channel)
			if err != nil {
				return nil, err
			}
			payload, err = websocket.EncryptMessage(ctx, data, sharedSecret)
			if err != nil {
				return nil, err
			}
		}

		response := &websocket.PusherResponse{
			Event:   name,
			Channel: channel,
			Data:    payload,
		}
		websocket.SendToChannelWithExclude(appID, channel, response, socketID)
		_ = websocket.PublishChannelMessage(ctx, channel, &websocket.TopicWResponse{
			AppID:           appID,
			Topic:           channel,
			ExcludeSocketID: socketID,
			PusherResponse:  response,
		})
		triggeredCount++
	}

	return &res.PusherMonitorDebugEventResult{
		AppId:          appID,
		TriggeredCount: triggeredCount,
		Channels:       channels,
		EventName:      name,
	}, nil
}

func (s *sPusherMonitor) getApp(ctx context.Context, appID string) (*entity.SystemApp, error) {
	var app *entity.SystemApp
	if err := service.SystemApp().Model(ctx).Where("app_id", appID).Where("status", 1).Scan(&app); err != nil {
		return nil, err
	}
	if g.IsEmpty(app) {
		return nil, myerror.ValidationFailed(ctx, fmt.Sprintf("app 不存在: %s", appID))
	}
	return app, nil
}

func (s *sPusherMonitor) getGroupMap(ctx context.Context, apps []*entity.SystemApp) (map[int64]string, error) {
	groupIDs := make([]int64, 0)
	seen := make(map[int64]struct{})
	for _, app := range apps {
		if _, ok := seen[app.GroupId]; ok {
			continue
		}
		seen[app.GroupId] = struct{}{}
		groupIDs = append(groupIDs, app.GroupId)
	}
	if len(groupIDs) == 0 {
		return map[int64]string{}, nil
	}

	var groups []*entity.SystemAppGroup
	if err := service.SystemAppGroup().Model(ctx).WhereIn("id", groupIDs).Scan(&groups); err != nil {
		return nil, err
	}
	groupMap := make(map[int64]string, len(groups))
	for _, group := range groups {
		groupMap[group.Id] = group.Name
	}
	return groupMap, nil
}

func (s *sPusherMonitor) listChannelItems(ctx context.Context, appID string) ([]res.PusherMonitorChannelItem, error) {
	channelNames := websocket.GetAllChannelsForApp(ctx, appID)
	items := make([]res.PusherMonitorChannelItem, 0, len(channelNames))
	for _, channelName := range channelNames {
		socketIDs := websocket.GetAllSocketIDByChannelForApp(ctx, appID, channelName)
		serverNames := make([]string, 0)
		serverSeen := make(map[string]struct{})
		for _, socketID := range socketIDs {
			serverName := websocket.GetServerNameBySocketId4Redis(ctx, socketID)
			if serverName == "" {
				continue
			}
			if _, ok := serverSeen[serverName]; ok {
				continue
			}
			serverSeen[serverName] = struct{}{}
			serverNames = append(serverNames, serverName)
		}
		sort.Strings(serverNames)

		item := res.PusherMonitorChannelItem{
			Name:              channelName,
			ChannelType:       s.channelTypeName(channelName),
			SubscriptionCount: len(socketIDs),
			UserCount:         len(websocket.GetUserIDsBySocketIDs(ctx, socketIDs)),
			ServerCount:       len(serverNames),
			ServerNames:       serverNames,
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *sPusherMonitor) channelTypeName(channel string) string {
	switch websocket.GetChannelType(channel) {
	case websocket.ChannelTypePrivate:
		return "private"
	case websocket.ChannelTypePresence:
		return "presence"
	case websocket.ChannelTypeEncrypted:
		return "encrypted"
	default:
		return "public"
	}
}

func (s *sPusherMonitor) buildConnectionItem(ctx context.Context, appID, socketID string) res.PusherMonitorConnectionItem {
	channels := websocket.GetAllChannelBySocketId(ctx, socketID)
	sort.Strings(channels)
	item := res.PusherMonitorConnectionItem{
		SocketId:     socketID,
		AppId:        appID,
		UserId:       websocket.GetUserIdBySocketId(ctx, socketID),
		ServerName:   websocket.GetServerNameBySocketId4Redis(ctx, socketID),
		ChannelCount: len(channels),
		Channels:     channels,
	}

	if heartbeat := websocket.GetSocketHeartbeatTime4Redis(ctx, socketID); heartbeat > 0 {
		item.LastHeartbeatAt = time.Unix(heartbeat, 0).Format(time.RFC3339)
	}
	if client := websocket.GetLocalClient(socketID); client != nil {
		item.Addr = client.Addr
		if client.FirstTime > 0 {
			item.ConnectedAt = time.Unix(client.FirstTime, 0).Format(time.RFC3339)
		}
	}
	return item
}
