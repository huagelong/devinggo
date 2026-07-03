import { requestClient } from '#/api/request';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace MonitorApi {
  export interface OnlineUserItem {
    id: number;
    username: string;
    nickname: string;
    app_id: number;
    login_ip: string;
    login_time: string;
  }

  export interface OnlineUserQuery {
    username?: string;
    app_id?: number;
    page?: number;
    page_size?: number;
  }

  export interface OnlineUserResponse extends PageResponse<OnlineUserItem> {}

  export interface KickPayload {
    id: number;
    app_id: number;
  }

  export interface CacheServerInfo {
    version?: string;
    clients?: string;
    redis_mode?: string;
    run_days?: string;
    port?: string;
    aof_enabled?: string;
    expired_keys?: string;
    sys_total_keys?: number;
    use_memory?: string;
    qps?: string;
    hit_rate?: number;
    blocked_clients?: string;
    rejected_conn?: string;
    memory_peak?: string;
    mem_fragment_ratio?: string;
    total_commands?: string;
  }

  export interface CacheInfo {
    server: CacheServerInfo;
    keys: string[];
  }

  export interface ViewCachePayload {
    key: string;
  }

  export interface ViewCacheResponse {
    content: string;
  }

  export interface DeleteCachePayload {
    key: string;
  }

  // Server Monitor
  export interface CpuInfo {
    num: number;
    usage: number;
    model: string;
  }

  export interface MemoryInfo {
    total: number;
    used: number;
    free: number;
    usage: number;
  }

  export interface DiskInfo {
    total: number;
    used: number;
    free: number;
    usage: number;
    mount_point: string;
    file_system: string;
  }

  export interface GoRuntimeInfo {
    go_version: string;
    goroutines: number;
    gc_stats: string;
    heap_alloc: number;
    heap_sys: number;
    stack_in_use: number;
  }

  export interface ServerInfoResponse {
    cpu: CpuInfo;
    memory: MemoryInfo;
    disks: DiskInfo[];
    go_runtime: GoRuntimeInfo;
    os: string;
    arch: string;
    hostname: string;
    uptime: number;
    server_time: string;
  }

  export interface DbMonitorGroup {
    groupName: string;
    isDefault: boolean;
    dbType: string;
  }

  export interface DbMonitorSummary {
    connectionStatus: string;
    currentConn: number;
    activeConn: number;
    databaseSize: number;
    hitRate: number;
    slowQueryEnabled: boolean;
  }

  export interface DbMonitorOverview {
    groupName: string;
    databaseName: string;
    version: string;
    maxConnections: number;
    currentConnections: number;
    activeConnections: number;
    idleConnections: number;
    waitingConnections: number;
    xactCommit: number;
    xactRollback: number;
    hitRate: number;
    databaseSize: number;
  }

  export interface DbMonitorTimeseriesPoint {
    time: string;
    tps: number;
    rollbackTps: number;
    avgQueryMs?: null | number;
    slowQueryCount?: null | number;
  }

  export interface DbMonitorCapabilities {
    hasPgStatStatements: boolean;
    availableMetrics: string[];
    unavailableMetrics: string[];
    slowQueryThreshold: number;
  }

  export interface DbMonitorAlert {
    level: string;
    message: string;
  }

  export interface DbMonitorResponse {
    summary: DbMonitorSummary;
    overview: DbMonitorOverview;
    timeseries: DbMonitorTimeseriesPoint[];
    capabilities: DbMonitorCapabilities;
    alerts: DbMonitorAlert[];
  }

  export interface QueueMonitorQueueItem {
    queue: string;
    memoryUsage: number;
    latency: number;
    size: number;
    groups: number;
    pending: number;
    active: number;
    scheduled: number;
    retry: number;
    archived: number;
    completed: number;
    aggregating: number;
    processed: number;
    failed: number;
    processedTotal: number;
    failedTotal: number;
    paused: boolean;
    timestamp: string;
  }

  export interface QueueMonitorWorkerItem {
    id: string;
    host: string;
    pid: number;
    concurrency: number;
    queues: Record<string, number>;
    strictPriority: boolean;
    started: string;
    status: string;
    activeWorkerCount: number;
  }

  export interface QueueMonitorOverviewResponse {
    queue: QueueMonitorQueueItem;
    workers: QueueMonitorWorkerItem[];
  }

  export interface QueueMonitorTaskItem {
    id: string;
    queue: string;
    type: string;
    payload: string;
    state: string;
    maxRetry: number;
    retried: number;
    lastErr: string;
    lastFailedAt: string;
    timeout: number;
    deadline: string;
    group: string;
    nextProcessAt: string;
    isOrphaned: boolean;
    retention: number;
    completedAt: string;
    result: string;
  }

  export interface QueueMonitorTaskQuery {
    queue: string;
    state: string;
    page: number;
    pageSize: number;
  }

  export interface QueueMonitorTaskResponse extends PageResponse<QueueMonitorTaskItem> {}

  export interface QueueMonitorSchedulerEntry {
    id: string;
    spec: string;
    taskType: string;
    queue: string;
    next: string;
    prev: string;
  }

  export interface PusherMonitorAppItem {
    id: number;
    groupId: number;
    groupName: string;
    appName: string;
    appId: string;
    appKey: string;
    status: number;
  }

  export interface PusherMonitorOverview {
    app: PusherMonitorAppItem;
    connectionCount: number;
    activeChannelCount: number;
    subscriptionCount: number;
    presenceChannelCount: number;
    presenceUserCount: number;
    recentEventCount: number;
    recentErrorCount: number;
  }

  export interface PusherMonitorChannelItem {
    name: string;
    channelType: string;
    subscriptionCount: number;
    userCount: number;
    serverCount: number;
    serverNames?: string[];
  }

  export interface PusherMonitorChannelDetail {
    channel: PusherMonitorChannelItem;
    socketIds: string[];
    userIds: string[];
  }

  export interface PusherMonitorConnectionItem {
    socketId: string;
    appId: string;
    userId: string;
    serverName: string;
    addr: string;
    channelCount: number;
    channels: string[];
    connectedAt: string;
    lastHeartbeatAt: string;
  }

  export interface PusherMonitorChannelsQuery extends Partial<PageQuery> {
    appId: string;
    keyword?: string;
    channelType?: string;
  }

  export interface PusherMonitorChannelsResponse extends PageResponse<PusherMonitorChannelItem> {}

  export interface PusherMonitorConnectionsQuery extends Partial<PageQuery> {
    appId: string;
    socketId?: string;
    userId?: string;
    channel?: string;
  }

  export interface PusherMonitorConnectionsResponse extends PageResponse<PusherMonitorConnectionItem> {}

  export interface PusherMonitorTerminateConnectionsPayload {
    socketIds: string[];
  }

  export interface PusherMonitorTerminateUsersPayload {
    userIds: string[];
  }

  export interface PusherMonitorTerminateResult {
    requestedCount: number;
    terminatedCount: number;
    socketIds: string[];
  }

  export interface PusherMonitorDebugEventPayload {
    name: string;
    channels: string[];
    data: string;
    socketId?: string;
  }

  export interface PusherMonitorDebugEventResult {
    appId: string;
    triggeredCount: number;
    channels: string[];
    eventName: string;
  }
}

