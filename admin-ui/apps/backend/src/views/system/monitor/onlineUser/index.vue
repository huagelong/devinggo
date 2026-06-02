<script lang="ts" setup>
import type { MonitorApi } from '#/api/system/monitor';

import { onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { Page } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';
import { getOnlineUserPageList, kickUser } from '#/api/system/monitor';
import { logger } from '#/utils/logger';

import { SearchIcon } from 'tdesign-icons-vue-next';
import { Button, Input, Space, Table } from 'tdesign-vue-next';

defineOptions({ name: 'SystemOnlineUser' });

const loading = ref(false);
const tableData = ref<MonitorApi.OnlineUserItem[]>([]);
const total = ref(0);
const searchUsername = ref('');

const columns = [
  { colKey: 'username', title: $t('system.monitor.onlineUser.username'), width: 180 },
  { colKey: 'nickname', title: $t('system.monitor.onlineUser.nickname'), width: 180 },
  { colKey: 'app_id', title: 'App ID', width: 120 },
  { colKey: 'login_ip', title: $t('system.monitor.onlineUser.loginIp'), width: 180 },
  { colKey: 'login_time', title: $t('system.monitor.onlineUser.loginTime'), width: 180 },
  { colKey: 'action', title: $t('common.action'), width: 120 },
];

async function fetchOnlineUsers() {
  loading.value = true;
  try {
    const params: MonitorApi.OnlineUserQuery = {
      page: 1,
      page_size: 20,
    };
    if (searchUsername.value) {
      params.username = searchUsername.value;
    }
    const response = await getOnlineUserPageList(params);
    tableData.value = response.items || [];
    total.value = Number(response.pageInfo?.total || response.total || 0);
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
  void fetchOnlineUsers();
}

function handleReset() {
  searchUsername.value = '';
  void fetchOnlineUsers();
}

onMounted(() => {
  void fetchOnlineUsers();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <div class="flex items-center gap-4">
          <Input
            v-model="searchUsername"
            :placeholder="$t('system.monitor.onlineUser.searchPlaceholder')"
            clearable
            class="w-64"
            @enter="handleSearch"
          />
          <Space>
            <Button theme="primary" @click="handleSearch">
              <template #icon><SearchIcon /></template>
              {{ $t('common.query') }}
            </Button>
            <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
          </Space>
        </div>
      </div>

      <div class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <Table
          :columns="columns"
          :data="tableData"
          :loading="loading"
          row-key="id"
          hover
          stripe
        >
          <template #action="{ row }">
            <Button
              size="small"
              theme="danger"
              variant="outline"
              @click="handleKick(row)"
            >
              {{ $t('system.monitor.onlineUser.forceLogout') }}
            </Button>
          </template>
        </Table>
      </div>
    </div>
  </Page>
</template>
