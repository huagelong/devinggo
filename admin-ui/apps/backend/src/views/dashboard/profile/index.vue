<script setup lang="ts">
import type { LogApi } from '#/api/system/log';

import { onMounted, reactive, ref } from 'vue';

// 导入用户信息相关的 Store
import { useUserStore } from '@vben/stores';

import {
  Button,
  Form,
  FormItem,
  Input,
  MessagePlugin,
  TabPanel,
  Tabs,
  Tag,
  Textarea,
  Timeline,
  TimelineItem,
} from 'tdesign-vue-next';

import { $t } from '#/locales';

import {
  getLoginLogListApi,
  getOperationLogListApi,
  modifyPasswordApi,
  updateUserInfoApi,
} from '#/api/core/profile';
import { getSystemInfoApi } from '#/api/core/user';
import { uploadImageFileApi } from '#/api/system/upload';
import { encryptPassword } from '#/store/auth';
import { logger } from '#/utils/logger';

const userStore = useUserStore();

// 左侧 Tabs
const leftTab = ref('info');
// 右侧 Tabs
const rightTab = ref('loginLog');

// 个人资料表单
const userInfoForm = reactive({
  username: '',
  nickname: '',
  phone: '',
  email: '',
  signed: '',
  avatar: '',
});

// 安全设置表单
const securityForm = reactive({
  oldPassword: '',
  newPassword: '',
  newPasswordConfirmation: '',
});

// 日志数据
const loginLogs = ref<LogApi.LoginLogItem[]>([]);
const operationLogs = ref<LogApi.OperLogItem[]>([]);

// 个人资料表单
async function fetchUserInfo() {
  try {
    const res = await getSystemInfoApi();
    if (res && res.user) {
      userInfoForm.username = res.user.username || '';
      userInfoForm.nickname = res.user.nickname || '';
      userInfoForm.phone = res.user.phone || '';
      userInfoForm.email = res.user.email || '';
      userInfoForm.signed = res.user.signed || '';
      userInfoForm.avatar = res.user.avatar || '';
    }
  } catch (error) {
    logger.error($t('common.profileUpdateFailed'), error);
  }
}

// 提交个人资料更新
async function handleUpdateInfo() {
  try {
    await updateUserInfoApi({
      nickname: userInfoForm.nickname,
      phone: userInfoForm.phone,
      email: userInfoForm.email,
      signed: userInfoForm.signed,
    });
    MessagePlugin.success($t('common.profileUpdateSuccess'));
    // 更新完成后重新获取数据
    fetchUserInfo();
  } catch {
    MessagePlugin.error($t('common.profileUpdateFailed'));
  }
}

// 提交修改密码
async function handleUpdatePassword() {
  if (securityForm.newPassword !== securityForm.newPasswordConfirmation) {
    MessagePlugin.error($t('common.passwordMismatch'));
    return;
  }
  try {
    const encrypted = {
      oldPassword: await encryptPassword(securityForm.oldPassword),
      newPassword: await encryptPassword(securityForm.newPassword),
      newPasswordConfirmation: await encryptPassword(
        securityForm.newPasswordConfirmation,
      ),
    };
    await modifyPasswordApi(encrypted);
    MessagePlugin.success($t('common.passwordChangeSuccess'));
    securityForm.oldPassword = '';
    securityForm.newPassword = '';
    securityForm.newPasswordConfirmation = '';
  } catch (error) {
    logger.error($t('common.passwordChangeFailed'), error);
  }
}

// 图片上传处理
function triggerUpload() {
  const fileInput = document.createElement('input');
  fileInput.type = 'file';
  fileInput.accept = 'image/*';
  fileInput.style.display = 'none';
  document.body.appendChild(fileInput);
  fileInput.addEventListener('change', async (e: Event) => {
    const file = (e.target as HTMLInputElement).files?.[0];
    if (!file) return;
    try {
      const res = await uploadImageFileApi(file);
      if (res && res.url) {
        userInfoForm.avatar = res.url;
        await updateUserInfoApi({
          avatar: res.url,
        });
        userStore.setUserInfo({
          ...userStore.userInfo,
          avatar: res.url,
        } as Parameters<typeof userStore.setUserInfo>[0]);
        MessagePlugin.success($t('common.avatarUploadSuccess'));
      }
    } catch (error) {
      logger.error($t('common.uploadFailed'), error);
      MessagePlugin.error($t('common.avatarUploadFailed'));
    } finally {
      fileInput.remove();
    }
  });
  fileInput.click();
}

