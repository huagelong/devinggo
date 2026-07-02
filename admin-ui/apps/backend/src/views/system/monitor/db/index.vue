<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { EchartsUI, useEcharts } from '@vben/plugins/echarts';
import type { EchartsUIType } from '@vben/plugins/echarts';

import { message } from '#/adapter/tdesign';
import { getDbMonitorGroups, getDbMonitorInfo } from '#/api/system/monitor';
import { logger } from '#/utils/logger';

import {
  Button,
  Card,
  Select,
  Space,
  Switch,
  TabPanel,
  Tabs,
  Tag,
} from 'tdesign-vue-next';

defineOptions({ name: 'SystemDbMonitor' });

type DbMonitorTab = 'capabilities' | 'overview' | 'timeseries';

const refreshIntervalOptions = [10, 30, 60].map((value) => ({
  label: `${value}s`,
  value,
}));

const loading = ref(false);
const groups = ref<MonitorApi.DbMonitorGroup[]>([]);
const activeGroup = ref('');
const activeTab = ref<DbMonitorTab>('overview');
const autoRefresh = ref(true);
const refreshInterval = ref(10);
const lastRefreshAt = ref('');
const monitorData = ref<MonitorApi.DbMonitorResponse | null>(null);
const refreshTimer = ref<ReturnType<typeof setInterval>>();

const tpsChartRef = ref<EchartsUIType>();
const latencyChartRef = ref<EchartsUIType>();
const { renderEcharts: renderTpsChart } = useEcharts(tpsChartRef);
const { renderEcharts: renderLatencyChart } = useEcharts(latencyChartRef);

let initialized = false;

const groupOptions = computed(() =>
  groups.value.map((item) => ({
    label: item.groupName,
    value: item.groupName,
  })),
);

const summaryCards = computed(() => {
  const summary = monitorData.value?.summary;
  if (!summary) {
    return [];
  }
  return [
    {
      label: $t('system.monitor.db.connectionStatus'),
      value: summary.connectionStatus,
      tone: 'text-emerald-600',
    },
    {
      label: $t('system.monitor.db.currentConn'),
      value: summary.currentConn,
      tone: 'text-sky-600',
    },
    {
      label: $t('system.monitor.db.activeConn'),
      value: summary.activeConn,
      tone: 'text-orange-600',
    },
    {
      label: $t('system.monitor.db.databaseSize'),
      value: formatBytes(summary.databaseSize),
      tone: 'text-violet-600',
    },
    {
      label: $t('system.monitor.db.hitRate'),
      value: `${summary.hitRate.toFixed(2)}%`,
      tone: 'text-cyan-600',
    },
    {
      label: $t('system.monitor.db.capabilities'),
      value: summary.slowQueryEnabled ? $t('common.enabled') : $t('common.disabled'),
      tone: summary.slowQueryEnabled ? 'text-primary' : 'text-muted-foreground',
    },
  ];
});

const overviewItems = computed(() => {
  const overview = monitorData.value?.overview;
  if (!overview) {
    return [];
  }
  return [
    { label: $t('system.monitor.db.group'), value: overview.groupName },
    { label: $t('system.monitor.db.databaseName'), value: overview.databaseName },
    { label: $t('system.monitor.db.version'), value: overview.version },
    { label: $t('system.monitor.db.maxConnections'), value: overview.maxConnections },
    { label: $t('system.monitor.db.currentConn'), value: overview.currentConnections },
    { label: $t('system.monitor.db.activeConn'), value: overview.activeConnections },
    { label: $t('system.monitor.db.idleConnections'), value: overview.idleConnections },
    { label: $t('system.monitor.db.waitingConnections'), value: overview.waitingConnections },
    { label: $t('system.monitor.db.xactCommit'), value: overview.xactCommit },
    { label: $t('system.monitor.db.xactRollback'), value: overview.xactRollback },
    { label: $t('system.monitor.db.hitRate'), value: `${overview.hitRate.toFixed(2)}%` },
    { label: $t('system.monitor.db.databaseSize'), value: formatBytes(overview.databaseSize) },
  ];
});

const capabilityData = computed(() => monitorData.value?.capabilities ?? null);
const alerts = computed(() => monitorData.value?.alerts ?? []);
const timeseries = computed(() => monitorData.value?.timeseries ?? []);

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

function formatRefreshTime(date = new Date()) {
  return date.toLocaleString();
}

function metricLabel(metric: string) {
  switch (metric) {
    case 'avgQueryMs': {
      return $t('system.monitor.db.avgQueryMs');
    }
    case 'connections': {
      return $t('system.monitor.db.currentConn');
    }
    case 'hitRate': {
      return $t('system.monitor.db.hitRate');
    }
    case 'rollbackTps': {
      return $t('system.monitor.db.rollbackTps');
    }
    case 'slowQueryCount': {
      return $t('system.monitor.db.slowQueryCount');
    }
    case 'tps': {
      return $t('system.monitor.db.tps');
    }
    default: {
      return metric;
    }
  }
}

