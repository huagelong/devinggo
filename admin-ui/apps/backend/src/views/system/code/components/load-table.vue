<script lang="ts" setup>
import type { GenerateApi } from '#/api/system/generate';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import { loadTable } from '#/api/system/generate';

import { Table } from 'tdesign-vue-next';

const emit = defineEmits<{
  success: [];
}>();

const tableData = ref<GenerateApi.TableColumn[]>([]);
const tableLoading = ref(false);
const submitLoading = ref(false);
const selectedNames = ref<string[]>([]);

const columns = [
  { colKey: 'selection', width: 60 },
  { colKey: 'name', title: $t('system.code.tableName'), width: 200 },
  { colKey: 'comment', title: $t('system.code.tableComment'), minWidth: 200 },
];

function open() {
  modalApi.open();
  selectedNames.value = [];
  tableData.value = [];
  void fetchTableList();
}

async function fetchTableList() {
  tableLoading.value = true;
  try {
    // 模拟数据，实际应该从API获取
    tableData.value = [];
  } catch (error) {
    logger.error(error);
    message.error($t('common.tableListFailed'));
  } finally {
    tableLoading.value = false;
  }
}

async function handleSubmit() {
  if (selectedNames.value.length === 0) {
    message.warning($t('common.selectTableFirst'));
    return;
  }

  submitLoading.value = true;
  try {
    const names = selectedNames.value.map((name) => ({
      name,
      comment: name,
      sourceName: name,
    }));
    await loadTable({ source: 'default', names });
    message.success($t('common.loadSuccess'));
    emit('success');
    modalApi.close();
  } catch (error) {
    logger.error(error);
    message.error($t('common.loadFailed'));
  } finally {
    submitLoading.value = false;
  }
}

function handleRowSelectionChange(keys: Array<number | string>) {
  selectedNames.value = keys.map((key) => String(key));
}

const [Modal, modalApi] = useVbenModal({
  onConfirm: handleSubmit,
  class: 'w-[800px]',
});

defineExpose({ open });
</script>

<template>
  <Modal :title="$t('system.code.loadTable')">
    <div class="flex flex-col gap-4">
      <div class="text-sm text-gray-500">
        {{ $t('system.code.loadTableTip') }}
      </div>

      <Table
        :columns="columns"
        :data="tableData"
        :loading="tableLoading"
        :row-selection="{
          type: 'multiple',
          selectedRowKeys: selectedNames,
          onChange: handleRowSelectionChange,
        }"
        row-key="name"
        hover
        stripe
      />

      <div class="text-sm text-gray-500">
        {{ $t('system.code.selectedTableCount', { count: selectedNames.length }) }}
      </div>
    </div>
  </Modal>
</template>
