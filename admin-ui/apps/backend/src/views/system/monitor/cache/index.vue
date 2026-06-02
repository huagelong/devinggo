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
  Col,
  Input,
  Row,
  Space,
  Table,
  Textarea,
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
  { colKey: 'name', title: $t('system.monitor.cache.keyName') },
  { colKey: 'action', title: $t('common.action'), width: 150, align: 'right' },
] satisfies PrimaryTableCol<TableRowData>[];

const filteredData = computed(() => {
  if (!searchKey.value) return cacheKeys.value.map((name) => ({ name }));
  const keyword = searchKey.value.toLowerCase();
  return cacheKeys.value
    .filter((name) => name.toLowerCase().includes(keyword))
    .map((name) => ({ name }));
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
    cacheContent.value = response.data?.content || '';
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

onMounted(() => {
  void fetchCacheInfo();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex flex-col gap-3">
      <!-- Redis Info Panel -->
      <Card :title="$t('system.monitor.cache.redisInfo')" class="w-full">
        <Row :gutter="24">
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.redisVersion') }}</div>
            <div class="text-base">{{ serverInfo.version || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.clientConnections') }}</div>
            <div class="text-base">{{ serverInfo.clients || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.runMode') }}</div>
            <div class="text-base">{{ serverInfo.redis_mode || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.runDays') }}</div>
            <div class="text-base">{{ serverInfo.run_days || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.port') }}</div>
            <div class="text-base">{{ serverInfo.port || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.aofStatus') }}</div>
            <div class="text-base">{{ serverInfo.aof_enabled || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.expiredKeys') }}</div>
            <div class="text-base">{{ serverInfo.expired_keys || '-' }}</div>
          </Col>
          <Col :span="6">
            <div class="mb-2 text-sm text-gray-500">{{ $t('system.monitor.cache.systemUsedKeys') }}</div>
            <div class="text-base">{{ serverInfo.sys_total_keys || '-' }}</div>
          </Col>
        </Row>
      </Card>

      <!-- Cache Operations -->
      <Card>
        <Space>
          <Button theme="danger" variant="outline" @click="handleClearAll">
            <template #icon><DeleteIcon /></template>
            {{ $t('system.monitor.cache.clearAll') }}
          </Button>
          <Button
            v-if="selectedKeys.length > 0"
            theme="danger"
            @click="handleBatchDelete"
          >
            {{ $t('common.batchDelete') }} ({{ selectedKeys.length }})
          </Button>
        </Space>
      </Card>

      <!-- Cache Data Management -->
      <Card :title="$t('system.monitor.cache.dataManagement')">
        <div class="flex gap-4">
          <!-- Left: Cache Key Table -->
          <div class="w-2/3">
            <Input
              v-model="searchKey"
              :placeholder="$t('system.monitor.cache.searchKeyPlaceholder')"
              clearable
              class="mb-3 w-full"
            >
              <template #prefix><SearchIcon /></template>
            </Input>

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

          <!-- Right: Cache Content -->
          <div class="w-1/3">
            <Textarea
              v-model="cacheContent"
              readonly
              :placeholder="$t('system.monitor.cache.contentPlaceholder')"
              class="h-full min-h-[400px] w-full"
            />
          </div>
        </div>
      </Card>
    </div>
  </Page>
</template>
