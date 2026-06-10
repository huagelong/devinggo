<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { computed, onMounted, ref, watch } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';
import { EchartsUI, useEcharts } from '@vben/plugins/echarts';
import type { EchartsUIType } from '@vben/plugins/echarts';

import { message } from '#/adapter/tdesign';
import {
  clearAllCache,
  deleteCacheKey,
  getCacheInfo,
  viewCache,
} from '#/api/system/monitor';
import { logger } from '#/utils/logger';

import { BrowseIcon, DeleteIcon, SearchIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  Card,
  Input,
  Popconfirm,
  Space,
  Table,
} from 'tdesign-vue-next';
import type { PrimaryTableCol, TableRowData } from 'tdesign-vue-next/es/table/type';

defineOptions({ name: 'SystemCache' });

const loading = ref(false);
const serverInfo = ref<MonitorApi.CacheServerInfo>({});
const cacheKeys = ref<string[]>([]);
const selectedKeys = ref<string[]>([]);
const searchKey = ref('');
const cacheContent = ref('');

const memoryChartRef = ref<EchartsUIType>();
const { renderEcharts: renderMemoryChart } = useEcharts(memoryChartRef);

interface StatCardItem {
  label: string;
  value: string | number;
  unit?: string;
  color?: string;
}

const statCards = computed<StatCardItem[]>(() => {
  const s = serverInfo.value;
  return [
    { label: $t('system.monitor.cache.redisVersion'), value: s.version || '-', color: 'text-blue-600' },
    { label: $t('system.monitor.cache.runMode'), value: s.redis_mode || '-', color: 'text-purple-600' },
    { label: $t('system.monitor.cache.port'), value: s.port || '-', color: 'text-gray-600' },
    { label: $t('system.monitor.cache.runDays'), value: s.run_days || '-', unit: '天', color: 'text-orange-600' },
    { label: $t('system.monitor.cache.clientConnections'), value: s.clients || '-', color: 'text-green-600' },
    { label: $t('system.monitor.cache.aofStatus'), value: s.aof_enabled || '-', color: 'text-red-600' },
    { label: $t('system.monitor.cache.systemUsedKeys'), value: s.sys_total_keys ?? '-', color: 'text-indigo-600' },
    { label: $t('system.monitor.cache.expiredKeys'), value: s.expired_keys || '-', color: 'text-yellow-600' },
    { label: $t('system.monitor.cache.qps'), value: s.qps || '-', unit: 'ops/s', color: 'text-cyan-600' },
    { label: $t('system.monitor.cache.hitRate'), value: s.hit_rate != null ? `${s.hit_rate.toFixed(2)}%` : '-', color: 'text-emerald-600' },
    { label: $t('system.monitor.cache.totalCommands'), value: s.total_commands || '-', color: 'text-pink-600' },
    { label: $t('system.monitor.cache.blockedClients'), value: s.blocked_clients || '-', color: 'text-rose-600' },
    { label: $t('system.monitor.cache.rejectedConn'), value: s.rejected_conn || '-', color: 'text-amber-600' },
    { label: $t('system.monitor.cache.memoryPeak'), value: s.memory_peak || '-', color: 'text-violet-600' },
    { label: $t('system.monitor.cache.memFragmentRatio'), value: s.mem_fragment_ratio || '-', color: 'text-teal-600' },
  ];
});

const memoryValue = computed(() => {
  const val = serverInfo.value.use_memory;
  if (!val) return 0;
  const num = parseFloat(val);
  if (Number.isNaN(num)) return 0;
  return num;
});

const memoryUnit = computed(() => {
  const val = serverInfo.value.use_memory;
  if (!val) return 'MB';
  const match = val.match(/[a-zA-Z]+$/);
  return match ? match[0].toUpperCase() : 'MB';
});

