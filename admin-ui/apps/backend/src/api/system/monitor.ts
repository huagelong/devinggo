import { requestClient } from '#/api/request';
import type { PageResponse } from '#/types/paging';

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
    sys_total_keys?: string;
    use_memory?: string;
  }

  export interface CacheInfo {
    server: CacheServerInfo;
    keys: string[];
  }

  export interface ViewCachePayload {
    key: string;
  }

  export interface ViewCacheResponse {
    data: {
      content: string;
    };
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