// 获取日志
async function fetchLogs() {
  try {
    const loginRes = await getLoginLogListApi({ page: 1, pageSize: 10 });
    if (loginRes && loginRes.items) {
      loginLogs.value = loginRes.items;
    }

    const opRes = await getOperationLogListApi({ page: 1, pageSize: 10 });
    if (opRes && opRes.items) {
      operationLogs.value = opRes.items;
    }
  } catch (error) {
    logger.error($t('common.logFetchFailed'), error);
  }
}

onMounted(() => {
  fetchUserInfo();
  fetchLogs();
});
</script>

<template>
  <div class="h-full p-4 overflow-auto bg-[var(--vben-color-background)]">
    <!-- 顶部 Banner -->
    <div
      class="relative flex flex-col items-center justify-center w-full h-48 overflow-hidden rounded-t-lg bg-blue-50 dark:bg-blue-900/20"
    >
      <!-- 虚拟背景装饰 -->
      <div class="absolute inset-0 pointer-events-none opacity-50">
        <!-- 类似设计图中的几个元素 -->
        <div
          class="absolute top-10 left-20 w-12 h-12 bg-teal-300 rounded-full blur-md"
        ></div>
        <div
          class="absolute bottom-10 left-40 w-6 h-6 bg-orange-500 rounded-full blur-sm"
        ></div>
        <div
          class="absolute top-20 right-20 w-16 h-4 bg-indigo-600 rounded rotate-45 blur-sm"
        ></div>
      </div>

      <!-- 头像和上传 -->
      <div class="relative z-10 z-20 mt-4 group">
        <div
          @click="triggerUpload"
          class="flex items-center justify-center w-24 h-24 overflow-hidden border-4 border-white rounded-full shadow-lg bg-gray-100 hover:bg-gray-200 cursor-pointer"
        >
          <img
            v-if="userInfoForm.avatar"
            :src="userInfoForm.avatar"
            class="object-cover w-full h-full"
          />
          <svg
            v-else
            xmlns="http://www.w3.org/2000/svg"
            class="w-10 h-10 text-gray-400"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>

          <div
            class="absolute inset-0 flex flex-col items-center justify-center text-white bg-black bg-opacity-50 opacity-0 group-hover:opacity-100 transition-opacity"
          >
            <span class="text-2xl">+</span>
            <span class="text-xs">{{ $t('page.profile.localUpload') }}</span>
          </div>
        </div>
      </div>

      <!-- 角色标签 -->
      <div class="z-10 mt-3 mb-2">
        <Tag
          v-for="role in userStore.userInfo?.roles"
          :key="role"
          theme="primary"
          shape="round"
          size="large"
        >
          {{ role === 'superAdmin' ? $t('page.profile.superAdmin') : role }}
        </Tag>
      </div>
    </div>

    <!-- 下方两列内容 -->
    <div class="flex flex-col gap-4 mt-4 md:flex-row">
      <!-- 左栏：个人资料、安全设置 -->
      <div
        class="flex-1 p-4 bg-white rounded shadow-sm dark:bg-[var(--vben-color-background-elevated)] min-h-[500px]"
      >
        <Tabs v-model="leftTab" class="h-full">
          <TabPanel value="info" :label="$t('page.profile.personalInfo')">
            <div class="pt-6 mt-4">
              <Form
                :data="userInfoForm"
                label-align="left"
                label-width="100px"
                @submit="handleUpdateInfo"
              >
                <FormItem :label="$t('page.profile.username')" name="username">
                  <Input v-model="userInfoForm.username" disabled />
                </FormItem>
                <FormItem :label="$t('page.profile.nickname')" name="nickname">
                  <Input
                    v-model="userInfoForm.nickname"
                    :placeholder="$t('page.profile.placeholder.nickname')"
                  />
                </FormItem>
                <FormItem :label="$t('page.profile.phone')" name="phone">
                  <Input
                    v-model="userInfoForm.phone"
                    :placeholder="$t('page.profile.placeholder.phone')"
                  />
                </FormItem>
                <FormItem :label="$t('page.profile.email')" name="email">
                  <Input
                    v-model="userInfoForm.email"
                    :placeholder="$t('page.profile.placeholder.email')"
                  />
                </FormItem>
                <FormItem :label="$t('page.profile.personalSign')" name="signed">
                  <Textarea
                    v-model="userInfoForm.signed"
                    :placeholder="$t('page.profile.placeholder.personalSign')"
                    :maxlength="255"
                    :autosize="{ minRows: 3, maxRows: 5 }"
                  />
                </FormItem>
                <FormItem>
                  <Button
                    theme="default"
                    type="submit"
                    class="bg-gray-800 text-white hover:bg-gray-700"
                  >
                    {{ $t('page.profile.save') }}
                  </Button>
                </FormItem>
              </Form>
            </div>
          </TabPanel>

          <TabPanel value="security" :label="$t('page.profile.securitySetting')">
            <div class="pt-6 mt-4">
              <Form
                :data="securityForm"
                label-align="left"
                label-width="100px"
                @submit="handleUpdatePassword"
              >
                <FormItem :label="$t('page.profile.oldPassword')" name="oldPassword" required-mark>
                  <Input
                    type="password"
                    v-model="securityForm.oldPassword"
                    :placeholder="$t('page.profile.placeholder.oldPassword')"
                  />
                </FormItem>
                <FormItem :label="$t('page.profile.newPassword')" name="newPassword" required-mark>
                  <Input
                    type="password"
                    v-model="securityForm.newPassword"
                    :placeholder="$t('page.profile.placeholder.newPassword')"
                  />
                </FormItem>
                <FormItem
                  :label="$t('page.profile.confirmPassword')"
                  name="newPasswordConfirmation"
                  required-mark
                >
                  <Input
                    type="password"
                    v-model="securityForm.newPasswordConfirmation"
                    :placeholder="$t('page.profile.placeholder.confirmPassword')"
                  />
                </FormItem>
                <FormItem>
                  <Button
                    theme="default"
                    type="submit"
                    class="bg-gray-800 text-white hover:bg-gray-700"
                  >
                    {{ $t('page.profile.save') }}
                  </Button>
                </FormItem>
              </Form>
            </div>
          </TabPanel>
        </Tabs>
      </div>

      <!-- 右栏：日志 -->
      <div
        class="flex-1 p-4 bg-white rounded shadow-sm dark:bg-[var(--vben-color-background-elevated)] min-h-[500px]"
      >
        <Tabs v-model="rightTab" class="h-full">
          <TabPanel value="loginLog" :label="$t('page.profile.loginLog')">
            <div class="pt-6 mt-4 overflow-y-auto max-h-[400px]">
              <Timeline>
                <TimelineItem
                  v-for="log in loginLogs"
                  :key="log.id"
                  theme="primary"
                >
                  <div
                    class="text-sm font-medium text-gray-800 dark:text-gray-200"
                  >
                    {{ $t('page.profile.loginLogTemplate', [log.login_time || log.created_at, log.status === 1 ? $t('page.profile.loginSuccess') : $t('page.profile.loginFailed')]) }}
                  </div>
                  <div class="mt-1 text-xs text-gray-500">
                    {{ $t('page.profile.ipLocation') }}: {{ log.ip_location || $t('page.profile.unknown') }}，{{ $t('page.profile.os') }}: {{ log.os || $t('page.profile.unknown') }}
                  </div>
                </TimelineItem>
                <div
                  v-if="loginLogs.length === 0"
                  class="text-center text-gray-400 py-10"
                >
                  {{ $t('page.profile.noLogs') }}
                </div>
              </Timeline>
            </div>
          </TabPanel>

          <TabPanel value="opLog" :label="$t('page.profile.operLog')">
            <div class="pt-6 mt-4 overflow-y-auto max-h-[400px]">
              <Timeline>
                <TimelineItem
                  v-for="log in operationLogs"
                  :key="log.id"
                  theme="primary"
                >
                  <div
                    class="text-sm font-medium text-gray-800 dark:text-gray-200"
                  >
                    {{ $t('page.profile.operLogTemplate', [log.created_at, log.service_name || $t('page.profile.operated')]) }}
                  </div>
                  <div class="mt-1 text-xs text-gray-500">
                    {{ $t('page.profile.ipLocation') }}: {{ log.ip_location || $t('page.profile.unknown') }}，{{ $t('page.profile.method') }}: {{ log.method }}，{{ $t('page.profile.path') }}: {{ log.router }}
                  </div>
                </TimelineItem>
                <div
                  v-if="operationLogs.length === 0"
                  class="text-center text-gray-400 py-10"
                >
                  {{ $t('page.profile.noLogs') }}
                </div>
              </Timeline>
            </div>
          </TabPanel>
        </Tabs>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 避免覆盖全局表单样式 */
</style>
