<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Button,
  Card,
  Dialog,
  Input,
  Popconfirm,
  Select,
  Switch,
  Table,
  TabPanel,
  Tabs,
  Tag,
  Textarea,
} from 'tdesign-vue-next';

import { message } from '#/adapter/tdesign';
import {
  debugPusherMonitorEvent,
  getPusherMonitorApps,
  getPusherMonitorChannelDetail,
  getPusherMonitorChannels,
  getPusherMonitorConnections,
  getPusherMonitorOverview,
  terminatePusherMonitorConnections,
  terminatePusherMonitorUsers,
} from '#/api/system/monitor';
import { logger } from '#/utils/logger';
import { formatQueueMonitorTime } from '#/utils/queue-monitor-time';

defineOptions({ name: 'SystemPusherMonitor' });

type PusherMonitorTab = 'channels' | 'connections';

const refreshIntervalOptions = [10, 30, 60].map((value) => ({
  label: `${value}s`,
  value,
}));

const appsLoading = ref(false);
const overviewLoading = ref(false);
const channelLoading = ref(false);
const channelDetailLoading = ref(false);
const connectionLoading = ref(false);
const terminateLoading = ref(false);
const debugSubmitting = ref(false);

const apps = ref<MonitorApi.PusherMonitorAppItem[]>([]);
const activeAppId = ref('');
const activeTab = ref<PusherMonitorTab>('channels');
const autoRefresh = ref(true);
const refreshInterval = ref(10);
const lastRefreshAt = ref('');

const overviewData = ref<MonitorApi.PusherMonitorOverview | null>(null);

const channelKeyword = ref('');
const activeChannelType = ref('');
const channelItems = ref<MonitorApi.PusherMonitorChannelItem[]>([]);
const activeChannelName = ref('');
const channelDetail = ref<MonitorApi.PusherMonitorChannelDetail | null>(null);
const channelPagination = ref({
  current: 1,
  pageSize: 10,
  total: 0,
  pageSizeOptions: [10, 20, 50],
  showJumper: true,
  showPageSize: true,
});

const connectionSocketId = ref('');
const connectionUserId = ref('');
const connectionChannel = ref('');
const connectionItems = ref<MonitorApi.PusherMonitorConnectionItem[]>([]);
const selectedConnectionKeys = ref<Array<number | string>>([]);
const connectionPagination = ref({
  current: 1,
  pageSize: 20,
  total: 0,
  pageSizeOptions: [10, 20, 50],
  showJumper: true,
  showPageSize: true,
});

const debugDialogVisible = ref(false);
const debugForm = ref({
  name: 'server-event',
  channelsText: '',
  data: '{}',
  socketId: '',
});

const refreshTimer = ref<ReturnType<typeof setInterval>>();

let initialized = false;

const appOptions = computed(() =>
  apps.value.map((item) => ({
    label: `${item.appName} (${item.appId})`,
    value: item.appId,
  })),
);

const selectedApp = computed(
  () => apps.value.find((item) => item.appId === activeAppId.value) ?? null,
);

const channelTypeOptions = computed(() => [
  { label: $t('common.all'), value: '' },
  { label: $t('system.monitor.pusher.channelTypes.public'), value: 'public' },
  { label: $t('system.monitor.pusher.channelTypes.private'), value: 'private' },
  { label: $t('system.monitor.pusher.channelTypes.presence'), value: 'presence' },
  { label: $t('system.monitor.pusher.channelTypes.encrypted'), value: 'encrypted' },
]);

const summaryCards = computed(() => {
  const overview = overviewData.value;
  if (!overview) {
    return [];
  }
  return [
    {
      label: $t('system.monitor.pusher.connectionCount'),
      value: overview.connectionCount,
      tone: 'text-sky-600',
    },
    {
      label: $t('system.monitor.pusher.activeChannelCount'),
      value: overview.activeChannelCount,
      tone: 'text-emerald-600',
    },
    {
      label: $t('system.monitor.pusher.subscriptionCount'),
      value: overview.subscriptionCount,
      tone: 'text-cyan-600',
    },
    {
      label: $t('system.monitor.pusher.presenceChannelCount'),
      value: overview.presenceChannelCount,
      tone: 'text-violet-600',
    },
    {
      label: $t('system.monitor.pusher.presenceUserCount'),
      value: overview.presenceUserCount,
      tone: 'text-orange-600',
    },
    {
      label: $t('system.monitor.pusher.recentEventCount'),
      value: overview.recentEventCount,
      tone: 'text-primary',
    },
    {
      label: $t('system.monitor.pusher.recentErrorCount'),
      value: overview.recentErrorCount,
      tone: 'text-rose-600',
    },
  ];
});

