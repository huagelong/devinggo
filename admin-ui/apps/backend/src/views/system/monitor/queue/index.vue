<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { computed, onMounted, onUnmounted, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import {
  getQueueMonitorOverview,
  getQueueMonitorQueues,
  getQueueMonitorSchedulerEntries,
  getQueueMonitorTasks,
} from '#/api/system/monitor';
import { logger } from '#/utils/logger';
import { formatQueueMonitorTime } from '#/utils/queue-monitor-time';

import {
  Button,
  Card,
  Select,
  Switch,
  TabPanel,
  Table,
  Tabs,
  Tag,
} from 'tdesign-vue-next';

defineOptions({ name: 'SystemQueueMonitor' });

type QueueMonitorTab = 'overview' | 'tasks' | 'schedule' | 'workers';

const refreshIntervalOptions = [10, 30, 60].map((value) => ({
  label: `${value}s`,
  value,
}));

const taskStateOptions = [
  'pending',
  'active',
  'scheduled',
  'retry',
  'archived',
  'completed',
].map((value) => ({
  label: $t(`system.monitor.queue.taskStates.${value}`),
  value,
}));

const loading = ref(false);
const taskLoading = ref(false);
const scheduleLoading = ref(false);
const queues = ref<MonitorApi.QueueMonitorQueueItem[]>([]);
const activeQueue = ref('');
const activeTab = ref<QueueMonitorTab>('overview');
const activeTaskState = ref('scheduled');
const autoRefresh = ref(true);
const refreshInterval = ref(10);
const lastRefreshAt = ref('');
const overviewData = ref<MonitorApi.QueueMonitorOverviewResponse | null>(null);
const schedulerEntries = ref<MonitorApi.QueueMonitorSchedulerEntry[]>([]);
const taskItems = ref<MonitorApi.QueueMonitorTaskItem[]>([]);
const taskPagination = ref({
  current: 1,
  pageSize: 20,
  total: 0,
  pageSizeOptions: [10, 20, 50],
  showJumper: true,
  showPageSize: true,
});
const refreshTimer = ref<ReturnType<typeof setInterval>>();

let initialized = false;

const queueOptions = computed(() =>
  queues.value.map((item) => ({
    label: item.queue,
    value: item.queue,
  })),
);

const currentQueue = computed(() => overviewData.value?.queue ?? null);

const summaryCards = computed(() => {
  const queue = currentQueue.value;
  if (!queue) {
    return [];
  }
  return [
    { label: $t('system.monitor.queue.pending'), value: queue.pending, tone: 'text-sky-600' },
    { label: $t('system.monitor.queue.active'), value: queue.active, tone: 'text-emerald-600' },
    { label: $t('system.monitor.queue.scheduled'), value: queue.scheduled, tone: 'text-violet-600' },
    { label: $t('system.monitor.queue.retry'), value: queue.retry, tone: 'text-amber-600' },
    { label: $t('system.monitor.queue.archived'), value: queue.archived, tone: 'text-rose-600' },
    { label: $t('system.monitor.queue.latency'), value: formatSeconds(queue.latency), tone: 'text-cyan-600' },
  ];
});

const overviewItems = computed(() => {
  const queue = currentQueue.value;
  if (!queue) {
    return [];
  }
  return [
    { label: $t('system.monitor.queue.queueName'), value: queue.queue },
    { label: $t('system.monitor.queue.size'), value: queue.size },
    { label: $t('system.monitor.queue.groups'), value: queue.groups },
    { label: $t('system.monitor.queue.memoryUsage'), value: formatBytes(queue.memoryUsage) },
    { label: $t('system.monitor.queue.processedToday'), value: queue.processed },
    { label: $t('system.monitor.queue.failedToday'), value: queue.failed },
    { label: $t('system.monitor.queue.processedTotal'), value: queue.processedTotal },
    { label: $t('system.monitor.queue.failedTotal'), value: queue.failedTotal },
    { label: $t('system.monitor.queue.completed'), value: queue.completed },
    { label: $t('system.monitor.queue.aggregating'), value: queue.aggregating },
    { label: $t('system.monitor.queue.paused'), value: queue.paused ? $t('common.yes') : $t('common.no') },
    { label: $t('system.monitor.queue.snapshotAt'), value: formatQueueMonitorTime(queue.timestamp) || '-' },
  ];
});

const filteredSchedulerEntries = computed(() => {
  if (!activeQueue.value) {
    return schedulerEntries.value;
  }
  return schedulerEntries.value.filter((item) => item.queue === activeQueue.value);
});

const workerItems = computed(() => overviewData.value?.workers ?? []);

const taskColumns = computed(() => [
  { colKey: 'type', title: $t('system.monitor.queue.taskType'), minWidth: 220 },
  { colKey: 'id', title: 'ID', minWidth: 220 },
  { colKey: 'state', title: $t('common.status'), width: 140 },
  { colKey: 'nextProcessAt', title: $t('system.monitor.queue.nextProcessAt'), width: 180 },
  { colKey: 'retried', title: $t('system.monitor.queue.retried'), width: 100 },
  { colKey: 'maxRetry', title: $t('system.monitor.queue.maxRetry'), width: 100 },
  { colKey: 'payload', title: $t('system.monitor.queue.payload'), minWidth: 280 },
  { colKey: 'lastErr', title: $t('system.monitor.queue.lastErr'), minWidth: 220 },
]);

const scheduleColumns = computed(() => [
  { colKey: 'taskType', title: $t('system.monitor.queue.taskType'), minWidth: 220 },
  { colKey: 'queue', title: $t('system.monitor.queue.queueName'), width: 140 },
  { colKey: 'spec', title: $t('system.monitor.queue.scheduleSpec'), minWidth: 180 },
  { colKey: 'next', title: $t('system.monitor.queue.nextEnqueueAt'), width: 180 },
  { colKey: 'prev', title: $t('system.monitor.queue.prevEnqueueAt'), width: 180 },
]);

const workerColumns = computed(() => [
  { colKey: 'host', title: $t('system.monitor.queue.workerHost'), minWidth: 180 },
  { colKey: 'pid', title: 'PID', width: 100 },
  { colKey: 'concurrency', title: $t('system.monitor.queue.concurrency'), width: 120 },
  { colKey: 'status', title: $t('common.status'), width: 120 },
  { colKey: 'activeWorkerCount', title: $t('system.monitor.queue.activeWorkerCount'), width: 140 },
  { colKey: 'queues', title: $t('system.monitor.queue.queueWeights'), minWidth: 220 },
  { colKey: 'started', title: $t('system.monitor.queue.startedAt'), width: 180 },
]);

function formatBytes(bytes: number): string {
  if (!bytes) {
    return '0 B';
  }
  const units = ['B', 'KB', 'MB', 'GB', 'TB'];
  let value = bytes;
  let index = 0;
  while (value >= 1024 && index < units.length - 1) {
    value /= 1024;
    index++;
  }
  return `${value.toFixed(value >= 10 || index === 0 ? 0 : 2)} ${units[index]}`;
}

function formatSeconds(seconds: number): string {
  if (!seconds) {
    return '0s';
  }
  if (seconds < 60) {
    return `${seconds}s`;
  }
  const minutes = Math.floor(seconds / 60);
  const remainSeconds = seconds % 60;
  return `${minutes}m ${remainSeconds}s`;
}

function formatRefreshTime(date = new Date()) {
  return formatQueueMonitorTime(date);
}

function queueWeightsLabel(queuesMap: Record<string, number>) {
  return Object.entries(queuesMap || {})
    .map(([name, weight]) => `${name}:${weight}`)
    .join(', ');
}

function payloadLabel(payload: string) {
  if (!payload) {
    return '-';
  }
  try {
    return JSON.stringify(JSON.parse(payload), null, 2);
  } catch {
    return payload;
  }
}

function taskStateTheme(state: string) {
  switch (state) {
    case 'active': {
      return 'success';
    }
    case 'scheduled': {
      return 'primary';
    }
    case 'retry': {
      return 'warning';
    }
    case 'archived': {
      return 'danger';
    }
    case 'completed': {
      return 'success';
    }
    default: {
      return 'default';
    }
  }
}

async function fetchQueues() {
  try {
    queues.value = await getQueueMonitorQueues();
    if (!activeQueue.value && queues.value.length > 0) {
      const firstQueue = queues.value[0];
      if (firstQueue) {
        activeQueue.value = firstQueue.queue;
      }
    }
  } catch (error) {
    logger.error(error);
    message.error($t('common.queueMonitorQueuesFailed'));
  }
}

async function fetchOverview(options?: { reset?: boolean }) {
  if (!activeQueue.value) {
    overviewData.value = null;
    return;
  }
  if (options?.reset) {
    overviewData.value = null;
  }
  loading.value = true;
  try {
    overviewData.value = await getQueueMonitorOverview(activeQueue.value);
    lastRefreshAt.value = formatRefreshTime();
  } catch (error) {
    logger.error(error);
    if (options?.reset) {
      overviewData.value = null;
    }
    message.error($t('common.queueMonitorOverviewFailed'));
  } finally {
    loading.value = false;
  }
}

async function fetchTasks() {
  if (!activeQueue.value) {
    taskItems.value = [];
    taskPagination.value.total = 0;
    return;
  }
  taskLoading.value = true;
  try {
    const response = await getQueueMonitorTasks({
      queue: activeQueue.value,
      state: activeTaskState.value,
      page: taskPagination.value.current,
      pageSize: taskPagination.value.pageSize,
    });
    taskItems.value = response.items ?? [];
    taskPagination.value.total = Number(response.pageInfo?.total || response.total || 0);
  } catch (error) {
    logger.error(error);
    message.error($t('common.queueMonitorTasksFailed'));
  } finally {
    taskLoading.value = false;
  }
}

async function fetchSchedulerEntries() {
  scheduleLoading.value = true;
  try {
    schedulerEntries.value = await getQueueMonitorSchedulerEntries();
  } catch (error) {
    logger.error(error);
    message.error($t('common.queueMonitorSchedulerFailed'));
  } finally {
    scheduleLoading.value = false;
  }
}

async function refreshAll(options?: { reset?: boolean }) {
  await Promise.all([
    fetchOverview(options),
    fetchTasks(),
    fetchSchedulerEntries(),
  ]);
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

function handleTaskPageChange(pageInfo: { current: number; pageSize: number }) {
  taskPagination.value.current = pageInfo.current;
  taskPagination.value.pageSize = pageInfo.pageSize;
  void fetchTasks();
}

async function handleManualRefresh() {
  await refreshAll();
}

watch(activeQueue, async () => {
  if (!initialized) {
    return;
  }
  taskPagination.value.current = 1;
  await refreshAll({ reset: true });
});

watch(activeTaskState, async () => {
  if (!initialized) {
    return;
  }
  taskPagination.value.current = 1;
  await fetchTasks();
});

watch([autoRefresh, refreshInterval], () => {
  restartAutoRefresh();
});

onMounted(async () => {
  await fetchQueues();
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
              <span class="h-2 w-2 rounded-full bg-sky-500" />
              <span>Asynq</span>
            </div>
            <div class="mt-2 flex flex-wrap items-center gap-3">
              <div class="text-2xl font-semibold text-foreground">
                {{ $t('system.monitor.queue.title') }}
              </div>
              <Tag theme="primary" variant="light">
                {{ activeQueue || '-' }}
              </Tag>
            </div>
            <div class="mt-2 text-sm text-muted-foreground">
              {{ $t('system.monitor.queue.description') }}
            </div>
          </div>

          <div class="grid gap-3 md:grid-cols-2 2xl:flex 2xl:flex-wrap 2xl:items-center 2xl:justify-end">
            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.queue.queueName') }}
              </div>
              <Select
                v-model="activeQueue"
                class="w-full min-w-52"
                :options="queueOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.queue.autoRefresh') }}
              </div>
              <div class="flex items-center justify-between rounded-lg border border-border/70 bg-muted/30 px-3 py-2">
                <span class="text-sm text-muted-foreground">{{ $t('system.monitor.queue.autoRefresh') }}</span>
                <Switch v-model="autoRefresh" size="small" />
              </div>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.queue.refreshInterval') }}
              </div>
              <Select
                v-model="refreshInterval"
                class="w-full min-w-32"
                :options="refreshIntervalOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.queue.manualRefresh') }}
              </div>
              <Button block theme="primary" @click="handleManualRefresh">
                {{ $t('common.refresh') }}
              </Button>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.queue.lastRefreshAt') }}
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
        class="grid grid-cols-1 gap-3 sm:grid-cols-2 xl:grid-cols-6"
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

      <Card v-if="queueOptions.length === 0">
        <div class="py-12 text-center text-muted-foreground">
          {{ $t('common.noData') }}
        </div>
      </Card>

      <Card v-else class="border-border/70">
        <Tabs v-model:value="activeTab" class="min-h-[420px]">
          <TabPanel value="overview" :label="$t('system.monitor.queue.overview')">
            <div
              v-if="overviewItems.length > 0"
              class="mt-4 grid grid-cols-1 gap-3 md:grid-cols-2 xl:grid-cols-3"
            >
              <div
                v-for="item in overviewItems"
                :key="item.label"
                class="rounded-xl border border-border/70 bg-muted/30 px-4 py-3"
              >
                <div class="text-xs uppercase tracking-wide text-muted-foreground/70">
                  {{ item.label }}
                </div>
                <div class="mt-2 text-base font-medium text-foreground">
                  {{ item.value }}
                </div>
              </div>
            </div>
            <div v-else class="py-12 text-center text-muted-foreground">
              {{ loading ? $t('common.loading') : $t('common.noData') }}
            </div>
          </TabPanel>

          <TabPanel value="tasks" :label="$t('system.monitor.queue.tasks')">
            <div class="mt-4 flex flex-col gap-4">
              <div class="flex flex-col gap-3 md:flex-row md:items-center md:justify-between">
                <div class="text-sm text-muted-foreground">
                  {{ $t('system.monitor.queue.taskState') }}
                </div>
                <Select
                  v-model="activeTaskState"
                  class="w-full md:w-56"
                  :options="taskStateOptions"
                />
              </div>

              <Table
                row-key="id"
                hover
                stripe
                :columns="taskColumns"
                :data="taskItems"
                :loading="taskLoading"
                :pagination="taskPagination"
                @page-change="handleTaskPageChange"
              >
                <template #state="{ row }">
                  <Tag :theme="taskStateTheme(row.state)" variant="light">
                    {{ $t(`system.monitor.queue.taskStates.${row.state}`) }}
                  </Tag>
                </template>
                <template #payload="{ row }">
                  <div class="max-w-[360px] whitespace-pre-wrap break-all text-xs text-muted-foreground">
                    {{ payloadLabel(row.payload) }}
                  </div>
                </template>
                <template #lastErr="{ row }">
                  <span :title="row.lastErr || '-'">{{ row.lastErr || '-' }}</span>
                </template>
                <template #nextProcessAt="{ row }">
                  {{ formatQueueMonitorTime(row.nextProcessAt) || '-' }}
                </template>
              </Table>
            </div>
          </TabPanel>

          <TabPanel value="schedule" :label="$t('system.monitor.queue.schedule')">
            <div class="mt-4 flex flex-col gap-4">
              <Table
                row-key="id"
                hover
                stripe
                :columns="scheduleColumns"
                :data="filteredSchedulerEntries"
                :loading="scheduleLoading"
              >
                <template #next="{ row }">
                  {{ formatQueueMonitorTime(row.next) || '-' }}
                </template>
                <template #prev="{ row }">
                  {{ formatQueueMonitorTime(row.prev) || '-' }}
                </template>
              </Table>
              <div
                v-if="!scheduleLoading && filteredSchedulerEntries.length === 0"
                class="rounded-xl border border-dashed border-border bg-muted/10 px-4 py-10 text-center text-muted-foreground"
              >
                {{ $t('common.noData') }}
              </div>
            </div>
          </TabPanel>

          <TabPanel value="workers" :label="$t('system.monitor.queue.workers')">
            <div class="mt-4 flex flex-col gap-4">
              <Table
                row-key="id"
                hover
                stripe
                :columns="workerColumns"
                :data="workerItems"
                :loading="loading"
              >
                <template #status="{ row }">
                  <Tag :theme="row.status === 'active' ? 'success' : 'warning'" variant="light">
                    {{ row.status || '-' }}
                  </Tag>
                </template>
                <template #queues="{ row }">
                  {{ queueWeightsLabel(row.queues) || '-' }}
                </template>
                <template #started="{ row }">
                  {{ formatQueueMonitorTime(row.started) || '-' }}
                </template>
              </Table>
              <div
                v-if="!loading && workerItems.length === 0"
                class="rounded-xl border border-dashed border-border bg-muted/10 px-4 py-10 text-center text-muted-foreground"
              >
                {{ $t('common.noData') }}
              </div>
            </div>
          </TabPanel>
        </Tabs>
      </Card>
    </div>
  </Page>
</template>
