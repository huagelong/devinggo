<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { UserApi } from '#/api/system/user';

import { computed, reactive, ref, watch } from 'vue';

import { $t } from '@vben/locales';

import {
  Button,
  Dialog,
  Form,
  FormItem,
  Input,
  Select,
  Table,
  TreeSelect,
} from 'tdesign-vue-next';
import type { PrimaryTableCol, TableRowData } from 'tdesign-vue-next/es/table/type';

import { getUserList } from '#/api/system/user';
import { getDeptTree } from '#/api/system/dept';
import { getRoleList } from '#/api/system/role';
import { getPostList } from '#/api/system/post';

interface Props {
  visible: boolean;
  selectedIds: number[];
}

interface Emits {
  (e: 'update:visible', value: boolean): void;
  (e: 'confirm', ids: number[]): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const loading = ref(false);
const userList = ref<UserApi.ListItem[]>([]);
const selectedRowKeys = ref<Array<number | string>>([]);
const deptTreeData = ref<any[]>([]);
const roleOptions = ref<Array<{ label: string; value: number }>>([]);
const postOptions = ref<Array<{ label: string; value: number }>>([]);

const searchForm = reactive({
  username: '',
  nickname: '',
  phone: '',
  email: '',
  dept_ids: [] as number[],
  role_id: undefined as number | undefined,
  post_id: undefined as number | undefined,
});

const pagination = reactive({
  current: 1,
  pageSize: 10,
  pageSizeOptions: [10, 20, 50, 100],
  showJumper: true,
  showPageSize: true,
  total: 0,
});

const columns = [
  { align: 'center' as const, colKey: 'row-select', type: 'multiple' as const, width: 50 },
  { align: 'center' as const, colKey: 'username', title: $t('system.user.username'), minWidth: 120 },
  { align: 'center' as const, colKey: 'nickname', title: $t('system.user.nickname'), minWidth: 120 },
  { align: 'center' as const, colKey: 'phone', title: $t('system.user.phone'), minWidth: 140 },
  { align: 'center' as const, colKey: 'email', title: $t('system.user.email'), minWidth: 180 },
  { align: 'center' as const, colKey: 'dept_name', title: $t('system.user.dept'), minWidth: 160 },
  { align: 'center' as const, colKey: 'role_name', title: $t('system.user.role'), minWidth: 140 },
  { align: 'center' as const, colKey: 'post_name', title: $t('system.user.post'), minWidth: 140 },
] satisfies PrimaryTableCol<TableRowData>[];

const dialogVisible = computed({
  get: () => props.visible,
  set: (value) => emit('update:visible', value),
});

async function fetchUserList(page = 1) {
  loading.value = true;
  try {
    const response = await getUserList({
      page,
      pageSize: pagination.pageSize,
      username: searchForm.username || undefined,
      nickname: searchForm.nickname || undefined,
      phone: searchForm.phone || undefined,
      email: searchForm.email || undefined,
      dept_ids: searchForm.dept_ids.length > 0 ? searchForm.dept_ids : undefined,
      role_id: searchForm.role_id,
      post_id: searchForm.post_id,
    });
    userList.value = response.items ?? [];
    pagination.total = Number(response.pageInfo?.total || response.total || 0);
  } catch (error) {
    logger.error(error);
  } finally {
    loading.value = false;
  }
}

async function fetchDeptTree() {
  try {
    const response = await getDeptTree();
    deptTreeData.value = response || [];
  } catch (error) {
    logger.error(error);
  }
}

async function fetchRoleOptions() {
  try {
    const response = await getRoleList();
    const items = Array.isArray(response) ? response : (response.items ?? []);
    roleOptions.value = items.map((item: any) => ({ label: item.name, value: item.id }));
  } catch (error) {
    logger.error(error);
  }
}

async function fetchPostOptions() {
  try {
    const response = await getPostList();
    const items = Array.isArray(response) ? response : (response.items ?? []);
    postOptions.value = items.map((item: any) => ({ label: item.name, value: item.id }));
  } catch (error) {
    logger.error(error);
  }
}

function handleSearch() {
  pagination.current = 1;
  void fetchUserList(1);
}

function handleReset() {
  searchForm.username = '';
  searchForm.nickname = '';
  searchForm.phone = '';
  searchForm.email = '';
  searchForm.dept_ids = [];
  searchForm.role_id = undefined;
  searchForm.post_id = undefined;
  pagination.current = 1;
  void fetchUserList(1);
}

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  pagination.current = pageInfo.current;
  pagination.pageSize = pageInfo.pageSize;
  void fetchUserList(pageInfo.current);
}