const activeChannelSockets = computed(() => channelDetail.value?.socketIds ?? []);
const activeChannelUsers = computed(() => channelDetail.value?.userIds ?? []);

const channelColumns = computed(() => [
  { colKey: 'name', title: $t('system.monitor.pusher.channelName'), minWidth: 280 },
  { colKey: 'channelType', title: $t('system.monitor.pusher.channelType'), width: 140 },
  { colKey: 'subscriptionCount', title: $t('system.monitor.pusher.subscriptionCount'), width: 130 },
  { colKey: 'userCount', title: $t('system.monitor.pusher.userCount'), width: 120 },
  { colKey: 'serverCount', title: $t('system.monitor.pusher.serverCount'), width: 120 },
  { colKey: 'action', title: $t('common.action'), width: 120, align: 'center' as const },
]);

const connectionColumns = computed(() => [
  { colKey: 'row-select', type: 'multiple' as const, width: 52 },
  { colKey: 'socketId', title: $t('system.monitor.pusher.socketId'), minWidth: 190 },
  { colKey: 'userId', title: $t('system.monitor.pusher.userId'), minWidth: 160 },
  { colKey: 'serverName', title: $t('system.monitor.pusher.serverName'), width: 160 },
  { colKey: 'channelCount', title: $t('system.monitor.pusher.channelCount'), width: 120 },
  { colKey: 'channels', title: $t('system.monitor.pusher.channels'), minWidth: 260 },
  { colKey: 'addr', title: $t('system.monitor.pusher.address'), minWidth: 180 },
  { colKey: 'connectedAt', title: $t('system.monitor.pusher.connectedAt'), width: 180 },
  { colKey: 'lastHeartbeatAt', title: $t('system.monitor.pusher.lastHeartbeatAt'), width: 180 },
]);

function channelTypeTheme(channelType: string) {
  switch (channelType) {
    case 'encrypted': {
      return 'warning';
    }
    case 'presence': {
      return 'success';
    }
    case 'private': {
      return 'primary';
    }
    default: {
      return 'default';
    }
  }
}

function channelPreview(channels: string[]) {
  if (!channels || channels.length === 0) {
    return [];
  }
  return channels.slice(0, 3);
}

function restChannelCount(channels: string[]) {
  return channels.length > 3 ? channels.length - 3 : 0;
}

function formatRefreshTime(date = new Date()) {
  return formatQueueMonitorTime(date);
}

function stopAutoRefresh() {
  if (refreshTimer.value) {
    clearInterval(refreshTimer.value);
    refreshTimer.value = undefined;
  }
}

function restartAutoRefresh() {
  stopAutoRefresh();
  if (!autoRefresh.value) {
    return;
  }
  refreshTimer.value = setInterval(() => {
    void refreshAll();
  }, refreshInterval.value * 1000);
}

function parseDebugChannels(text: string) {
  return text
    .split(/[\n,]/)
    .map((item) => item.trim())
    .filter(Boolean);
}

async function fetchApps() {
  appsLoading.value = true;
  try {
    apps.value = await getPusherMonitorApps();
    if (!activeAppId.value && apps.value.length > 0) {
      activeAppId.value = apps.value[0]?.appId ?? '';
      return;
    }
    if (activeAppId.value && !apps.value.some((item) => item.appId === activeAppId.value)) {
      activeAppId.value = apps.value[0]?.appId ?? '';
    }
  } catch (error) {
    logger.error(error);
    message.error($t('common.pusherMonitorAppsFailed'));
  } finally {
    appsLoading.value = false;
  }
}

