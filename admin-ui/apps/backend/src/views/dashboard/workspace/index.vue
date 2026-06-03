<script lang="ts" setup>
import type { MessageApi } from '#/api/core/message';
import type { NoticeApi } from '#/api/system/notice';

import { computed, onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

import { $t } from '@vben/locales';
import { useAccessStore, useUserStore } from '@vben/stores';
import { VbenIcon } from '@vben/common-ui';

import {
  getQueueMessageReceiveListApi,
  updateQueueMessageReadStatusApi,
} from '#/api/core/message';
import { getNoticePageList } from '#/api/system/notice';
import { logger } from '#/utils/logger';

const userStore = useUserStore();
const accessStore = useAccessStore();
const router = useRouter();

// 加载状态
const loading = ref({
  messages: false,
  notices: false,
});

// 未读消息列表
const unreadMessages = ref<MessageApi.QueueMessageItem[]>([]);
const unreadCount = ref(0);

// 系统公告列表
const noticeList = ref<NoticeApi.ListItem[]>([]);

// 快捷导航候选配置（通过菜单原始标题匹配，自动获取正确路径）
const quickNavCandidates = [
  {
    color: '#3b82f6',
    description: $t('system.user.title'),
    icon: 'lucide:users',
    matchTitle: '用户管理',
    title: $t('system.user.title'),
  },
  {
    color: '#10b981',
    description: $t('system.role.title'),
    icon: 'lucide:shield',
    matchTitle: '角色管理',
    title: $t('system.role.title'),
  },
  {
    color: '#f59e0b',
    description: $t('system.dept.title'),
    icon: 'lucide:building',
    matchTitle: '部门管理',
    title: $t('system.dept.title'),
  },
  {
    color: '#8b5cf6',
    description: $t('system.menu.manageTitle'),
    icon: 'lucide:menu',
    matchTitle: '菜单管理',
    title: $t('system.menu.manageTitle'),
  },
  {
    color: '#ef4444',
    description: $t('system.dict.title'),
    icon: 'lucide:book-open',
    matchTitle: '数据字典',
    title: $t('system.dict.title'),
  },
  {
    color: '#06b6d4',
    description: $t('system.notice.manageTitle'),
    icon: 'lucide:bell',
    matchTitle: '系统公告',
    title: $t('system.notice.manageTitle'),
  },
  {
    color: '#ec4899',
    description: $t('system.config.title'),
    icon: 'lucide:settings',
    matchTitle: '配置',
    title: $t('system.config.title'),
  },
  {
    color: '#6366f1',
    description: $t('system.logs.title'),
    icon: 'lucide:file-text',
    matchTitle: '登录日志',
    title: $t('system.logs.title'),
  },
];

// 从用户菜单中查找指定标题的菜单项
function findMenuByTitle(menus: any[], title: string): any | undefined {
  for (const menu of menus) {
    if (menu.name === title) {
      return menu;
    }
    if (menu.children?.length > 0) {
      const found = findMenuByTitle(menu.children, title);
      if (found) return found;
    }
  }
  return undefined;
}

// 根据用户权限过滤后的快捷导航
const quickNavItems = computed(() => {
  return quickNavCandidates
    .map((candidate) => {
      const menu = findMenuByTitle(accessStore.accessMenus, candidate.matchTitle);
      if (menu) {
        return {
          ...candidate,
          path: menu.path,
          icon: menu.icon || candidate.icon,
        };
      }
      return null;
    })
    .filter((item): item is NonNullable<typeof item> => item !== null);
});

// 获取未读消息
async function fetchUnreadMessages() {
  loading.value.messages = true;
  try {
    const res = await getQueueMessageReceiveListApi({
      page: 1,
      pageSize: 5,
      read_status: '1',
    });
    unreadMessages.value = res.items || [];
    unreadCount.value = res.total || 0;
  } catch (error) {
    logger.error('Failed to fetch unread messages', error);
  } finally {
    loading.value.messages = false;
  }
}

// 获取系统公告
async function fetchNoticeList() {
  loading.value.notices = true;
  try {
    const res = await getNoticePageList({
      page: 1,
      pageSize: 5,
    });
    noticeList.value = res.items || [];
  } catch (error) {
    logger.error('Failed to fetch notice list', error);
  } finally {
    loading.value.notices = false;
  }
}

// 标记消息为已读
async function markAsRead(message: MessageApi.QueueMessageItem) {
  try {
    await updateQueueMessageReadStatusApi({ ids: [message.id] });
    // 刷新列表
    await fetchUnreadMessages();
  } catch (error) {
    logger.error('Failed to mark message as read', error);
  }
}

// 跳转到消息中心
function goToMessageCenter() {
  router.push('/dashboard/message');
}

// 跳转到指定路径
function navigateTo(path: string) {
  router.push(path);
}

// 获取指定标题菜单的路径
function getMenuPath(title: string): string | undefined {
  const menu = findMenuByTitle(accessStore.accessMenus, title);
  return menu?.path;
}

// 公告页面路径
const noticePath = computed(() => getMenuPath('系统公告'));

// 格式化时间
function formatTime(time: string) {
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
  return date.toLocaleDateString();
}

// 获取问候语
function getGreeting() {
  const hour = new Date().getHours();
  if (hour < 6) return '凌晨好';
  if (hour < 9) return '早上好';
  if (hour < 12) return '上午好';
  if (hour < 14) return '中午好';
  if (hour < 18) return '下午好';
  return '晚上好';
}

onMounted(() => {
  fetchUnreadMessages();
  fetchNoticeList();
});
</script>

<template>
  <div class="p-5">
    <!-- 欢迎区 -->
    <div class="mb-6 flex items-center gap-4 rounded-lg bg-white p-6 shadow-sm">
      <div class="h-16 w-16 overflow-hidden rounded-full bg-gray-100">
        <img
          v-if="userStore.userInfo?.avatar"
          :src="userStore.userInfo.avatar"
          alt="avatar"
          class="h-full w-full object-cover"
        />
        <div
          v-else
          class="flex h-full w-full items-center justify-center text-2xl text-gray-400"
        >
          {{ userStore.userInfo?.realName?.charAt(0) || 'U' }}
        </div>
      </div>
      <div>
        <h1 class="text-xl font-semibold text-gray-900">
          {{ getGreeting() }}，{{ userStore.userInfo?.realName || userStore.userInfo?.username }}
        </h1>
        <p class="mt-1 text-sm text-gray-500">
          {{ $t('dashboard.workspace.welcomeDesc') }}
        </p>
      </div>
      <div class="ml-auto text-right">
        <div class="text-sm text-gray-500">{{ $t('dashboard.workspace.today') }}</div>
        <div class="text-lg font-semibold text-gray-900">
          {{ new Date().toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric', weekday: 'long' }) }}
        </div>
      </div>
    </div>

    <div class="flex flex-col gap-6 lg:flex-row">
      <!-- 左侧区域 -->
      <div class="flex-1">
        <!-- 快捷入口 -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">
            {{ $t('dashboard.workspace.quickNav') }}
          </h2>
          <div v-if="quickNavItems.length === 0" class="py-8 text-center text-gray-500">
            {{ $t('dashboard.workspace.noQuickNav') }}
          </div>
          <div v-else class="grid grid-cols-2 gap-4 sm:grid-cols-3 lg:grid-cols-4">
            <div
              v-for="item in quickNavItems"
              :key="item.path"
              class="group cursor-pointer rounded-lg border border-gray-100 p-4 transition-all hover:shadow-md"
              @click="navigateTo(item.path)"
            >
              <div
                class="mb-3 flex h-10 w-10 items-center justify-center rounded-lg"
                :style="{ backgroundColor: item.color + '15' }"
              >
                <VbenIcon
                  :icon="item.icon"
                  class="h-5 w-5"
                  :style="{ color: item.color }"
                  fallback
                />
              </div>
              <div class="text-sm font-medium text-gray-900">{{ item.title }}</div>
              <div class="mt-1 text-xs text-gray-500">{{ item.description }}</div>
            </div>
          </div>
        </div>

        <!-- 未读消息 -->
        <div class="rounded-lg bg-white p-6 shadow-sm">
          <div class="mb-4 flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-900">
              {{ $t('dashboard.workspace.myMessages') }}
            </h2>
            <div class="flex items-center gap-2">
              <span
                v-if="unreadCount > 0"
                class="rounded-full bg-red-100 px-2 py-0.5 text-xs font-medium text-red-600"
              >
                {{ unreadCount }}{{ $t('dashboard.workspace.unreadCount') }}
              </span>
              <button
                class="text-sm text-blue-600 hover:text-blue-700"
                @click="goToMessageCenter"
              >
                {{ $t('dashboard.workspace.viewAll') }}
              </button>
            </div>
          </div>

          <div v-if="loading.messages" class="py-8 text-center text-gray-500">
            {{ $t('common.loading') }}
          </div>

          <div v-else-if="unreadMessages.length === 0" class="py-8 text-center text-gray-500">
            {{ $t('dashboard.workspace.noUnreadMessages') }}
          </div>

          <div v-else class="space-y-3">
            <div
              v-for="message in unreadMessages"
              :key="message.id"
              class="group flex items-start gap-3 rounded-lg border border-gray-100 p-3 transition-all hover:bg-gray-50"
            >
              <div class="mt-0.5 h-2 w-2 flex-shrink-0 rounded-full bg-red-500"></div>
              <div class="min-w-0 flex-1">
                <div class="flex items-center justify-between">
                  <h3 class="truncate text-sm font-medium text-gray-900">
                    {{ message.title }}
                  </h3>
                  <span class="ml-2 flex-shrink-0 text-xs text-gray-400">
                    {{ formatTime(message.created_at) }}
                  </span>
                </div>
                <p class="mt-1 line-clamp-2 text-xs text-gray-500">
                  {{ message.content }}
                </p>
              </div>
              <button
                class="ml-2 flex-shrink-0 rounded px-2 py-1 text-xs text-blue-600 opacity-0 transition-opacity hover:bg-blue-50 group-hover:opacity-100"
                @click.stop="markAsRead(message)"
              >
                {{ $t('dashboard.workspace.markAsRead') }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- 右侧区域 -->
      <div class="w-full lg:w-80">
        <!-- 个人信息卡片 -->
        <div class="mb-6 rounded-lg bg-white p-6 shadow-sm">
          <h2 class="mb-4 text-lg font-semibold text-gray-900">
            {{ $t('dashboard.workspace.myInfo') }}
          </h2>
          <div class="space-y-3">
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">{{ $t('system.user.username') }}</span>
              <span class="font-medium text-gray-900">{{ userStore.userInfo?.username }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">{{ $t('system.user.nickname') }}</span>
              <span class="font-medium text-gray-900">{{ userStore.userInfo?.realName || '-' }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">{{ $t('system.user.role') }}</span>
              <span class="font-medium text-gray-900">{{ userStore.userInfo?.roles?.join(', ') || '-' }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span class="text-gray-500">{{ $t('system.user.lastLogin') }}</span>
              <span class="font-medium text-gray-900">{{ $t('dashboard.workspace.justNow') }}</span>
            </div>
          </div>
          <button
            class="mt-4 w-full rounded-lg bg-blue-600 py-2 text-sm font-medium text-white transition-colors hover:bg-blue-700"
            @click="navigateTo('/dashboard/profile')"
          >
            {{ $t('dashboard.workspace.editProfile') }}
          </button>
        </div>

        <!-- 系统公告 -->
        <div class="rounded-lg bg-white p-6 shadow-sm">
          <div class="mb-4 flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-900">
              {{ $t('dashboard.workspace.systemNotice') }}
            </h2>
            <button
              v-if="noticePath"
              class="text-sm text-blue-600 hover:text-blue-700"
              @click="navigateTo(noticePath)"
            >
              {{ $t('dashboard.workspace.more') }}
            </button>
          </div>
          <div v-if="loading.notices" class="py-8 text-center text-gray-500">
            {{ $t('common.loading') }}
          </div>

          <div v-else-if="noticeList.length === 0" class="py-8 text-center text-gray-500">
            {{ $t('dashboard.workspace.noNotice') }}
          </div>

          <div v-else class="space-y-3">
            <div
              v-for="notice in noticeList"
              :key="notice.id"
              :class="[
                'rounded-lg border border-gray-100 p-3 transition-all',
                noticePath ? 'cursor-pointer hover:bg-gray-50' : '',
              ]"
              @click="noticePath && navigateTo(noticePath)"
            >
              <div class="flex items-center gap-2">
                <span class="rounded bg-blue-100 px-1.5 py-0.5 text-xs font-medium text-blue-600">
                  {{ $t('dashboard.workspace.announcement') }}
                </span>
                <span class="truncate text-sm text-gray-900">{{ notice.title }}</span>
              </div>
              <p class="mt-1 text-xs text-gray-500">{{ notice.created_at ? new Date(notice.created_at).toLocaleDateString() : '' }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
