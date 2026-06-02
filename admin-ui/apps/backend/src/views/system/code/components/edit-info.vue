<script lang="ts" setup>
import type { FieldConfigRow } from '../model';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import { readTable, updateCode } from '#/api/system/generate';

import {
  Button,
  Checkbox,
  CheckboxGroup,
  Form,
  FormItem,
  Input,
  Select,
  Switch,
  TabPanel,
  Table,
  Tabs,
} from 'tdesign-vue-next';

import {
  componentTypeOptions,
  generateTypeOptions,
  menuButtonOptions,
  queryTypeOptions,
  tplTypeOptions,
  viewTypeOptions,
} from '../schemas';

const emit = defineEmits<{
  success: [];
}>();

const loading = ref(false);
const submitLoading = ref(false);
const activeTab = ref('base_config');
const tableName = ref('');
const tableComment = ref('');
const remark = ref('');
const moduleName = ref('');
const type = ref<'single' | 'tree'>('single');
const menuName = ref('');
const componentType = ref<number>(1);
const tplType = ref('default');
const treeId = ref('');
const treeParentId = ref('');
const treeName = ref('');
const tagId = ref('');
const tagName = ref('');
const tagViewName = ref('');
const fieldColumns = [
  { colKey: 'sort', title: $t('system.code.field.sort'), width: 70 },
  { colKey: 'column_name', title: $t('system.code.field.name'), width: 150 },
  { colKey: 'column_comment', title: $t('system.code.field.comment'), width: 150 },
  { colKey: 'column_type', title: $t('system.code.field.type'), width: 120 },
  { colKey: 'is_required', title: $t('system.code.field.required'), width: 70 },
  { colKey: 'is_insert', title: $t('system.code.field.insert'), width: 70 },
  { colKey: 'is_edit', title: $t('system.code.field.edit'), width: 70 },
  { colKey: 'is_list', title: $t('system.code.field.list'), width: 70 },
  { colKey: 'is_query', title: $t('system.code.field.query'), width: 70 },
  { colKey: 'is_sort', title: $t('system.code.field.sort'), width: 70 },
  { colKey: 'query_type', title: $t('system.code.field.queryType'), width: 100 },
  { colKey: 'view_type', title: $t('system.code.field.viewType'), width: 120 },
  { colKey: 'action', title: $t('common.action'), width: 100 },
];

const fieldList = ref<FieldConfigRow[]>([]);
const menuButtons = ref<string[]>(['save', 'update', 'read', 'delete']);

async function open(id: number) {
  loading.value = true;
  try {
    await readTable(id);
    tableName.value = '';
    tableComment.value = '';
    fieldList.value = [];
    activeTab.value = 'base_config';
    modalApi.setState({ title: $t('system.code.editTitle') });
    modalApi.open();
  } catch (error) {
    logger.error(error);
    message.error($t('common.tableInfoFailed'));
  } finally {
    loading.value = false;
  }
}

function handleEditField(row: FieldConfigRow) {
  void row;
  message.info($t('common.advancedConfigWIP'));
}

function handleAllChecked(checked: boolean, key: keyof FieldConfigRow) {
  if (checked) {
    fieldList.value.forEach((f) => {
      (f as any)[key] = 1;
    });
  } else {
    fieldList.value.forEach((f) => {
      (f as any)[key] = 2;
    });
  }
}

async function handleSubmit() {
  submitLoading.value = true;
  try {
    const payload: any = {
      id: 0,
      type: type.value,
      menu_buttons: menuButtons.value,
      fields: fieldList.value,
    };
    await updateCode(payload);
    message.success($t('common.saveSuccess'));
    emit('success');
    modalApi.close();
  } catch (error) {
    logger.error(error);
    message.error($t('common.saveFailed'));
  } finally {
    submitLoading.value = false;
  }
}

const [Modal, modalApi] = useVbenModal({
  onConfirm: handleSubmit,
  class: 'w-[1100px]',
});

defineExpose({ open });
</script>

