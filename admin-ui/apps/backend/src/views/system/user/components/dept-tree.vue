<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DeptApi } from '#/api/system/dept';
import type { IdType } from '#/types/common';

import { computed, onMounted, ref } from 'vue';

import { $t } from '#/locales';

import {
  CaretDownSmallIcon,
  CaretRightSmallIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import { Input, Tree } from 'tdesign-vue-next';

import { getDeptTree } from '#/api/system/dept';

const emit = defineEmits(['select']);

const treeData = ref<DeptApi.TreeNode[]>([]);
const searchText = ref('');

const treeKeys = {
  value: 'id',
  label: 'label',
  children: 'children',
};

// 展开的节点
const expanded = ref<IdType[]>([]);
const isFolding = ref(true);

const filter = computed(() => {
  if (!searchText.value) {
    return undefined;
  }

  return (node: any) => {
    const label = node.label ?? '';
    return label.includes(searchText.value);
  };
});

const expandButtonText = computed(() =>
  expanded.value.length > 0 ? $t('common.expandCollapse') : $t('common.collapse'),
);

async function fetchDeptTree() {
  try {
    const res = await getDeptTree();
    treeData.value = res || [];
    // 默认展开所有节点
    const ids: IdType[] = [];
    const traverse = (nodes: DeptApi.TreeNode[]) => {
      nodes.forEach((node) => {
        ids.push(node.id);
        if (node.children) {
          traverse(node.children);
        }
      });
    };
    traverse(treeData.value);
    expanded.value = ids;
  } catch (error) {
    logger.error('Failed to fetch dept tree', error);
  }
}

function handleNodeClick(context: unknown) {
  const nodeValue =
    (context as { node?: { value?: IdType } })?.node?.value ?? '';
  emit('select', nodeValue);
}

function toggleExpand() {
  if (isFolding.value) {
    expanded.value = [];
  } else {
    const ids: IdType[] = [];
    const traverse = (nodes: DeptApi.TreeNode[]) => {
      nodes.forEach((node) => {
        ids.push(node.id);
        if (node.children) {
          traverse(node.children);
        }
      });
    };
    traverse(treeData.value);
    expanded.value = ids;
  }
  isFolding.value = !isFolding.value;
}

onMounted(() => {
  fetchDeptTree();
});
</script>

<template>
  <div class="flex h-full w-56 flex-col bg-white">
    <div
      class="flex items-center justify-between gap-2 border-b border-gray-100 p-3"
    >
      <Input v-model="searchText" :placeholder="$t('system.user.searchDept')" class="flex-1 min-w-0">
        <template #prefixIcon>
          <SearchIcon />
        </template>
      </Input>
      <div
        class="cursor-pointer whitespace-nowrap text-[14px] text-blue-500 hover:text-blue-600"
        @click="toggleExpand"
      >
        {{ expandButtonText }}
      </div>
    </div>
    <div class="custom-tree-wrap flex-1 overflow-auto p-2">
      <Tree
        v-model:expanded="expanded"
        :data="treeData"
        :filter="filter"
        :keys="treeKeys"
        hover
        activable
        line
        allow-fold-node-on-filter
        @click="handleNodeClick"
      >
        <template #icon="{ node }">
          <CaretDownSmallIcon v-if="node.expanded" />
          <CaretRightSmallIcon v-else />
        </template>
      </Tree>
    </div>
  </div>
</template>

<style scoped>
.custom-tree-wrap {
  :deep(.t-tree) {
    --td-tree-hover-color: #f3f4f6;
  }
}
</style>