export function getOnlineUserPageList(params: MonitorApi.OnlineUserQuery) {
  return requestClient.get<MonitorApi.OnlineUserResponse>(
    '/system/onlineUser/index',
    { params },
  );
}

export function kickUser(data: MonitorApi.KickPayload) {
  return requestClient.post<void>('/system/onlineUser/kick', data);
}

export function getCacheInfo() {
  return requestClient.get<MonitorApi.CacheInfo>('/system/cache/monitor');
}

export function viewCache(data: MonitorApi.ViewCachePayload) {
  return requestClient.post<MonitorApi.ViewCacheResponse>(
    '/system/cache/view',
    data,
  );
}

export function deleteCacheKey(data: MonitorApi.DeleteCachePayload) {
  return requestClient.delete<void>('/system/cache/delete', { data });
}

export function clearAllCache() {
  return requestClient.delete<void>('/system/cache/clear');
}

// Server Monitor APIs
export function getServerInfo() {
  return requestClient.get<MonitorApi.ServerInfoResponse>(
    '/system/server/monitor',
  );
}

export function getDbMonitorGroups() {
  return requestClient.get<MonitorApi.DbMonitorGroup[]>(
    '/system/dbMonitor/groups',
  );
}

export function getDbMonitorInfo(groupName: string) {
  return requestClient.get<MonitorApi.DbMonitorResponse>(
    '/system/dbMonitor/monitor',
    {
      params: { groupName },
    },
  );
}

