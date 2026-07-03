package websocket

import "context"

func FilterSocketIDsByApp(ctx context.Context, appID string, socketIDs []string) []string {
	if appID == "" || len(socketIDs) == 0 {
		return socketIDs
	}

	filtered := make([]string, 0, len(socketIDs))
	for _, socketID := range socketIDs {
		if GetAppIDBySocketId4Redis(ctx, socketID) == appID {
			filtered = append(filtered, socketID)
		}
	}
	return filtered
}

func GetAllSocketIDByChannelForApp(ctx context.Context, appID, channel string) []string {
	return FilterSocketIDsByApp(ctx, appID, GetAllSocketIdByChannel4Redis(ctx, channel))
}

func GetAllChannelsForApp(ctx context.Context, appID string) []string {
	allChannels := GetAllChannels(ctx)
	filtered := make([]string, 0, len(allChannels))
	for _, channel := range allChannels {
		if len(GetAllSocketIDByChannelForApp(ctx, appID, channel)) > 0 {
			filtered = append(filtered, channel)
		}
	}
	return filtered
}

func GetAllSocketIDsByApp(ctx context.Context, appID string) []string {
	if appID == "" {
		return nil
	}
	heartbeatMap := getSocketHeartbeatMap4Redis(ctx)
	socketIDs := make([]string, 0, len(heartbeatMap))
	for socketID := range heartbeatMap {
		if GetAppIDBySocketId4Redis(ctx, socketID) == appID {
			socketIDs = append(socketIDs, socketID)
		}
	}
	return socketIDs
}

func GetUserIDsBySocketIDs(ctx context.Context, socketIDs []string) []string {
	seen := make(map[string]struct{}, len(socketIDs))
	userIDs := make([]string, 0, len(socketIDs))
	for _, socketID := range socketIDs {
		userID := GetUserIdBySocketId(ctx, socketID)
		if userID == "" {
			continue
		}
		if _, ok := seen[userID]; ok {
			continue
		}
		seen[userID] = struct{}{}
		userIDs = append(userIDs, userID)
	}
	return userIDs
}