async function fetchOverview(options?: { reset?: boolean }) {
  if (!activeAppId.value) {
    overviewData.value = null;
    return;
  }
  if (options?.reset) {
    overviewData.value = null;
    lastRefreshAt.value = '';
  }
  overviewLoading.value = true;
  try {
    overviewData.value = await getPusherMonitorOverview(activeAppId.value);
    lastRefreshAt.value = formatRefreshTime();
  } catch (error) {
    logger.error(error);
    if (options?.reset) {
      overviewData.value = null;
    }
    message.error($t('common.pusherMonitorOverviewFailed'));
  } finally {
    overviewLoading.value = false;
  }
}

async function fetchChannelDetail(channelName = activeChannelName.value) {
  if (!activeAppId.value || !channelName) {
    channelDetail.value = null;
    return;
  }
  channelDetailLoading.value = true;
  try {
    channelDetail.value = await getPusherMonitorChannelDetail(activeAppId.value, channelName);
  } catch (error) {
    logger.error(error);
    channelDetail.value = null;
    message.error($t('common.pusherMonitorChannelDetailFailed'));
  } finally {
    channelDetailLoading.value = false;
  }
}

async function fetchChannels() {
  if (!activeAppId.value) {
    channelItems.value = [];
    channelDetail.value = null;
    channelPagination.value.total = 0;
    return;
  }
  channelLoading.value = true;
  try {
    const response = await getPusherMonitorChannels({
      appId: activeAppId.value,
      keyword: channelKeyword.value || undefined,
      channelType: activeChannelType.value || undefined,
      page: channelPagination.value.current,
      pageSize: channelPagination.value.pageSize,
    });
    channelItems.value = response.items ?? [];
    channelPagination.value.total = Number(response.pageInfo?.total || response.total || 0);

    const preferredChannel =
      channelItems.value.find((item) => item.name === activeChannelName.value)?.name
      ?? channelItems.value[0]?.name
      ?? '';

    if (!preferredChannel) {
      activeChannelName.value = '';
      channelDetail.value = null;
      return;
    }

    if (preferredChannel !== activeChannelName.value) {
      activeChannelName.value = preferredChannel;
    }
    await fetchChannelDetail(preferredChannel);
  } catch (error) {
    logger.error(error);
    channelItems.value = [];
    channelDetail.value = null;
    channelPagination.value.total = 0;
    message.error($t('common.pusherMonitorChannelsFailed'));
  } finally {
    channelLoading.value = false;
  }
}

async function fetchConnections() {
  if (!activeAppId.value) {
    connectionItems.value = [];
    connectionPagination.value.total = 0;
    return;
  }
  connectionLoading.value = true;
  try {
    const response = await getPusherMonitorConnections({
      appId: activeAppId.value,
      socketId: connectionSocketId.value || undefined,
      userId: connectionUserId.value || undefined,
      channel: connectionChannel.value || undefined,
      page: connectionPagination.value.current,
      pageSize: connectionPagination.value.pageSize,
    });
    connectionItems.value = response.items ?? [];
    connectionPagination.value.total = Number(response.pageInfo?.total || response.total || 0);
  } catch (error) {
    logger.error(error);
    connectionItems.value = [];
    connectionPagination.value.total = 0;
    message.error($t('common.pusherMonitorConnectionsFailed'));
  } finally {
    connectionLoading.value = false;
  }
}

async function refreshAll(options?: { reset?: boolean }) {
  await Promise.all([
    fetchOverview(options),
    fetchChannels(),
    fetchConnections(),
  ]);
}

async function handleSelectChannel(channelName: string) {
  if (!channelName) {
    return;
  }
  activeChannelName.value = channelName;
  await fetchChannelDetail(channelName);
}

function handleChannelPageChange(pageInfo: { current: number; pageSize: number }) {
  channelPagination.value.current = pageInfo.current;
  channelPagination.value.pageSize = pageInfo.pageSize;
  void fetchChannels();
}

function handleConnectionPageChange(pageInfo: { current: number; pageSize: number }) {
  connectionPagination.value.current = pageInfo.current;
  connectionPagination.value.pageSize = pageInfo.pageSize;
  void fetchConnections();
}

