package res

type PusherMonitorAppItem struct {
	Id        int64  `json:"id"`
	GroupId   int64  `json:"groupId"`
	GroupName string `json:"groupName"`
	AppName   string `json:"appName"`
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	Status    int    `json:"status"`
}

type PusherMonitorOverview struct {
	App                  PusherMonitorAppItem `json:"app"`
	ConnectionCount      int                  `json:"connectionCount"`
	ActiveChannelCount   int                  `json:"activeChannelCount"`
	SubscriptionCount    int                  `json:"subscriptionCount"`
	PresenceChannelCount int                  `json:"presenceChannelCount"`
	PresenceUserCount    int                  `json:"presenceUserCount"`
	RecentEventCount     int                  `json:"recentEventCount"`
	RecentErrorCount     int                  `json:"recentErrorCount"`
}

type PusherMonitorChannelItem struct {
	Name              string   `json:"name"`
	ChannelType       string   `json:"channelType"`
	SubscriptionCount int      `json:"subscriptionCount"`
	UserCount         int      `json:"userCount"`
	ServerCount       int      `json:"serverCount"`
	ServerNames       []string `json:"serverNames,omitempty"`
}

type PusherMonitorChannelDetail struct {
	Channel   PusherMonitorChannelItem `json:"channel"`
	SocketIds []string                 `json:"socketIds"`
	UserIds   []string                 `json:"userIds"`
}

type PusherMonitorConnectionItem struct {
	SocketId        string   `json:"socketId"`
	AppId           string   `json:"appId"`
	UserId          string   `json:"userId"`
	ServerName      string   `json:"serverName"`
	Addr            string   `json:"addr"`
	ChannelCount    int      `json:"channelCount"`
	Channels        []string `json:"channels"`
	ConnectedAt     string   `json:"connectedAt"`
	LastHeartbeatAt string   `json:"lastHeartbeatAt"`
}

type PusherMonitorTerminateResult struct {
	RequestedCount  int      `json:"requestedCount"`
	TerminatedCount int      `json:"terminatedCount"`
	SocketIds       []string `json:"socketIds"`
}

type PusherMonitorDebugEventResult struct {
	AppId          string   `json:"appId"`
	TriggeredCount int      `json:"triggeredCount"`
	Channels       []string `json:"channels"`
	EventName      string   `json:"eventName"`
}