function alertTheme(level: string) {
  switch (level) {
    case 'error': {
      return 'danger';
    }
    case 'warning': {
      return 'warning';
    }
    default: {
      return 'primary';
    }
  }
}

function renderCharts() {
  const points = timeseries.value;
  renderTpsChart({
    color: ['#0ea5e9', '#f97316'],
    tooltip: { trigger: 'axis' },
    legend: {
      bottom: 0,
      data: [$t('system.monitor.db.tps'), $t('system.monitor.db.rollbackTps')],
    },
    grid: { top: 24, right: 16, bottom: 48, left: 40 },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: points.map((item) => item.time),
    },
    yAxis: { type: 'value' },
    series: [
      {
        type: 'line',
        smooth: true,
        symbol: 'none',
        name: $t('system.monitor.db.tps'),
        data: points.map((item) => Number(item.tps.toFixed(2))),
      },
      {
        type: 'line',
        smooth: true,
        symbol: 'none',
        name: $t('system.monitor.db.rollbackTps'),
        data: points.map((item) => Number(item.rollbackTps.toFixed(2))),
      },
    ],
  });

  renderLatencyChart({
    color: ['#8b5cf6', '#22c55e'],
    tooltip: { trigger: 'axis' },
    legend: {
      bottom: 0,
      data: [$t('system.monitor.db.avgQueryMs'), $t('system.monitor.db.slowQueryCount')],
    },
    grid: { top: 24, right: 16, bottom: 48, left: 40 },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: points.map((item) => item.time),
    },
    yAxis: [
      { type: 'value', name: 'ms' },
      { type: 'value', name: 'count' },
    ],
    series: [
      {
        type: 'line',
        smooth: true,
        symbol: 'none',
        name: $t('system.monitor.db.avgQueryMs'),
        data: points.map((item) => item.avgQueryMs == null ? null : Number(item.avgQueryMs.toFixed(2))),
      },
      {
        type: 'bar',
        name: $t('system.monitor.db.slowQueryCount'),
        yAxisIndex: 1,
        data: points.map((item) => item.slowQueryCount ?? null),
      },
    ],
  });
}

async function fetchGroups() {
  try {
    groups.value = await getDbMonitorGroups();
    if (!activeGroup.value && groups.value.length > 0) {
      const firstGroup = groups.value[0];
      if (firstGroup) {
        activeGroup.value = groups.value.find((item) => item.isDefault)?.groupName ?? firstGroup.groupName;
      }
    }
  } catch (error) {
    logger.error(error);
    message.error($t('common.dbMonitorGroupsFailed'));
  }
}

async function fetchMonitor(options?: { reset?: boolean }) {
  if (!activeGroup.value) {
    return;
  }

  if (options?.reset) {
    monitorData.value = null;
    lastRefreshAt.value = '';
  }

  loading.value = true;
  try {
    monitorData.value = await getDbMonitorInfo(activeGroup.value);
    lastRefreshAt.value = formatRefreshTime();
    if (activeTab.value === 'timeseries') {
      await nextTick();
      renderCharts();
    }
  } catch (error) {
    logger.error(error);
    if (options?.reset) {
      monitorData.value = null;
    }
    message.error($t('common.dbMonitorInfoFailed'));
  } finally {
    loading.value = false;
  }
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
    void fetchMonitor();
  }, refreshInterval.value * 1000);
}

async function handleManualRefresh() {
  await fetchMonitor();
}

watch(activeGroup, async () => {
  if (!initialized) {
    return;
  }
  await fetchMonitor({ reset: true });
});

watch([autoRefresh, refreshInterval], () => {
  restartAutoRefresh();
});

watch(activeTab, async (value) => {
  if (value !== 'timeseries' || timeseries.value.length === 0) {
    return;
  }
  await nextTick();
  renderCharts();
});