function renderMemoryGauge() {
  const value = memoryValue.value;
  const unit = memoryUnit.value;
  let displayValue = value;
  let displayUnit = unit;

  if (unit === 'B' || (!unit && value > 1024 * 1024 * 1024)) {
    displayValue = value / 1024 / 1024 / 1024;
    displayUnit = 'GB';
  } else if (unit === 'K' || unit === 'KB' || (!unit && value > 1024 * 1024)) {
    displayValue = value / 1024 / 1024;
    displayUnit = 'MB';
  }

  const maxValue = displayValue > 0 ? Math.ceil(displayValue * 2.5) : 10;

  renderMemoryChart({
    series: [
      {
        type: 'gauge',
        startAngle: 200,
        endAngle: -20,
        min: 0,
        max: maxValue,
        splitNumber: 10,
        itemStyle: {
          color: '#5ab1ef',
        },
        progress: {
          show: true,
          width: 12,
        },
        pointer: {
          show: true,
          width: 4,
          length: '60%',
        },
        axisLine: {
          lineStyle: {
            width: 12,
          },
        },
        axisTick: {
          distance: -18,
          splitNumber: 5,
          lineStyle: {
            width: 1,
            color: '#999',
          },
        },
        splitLine: {
          distance: -22,
          length: 8,
          lineStyle: {
            width: 2,
            color: '#999',
          },
        },
        axisLabel: {
          distance: -12,
          color: '#999',
          fontSize: 10,
        },
        anchor: {
          show: true,
          size: 15,
          itemStyle: {
            borderColor: '#5ab1ef',
            borderWidth: 2,
          },
        },
        title: {
          show: true,
          offsetCenter: [0, '70%'],
          fontSize: 12,
          color: '#666',
        },
        detail: {
          valueAnimation: true,
          fontSize: 24,
          fontWeight: 'bold',
          offsetCenter: [0, '40%'],
          formatter: `{value}`,
          color: '#333',
        },
        data: [
          {
            value: Number(displayValue.toFixed(2)),
            name: `Redis${$t('system.monitor.cache.memoryUsage') || '占用内存'} (${displayUnit})`,
          },
        ],
      },
    ],
  });
}

watch(() => serverInfo.value.use_memory, () => {
  if (serverInfo.value.use_memory) {
    renderMemoryGauge();
  }
}, { immediate: true });

const columns = [
  { colKey: 'name', title: $t('system.monitor.cache.keyName'), width: 'auto' },
  { colKey: 'action', title: $t('common.action'), width: 150, align: 'right' },
] satisfies PrimaryTableCol<TableRowData>[];

const filteredData = computed(() => {
  if (!searchKey.value) return cacheKeys.value.map((name) => ({ name }));
  const keyword = searchKey.value.toLowerCase();
  return cacheKeys.value
    .filter((name) => name.toLowerCase().includes(keyword))
    .map((name) => ({ name }));
});

const formattedContent = computed(() => {
  if (!cacheContent.value) return '';
  try {
    const parsed = JSON.parse(cacheContent.value);
    return JSON.stringify(parsed, null, 2);
  } catch {
    return cacheContent.value;
  }
});

const isJsonContent = computed(() => {
  if (!cacheContent.value) return false;
  try {
    JSON.parse(cacheContent.value);
    return true;
  } catch {
    return false;
  }
});

async function fetchCacheInfo() {
  loading.value = true;
  try {
    const response = await getCacheInfo();
    serverInfo.value = response.server || {};
    cacheKeys.value = response.keys || [];
  } catch (error) {
    logger.error(error);
    message.error($t('common.cacheInfoFailed'));
  } finally {
    loading.value = false;
  }
}

async function handleViewKey(key: string) {
  try {
    const response = await viewCache({ key });
    cacheContent.value = response?.content || '';
  } catch (error) {
    logger.error(error);
    message.error($t('common.cacheViewFailed'));
  }
}

async function handleDeleteKey(key: string) {
  try {
    await deleteCacheKey({ key });
    message.success($t('common.deleteSuccess'));
    if (cacheContent.value && selectedKeys.value.includes(key)) {
      cacheContent.value = '';
    }
    await fetchCacheInfo();
  } catch (error) {
    logger.error(error);
    message.error($t('common.deleteFailed'));
  }
}

async function handleClearAll() {
  try {
    await clearAllCache();
    message.success($t('common.clearCacheSuccess'));
    cacheContent.value = '';
    await fetchCacheInfo();
  } catch (error) {
    logger.error(error);
    message.error($t('common.clearCacheFailed'));
  }
}

async function handleBatchDelete() {
  if (selectedKeys.value.length === 0) {
    message.warning($t('common.selectCacheFirst'));
    return;
  }

  try {
    for (const key of selectedKeys.value) {
      await deleteCacheKey({ key });
    }
    message.success($t('common.batchDeleteSuccess'));
    selectedKeys.value = [];
    cacheContent.value = '';
    await fetchCacheInfo();
  } catch (error) {
    logger.error(error);
    message.error($t('common.batchDeleteFailed'));
  }
}

function handleSelectChange(value: Array<number | string>) {
  selectedKeys.value = value.map((item) => String(item));
}

