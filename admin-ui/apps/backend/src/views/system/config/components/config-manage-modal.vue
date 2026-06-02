<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { ConfigApi } from '#/api/system/config';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { nextTick, reactive, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { SearchIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  Form,
  FormItem,
  Input,
  MessagePlugin,
  Popconfirm,
  Select,
  Space,
  Table,
} from 'tdesign-vue-next';

import {
  deleteConfig,
  getConfigList,
  getConfigGroupList,
} from '#/api/system/config';

import ConfigFormModal from './config-form-modal.vue';

const configFormModalRef = ref<{ open: (options?: { data?: ConfigApi.ConfigSubmitPayload }) => void }>();

const searchForm = reactive({
  key: '',
  name: '',
});
const tableData = ref<ConfigApi.ConfigItem[]>([]);
const loading = ref(false);
const currentGroupId = ref<number>();
const groupOptions = ref<DictOption<number>[]>([]);

const [Modal, modalApi] = useVbenModal({
  footer: false,
  class: 'w-[1120px] max-w-[96vw]',
});

function buildParams() {
  const params: ConfigApi.ConfigListQuery = {
    group_id: currentGroupId.value,
  };
  if (searchForm.name) params.name = searchForm.name;
  if (searchForm.key) params.key = searchForm.key;
  return params;
}

async function fetchTableData() {
  if (!currentGroupId.value) return;
  loading.value = true;
  try {
    const list = await getConfigList(buildParams());
    tableData.value = list;
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.configDataLoadFailed2'));
  } finally {
    loading.value = false;
  }
}

async function fetchGroupOptions() {
  const list = await getConfigGroupList();
  groupOptions.value = list.map((item) => ({
    label: item.name,
    value: item.id,
  }));
}

function handleSearch() {
  void fetchTableData();
}

function handleReset() {
  searchForm.key = '';
  searchForm.name = '';
  void fetchTableData();
}

function handleAdd() {
  configFormModalRef.value?.open({
    data: { group_id: currentGroupId.value } as ConfigApi.ConfigSubmitPayload,
  });
}

function handleEdit(row: ConfigApi.ConfigItem) {
  configFormModalRef.value?.open({
    data: {
      ...row,
      group_id: row.group_id,
      value: row.value as string,
    } as ConfigApi.ConfigSubmitPayload,
  });
}

async function handleDelete(row: ConfigApi.ConfigItem) {
  try {
    await deleteConfig({ ids: [row.id] });
    MessagePlugin.success($t('common.deleteSuccess'));
    await fetchTableData();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.deleteFailed'));
  }
}

async function open(groupId: number) {
  currentGroupId.value = groupId;
  searchForm.key = '';
  searchForm.name = '';
  await fetchGroupOptions();
  modalApi.setState({ title: $t('system.config.manageTitle') });
  modalApi.open();
  await nextTick();
  await fetchTableData();
}

defineExpose({
  open,
  refresh: fetchTableData,
});
</script>

<template>
  <Modal>
    <div class="config-manage-modal flex flex-col gap-4">
      <Form
        :data="searchForm"
        class="config-search-form"
        label-width="90px"
        layout="inline"
        colon
      >
        <div class="grid grid-cols-1 gap-x-4 gap-y-3 md:grid-cols-2 xl:grid-cols-3">
          <FormItem :label="$t('system.config.name')" name="name">
            <Input
              v-model="searchForm.name"
              :placeholder="$t('system.config.placeholder.enterConfigName')"
              clearable
            />
          </FormItem>
          <FormItem :label="$t('system.config.code')" name="key">
            <Input
              v-model="searchForm.key"
              :placeholder="$t('system.config.placeholder.enterConfigCode')"
              clearable
            />
          </FormItem>
          <FormItem :label="$t('system.config.currentGroup')">
            <Select
              v-model="currentGroupId"
              :options="groupOptions"
              :placeholder="$t('system.config.placeholder.selectGroup')"
              class="w-full"
              @change="handleSearch"
            />
          </FormItem>
        </div>
        <div class="config-search-actions">
          <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
          <Button theme="primary" @click="handleSearch">
            <template #icon><SearchIcon /></template>
            {{ $t('common.query') }}
          </Button>
        </div>
      </Form>

      <div class="manage-table-card">
        <div class="mb-3 flex items-center justify-between">
          <Space>
            <Button theme="primary" @click="handleAdd">{{ $t('system.config.addConfigTitle') }}</Button>
          </Space>
        </div>

        <Table
          :columns="[
            { colKey: 'name', title: $t('system.config.name'), width: 220 },
            { colKey: 'key', title: $t('system.config.code'), width: 180 },
            { colKey: 'value', title: $t('system.config.value'), minWidth: 200 },
            { colKey: 'input_type', title: $t('system.config.inputComponent'), width: 120 },
            { colKey: 'sort', title: $t('common.sort'), width: 80 },
            { colKey: 'action', title: $t('common.action'), width: 180, align: 'center' },
          ]"
          :data="tableData"
          :loading="loading"
          row-key="id"
          hover
          stripe
          table-layout="fixed"
        >
          <template #action="{ row }">
            <div class="flex items-center justify-center gap-2">
              <Button size="small" theme="primary" variant="outline" @click="handleEdit(row)">
                {{ $t('common.edit') }}
              </Button>
              <Popconfirm :content="$t('system.config.confirmDelete')" @confirm="handleDelete(row)">
                <Button size="small" theme="danger" variant="outline">
                  {{ $t('common.delete') }}
                </Button>
              </Popconfirm>
            </div>
          </template>
        </Table>
      </div>
    </div>

    <ConfigFormModal ref="configFormModalRef" @success="fetchTableData" />
  </Modal>
</template>

<style scoped>
.config-manage-modal {
  padding: 2px 2px 6px;
}

.config-search-form {
  border: 1px solid var(--td-component-border, #e7e7e7);
  border-radius: 10px;
  padding: 14px 14px 10px;
  background: #fafcff;
}

.config-search-form :deep(.t-form__item) {
  width: 100%;
  margin-right: 0;
  margin-bottom: 0;
}

.config-search-form :deep(.t-form__controls) {
  flex: 1;
  min-width: 0;
}

.config-search-actions {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  padding-top: 10px;
}

.manage-table-card {
  border: 1px solid var(--td-component-border, #e7e7e7);
  border-radius: 10px;
  background: #fff;
  padding: 14px;
}

@media (max-width: 1023px) {
  .config-search-actions {
    justify-content: flex-start;
    flex-wrap: wrap;
  }
}
</style>