function handleSelectChange(keys: Array<number | string>) {
  selectedRowKeys.value = keys;
}

function handleConfirm() {
  const ids = selectedRowKeys.value
    .map((key) => Number(key))
    .filter((id) => !Number.isNaN(id));
  emit('confirm', ids);
  emit('update:visible', false);
}

watch(
  () => props.visible,
  (visible) => {
    if (visible) {
      selectedRowKeys.value = props.selectedIds.map(String);
      void fetchUserList(1);
      void fetchDeptTree();
      void fetchRoleOptions();
      void fetchPostOptions();
    }
  },
);
</script>

<template>
  <Dialog
    v-model:visible="dialogVisible"
    :header="$t('system.notice.selectUser')"
    width="1200px"
    :close-on-overlay-click="false"
    :footer="false"
  >
    <div class="space-y-4">
      <div class="rounded-md bg-card p-4">
        <Form :data="searchForm" label-width="80px" layout="inline" colon>
          <div class="grid grid-cols-4 gap-x-4 gap-y-3">
            <FormItem :label="$t('system.user.username')" name="username">
              <Input
                :value="searchForm.username"
                :placeholder="$t('ui.placeholder.input', [$t('system.user.username')])"
                clearable
                @change="(val: any) => (searchForm.username = val)"
              />
            </FormItem>
            <FormItem :label="$t('system.user.nickname')" name="nickname">
              <Input
                :value="searchForm.nickname"
                :placeholder="$t('ui.placeholder.input', [$t('system.user.nickname')])"
                clearable
                @change="(val: any) => (searchForm.nickname = val)"
              />
            </FormItem>
            <FormItem :label="$t('system.user.phone')" name="phone">
              <Input
                :value="searchForm.phone"
                :placeholder="$t('ui.placeholder.input', [$t('system.user.phone')])"
                clearable
                @change="(val: any) => (searchForm.phone = val)"
              />
            </FormItem>
            <FormItem :label="$t('system.user.email')" name="email">
              <Input
                :value="searchForm.email"
                :placeholder="$t('ui.placeholder.input', [$t('system.user.email')])"
                clearable
                @change="(val: any) => (searchForm.email = val)"
              />
            </FormItem>
            <FormItem :label="$t('system.user.dept')" name="dept_ids">
              <TreeSelect
                :value="searchForm.dept_ids"
                :data="deptTreeData"
                :keys="{ value: 'id', label: 'label', children: 'children' }"
                :multiple="true"
                :tree-props="{ checkStrictly: true }"
                :placeholder="$t('ui.placeholder.select', [$t('system.user.dept')])"
                clearable
                class="w-full"
                @change="(val: any) => (searchForm.dept_ids = val || [])"
              />
            </FormItem>
            <FormItem :label="$t('system.user.role')" name="role_id">
              <Select
                :value="searchForm.role_id"
                :options="roleOptions"
                :placeholder="$t('ui.placeholder.select', [$t('system.user.role')])"
                clearable
                class="w-full"
                @change="(val: any) => (searchForm.role_id = val)"
              />
            </FormItem>
            <FormItem :label="$t('system.user.post')" name="post_id">
              <Select
                :value="searchForm.post_id"
                :options="postOptions"
                :placeholder="$t('ui.placeholder.select', [$t('system.user.post')])"
                clearable
                class="w-full"
                @change="(val: any) => (searchForm.post_id = val)"
              />
            </FormItem>
            <div class="flex justify-end gap-2 pb-[7px]">
              <Button theme="default" @click="handleReset">
                {{ $t('common.reset') }}
              </Button>
              <Button theme="primary" @click="handleSearch">
                {{ $t('common.query') }}
              </Button>
            </div>
          </div>
        </Form>
      </div>

      <Table
        :columns="columns"
        :data="userList"
        :loading="loading"
        :pagination="pagination"
        :selected-row-keys="selectedRowKeys"
        row-key="id"
        hover
        stripe
        :reserve-selected-row-on-paginate="true"
        @page-change="handlePageChange"
        @select-change="handleSelectChange"
      />

      <div class="flex justify-end gap-2">
        <Button theme="default" variant="outline" @click="dialogVisible = false">
          {{ $t('common.cancel') }}
        </Button>
        <Button theme="primary" @click="handleConfirm">
          {{ $t('common.confirm') }}
        </Button>
      </div>
    </div>
  </Dialog>
</template>
