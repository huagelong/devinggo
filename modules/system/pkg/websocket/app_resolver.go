package websocket

import (
	"context"
	"fmt"

	"devinggo/internal/model/entity"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
)

func ResolvePusherAppByID(ctx context.Context, appID string) (*PusherAuthCredentials, error) {
	if appID == "" {
		return nil, fmt.Errorf("app_id is required")
	}

	var app *entity.SystemApp
	if err := service.SystemApp().Model(ctx).Where("app_id", appID).Where("status", 1).Scan(&app); err != nil {
		return nil, err
	}
	if g.IsEmpty(app) {
		return nil, fmt.Errorf("pusher app not found: %s", appID)
	}

	return &PusherAuthCredentials{
		AppID:     app.AppId,
		AppKey:    app.AppKey,
		AppSecret: app.AppSecret,
	}, nil
}

func ResolvePusherAppByKey(ctx context.Context, appKey string) (*PusherAuthCredentials, error) {
	if appKey == "" {
		return nil, fmt.Errorf("app_key is required")
	}

	var app *entity.SystemApp
	if err := service.SystemApp().Model(ctx).Where("app_key", appKey).Where("status", 1).Scan(&app); err != nil {
		return nil, err
	}
	if g.IsEmpty(app) {
		return nil, fmt.Errorf("pusher app not found for key: %s", appKey)
	}

	return &PusherAuthCredentials{
		AppID:     app.AppId,
		AppKey:    app.AppKey,
		AppSecret: app.AppSecret,
	}, nil
}

func ResolvePusherAppBySocketID(ctx context.Context, socketID string) (*PusherAuthCredentials, error) {
	appID := GetAppIDBySocketId4Redis(ctx, socketID)
	if appID == "" {
		return nil, fmt.Errorf("socket app not found: %s", socketID)
	}
	return ResolvePusherAppByID(ctx, appID)
}