function handleConnectionSelectChange(keys: Array<number | string>) {
  selectedConnectionKeys.value = keys;
}

function handleChannelSearch() {
  channelPagination.value.current = 1;
  void fetchChannels();
}

function handleChannelReset() {
  channelKeyword.value = '';
  activeChannelType.value = '';
  channelPagination.value.current = 1;
  void fetchChannels();
}

function handleConnectionSearch() {
  connectionPagination.value.current = 1;
  void fetchConnections();
}

function handleConnectionReset() {
  connectionSocketId.value = '';
  connectionUserId.value = '';
  connectionChannel.value = '';
  connectionPagination.value.current = 1;
  void fetchConnections();
}

async function handleManualRefresh() {
  await refreshAll();
}

function openDebugDialog(channelName?: string) {
  debugForm.value = {
    name: 'server-event',
    channelsText: channelName || activeChannelName.value || '',
    data: '{}',
    socketId: '',
  };
  debugDialogVisible.value = true;
}

function closeDebugDialog() {
  debugDialogVisible.value = false;
}

function handleDebugDialogVisibleChange(value: boolean) {
  debugDialogVisible.value = value;
}

async function handleSubmitDebugEvent() {
  if (!activeAppId.value) {
    return;
  }
  const channels = parseDebugChannels(debugForm.value.channelsText);
  if (channels.length === 0) {
    message.warning($t('system.monitor.pusher.targetChannelsRequired'));
    return;
  }

  debugSubmitting.value = true;
  try {
    await debugPusherMonitorEvent(activeAppId.value, {
      name: debugForm.value.name,
      channels,
      data: debugForm.value.data,
      socketId: debugForm.value.socketId || undefined,
    });
    message.success($t('common.executeSuccess'));
    debugDialogVisible.value = false;
    await refreshAll();
  } catch (error) {
    logger.error(error);
    message.error($t('common.pusherMonitorDebugFailed'));
  } finally {
    debugSubmitting.value = false;
  }
}

async function handleTerminateConnections(socketIds: string[]) {
  if (!activeAppId.value || socketIds.length === 0) {
    return;
  }
  terminateLoading.value = true;
  try {
    await terminatePusherMonitorConnections(activeAppId.value, { socketIds });
    selectedConnectionKeys.value = [];
    message.success($t('common.operationSuccess'));
    await refreshAll();
  } catch (error) {
    logger.error(error);
    message.error($t('common.pusherMonitorTerminateFailed'));
  } finally {
    terminateLoading.value = false;
  }
}

async function handleTerminateUsers(userIds: string[]) {
  if (!activeAppId.value || userIds.length === 0) {
    return;
  }
  terminateLoading.value = true;
  try {
    await terminatePusherMonitorUsers(activeAppId.value, { userIds });
    selectedConnectionKeys.value = [];
    message.success($t('common.operationSuccess'));
    await refreshAll();
  } catch (error) {
    logger.error(error);
    message.error($t('common.pusherMonitorTerminateFailed'));
  } finally {
    terminateLoading.value = false;
  }
}

function handleBatchTerminateSelectedConnections() {
  const socketIds = selectedConnectionKeys.value.map(String);
  void handleTerminateConnections(socketIds);
}

function handleTerminateActiveChannelConnections() {
  void handleTerminateConnections(activeChannelSockets.value);
}

function handleTerminateActiveChannelUsers() {
  void handleTerminateUsers(activeChannelUsers.value);
}

watch(activeAppId, async () => {
  if (!initialized) {
    return;
  }
  channelPagination.value.current = 1;
  connectionPagination.value.current = 1;
  selectedConnectionKeys.value = [];
  activeChannelName.value = '';
  channelDetail.value = null;
  await refreshAll({ reset: true });
});

watch([autoRefresh, refreshInterval], () => {
  restartAutoRefresh();
});

onMounted(async () => {
  await fetchApps();
  await refreshAll({ reset: true });
  initialized = true;
  restartAutoRefresh();
});