function handleCopyContent() {
  if (!cacheContent.value) return;
  navigator.clipboard.writeText(formattedContent.value).then(() => {
    message.success('复制成功');
  }).catch(() => {
    message.error('复制失败');
  });
}

onMounted(() => {
  void fetchCacheInfo();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <!-- Redis Info Panel -->
      <div class="flex gap-3">
        <Card :title="$t('system.monitor.cache.redisInfo')" class="flex-1" size="small">
          <div class="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-5">
            <div
              v-for="(card, index) in statCards"
              :key="index"
              class="flex flex-col justify-center rounded-lg border border-gray-100 bg-gray-50/50 px-3 py-2 transition-colors hover:bg-gray-100/50"
            >
              <span class="text-xs text-gray-400">{{ card.label }}</span>
              <div class="mt-0.5 flex items-baseline gap-1">
                <span class="text-base font-bold" :class="card.color || 'text-gray-700'">{{ card.value }}</span>
                <span v-if="card.unit" class="text-xs text-gray-400">{{ card.unit }}</span>
              </div>
            </div>
          </div>
        </Card>
        <Card :title="$t('system.monitor.cache.memoryUsage') || 'Redis占用内存'" class="w-64 shrink-0" size="small">
          <EchartsUI ref="memoryChartRef" class="h-48" />
        </Card>
      </div>

      <!-- Cache Data Management -->
      <Card :title="$t('system.monitor.cache.dataManagement')" class="flex min-h-0 flex-1 flex-col">
        <template #actions>
          <Space>
            <Button size="small" theme="danger" variant="outline" @click="handleClearAll">
              <template #icon><DeleteIcon /></template>
              {{ $t('system.monitor.cache.clearAll') }}
            </Button>
            <Popconfirm
              v-if="selectedKeys.length > 0"
              :content="$t('common.batchDeleteDataConfirm')"
              @confirm="handleBatchDelete"
            >
              <Button size="small" theme="danger">
                {{ $t('common.batchDelete') }} ({{ selectedKeys.length }})
              </Button>
            </Popconfirm>
          </Space>
        </template>

        <div class="flex h-full gap-4">
          <!-- Left: Cache Key Table -->
          <div class="flex min-h-0 w-2/3 flex-col">
            <Input
              v-model="searchKey"
              :placeholder="$t('system.monitor.cache.searchKeyPlaceholder')"
              clearable
              class="mb-3 w-full"
            >
              <template #prefix><SearchIcon /></template>
            </Input>

            <div class="min-h-0 flex-1 overflow-hidden">
              <Table
                :columns="columns"
                :data="filteredData"
                :loading="loading"
                row-key="name"
                :row-selection="{
                  type: 'checkbox',
                  showCheckedAll: true,
                }"
                hover
                stripe
                @select-change="handleSelectChange"
              >
                <template #action="{ row }">
                  <Space>
                    <Button
                      size="small"
                      theme="primary"
                      variant="outline"
                      @click="handleViewKey(row?.name)"
                    >
                      <template #icon><BrowseIcon /></template>
                      {{ $t('common.detail') }}
                    </Button>
                    <Button
                      size="small"
                      theme="danger"
                      variant="outline"
                      @click="handleDeleteKey(row?.name)"
                    >
                      <template #icon><DeleteIcon /></template>
                      {{ $t('common.delete') }}
                    </Button>
                  </Space>
                </template>
              </Table>
            </div>
          </div>

          <!-- Right: Cache Content -->
          <div class="flex min-h-0 w-1/3 flex-col gap-2">
            <div class="flex items-center justify-between">
              <span class="text-sm font-medium text-gray-600">
                缓存内容
                <span v-if="isJsonContent" class="ml-1 text-xs text-green-500">(JSON)</span>
              </span>
              <Button
                v-if="cacheContent"
                size="small"
                variant="text"
                theme="primary"
                @click="handleCopyContent"
              >
                复制
              </Button>
            </div>
            <div class="min-h-0 flex-1 overflow-auto rounded border border-gray-200 bg-gray-50 p-3">
              <pre
                v-if="formattedContent"
                class="m-0 font-mono text-xs leading-relaxed text-gray-700"
              >{{ formattedContent }}</pre>
              <div v-else class="flex h-full items-center justify-center text-sm text-gray-400">
                {{ $t('system.monitor.cache.contentPlaceholder') }}
              </div>
            </div>
          </div>
        </div>
      </Card>
    </div>
  </Page>
</template>

<style scoped>
pre {
  white-space: pre-wrap;
  word-break: break-all;
}
</style>