export function getQueueMonitorQueues() {
  return requestClient.get<MonitorApi.QueueMonitorQueueItem[]>(
    '/system/queueMonitor/queues',
  );
}

export function getQueueMonitorOverview(queue: string) {
  return requestClient.get<MonitorApi.QueueMonitorOverviewResponse>(
    '/system/queueMonitor/overview',
    {
      params: { queue },
    },
  );
}

export function getQueueMonitorTasks(params: MonitorApi.QueueMonitorTaskQuery) {
  return requestClient.get<MonitorApi.QueueMonitorTaskResponse>(
    '/system/queueMonitor/tasks',
    {
      params,
    },
  );
}

export function getQueueMonitorSchedulerEntries() {
  return requestClient.get<MonitorApi.QueueMonitorSchedulerEntry[]>(
    '/system/queueMonitor/schedulerEntries',
  );
}

export function getPusherMonitorApps() {
  return requestClient.get<MonitorApi.PusherMonitorAppItem[]>(
    '/system/pusherMonitor/apps',
  );
}

export function getPusherMonitorOverview(appId: string) {
  return requestClient.get<MonitorApi.PusherMonitorOverview>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/overview`,
  );
}

export function getPusherMonitorChannels(
  params: MonitorApi.PusherMonitorChannelsQuery,
) {
  const { appId, ...query } = params;
  return requestClient.get<MonitorApi.PusherMonitorChannelsResponse>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/channels`,
    {
      params: query,
    },
  );
}

export function getPusherMonitorChannelDetail(appId: string, channelName: string) {
  return requestClient.get<MonitorApi.PusherMonitorChannelDetail>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/channels/${encodeURIComponent(channelName)}/detail`,
  );
}

export function getPusherMonitorConnections(
  params: MonitorApi.PusherMonitorConnectionsQuery,
) {
  const { appId, ...query } = params;
  return requestClient.get<MonitorApi.PusherMonitorConnectionsResponse>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/connections`,
    {
      params: query,
    },
  );
}

export function terminatePusherMonitorConnections(
  appId: string,
  data: MonitorApi.PusherMonitorTerminateConnectionsPayload,
) {
  return requestClient.post<MonitorApi.PusherMonitorTerminateResult>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/connections/terminate-batch`,
    data,
  );
}

export function terminatePusherMonitorUsers(
  appId: string,
  data: MonitorApi.PusherMonitorTerminateUsersPayload,
) {
  return requestClient.post<MonitorApi.PusherMonitorTerminateResult>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/users/terminate-batch`,
    data,
  );
}

export function debugPusherMonitorEvent(
  appId: string,
  data: MonitorApi.PusherMonitorDebugEventPayload,
) {
  return requestClient.post<MonitorApi.PusherMonitorDebugEventResult>(
    `/system/pusherMonitor/apps/${encodeURIComponent(appId)}/events/debug`,
    data,
  );
}
