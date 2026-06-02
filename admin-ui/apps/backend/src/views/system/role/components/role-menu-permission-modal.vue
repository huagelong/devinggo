<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { MenuApi } from '#/api/system/menu';
import type { RoleApi } from '#/api/system/role';

import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import {
  Checkbox,
  Form,
  FormItem,
  Input,
  Space,
  Tree,
  MessagePlugin,
} from 'tdesign-vue-next';

import { getMenuTreeOptions } from '#/api/system/menu';
import {
  getMenuByRole,
  updateRoleMenuPermission,
} from '#/api/system/role';

const emit = defineEmits(['success']);

const loading = ref(false);
const currentRole = ref<null | RoleApi.ListItem>(null);
const menuList = ref<MenuApi.TreeOptionItem[]>([]);
const checkedKeys = ref<Array<number | string>>([]);
const expandedKeys = ref<Array<number | string>>([]);
const cancelLinkage = ref(false);
const searchText = ref('');

const treeKeys = {
  value: 'id',
  label: 'label',
  children: 'children',
};

function flattenTreeIds(nodes: MenuApi.TreeOptionItem[]) {
  const ids: Array<number | string> = [];
  const travel = (items: MenuApi.TreeOptionItem[]) => {
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

const allNodeKeys = computed(() => flattenTreeIds(menuList.value));

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

function extractMenuIds(list?: RoleApi.RoleMenuPermission[]) {
  return (
    list?.flatMap((item) => item.menus?.map((menu) => Number(menu.id)) ?? []) ??
    []
  );
}

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    const role = currentRole.value;
    if (!role) return;

    try {
      modalApi.setState({ confirmLoading: true });
      await updateRoleMenuPermission(role.id, {
        menu_ids: checkedKeys.value.map((item) => Number(item)),
      });
      MessagePlugin.success($t('common.menuPermissionUpdateSuccess'));
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
  menuList.value = [];
  checkedKeys.value = [];
  expandedKeys.value = [];
  cancelLinkage.value = false;

  modalApi.setState({
    title: $t('system.role.menuPermission'),
  });
  modalApi.open();

  loading.value = true;
  try {
    const [menuTree, relationData] = await Promise.all([
      getMenuTreeOptions({ scope: true }).catch(
        () => [] as MenuApi.TreeOptionItem[],
      ),
      getMenuByRole(role.id).catch(() => [] as RoleApi.RoleMenuPermission[]),
    ]);

    menuList.value = menuTree;
    checkedKeys.value = extractMenuIds(relationData);
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
      <FormItem :label="$t('system.role.searchMenu')" name="search">
        <Input v-model="searchText" :placeholder="$t('common.filterMenu')" clearable />
      </FormItem>
      <FormItem :label="$t('system.role.menuList')" name="menu_ids">
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
              {{ $t('common.menuLoading') }}
            </div>
            <Tree
              v-else
              v-model="checkedKeys"
              v-model:expanded="expandedKeys"
              :check-strictly="cancelLinkage"
              :data="menuList"
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