onUnmounted(() => {
  stopAutoRefresh();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex flex-col gap-4">
      <Card class="border-border/70 bg-background/95 shadow-sm">
        <div class="flex flex-col gap-4 xl:flex-row xl:items-center xl:justify-between">
          <div class="min-w-0">
            <div class="flex items-center gap-2 text-[11px] uppercase tracking-[0.28em] text-muted-foreground/70">
              <span class="h-2 w-2 rounded-full bg-cyan-500"></span>
              <span>Pusher Protocol</span>
            </div>
            <div class="mt-2 flex flex-wrap items-center gap-3">
              <div class="text-2xl font-semibold text-foreground">
                {{ $t('system.monitor.pusher.title') }}
              </div>
              <Tag theme="primary" variant="light">
                {{ selectedApp?.appName || '-' }}
              </Tag>
              <Tag v-if="selectedApp?.groupName" theme="default" variant="light">
                {{ selectedApp?.groupName }}
              </Tag>
            </div>
            <div class="mt-2 text-sm text-muted-foreground">
              {{ $t('system.monitor.pusher.description') }}
            </div>
            <div class="mt-2 flex flex-wrap gap-2 text-xs text-muted-foreground">
              <span>{{ $t('system.monitor.pusher.appId') }}: {{ selectedApp?.appId || '-' }}</span>
              <span>{{ $t('system.monitor.pusher.appKey') }}: {{ selectedApp?.appKey || '-' }}</span>
            </div>
          </div>

          <div class="grid gap-3 md:grid-cols-2 2xl:flex 2xl:flex-wrap 2xl:items-center 2xl:justify-end">
            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.pusher.app') }}
              </div>
              <Select
                v-model="activeAppId"
                :loading="appsLoading"
                class="w-full min-w-60"
                :options="appOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.pusher.autoRefresh') }}
              </div>
              <div class="flex items-center justify-between rounded-lg border border-border/70 bg-muted/30 px-3 py-2">
                <span class="text-sm text-muted-foreground">{{ $t('system.monitor.pusher.autoRefresh') }}</span>
                <Switch v-model="autoRefresh" size="small" />
              </div>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.pusher.refreshInterval') }}
              </div>
              <Select
                v-model="refreshInterval"
                class="w-full min-w-32"
                :options="refreshIntervalOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.pusher.manualRefresh') }}
              </div>
              <Button block theme="primary" @click="handleManualRefresh">
                {{ $t('common.refresh') }}
              </Button>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.pusher.lastRefreshAt') }}
              </div>
              <div class="rounded-lg border border-border/70 bg-muted/20 px-3 py-2 text-sm text-muted-foreground">
                <span class="font-medium text-foreground">{{ lastRefreshAt || '-' }}</span>
              </div>
            </div>
          </div>
        </div>
      </Card>

      <div
        v-if="summaryCards.length > 0"
        class="grid grid-cols-1 gap-3 sm:grid-cols-2 xl:grid-cols-4 2xl:grid-cols-7"
      >
        <Card
          v-for="card in summaryCards"
          :key="card.label"
          class="border-border/70 bg-background/90"
          size="small"
        >
          <div class="text-xs uppercase tracking-wide text-muted-foreground/70">
            {{ card.label }}
          </div>
          <div class="mt-2 text-xl font-semibold" :class="card.tone">
            {{ card.value }}
          </div>
        </Card>
      </div>

      <Card v-if="appOptions.length === 0">
        <div class="py-12 text-center text-muted-foreground">
          {{ $t('common.noData') }}
        </div>
      </Card>

      <Card v-else class="border-border/70">
        <Tabs v-model:value="activeTab" class="min-h-[520px]">
          <TabPanel value="channels" :label="$t('system.monitor.pusher.channelsTab')">
            <div class="mt-4 grid gap-4 xl:grid-cols-[minmax(0,1.45fr)_360px]">
              <div class="flex min-h-0 flex-col gap-4">
                <div class="grid gap-3 md:grid-cols-[minmax(0,1fr)_180px_120px_120px]">
                  <Input
                    v-model="channelKeyword"
                    :placeholder="$t('system.monitor.pusher.searchChannelPlaceholder')"
                    clearable
                  />
                  <Select
                    v-model="activeChannelType"
                    :options="channelTypeOptions"
                  />
                  <Button theme="default" @click="handleChannelReset">
                    {{ $t('common.reset') }}
                  </Button>
                  <Button theme="primary" @click="handleChannelSearch">
                    {{ $t('common.query') }}
                  </Button>
                </div>

                <Table
                  row-key="name"
                  hover
                  stripe
                  :columns="channelColumns"
                  :data="channelItems"
                  :loading="channelLoading || overviewLoading"
                  :pagination="channelPagination"
                  @page-change="handleChannelPageChange"
                >
                  <template #channelType="{ row }">
                    <Tag :theme="channelTypeTheme(row.channelType)" variant="light">
                      {{ $t(`system.monitor.pusher.channelTypes.${row.channelType}`) }}
                    </Tag>
                  </template>
                  <template #action="{ row }">
                    <Button
                      size="small"
                      :theme="row.name === activeChannelName ? 'primary' : 'default'"
                      :variant="row.name === activeChannelName ? 'base' : 'outline'"
                      @click="handleSelectChannel(row.name)"
                    >
                      {{ row.name === activeChannelName ? $t('system.monitor.pusher.activeChannel') : $t('common.detail') }}
                    </Button>
                  </template>
                </Table>
              </div>

              <Card
                :title="$t('system.monitor.pusher.channelDetail')"
                class="border-border/70 bg-muted/10"
                size="small"
              >
                <div v-if="channelDetail" class="flex flex-col gap-4">
                  <div class="rounded-xl border border-border/70 bg-background px-4 py-3">
                    <div class="text-sm font-medium text-foreground">
                      {{ channelDetail.channel.name }}
                    </div>
                    <div class="mt-2 flex flex-wrap gap-2">
                      <Tag :theme="channelTypeTheme(channelDetail.channel.channelType)" variant="light">
                        {{ $t(`system.monitor.pusher.channelTypes.${channelDetail.channel.channelType}`) }}
                      </Tag>
                      <Tag theme="primary" variant="light">
                        {{ $t('system.monitor.pusher.subscriptionCount') }} {{ channelDetail.channel.subscriptionCount }}
                      </Tag>
                      <Tag theme="success" variant="light">
                        {{ $t('system.monitor.pusher.userCount') }} {{ channelDetail.channel.userCount }}
                      </Tag>
                      <Tag theme="default" variant="light">
                        {{ $t('system.monitor.pusher.serverCount') }} {{ channelDetail.channel.serverCount }}
                      </Tag>
                    </div>
                  </div>

                  <div class="rounded-xl border border-border/70 bg-background px-4 py-3">
                    <div class="mb-2 text-sm font-medium text-foreground">
                      {{ $t('system.monitor.pusher.quickActions') }}
                    </div>
                    <div class="grid gap-2">
                      <Button theme="primary" variant="outline" @click="openDebugDialog(channelDetail.channel.name)">
                        {{ $t('system.monitor.pusher.debugEvent') }}
                      </Button>

                      <Popconfirm
                        v-if="activeChannelSockets.length > 0"
                        :content="$t('system.monitor.pusher.confirmTerminateChannelConnections')"
                        @confirm="handleTerminateActiveChannelConnections"
                      >
                        <Button :loading="terminateLoading" theme="danger" variant="outline">
                          {{ $t('system.monitor.pusher.disconnectChannelConnections') }} ({{ activeChannelSockets.length }})
                        </Button>
                      </Popconfirm>
                      <Button v-else disabled theme="danger" variant="outline">
                        {{ $t('system.monitor.pusher.disconnectChannelConnections') }}
                      </Button>

                      <Popconfirm
                        v-if="activeChannelUsers.length > 0"
                        :content="$t('system.monitor.pusher.confirmTerminateChannelUsers')"
                        @confirm="handleTerminateActiveChannelUsers"
                      >
                        <Button :loading="terminateLoading" theme="warning" variant="outline">
                          {{ $t('system.monitor.pusher.disconnectChannelUsers') }} ({{ activeChannelUsers.length }})
                        </Button>
                      </Popconfirm>
                      <Button v-else disabled theme="warning" variant="outline">
                        {{ $t('system.monitor.pusher.disconnectChannelUsers') }}
                      </Button>
                    </div>
                  </div>

                  <div class="rounded-xl border border-border/70 bg-background px-4 py-3">
                    <div class="mb-2 text-sm font-medium text-foreground">
                      {{ $t('system.monitor.pusher.serverNames') }}
                    </div>
                    <div class="flex flex-wrap gap-2">
                      <Tag
                        v-for="serverName in channelDetail.channel.serverNames || []"
                        :key="serverName"
                        theme="default"
                        variant="light"
                      >
                        {{ serverName }}
                      </Tag>
                      <span
                        v-if="(channelDetail.channel.serverNames || []).length === 0"
                        class="text-sm text-muted-foreground"
                      >
                        {{ $t('common.noData') }}
                      </span>
                    </div>
                  </div>

                  <div class="grid gap-4 md:grid-cols-2 xl:grid-cols-1 2xl:grid-cols-2">
                    <div class="rounded-xl border border-border/70 bg-background px-4 py-3">
                      <div class="mb-2 text-sm font-medium text-foreground">
                        {{ $t('system.monitor.pusher.channelSockets') }}
                      </div>
                      <div class="max-h-40 overflow-auto">
                        <div class="flex flex-wrap gap-2">
                          <Tag
                            v-for="socketId in channelDetail.socketIds"
                            :key="socketId"
                            theme="primary"
                            variant="light"
                          >
                            {{ socketId }}
                          </Tag>
                          <span
                            v-if="channelDetail.socketIds.length === 0"
                            class="text-sm text-muted-foreground"
                          >
                            {{ $t('common.noData') }}
                          </span>
                        </div>
                      </div>
                    </div>

                    <div class="rounded-xl border border-border/70 bg-background px-4 py-3">
                      <div class="mb-2 text-sm font-medium text-foreground">
                        {{ $t('system.monitor.pusher.channelUsers') }}
                      </div>
                      <div class="max-h-40 overflow-auto">
                        <div class="flex flex-wrap gap-2">
                          <Tag
                            v-for="userId in channelDetail.userIds"
                            :key="userId"
                            theme="success"
                            variant="light"
                          >
                            {{ userId }}
                          </Tag>
                          <span
                            v-if="channelDetail.userIds.length === 0"
                            class="text-sm text-muted-foreground"
                          >
                            {{ $t('common.noData') }}
                          </span>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>

                <div
                  v-else
                  class="flex min-h-[420px] items-center justify-center text-sm text-muted-foreground"
                >
                  {{ channelDetailLoading ? $t('common.loading') : $t('system.monitor.pusher.selectPrompt') }}
                </div>
              </Card>
            </div>
          </TabPanel>

          <TabPanel value="connections" :label="$t('system.monitor.pusher.connectionsTab')">
            <div class="mt-4 flex flex-col gap-4">
              <div class="grid gap-3 xl:grid-cols-[minmax(0,1fr)_minmax(0,1fr)_minmax(0,1fr)_120px_120px_auto]">
                <Input
                  v-model="connectionSocketId"
                  :placeholder="$t('system.monitor.pusher.searchSocketPlaceholder')"
                  clearable
                />
                <Input
                  v-model="connectionUserId"
                  :placeholder="$t('system.monitor.pusher.searchUserPlaceholder')"
                  clearable
                />
                <Input
                  v-model="connectionChannel"
                  :placeholder="$t('system.monitor.pusher.searchConnectionChannelPlaceholder')"
                  clearable
                />
                <Button theme="default" @click="handleConnectionReset">
                  {{ $t('common.reset') }}
                </Button>
                <Button theme="primary" @click="handleConnectionSearch">
                  {{ $t('common.query') }}
                </Button>
                <div class="flex justify-start xl:justify-end">
                  <Popconfirm
                    v-if="selectedConnectionKeys.length > 0"
                    :content="$t('system.monitor.pusher.confirmTerminateSelectedConnections')"
                    @confirm="handleBatchTerminateSelectedConnections"
                  >
                    <Button :loading="terminateLoading" theme="danger">
                      {{ $t('system.monitor.pusher.batchDisconnect') }} ({{ selectedConnectionKeys.length }})
                    </Button>
                  </Popconfirm>
                  <Button v-else disabled theme="danger">
                    {{ $t('system.monitor.pusher.batchDisconnect') }}
                  </Button>
                </div>
              </div>

              <Table
                row-key="socketId"
                hover
                stripe
                :columns="connectionColumns"
                :data="connectionItems"
                :loading="connectionLoading"
                :pagination="connectionPagination"
                :selected-row-keys="selectedConnectionKeys"
                @page-change="handleConnectionPageChange"
                @select-change="handleConnectionSelectChange"
              >
                <template #userId="{ row }">
                  <span>{{ row.userId || '-' }}</span>
                </template>
                <template #addr="{ row }">
                  <span>{{ row.addr || '-' }}</span>
                </template>
                <template #connectedAt="{ row }">
                  {{ formatQueueMonitorTime(row.connectedAt) || '-' }}
                </template>
                <template #lastHeartbeatAt="{ row }">
                  {{ formatQueueMonitorTime(row.lastHeartbeatAt) || '-' }}
                </template>
                <template #channels="{ row }">
                  <div class="flex flex-wrap gap-2">
                    <Tag
                      v-for="channel in channelPreview(row.channels)"
                      :key="`${row.socketId}-${channel}`"
                      size="small"
                      theme="default"
                      variant="light"
                    >
                      {{ channel }}
                    </Tag>
                    <Tag
                      v-if="restChannelCount(row.channels) > 0"
                      size="small"
                      theme="primary"
                      variant="light"
                    >
                      +{{ restChannelCount(row.channels) }}
                    </Tag>
                    <span v-if="row.channels.length === 0" class="text-sm text-muted-foreground">-</span>
                  </div>
                </template>
              </Table>
            </div>
          </TabPanel>
        </Tabs>

        <div class="mt-4 text-xs text-muted-foreground">
          {{ $t('system.monitor.pusher.runtimeNotice') }}
        </div>
      </Card>
    </div>

    <Dialog
      :visible="debugDialogVisible"
      width="720px"
      :header="$t('system.monitor.pusher.eventDebugTitle')"
      destroy-on-close
      @update:visible="handleDebugDialogVisibleChange"
    >
      <div class="grid gap-4 py-2">
        <div class="grid gap-2">
          <div class="text-sm font-medium text-foreground">{{ $t('system.monitor.pusher.eventName') }}</div>
          <Input v-model="debugForm.name" />
        </div>

        <div class="grid gap-2">
          <div class="text-sm font-medium text-foreground">{{ $t('system.monitor.pusher.targetChannels') }}</div>
          <Textarea
            v-model="debugForm.channelsText"
            :placeholder="$t('system.monitor.pusher.targetChannelsPlaceholder')"
            :autosize="{ minRows: 3, maxRows: 6 }"
          />
        </div>

        <div class="grid gap-2">
          <div class="text-sm font-medium text-foreground">{{ $t('system.monitor.pusher.eventData') }}</div>
          <Textarea
            v-model="debugForm.data"
            :autosize="{ minRows: 6, maxRows: 12 }"
          />
        </div>

        <div class="grid gap-2">
          <div class="text-sm font-medium text-foreground">{{ $t('system.monitor.pusher.excludeSocketId') }}</div>
          <Input
            v-model="debugForm.socketId"
            :placeholder="$t('system.monitor.pusher.excludeSocketPlaceholder')"
          />
        </div>

        <div class="rounded-xl border border-border/70 bg-muted/20 px-4 py-3 text-sm text-muted-foreground">
          {{ $t('system.monitor.pusher.debugTips') }}
        </div>
      </div>

      <template #footer>
        <div class="flex justify-end gap-2">
          <Button theme="default" @click="closeDebugDialog">{{ $t('common.cancel') }}</Button>
          <Button theme="primary" :loading="debugSubmitting" @click="handleSubmitDebugEvent">
            {{ $t('common.execute') }}
          </Button>
        </div>
      </template>
    </Dialog>
  </Page>
</template>
