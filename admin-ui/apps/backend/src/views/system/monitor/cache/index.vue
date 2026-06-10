<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { computed, onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

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
      <Card :title="$t('system.monitor.cache.redisInfo')" class="w-full" size="small">
        <div class="grid grid-cols-2 gap-x-8 gap-y-2 sm:grid-cols-4">
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.redisVersion') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.version || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.clientConnections') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.clients || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.runMode') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.redis_mode || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.runDays') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.run_days || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.port') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.port || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.aofStatus') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.aof_enabled || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.expiredKeys') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.expired_keys || '-' }}</span>
          </div>
          <div class="flex items-center gap-2 overflow-hidden">
            <span class="text-xs text-gray-400 whitespace-nowrap">{{ $t('system.monitor.cache.systemUsedKeys') }}</span>
            <span class="truncate text-sm font-medium">{{ serverInfo.sys_total_keys || '-' }}</span>
          </div>
        </div>
      </Card>

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
