<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { onMounted, onUnmounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import { getOnlineUserPageList, kickUser } from '#/api/system/monitor';
import { getAppPageList } from '#/api/system/app';
import type { AppApi } from '#/api/system/app';
import { logger } from '#/utils/logger';

import { SearchIcon, DeleteIcon, FullscreenIcon, FullscreenExitIcon, RefreshIcon } from 'tdesign-icons-vue-next';
import { Button, Form, FormItem, Input, Select, Table, Tooltip } from 'tdesign-vue-next';

defineOptions({ name: 'SystemOnlineUser' });

const loading = ref(false);
const tableData = ref<MonitorApi.OnlineUserItem[]>([]);
const searchForm = ref({
  username: '',
  app_id: undefined as number | undefined,
});

const appOptions = ref<{ label: string; value: number }[]>([]);
const appMap = ref<Record<string | number, string>>({});

const tableContainerRef = ref<HTMLElement>();
const isFullscreen = ref(false);

const pagination = ref({
  current: 1,
  pageSize: 20,
  total: 0,
  pageSizeOptions: [],
  showJumper: false,
  showPageSize: false,
});

const columns = [
  { colKey: 'username', title: '用户账户', width: 180 },
  { colKey: 'nickname', title: '用户昵称', width: 180 },
  { colKey: 'app_id', title: 'App', width: 120 },
  { colKey: 'login_ip', title: $t('system.monitor.onlineUser.loginIp'), width: 180 },
  { colKey: 'login_time', title: $t('system.monitor.onlineUser.loginTime'), width: 180 },
  { colKey: 'action', title: $t('common.action'), width: 120 },
];

async function fetchAppList() {
  try {
    const response = await getAppPageList({ page: 1, pageSize: 1000 });
    const apps = response.items || [];
    appOptions.value = apps.map((app: AppApi.ListItem) => ({
      label: app.app_name || app.app_id || String(app.id),
      value: app.id,
    }));
    appMap.value = apps.reduce((map: Record<string | number, string>, app: AppApi.ListItem) => {
      map[app.id] = app.app_name || app.app_id || String(app.id);
      return map;
    }, {});
  } catch (error) {
    logger.error(error);
  }
}

async function fetchOnlineUsers() {
  loading.value = true;
  try {
    const params: MonitorApi.OnlineUserQuery = {
      page: pagination.value.current,
      page_size: pagination.value.pageSize,
    };
    if (searchForm.value.username) {
      params.username = searchForm.value.username;
    }
    if (searchForm.value.app_id !== undefined) {
      params.app_id = searchForm.value.app_id;
    }
    const response = await getOnlineUserPageList(params);
    tableData.value = response.items || [];
    pagination.value.total = Number(response.pageInfo?.total || response.total || 0);
    console.log('在线用户数据:', response.items);
  } catch (error) {
    logger.error(error);
    message.error($t('common.onlineUserFetchFailed'));
  } finally {
    loading.value = false;
  }
}

async function handleKick(row: MonitorApi.OnlineUserItem) {
  try {
    await kickUser({ id: row.id, app_id: row.app_id });
    message.success($t('common.forceLogoutSuccess'));
    await fetchOnlineUsers();
  } catch (error) {
    logger.error(error);
    message.error($t('common.forceLogoutFailed'));
  }
}

function handleSearch() {
  pagination.value.current = 1;
  void fetchOnlineUsers();
}

function handleReset() {
  searchForm.value = { username: '', app_id: undefined };
  pagination.value.current = 1;
  void fetchOnlineUsers();
}

function handlePageChange(pageInfo: { current: number; pageSize: number }) {
  pagination.value.current = pageInfo.current;
  pagination.value.pageSize = pageInfo.pageSize;
  void fetchOnlineUsers();
}

function handleFullscreenChange() {
  isFullscreen.value = !!document.fullscreenElement;
}

function toggleFullscreen() {
  if (document.fullscreenElement) {
    document.exitFullscreen();
    return;
  }
  tableContainerRef.value?.requestFullscreen();
}

onMounted(() => {
  void fetchAppList();
  void fetchOnlineUsers();
  document.addEventListener('fullscreenchange', handleFullscreenChange);
});

onUnmounted(() => {
  document.removeEventListener('fullscreenchange', handleFullscreenChange);
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="100px" layout="inline" colon>
          <div class="grid grid-cols-2 gap-x-4 gap-y-3 w-full">
            <FormItem label="用户账户" name="username">
              <Input
                v-model="searchForm.username"
                placeholder="请输入用户账户"
                clearable
                @enter="handleSearch"
              />
            </FormItem>
            <FormItem label="App" name="app_id">
              <Select
                v-model="searchForm.app_id"
                :options="appOptions"
                placeholder="请选择App"
                clearable
                class="w-full"
              />
            </FormItem>
          </div>
          <div class="flex justify-end gap-2 pt-2">
            <Button theme="default" @click="handleReset">
              <template #icon><DeleteIcon /></template>
              重置
            </Button>
            <Button theme="primary" @click="handleSearch">
              <template #icon><SearchIcon /></template>
              搜索
            </Button>
          </div>
        </Form>
      </div>

      <div ref="tableContainerRef" class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <div class="mb-3 flex items-center justify-end gap-2">
          <Tooltip :content="isFullscreen ? $t('common.exitFullscreen') : $t('common.fullscreen')">
            <Button shape="square" variant="outline" @click="toggleFullscreen">
              <template #icon>
                <FullscreenExitIcon v-if="isFullscreen" />
                <FullscreenIcon v-else />
              </template>
            </Button>
          </Tooltip>
          <Button shape="square" variant="outline" @click="fetchOnlineUsers">
            <template #icon><RefreshIcon /></template>
          </Button>
        </div>
        <div class="min-h-0 flex-1 overflow-hidden">
          <Table
            :columns="columns"
            :data="tableData"
            :loading="loading"
            :pagination="pagination"
            row-key="id"
            hover
            stripe
            @page-change="handlePageChange"
          >
          <template #app_id="{ row }">
            {{ appMap[row.app_id] || row.app_id || '-' }}
          </template>
          <template #action="{ row }">
            <Button
              size="small"
              theme="primary"
              variant="text"
              @click="handleKick(row)"
            >
              强制退出
            </Button>
          </template>
        </Table>
        </div>
      </div>
    </div>
  </Page>
</template>
