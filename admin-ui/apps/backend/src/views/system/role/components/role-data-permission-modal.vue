<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DeptApi } from '#/api/system/dept';
import type { RoleApi } from '#/api/system/role';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Checkbox,
  Form,
  FormItem,
  Input,
  Select,
  Space,
  Tree,
  MessagePlugin,
} from 'tdesign-vue-next';

import { getDeptTree } from '#/api/system/dept';
import {
  getDeptByRole,
  updateRoleDataPermission,
} from '#/api/system/role';

import { roleDataScopeOptions } from '../schemas';

const emit = defineEmits(['success']);

const loading = ref(false);
const currentRole = ref<null | RoleApi.ListItem>(null);
const deptList = ref<DeptApi.TreeNode[]>([]);
const checkedKeys = ref<Array<number | string>>([]);
const expandedKeys = ref<Array<number | string>>([]);
const cancelLinkage = ref(false);
const searchText = ref('');
const dataScope = ref<number>(1);

const treeKeys = {
  value: 'id',
  label: 'label',
  children: 'children',
};

function flattenTreeIds(nodes: DeptApi.TreeNode[]) {
  const ids: Array<number | string> = [];
  const travel = (items: DeptApi.TreeNode[]) => {
    items.forEach((item) => {
      ids.push(item.id);
      if (item.children?.length) {
        travel(item.children);
      }
    });
  };
  travel(nodes);
  return ids;
}

const allNodeKeys = computed(() => flattenTreeIds(deptList.value));
const shouldShowDeptTree = computed(() => Number(dataScope.value) === 2);

function handleExpand(checked: boolean) {
  expandedKeys.value = checked ? [...allNodeKeys.value] : [];
}

function handleSelect(checked: boolean) {
  checkedKeys.value = checked ? [...allNodeKeys.value] : [];
}

function handleLinkage(checked: boolean) {
  cancelLinkage.value = checked;
}

function handleTreeClick(context: { node?: { value?: number | string } }) {
  const nodeValue = context?.node?.value;
  if (nodeValue === undefined) return;
  expandedKeys.value = expandedKeys.value.includes(nodeValue)
    ? expandedKeys.value.filter((item) => item !== nodeValue)
    : [...expandedKeys.value, nodeValue];
}

function filterTreeNode(node: { label?: string }) {
  const label = node?.label ?? '';
  return !searchText.value || label.includes(searchText.value);
}

function extractDeptIds(list?: RoleApi.RoleDeptPermission[]) {
  return (
    list?.flatMap((item) => item.depts?.map((dept) => Number(dept.id)) ?? []) ??
    []
  );
}

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    const role = currentRole.value;
    if (!role) return;

    try {
      modalApi.setState({ confirmLoading: true });
      await updateRoleDataPermission(role.id, {
        data_scope: Number(dataScope.value),
        dept_ids: checkedKeys.value.map((item) => Number(item)),
      });
      MessagePlugin.success($t('common.dataPermissionUpdateSuccess'));
      emit('success');
      modalApi.close();
    } catch (error) {
      logger.error(error);
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  class: 'w-[960px] max-w-[94vw]',
});

async function open(role: RoleApi.ListItem) {
  currentRole.value = role;
  searchText.value = '';
  deptList.value = [];
  checkedKeys.value = [];
  expandedKeys.value = [];
  cancelLinkage.value = false;
  dataScope.value = Number(role.data_scope ?? 1);

  modalApi.setState({
    title: $t('system.role.dataPermission'),
  });
  modalApi.open();

  loading.value = true;
  try {
    const [deptTree, relationData] = await Promise.all([
      getDeptTree().catch(() => [] as DeptApi.TreeNode[]),
      getDeptByRole(role.id).catch(() => [] as RoleApi.RoleDeptPermission[]),
    ]);

    deptList.value = deptTree;
    checkedKeys.value = extractDeptIds(relationData);
    if (checkedKeys.value.length > 0) {
      cancelLinkage.value = true;
    }
  } finally {
    loading.value = false;
  }
}

defineExpose({
  open,
});
</script>

<template>
  <Modal>
    <Form
      :data="currentRole ?? {}"
      class="permission-form"
      label-width="100px"
      layout="inline"
      colon
    >
      <FormItem :label="$t('system.role.name')" name="name">
        <Input :model-value="currentRole?.name ?? ''" disabled />
      </FormItem>
      <FormItem :label="$t('system.role.code')" name="code">
        <Input :model-value="currentRole?.code ?? ''" disabled />
      </FormItem>
      <FormItem :label="$t('system.role.dataScope')" name="data_scope">
        <Select
          v-model="dataScope"
          :options="roleDataScopeOptions"
          :placeholder="$t('ui.placeholder.select', [$t('system.role.dataScope')])"
        />
      </FormItem>

      <template v-if="shouldShowDeptTree">
        <FormItem :label="$t('system.role.searchDept')" name="search">
          <Input v-model="searchText" :placeholder="$t('common.filterDept')" clearable />
        </FormItem>
        <FormItem :label="$t('system.role.deptList')" name="dept_ids">
          <div class="w-full">
            <Space class="mb-3">
              <Checkbox @change="handleExpand">{{ $t('common.expandCollapse') }}</Checkbox>
              <Checkbox @change="handleSelect">{{ $t('common.selectAllNone') }}</Checkbox>
              <Checkbox
                :checked="cancelLinkage"
                @change="handleLinkage"
              >
                {{ $t('common.disableParentChildLink') }}
              </Checkbox>
            </Space>
            <div class="tree-container">
              <div
                v-if="loading"
                class="flex h-[320px] items-center justify-center text-sm text-gray-500"
              >
                {{ $t('common.deptLoading') }}
              </div>
              <Tree
                v-else
                v-model="checkedKeys"
                v-model:expanded="expandedKeys"
                :check-strictly="cancelLinkage"
                :data="deptList"
                :filter="filterTreeNode"
                :keys="treeKeys"
                :value-mode="'all'"
                checkable
                hover
                line
                @click="handleTreeClick"
              />
            </div>
          </div>
        </FormItem>
      </template>
    </Form>
  </Modal>
</template>

<style scoped>
.permission-form :deep(.t-form__item) {
  width: 100%;
  margin-right: 0;
}

.permission-form :deep(.t-form__controls) {
  flex: 1;
  min-width: 0;
}

.tree-container {
  border: 1px solid var(--td-component-border);
  border-radius: 6px;
  max-height: 360px;
  min-height: 320px;
  overflow: auto;
  padding: 8px;
}
</style>