onMounted(async () => {
  await fetchGroups();
  await fetchMonitor({ reset: true });
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
      <Card class="overflow-hidden border-0 bg-gradient-to-r from-slate-900 via-slate-800 to-cyan-900 text-white shadow-sm">
        <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <div class="text-xs uppercase tracking-[0.28em] text-white/60">
              PostgreSQL
            </div>
            <div class="mt-2 text-2xl font-semibold">
              {{ $t('system.monitor.db.title') }}
            </div>
            <div class="mt-1 text-sm text-white/70">
              {{ activeGroup || '-' }}
            </div>
          </div>
          <div class="flex flex-col gap-3 sm:flex-row sm:flex-wrap sm:items-center sm:justify-end">
            <Select
              v-model="activeGroup"
              class="w-full min-w-40 sm:w-44"
              :options="groupOptions"
            />
            <Space align="center">
              <span class="text-sm text-white/70">{{ $t('common.refresh') }}</span>
              <Switch v-model="autoRefresh" size="small" />
            </Space>
            <Select
              v-model="refreshInterval"
              class="w-full min-w-24 sm:w-28"
              :options="refreshIntervalOptions"
            />
            <Button theme="primary" @click="handleManualRefresh">
              {{ $t('common.refresh') }}
            </Button>
            <Tag theme="primary" variant="light-outline">
              {{ $t('system.monitor.db.lastRefreshAt') }}: {{ lastRefreshAt || '-' }}
            </Tag>
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

      <Card v-if="groupOptions.length === 0">
        <div class="py-12 text-center text-muted-foreground">
          {{ $t('common.noData') }}
        </div>
      </Card>

      <Card v-else class="border-border/70">
        <Tabs v-model:value="activeTab" class="min-h-[420px]">
          <TabPanel value="overview" :label="$t('system.monitor.db.overview')">
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

          <TabPanel value="timeseries" :label="$t('system.monitor.db.timeseries')">
            <div
              v-if="timeseries.length === 0"
              class="mt-4 flex min-h-80 items-center justify-center rounded-xl border border-dashed border-border bg-muted/20 text-muted-foreground"
            >
              {{ $t('system.monitor.db.dataCollecting') }}
            </div>
            <div v-else class="mt-4 grid grid-cols-1 gap-4 xl:grid-cols-2">
              <Card :title="$t('system.monitor.db.tps')" size="small">
                <EchartsUI ref="tpsChartRef" class="h-80" />
              </Card>
              <Card :title="$t('system.monitor.db.avgQueryMs')" size="small">
                <div
                  v-if="!capabilityData?.hasPgStatStatements"
                  class="flex h-80 items-center justify-center rounded-lg border border-dashed border-border bg-muted/20 text-sm text-muted-foreground"
                >
                  {{ $t('system.monitor.db.pgStatStatementsMissing') }}
                </div>
                <EchartsUI v-else ref="latencyChartRef" class="h-80" />
              </Card>
            </div>
          </TabPanel>

          <TabPanel value="capabilities" :label="$t('system.monitor.db.capabilities')">
            <div v-if="capabilityData" class="mt-4 flex flex-col gap-4">
              <Card size="small">
                <div class="flex flex-wrap items-center gap-3">
                  <Tag
                    :theme="capabilityData.hasPgStatStatements ? 'success' : 'warning'"
                    variant="light"
                  >
                    pg_stat_statements: {{ capabilityData.hasPgStatStatements ? 'ON' : 'OFF' }}
                  </Tag>
                  <Tag theme="primary" variant="light">
                    threshold: {{ capabilityData.slowQueryThreshold }}ms
                  </Tag>
                </div>
              </Card>

              <Card size="small">
                <div class="text-sm font-medium">
                  {{ $t('common.enabled') }}
                </div>
                <div class="mt-3 flex flex-wrap gap-2">
                  <Tag
                    v-for="metric in capabilityData.availableMetrics"
                    :key="metric"
                    theme="success"
                    variant="light"
                  >
                    {{ metricLabel(metric) }}
                  </Tag>
                </div>
              </Card>

              <Card size="small">
                <div class="text-sm font-medium">
                  {{ $t('common.disabled') }}
                </div>
                <div
                  v-if="capabilityData.unavailableMetrics.length > 0"
                  class="mt-3 flex flex-wrap gap-2"
                >
                  <Tag
                    v-for="metric in capabilityData.unavailableMetrics"
                    :key="metric"
                    theme="warning"
                    variant="light"
                  >
                    {{ metricLabel(metric) }}
                  </Tag>
                </div>
                <div v-else class="mt-3 text-sm text-muted-foreground">
                  {{ $t('common.noData') }}
                </div>
              </Card>

              <Card
                v-if="!capabilityData.hasPgStatStatements || alerts.length > 0"
                size="small"
              >
                <div class="flex flex-col gap-2">
                  <Tag
                    v-if="!capabilityData.hasPgStatStatements"
                    theme="warning"
                    variant="light"
                  >
                    {{ $t('common.dbMonitorCapabilityMissing') }}
                  </Tag>
                  <Tag
                    v-for="alert in alerts"
                    :key="`${alert.level}-${alert.message}`"
                    :theme="alertTheme(alert.level)"
                    variant="light"
                  >
                    {{ alert.message }}
                  </Tag>
                </div>
              </Card>
            </div>
            <div v-else class="py-12 text-center text-muted-foreground">
              {{ loading ? $t('common.loading') : $t('common.noData') }}
            </div>
          </TabPanel>
        </Tabs>
      </Card>
    </div>
  </Page>
</template>
