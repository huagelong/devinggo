<script lang="ts" setup>
import type { DataMaintainApi } from '#/api/system/data-maintain';

import { ref } from 'vue';

import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';
import { getDataMaintainDetailed } from '#/api/system/data-maintain';

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
const detailColumns = ref<Array<{ field: string; type?: string; comment?: string }>>([]);

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
    detailColumns.value = Object.values(response || {}).map((item) => ({
      comment: item.comment,
      field: item.field,
      type: item.type,
    }));
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

defineExpose({
  close,
  open,
});
</script>

<template>
  <div
    v-if="visible"
    class="mt-3 rounded-md border border-gray-100 bg-gray-50 p-4"
  >
    <div class="mb-2 flex items-center justify-between">
      <div class="text-sm font-medium text-gray-700">
        {{ $t('system.dataMaintain.tableDetail', [currentTable?.name || '-']) }}
      </div>
      <Button size="small" variant="text" @click="close">{{ $t('common.collapse') }}</Button>
    </div>

    <div class="mb-3 grid grid-cols-3 gap-3 text-sm text-gray-600">
      <div>{{ $t('system.dataMaintain.engine') }}：{{ currentTable?.engine || '-' }}</div>
      <div>{{ $t('system.dataMaintain.collation') }}：{{ currentTable?.collation || '-' }}</div>
      <div>{{ $t('system.dataMaintain.rows') }}：{{ currentTable?.rows ?? '-' }}</div>
    </div>

    <div v-if="!hasDetailedApi" class="text-sm text-gray-500">
      {{ $t('system.dataMaintain.detailApiNotAvailable') }}
    </div>

    <Table
      v-else
      row-key="field"
      size="small"
      :loading="loading"
      :data="detailColumns"
      :columns="[
        { colKey: 'field', title: $t('system.dataMaintain.fieldColName'), width: 220 },
        { colKey: 'type', title: $t('system.dataMaintain.fieldColType'), width: 180 },
        { colKey: 'comment', title: $t('system.dataMaintain.fieldColComment'), minWidth: 240 },
      ]"
    />

    <div class="mt-3 flex items-center gap-2 text-xs text-gray-500">
      <Tag theme="warning" variant="light">{{ $t('system.dataMaintain.capabilityReserved') }}</Tag>
      <span>{{ $t('system.dataMaintain.capabilityReservedDesc') }}</span>
    </div>
  </div>
</template>
