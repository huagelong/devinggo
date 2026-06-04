<script lang="ts" setup>
import type { NotificationItem } from '@vben/layouts';
import type { MessageApi } from '#/api/core/message';

import { computed, onMounted, onUnmounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

import { AuthenticationLoginExpiredModal } from '@vben/common-ui';
import { useWatermark } from '@vben/hooks';
import { BookOpenText, CircleHelp, SvgGithubIcon } from '@vben/icons';
import {
  BasicLayout,
  LockScreen,
  Notification,
  UserDropdown,
} from '@vben/layouts';
import { preferences } from '@vben/preferences';
import { useAccessStore, useUserStore } from '@vben/stores';
import { openWindow } from '@vben/utils';

import { dialog, message } from '#/adapter/tdesign';
import { Dialog } from 'tdesign-vue-next';
import { $t } from '#/locales';
import { useAuthStore } from '#/store';
import { clearAllCache } from '#/api/system/monitor';
import {
  deleteQueueMessageApi,
  getQueueMessageReceiveListApi,
  updateQueueMessageReadStatusApi,
} from '#/api/core/message';
import { useRealtimeNotifications } from '#/composables/pusher';
import { sanitizeHtml } from '#/utils/sanitize';
import { logger } from '#/utils/logger';
import LoginForm from '#/views/_core/authentication/login.vue';

const router = useRouter();

const {
  start: startRealtime,
  stop: stopRealtime,
  latestNotification,
} = useRealtimeNotifications();

// 通知消息列表
const notifications = ref<NotificationItem[]>([]);
const rawMessages = ref<MessageApi.QueueMessageItem[]>([]);

// 弹窗相关
const detailVisible = ref(false);
const detailData = ref<MessageApi.QueueMessageItem | null>(null);

// 定时器
let pollTimer: ReturnType<typeof setInterval> | null = null;
const userStore = useUserStore();
const authStore = useAuthStore();
const accessStore = useAccessStore();
const { destroyWatermark, updateWatermark } = useWatermark();

const showDot = computed(() =>
  notifications.value.some((item) => !item.isRead),
);

// 格式化相对时间
function formatRelativeTime(time: string): string {
  if (!time) return '';
  const date = new Date(time);
  const now = new Date();
  const diff = now.getTime() - date.getTime();
  const minutes = Math.floor(diff / 60000);
  const hours = Math.floor(diff / 3600000);
  const days = Math.floor(diff / 86400000);

  if (minutes < 1) return '刚刚';
  if (minutes < 60) return `${minutes}分钟前`;
  if (hours < 24) return `${hours}小时前`;
  if (days < 30) return `${days}天前`;
  return date.toLocaleDateString('zh-CN');
}

// 获取未读消息
async function fetchMessages() {
  try {
    const res = await getQueueMessageReceiveListApi({
      page: 1,
      pageSize: 10,
      read_status: '1',
    });
    const items = res?.items ?? [];
    rawMessages.value = items;
    notifications.value = items.map((item) => ({
      id: item.id,
      avatar: item.send_user?.avatar || preferences.app.defaultAvatar,
      date: formatRelativeTime(item.created_at),
      isRead: item.read_status === 2,
      message: item.content
        ? item.content.replace(/<[^>]*>/g, '').substring(0, 50)
        : '',
      title: item.title,
    }));
  } catch (error) {
    logger.error('Failed to fetch notifications', error);
  }
}

// 查看详情
async function handleViewDetail(item: NotificationItem) {
  const msg = rawMessages.value.find((m) => m.id === item.id);

  if (msg) {
    detailData.value = msg;
    detailVisible.value = true;
    // 自动标记为已读
    try {
      await updateQueueMessageReadStatusApi({ ids: [item.id as number] });
      await fetchMessages();
    } catch (error) {
      logger.error('Failed to mark message as read', error);
    }
  }
}

// 标记已读
async function markRead(id: number | string) {
  try {
    await updateQueueMessageReadStatusApi({ ids: [id as number] });
    await fetchMessages();
  } catch (error) {
    logger.error('Failed to mark message as read', error);
  }
}

// 删除单条
async function remove(id: number | string) {
  try {
    await deleteQueueMessageApi({ ids: [id as number] });
    message.success($t('common.deleteSuccess'));
    await fetchMessages();
  } catch (error) {
    logger.error('Failed to delete message', error);
  }
}

// 全部已读
async function handleMakeAll() {
  const unreadIds = notifications.value
    .filter((item) => !item.isRead)
    .map((item) => item.id as number);
  if (unreadIds.length === 0) return;
  try {
    await updateQueueMessageReadStatusApi({ ids: unreadIds });
    await fetchMessages();
  } catch (error) {
    logger.error('Failed to mark all as read', error);
  }
}

// 清空
async function handleNoticeClear() {
  const ids = notifications.value.map((item) => item.id as number);
  if (ids.length === 0) return;
  try {
    await deleteQueueMessageApi({ ids });
    message.success($t('common.deleteSuccess'));
    await fetchMessages();
  } catch (error) {
    logger.error('Failed to clear notifications', error);
  }
}

// 查看全部
function handleViewAll() {
  router.push('/dashboard/message');
}

// 轮询刷新
function startPolling() {
  pollTimer = setInterval(() => {
    fetchMessages();
  }, 60000);
}

function stopPolling() {
  if (pollTimer) {
    clearInterval(pollTimer);
    pollTimer = null;
  }
}

onMounted(() => {
  fetchMessages();
  startPolling();
  try {
    startRealtime();
  } catch (error) {
    logger.error('Failed to start realtime notifications', error);
  }
});

onUnmounted(() => {
  stopPolling();
  stopRealtime();
});

watch(latestNotification, (notification) => {
  if (notification) {
    fetchMessages();
  }
});

const menus = computed(() => [
  {
    handler: () => {
      router.push('/dashboard/profile');
    },
    icon: 'lucide:user',
    text: $t('page.auth.profile'),
  },
  {
    handler: () => {
      openWindow('https://devinggo.devinghub.com', {
        target: '_blank',
      });
    },
    icon: BookOpenText,
    text: $t('ui.widgets.document'),
  },
  {
    handler: () => {
      openWindow('https://github.com/huagelong/devinggo', {
        target: '_blank',
      });
    },
    icon: SvgGithubIcon,
    text: 'GitHub',
  },
  {
    handler: () => {
      openWindow('https://github.com/huagelong/devinggo/issues', {
        target: '_blank',
      });
    },
    icon: CircleHelp,
    text: $t('ui.widgets.qa'),
  },
]);

const avatar = computed(() => {
  return userStore.userInfo?.avatar ?? preferences.app.defaultAvatar;
});

async function handleLogout() {
  await authStore.logout(false);
}

function handleClearAllCache() {
  const dialogInstance = dialog.confirm({
    body: $t('common.confirmClearAllCache'),
    header: $t('common.prompt'),
    onClose: () => dialogInstance.hide(),
    onConfirm: async () => {
      try {
        await clearAllCache();
        message.success($t('common.clearCacheSuccess2'));
      } catch {
        message.error($t('common.clearCacheFailed2'));
      } finally {
        dialogInstance.hide();
      }
    },
  });
}

watch(
  () => ({
    enable: preferences.app.watermark,
    content: preferences.app.watermarkContent,
  }),
  async ({ enable, content }) => {
    if (enable) {
      await updateWatermark({
        content:
          content ||
          `${userStore.userInfo?.username} - ${userStore.userInfo?.realName}\n${new Date().toLocaleDateString(preferences.app.locale)}`,
      });
    } else {
      destroyWatermark();
    }
  },
  {
    immediate: true,
  },
);
</script>

<template>
  <BasicLayout @clear-preferences-and-logout="handleLogout">
    <template #user-dropdown>
      <UserDropdown
        :avatar
        :menus
        :text="userStore.userInfo?.realName"
        :description="userStore.userInfo?.email ?? ''"
        tag-text="Pro"
        @logout="handleLogout"
        @clearAllCache="handleClearAllCache"
      />
    </template>
    <template #notification>
      <Notification
        :dot="showDot"
        :notifications="notifications"
        @clear="handleNoticeClear"
        @read="(item) => item.id && markRead(item.id)"
        @remove="(item) => item.id && remove(item.id)"
        @make-all="handleMakeAll"
        @view-all="handleViewAll"
        @click="handleViewDetail"
      />
    </template>
    <template #extra>
      <AuthenticationLoginExpiredModal
        v-model:open="accessStore.loginExpired"
        :avatar
      >
        <LoginForm />
      </AuthenticationLoginExpiredModal>
    </template>
    <template #lock-screen>
      <LockScreen :avatar @to-login="handleLogout" />
    </template>
  </BasicLayout>

  <!-- 消息详情弹窗 -->
  <Dialog
    v-model:visible="detailVisible"
    :header="detailData?.title || ''"
    width="800px"
    :footer="false"
    destroy-on-close
    placement="center"
  >
    <div class="flex flex-col gap-4 py-4 px-2">
      <div class="flex justify-between text-gray-500 text-sm border-b pb-2">
        <span>{{ detailData?.send_user?.nickname || '' }}</span>
        <span>{{ detailData?.created_at || '' }}</span>
      </div>
      <div
        class="bg-gray-50 p-4 rounded min-h-[200px]"
        v-html="sanitizeHtml(detailData?.content || '')"
      ></div>
    </div>
  </Dialog>
</template>