<template>
  <Modal>
    <div class="flex flex-col gap-4">
      <Tabs v-model:value="activeTab">
        <TabPanel value="base_config" :label="$t('system.code.baseConfig')">
          <Form :label-width="120" layout="inline" colon>
            <div class="grid grid-cols-2 gap-x-4 gap-y-3">
              <FormItem :label="$t('system.code.tableName')">
                <Input v-model="tableName" disabled />
              </FormItem>
              <FormItem :label="$t('system.code.tableComment')" name="table_comment">
                <Input v-model="tableComment" :placeholder="$t('ui.placeholder.input')" />
              </FormItem>
              <FormItem :label="$t('common.remark')">
                <Input v-model="remark" :placeholder="$t('ui.placeholder.input')" />
              </FormItem>
              <FormItem :label="$t('system.code.moduleName')" name="module_name">
                <Input v-model="moduleName" :placeholder="$t('ui.placeholder.input')" />
              </FormItem>
              <FormItem :label="$t('system.code.genType')" name="type">
                <Select v-model="type" :options="generateTypeOptions" />
              </FormItem>
              <FormItem :label="$t('system.code.menuName')" name="menu_name">
                <Input v-model="menuName" :placeholder="$t('ui.placeholder.input')" />
              </FormItem>
              <FormItem :label="$t('system.code.componentStyle')" name="component_type">
                <Select v-model="componentType" :options="componentTypeOptions" />
              </FormItem>
              <FormItem :label="$t('system.code.tplType')" name="tpl_type">
                <Select v-model="tplType" :options="tplTypeOptions" />
              </FormItem>
            </div>
            <template v-if="type === 'tree'">
              <div class="mt-4 border-t pt-4">
                <div class="mb-2 text-sm font-medium text-gray-500">{{ $t('system.code.treeConfig') }}</div>
                <div class="grid grid-cols-3 gap-x-4 gap-y-3">
                  <FormItem :label="$t('system.code.treeId')" name="tree_id">
                    <Input v-model="treeId" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                  <FormItem :label="$t('system.code.treeParentId')" name="tree_parent_id">
                    <Input v-model="treeParentId" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                  <FormItem :label="$t('system.code.treeName')" name="tree_name">
                    <Input v-model="treeName" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                </div>
              </div>
            </template>
            <template v-if="componentType === 3">
              <div class="mt-4 border-t pt-4">
                <div class="mb-2 text-sm font-medium text-gray-500">{{ $t('system.code.tagConfig') }}</div>
                <div class="grid grid-cols-3 gap-x-4 gap-y-3">
                  <FormItem :label="$t('system.code.tagId')" name="tag_id">
                    <Input v-model="tagId" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                  <FormItem :label="$t('system.code.tagName')" name="tag_name">
                    <Input v-model="tagName" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                  <FormItem :label="$t('system.code.tagViewName')" name="tag_view_name">
                    <Input v-model="tagViewName" :placeholder="$t('ui.placeholder.input')" />
                  </FormItem>
                </div>
              </div>
            </template>
          </Form>
        </TabPanel>
        <TabPanel value="field_config" :label="$t('system.code.fieldConfig')">
          <div class="mb-2 flex items-center gap-2">
            <Checkbox
              :checked="fieldList.every((f) => f.is_required === 1)"
              @change="(checked: boolean) => handleAllChecked(checked, 'is_required')"
            >
              {{ $t('system.code.field.selectAllRequired') }}
            </Checkbox>
            <Checkbox
              :checked="fieldList.every((f) => f.is_insert === 1)"
              @change="(checked: boolean) => handleAllChecked(checked, 'is_insert')"
            >
              {{ $t('system.code.field.selectAllInsert') }}
            </Checkbox>
            <Checkbox
              :checked="fieldList.every((f) => f.is_edit === 1)"
              @change="(checked: boolean) => handleAllChecked(checked, 'is_edit')"
            >
              {{ $t('system.code.field.selectAllEdit') }}
            </Checkbox>
            <Checkbox
              :checked="fieldList.every((f) => f.is_list === 1)"
              @change="(checked: boolean) => handleAllChecked(checked, 'is_list')"
            >
              {{ $t('system.code.field.selectAllList') }}
            </Checkbox>
            <Checkbox
              :checked="fieldList.every((f) => f.is_query === 1)"
              @change="(checked: boolean) => handleAllChecked(checked, 'is_query')"
            >
              {{ $t('system.code.field.selectAllQuery') }}
            </Checkbox>
          </div>

          <Table
            :columns="fieldColumns"
            :data="fieldList"
            row-key="column_name"
            hover
            stripe
          >
            <template #is_required="{ row }">
              <Switch v-model="row.is_required" :value="row.is_required === 1" />
            </template>
            <template #is_insert="{ row }">
              <Switch v-model="row.is_insert" :value="row.is_insert === 1" />
            </template>
            <template #is_edit="{ row }">
              <Switch v-model="row.is_edit" :value="row.is_edit === 1" />
            </template>
            <template #is_list="{ row }">
              <Switch v-model="row.is_list" :value="row.is_list === 1" />
            </template>
            <template #is_query="{ row }">
              <Switch v-model="row.is_query" :value="row.is_query === 1" />
            </template>
            <template #is_sort="{ row }">
              <Switch v-model="row.is_sort" :value="row.is_sort === 1" />
            </template>
            <template #query_type="{ row }">
              <Select
                v-model="row.query_type"
                :options="queryTypeOptions"
                size="small"
              />
            </template>
            <template #view_type="{ row }">
              <Select
                v-model="row.view_type"
                :options="viewTypeOptions"
                size="small"
              />
            </template>
            <template #action="{ row }">
              <Button size="small" theme="primary" @click="handleEditField(row)">
                {{ $t('system.code.baseConfig') }}
              </Button>
            </template>
          </Table>
        </TabPanel>
        <TabPanel value="menu_config" :label="$t('system.code.menuConfig')">
          <div class="py-4">
            <CheckboxGroup v-model="menuButtons" :options="menuButtonOptions" />
          </div>
        </TabPanel>
      </Tabs>
    </div>
  </Modal>
</template>
