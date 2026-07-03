<script lang="ts" setup>
import type { DataMaintainApi } from '#/api/system/data-maintain';

import { computed, ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { getDataMaintainDetailed } from '#/api/system/data-maintain';
import { logger } from '#/utils/logger';

import { Button, Table, Tag } from 'tdesign-vue-next';

interface OpenOptions {
  groupName?: string;
  hasDetailedApi: boolean;
  row: DataMaintainApi.ListItem;
}

const visible = ref(false);
const loading = ref(false);
const hasDetailedApi = ref(false);
const currentTable = ref<DataMaintainApi.ListItem>();
const detailColumns = ref<DataMaintainApi.ColumnItem[]>([]);

const detailTableColumns = computed(() => [
  { colKey: 'field', title: $t('system.dataMaintain.fieldColName'), width: 220 },
  { colKey: 'type', title: $t('system.dataMaintain.fieldColType'), width: 180 },
  { colKey: 'nullable', title: $t('system.dataMaintain.fieldColNullable'), width: 120 },
  { colKey: 'default_value', title: $t('system.dataMaintain.fieldColDefault'), minWidth: 180 },
  { colKey: 'key', title: $t('system.dataMaintain.fieldColKey'), width: 140 },
  { colKey: 'comment', title: $t('system.dataMaintain.fieldColComment'), minWidth: 220 },
]);

async function open(options: OpenOptions) {
  currentTable.value = options.row;
  hasDetailedApi.value = options.hasDetailedApi;
  visible.value = true;
  detailColumns.value = [];

  if (!options.hasDetailedApi) {
    return;
  }

  loading.value = true;
  try {
    const response = await getDataMaintainDetailed({
      group_name: options.groupName,
      table_name: options.row.name,
    });
    detailColumns.value = response.items ?? [];
  } catch (error) {
    logger.error(error);
    message.error($t('common.fieldDetailFailed'));
  } finally {
    loading.value = false;
  }
}

function close() {
  visible.value = false;
}

function formatBytes(bytes?: number) {
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

defineExpose({
  close,
  open,
});
</script>

<template>
  <div
    v-if="visible"
    class="mt-3 rounded-md border border-border bg-muted/50 p-4"
  >
    <div class="mb-2 flex items-center justify-between">
      <div class="text-sm font-medium text-foreground">
        {{ $t('system.dataMaintain.tableDetail', [currentTable?.name || '-']) }}
      </div>
      <Button size="small" variant="text" @click="close">{{ $t('common.collapse') }}</Button>
    </div>

    <div class="mb-3 grid grid-cols-2 gap-3 text-sm text-muted-foreground xl:grid-cols-5">
      <div>{{ $t('system.dataMaintain.engine') }}：{{ currentTable?.engine || '-' }}</div>
      <div>{{ $t('system.dataMaintain.collation') }}：{{ currentTable?.collation || '-' }}</div>
      <div>{{ $t('system.dataMaintain.rows') }}：{{ currentTable?.rows ?? '-' }}</div>
      <div>{{ $t('system.dataMaintain.tableSize') }}：{{ formatBytes(currentTable?.data_length) }}</div>
      <div>{{ $t('common.updateTime') }}：{{ currentTable?.update_time || '-' }}</div>
    </div>

    <div v-if="!hasDetailedApi" class="text-sm text-muted-foreground">
      {{ $t('system.dataMaintain.noFieldApi') }}
    </div>

    <Table
      v-else
      row-key="field"
      size="small"
      :loading="loading"
      :data="detailColumns"
      :columns="detailTableColumns"
    >
      <template #nullable="{ row }">
        <Tag :theme="row.nullable ? 'success' : 'default'" variant="light">
          {{ row.nullable ? $t('system.dataMaintain.yes') : $t('system.dataMaintain.no') }}
        </Tag>
      </template>
      <template #default_value="{ row }">
        {{ row.default_value || '-' }}
      </template>
      <template #key="{ row }">
        <Tag :theme="fieldKeyTheme(row.key)" variant="light">
          {{ fieldKeyLabel(row.key) }}
        </Tag>
      </template>
      <template #comment="{ row }">
        {{ row.comment || '-' }}
      </template>
    </Table>
  </div>
</template>
