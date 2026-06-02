<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { onMounted, onUnmounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { getServerInfo } from '#/api/system/monitor';
import { logger } from '#/utils/logger';

import {
  CpuIcon,
  DesktopIcon,
  ServerIcon,
} from 'tdesign-icons-vue-next';
import {
  Card,
  Col,
  Progress,
  Row,
  Space,
  Tag,
} from 'tdesign-vue-next';

defineOptions({ name: 'SystemServer' });

const loading = ref(false);
const serverInfo = ref<MonitorApi.ServerInfoResponse | null>(null);
const hasServerApi = ref(true);
const refreshTimer = ref<ReturnType<typeof setInterval>>();

function formatBytes(bytes: number): string {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${Number.parseFloat((bytes / k ** i).toFixed(2))} ${sizes[i]}`;
}

function formatUptime(seconds: number): string {
  const days = Math.floor(seconds / 86400);
  const hours = Math.floor((seconds % 86400) / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
    const parts: string[] = [];
    if (days > 0) parts.push(`${days}${$t('system.monitor.server.day')}`);
    if (hours > 0) parts.push(`${hours}${$t('system.monitor.server.hour')}`);
    if (minutes > 0) parts.push(`${minutes}${$t('system.monitor.server.minute')}`);
    return parts.join(' ') || `${seconds}${$t('system.monitor.server.second')}`;
}

async function fetchServerInfo() {
  if (!hasServerApi.value) return;

  loading.value = true;
  try {
    const response = await getServerInfo();
    serverInfo.value = response;
  } catch (error: unknown) {
    if ((error as { response?: { status?: number } })?.response?.status === 404) {
      hasServerApi.value = false;
      message.info($t('common.serverMonitorNotAvailable'));
    } else {
      logger.error(error);
      message.error($t('common.serverInfoFailed'));
    }
  } finally {
    loading.value = false;
  }
}

function startAutoRefresh() {
  refreshTimer.value = setInterval(() => {
    void fetchServerInfo();
  }, 10000);
}

function stopAutoRefresh() {
  if (refreshTimer.value) {
    clearInterval(refreshTimer.value);
    refreshTimer.value = undefined;
  }
}

onMounted(() => {
  void fetchServerInfo();
  startAutoRefresh();
});

onUnmounted(() => {
  stopAutoRefresh();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex flex-col gap-3">
      <!-- API Not Available Notice -->
      <Card v-if="!hasServerApi">
        <div class="flex flex-col items-center justify-center py-12 text-gray-400">
          <ServerIcon class="mb-4 text-6xl" />
          <div class="text-lg">{{ $t('system.monitor.server.apiNotAvailable') }}</div>
          <div class="text-sm">{{ $t('system.monitor.server.apiNotAvailableDesc') }}</div>
        </div>
      </Card>

      <!-- Server Info Content -->
      <template v-else>
        <!-- System Overview -->
        <Card :title="$t('system.monitor.server.systemOverview')">
          <template #actions>
            <Tag
              :theme="loading ? 'default' : 'success'"
              variant="light"
            >
              {{ loading ? $t('system.monitor.server.refreshing') : $t('system.monitor.server.realtimeMonitoring') }}
            </Tag>
          </template>
          <Row :gutter="24">
            <Col :span="6">
              <div class="mb-2 flex items-center gap-2 text-sm text-gray-500">
                <DesktopIcon />
                {{ $t('system.monitor.server.os') }}
              </div>
              <div class="text-base">{{ serverInfo?.os || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 flex items-center gap-2 text-sm text-gray-500">
                <ServerIcon />
                {{ $t('system.monitor.server.arch') }}
              </div>
              <div class="text-base">{{ serverInfo?.arch || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.hostname') }}</div>
              <div class="text-base">{{ serverInfo?.hostname || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.uptime') }}</div>
              <div class="text-base">
                {{ serverInfo?.uptime ? formatUptime(serverInfo.uptime) : '-' }}
              </div>
            </Col>
          </Row>
          <Row :gutter="24" class="mt-4">
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.serverTime') }}</div>
              <div class="text-base">{{ serverInfo?.server_time || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.goVersion') }}</div>
              <div class="text-base">{{ serverInfo?.go_runtime?.go_version || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">Goroutines</div>
              <div class="text-base">{{ serverInfo?.go_runtime?.goroutines || '-' }}</div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.gcStats') }}</div>
              <div class="text-base">{{ serverInfo?.go_runtime?.gc_stats || '-' }}</div>
            </Col>
          </Row>
        </Card>

        <!-- CPU & Memory -->
        <Row :gutter="16">
          <Col :span="12">
            <Card :title="$t('system.monitor.server.cpuUsage')">
              <div class="flex flex-col items-center py-4">
                <Progress
                  :percentage="serverInfo?.cpu?.usage ?? 0"
                  :size="'large'"
                  :theme="((
                    serverInfo?.cpu?.usage ?? 0
                  ) > 80
                    ? 'danger'
                    : (serverInfo?.cpu?.usage ?? 0) > 60
                      ? 'warning'
                      : 'primary') as any"
                  :label="`${(serverInfo?.cpu?.usage ?? 0).toFixed(1)}%`"
                />
                <div class="mt-4 text-sm text-gray-500">
                  <Space>
                    <span>
                      <CpuIcon />
                      {{ $t('system.monitor.server.cores') }}: {{ serverInfo?.cpu?.num ?? '-' }}
                    </span>
                    <span>{{ $t('system.monitor.server.model') }}: {{ serverInfo?.cpu?.model ?? '-' }}</span>
                  </Space>
                </div>
              </div>
            </Card>
          </Col>
          <Col :span="12">
            <Card :title="$t('system.monitor.server.memoryUsage')">
              <div class="flex flex-col items-center py-4">
                <Progress
                  :percentage="serverInfo?.memory?.usage ?? 0"
                  :size="'large'"
                  :theme="((
                    serverInfo?.memory?.usage ?? 0
                  ) > 80
                    ? 'danger'
                    : (serverInfo?.memory?.usage ?? 0) > 60
                      ? 'warning'
                      : 'primary') as any"
                  :label="`${(serverInfo?.memory?.usage ?? 0).toFixed(1)}%`"
                />
                <div class="mt-4 text-sm text-gray-500">
                  <Space>
                    <span>
                      {{ $t('system.monitor.server.used') }}: {{ formatBytes(serverInfo?.memory?.used ?? 0) }}
                    </span>
                    <span>
                      {{ $t('system.monitor.server.total') }}: {{ formatBytes(serverInfo?.memory?.total ?? 0) }}
                    </span>
                    <span>
                      {{ $t('system.monitor.server.available') }}: {{ formatBytes(serverInfo?.memory?.free ?? 0) }}
                    </span>
                  </Space>
                </div>
              </div>
            </Card>
          </Col>
        </Row>

        <!-- Disk Info -->
        <Card :title="$t('system.monitor.server.diskInfo')">
          <Row :gutter="16">
            <Col
              v-for="(disk, index) in serverInfo?.disks ?? []"
              :key="index"
              :span="8"
            >
              <Card :bordered="false" class="bg-gray-50">
                <div class="mb-2 text-sm font-medium text-gray-600">
                  {{ disk.mount_point }}
                </div>
                <Progress
                  :percentage="disk.usage"
                  :theme="(disk.usage > 80
                      ? 'danger'
                      : disk.usage > 60
                        ? 'warning'
                        : 'primary') as any"
                  :label="`${disk.usage.toFixed(1)}%`"
                />
                <div class="mt-2 text-xs text-gray-500">
                  <Space>
                    <span>{{ $t('system.monitor.server.fileSystem') }}: {{ disk.file_system }}</span>
                    <span>
                      {{ $t('system.monitor.server.total') }}: {{ formatBytes(disk.total) }}
                    </span>
                    <span>
                      {{ $t('system.monitor.server.used') }}: {{ formatBytes(disk.used) }}
                    </span>
                  </Space>
                </div>
              </Card>
            </Col>
          </Row>
          <div
            v-if="!serverInfo?.disks?.length"
            class="py-8 text-center text-gray-400"
          >
            {{ $t('system.monitor.server.noDiskInfo') }}
          </div>
        </Card>

        <!-- Go Runtime -->
        <Card :title="$t('system.monitor.server.goRuntime')">
          <Row :gutter="24">
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.heapAlloc') }}</div>
              <div class="text-base">
                {{ formatBytes(serverInfo?.go_runtime?.heap_alloc ?? 0) }}
              </div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.heapSysMemory') }}</div>
              <div class="text-base">
                {{ formatBytes(serverInfo?.go_runtime?.heap_sys ?? 0) }}
              </div>
            </Col>
            <Col :span="6">
              <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.server.stackInUse') }}</div>
              <div class="text-base">
                {{ formatBytes(serverInfo?.go_runtime?.stack_in_use ?? 0) }}
              </div>
            </Col>
          </Row>
        </Card>
      </template>
    </div>
  </Page>
</template>
