<script lang="ts" setup>
import type { DataMaintainApi } from '#/api/system/data-maintain';
import type { MonitorApi } from '#/api/system/monitor';

import { computed, nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { EchartsUI, useEcharts } from '@vben/plugins/echarts';
import type { EchartsUIType } from '@vben/plugins/echarts';

import { message } from '#/adapter/tdesign';
import { getDataMaintainDetailed, getDataMaintainPageList } from '#/api/system/data-maintain';
import { getDbMonitorGroups, getDbMonitorInfo } from '#/api/system/monitor';
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
import type { TableRowData } from 'tdesign-vue-next';

defineOptions({ name: 'SystemDbMonitor' });

type DbMonitorTab = 'capabilities' | 'overview' | 'tables' | 'timeseries';

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
const tableInfoLoading = ref(false);
const tableInfoList = ref<DataMaintainApi.ListItem[]>([]);
const expandedTableKeys = ref<Array<string | number>>([]);
const tableDetailLoadingMap = ref<Record<string, boolean>>({});
const tableDetailMap = ref<Record<string, DataMaintainApi.ColumnItem[]>>({});
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
const tableCount = computed(() => tableInfoList.value.length);
const timeseries = computed(() => monitorData.value?.timeseries ?? []);

const tableColumns = computed(() => [
  {
    colKey: 'name',
    title: $t('system.dataMaintain.tableName'),
    minWidth: 220,
  },
  {
    colKey: 'comment',
    title: $t('system.dataMaintain.tableComment'),
    minWidth: 260,
  },
  {
    colKey: 'engine',
    title: $t('system.dataMaintain.engine'),
    width: 140,
  },
  {
    colKey: 'collation',
    title: $t('system.dataMaintain.collation'),
    width: 180,
  },
  {
    colKey: 'rows',
    title: $t('system.dataMaintain.rows'),
    width: 120,
  },
  {
    colKey: 'data_length',
    title: $t('system.dataMaintain.tableSize'),
    width: 150,
  },
]);

const detailTableColumns = computed(() => [
  { colKey: 'field', title: $t('system.dataMaintain.fieldColName'), width: 220 },
  { colKey: 'type', title: $t('system.dataMaintain.fieldColType'), width: 180 },
  { colKey: 'nullable', title: $t('system.dataMaintain.fieldColNullable'), width: 120 },
  { colKey: 'default_value', title: $t('system.dataMaintain.fieldColDefault'), minWidth: 180 },
  { colKey: 'key', title: $t('system.dataMaintain.fieldColKey'), width: 140 },
  { colKey: 'comment', title: $t('system.dataMaintain.fieldColComment'), minWidth: 220 },
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

function formatRefreshTime(date = new Date()) {
  return formatQueueMonitorTime(date);
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

function fieldKeyLabel(key?: string) {
  switch (key) {
    case 'PRI': {
      return $t('system.dataMaintain.keyPrimary');
    }
    case 'UNI': {
      return $t('system.dataMaintain.keyUnique');
    }
    case 'MUL': {
      return $t('system.dataMaintain.keyIndex');
    }
    default: {
      return $t('system.dataMaintain.keyNone');
    }
  }
}

function fieldKeyTheme(key?: string) {
  switch (key) {
    case 'PRI': {
      return 'danger';
    }
    case 'UNI': {
      return 'success';
    }
    case 'MUL': {
      return 'primary';
    }
    default: {
      return 'default';
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

async function fetchTableInfo(options?: { reset?: boolean }) {
  if (!activeGroup.value) {
    return;
  }

  if (options?.reset) {
    tableInfoList.value = [];
    expandedTableKeys.value = [];
    tableDetailMap.value = {};
    tableDetailLoadingMap.value = {};
  }

  tableInfoLoading.value = true;
  try {
    const response = await getDataMaintainPageList({
      group_name: activeGroup.value,
      page: 1,
      pageSize: 500,
    });
    tableInfoList.value = response.items ?? [];
  } catch (error) {
    logger.error(error);
    if (options?.reset) {
      tableInfoList.value = [];
    }
    message.error($t('common.listLoadFailed'));
  } finally {
    tableInfoLoading.value = false;
  }
}

async function loadTableDetail(tableName: string) {
  if (!activeGroup.value || tableDetailMap.value[tableName] || tableDetailLoadingMap.value[tableName]) {
    return;
  }

  tableDetailLoadingMap.value = {
    ...tableDetailLoadingMap.value,
    [tableName]: true,
  };
  try {
    const response = await getDataMaintainDetailed({
      group_name: activeGroup.value,
      table_name: tableName,
    });
    tableDetailMap.value = {
      ...tableDetailMap.value,
      [tableName]: response.items ?? [],
    };
  } catch (error) {
    logger.error(error);
    message.error($t('common.fieldDetailFailed'));
  } finally {
    tableDetailLoadingMap.value = {
      ...tableDetailLoadingMap.value,
      [tableName]: false,
    };
  }
}

function handleTableExpandChange(
  expandedRowKeys: Array<string | number>,
  options: { currentRowData: TableRowData },
) {
  expandedTableKeys.value = expandedRowKeys;
  const currentTableName = (options.currentRowData as DataMaintainApi.ListItem | undefined)?.name;
  if (!currentTableName || !expandedRowKeys.includes(currentTableName)) {
    return;
  }
  void loadTableDetail(currentTableName);
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
  await Promise.all([
    fetchMonitor(),
    fetchTableInfo(),
  ]);
}

watch(activeGroup, async () => {
  if (!initialized) {
    return;
  }
  await Promise.all([
    fetchMonitor({ reset: true }),
    fetchTableInfo({ reset: true }),
  ]);
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
  await Promise.all([
    fetchMonitor({ reset: true }),
    fetchTableInfo({ reset: true }),
  ]);
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
              <span class="h-2 w-2 rounded-full bg-emerald-500" />
              <span>PostgreSQL</span>
            </div>
            <div class="mt-2 flex flex-wrap items-center gap-3">
              <div class="text-2xl font-semibold text-foreground">
                {{ $t('system.monitor.db.title') }}
              </div>
              <Tag theme="primary" variant="light">
                {{ activeGroup || '-' }}
              </Tag>
            </div>
            <div class="mt-2 text-sm text-muted-foreground">
              {{ $t('system.monitor.db.title') }}
              · PostgreSQL
            </div>
          </div>

          <div class="grid gap-3 md:grid-cols-2 2xl:flex 2xl:flex-wrap 2xl:items-center 2xl:justify-end">
            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.db.group') }}
              </div>
              <Select
                v-model="activeGroup"
                class="w-full min-w-52"
                :options="groupOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.db.autoRefresh') }}
              </div>
              <div class="flex items-center justify-between rounded-lg border border-border/70 bg-muted/30 px-3 py-2">
                <span class="text-sm text-muted-foreground">{{ $t('system.monitor.db.autoRefresh') }}</span>
                <Switch v-model="autoRefresh" size="small" />
              </div>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.db.refreshInterval') }}
              </div>
              <Select
                v-model="refreshInterval"
                class="w-full min-w-32"
                :options="refreshIntervalOptions"
              />
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.db.manualRefresh') }}
              </div>
              <Button block theme="primary" @click="handleManualRefresh">
                {{ $t('common.refresh') }}
              </Button>
            </div>

            <div class="flex flex-col gap-2">
              <div class="text-xs font-medium text-muted-foreground">
                {{ $t('system.monitor.db.lastRefreshAt') }}
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

          <TabPanel value="tables" :label="$t('system.monitor.db.tables')">
            <div class="mt-4 flex flex-col gap-4">
              <div class="flex items-center justify-between rounded-xl border border-border/70 bg-muted/20 px-4 py-3">
                <div>
                  <div class="text-sm font-medium text-foreground">
                    {{ $t('system.monitor.db.tables') }}
                  </div>
                  <div class="mt-1 text-sm text-muted-foreground">
                    {{ $t('system.monitor.db.group') }}: {{ activeGroup || '-' }}
                  </div>
                </div>
                <Tag theme="primary" variant="light">
                  {{ tableCount }} {{ $t('system.monitor.db.tableCountUnit') }}
                </Tag>
              </div>

              <Table
                row-key="name"
                expand-on-row-click
                hover
                stripe
                :columns="tableColumns"
                :data="tableInfoList"
                :expanded-row-keys="expandedTableKeys"
                :loading="tableInfoLoading"
                size="small"
                @expand-change="handleTableExpandChange"
              >
                <template #comment="{ row }">
                  <span :title="row.comment || '-'">{{ row.comment || '-' }}</span>
                </template>
                <template #data_length="{ row }">
                  {{ formatBytes(row.data_length ?? 0) }}
                </template>
                <template #expandedRow="{ row }">
                  <div class="rounded-xl border border-border/70 bg-muted/10 p-4">
                    <div class="mb-3 grid grid-cols-2 gap-3 text-sm text-muted-foreground xl:grid-cols-5">
                      <div>{{ $t('system.dataMaintain.engine') }}：{{ row.engine || '-' }}</div>
                      <div>{{ $t('system.dataMaintain.collation') }}：{{ row.collation || '-' }}</div>
                      <div>{{ $t('system.dataMaintain.rows') }}：{{ row.rows ?? '-' }}</div>
                      <div>{{ $t('system.dataMaintain.tableSize') }}：{{ formatBytes(row.data_length ?? 0) }}</div>
                      <div>{{ $t('common.updateTime') }}：{{ formatQueueMonitorTime(row.update_time) || '-' }}</div>
                    </div>

                    <Table
                      row-key="field"
                      size="small"
                      :bordered="false"
                      :columns="detailTableColumns"
                      :data="tableDetailMap[row.name] ?? []"
                      :loading="!!tableDetailLoadingMap[row.name]"
                    >
                      <template #nullable="{ row: columnRow }">
                        <Tag :theme="columnRow.nullable ? 'success' : 'default'" variant="light">
                          {{ columnRow.nullable ? $t('system.dataMaintain.yes') : $t('system.dataMaintain.no') }}
                        </Tag>
                      </template>
                      <template #default_value="{ row: columnRow }">
                        {{ columnRow.default_value || '-' }}
                      </template>
                      <template #key="{ row: columnRow }">
                        <Tag :theme="fieldKeyTheme(columnRow.key)" variant="light">
                          {{ fieldKeyLabel(columnRow.key) }}
                        </Tag>
                      </template>
                      <template #comment="{ row: columnRow }">
                        {{ columnRow.comment || '-' }}
                      </template>
                    </Table>
                  </div>
                </template>
              </Table>

              <div
                v-if="!tableInfoLoading && tableInfoList.length === 0"
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
