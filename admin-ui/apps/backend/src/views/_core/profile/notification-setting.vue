<script setup lang="ts">
import { computed } from 'vue';

import { $t } from '@vben/locales';
import { ProfileNotificationSetting } from '@vben/common-ui';

const formSchema = computed(() => {
  return [
    {
      value: true,
      fieldName: 'accountPassword',
      label: $t('page.profile.accountPasswordNotify'),
      description: $t('page.profile.accountPasswordNotifyDesc'),
    },
    {
      value: true,
      fieldName: 'systemMessage',
      label: $t('page.profile.systemMessage'),
      description: $t('page.profile.systemMessageDesc'),
    },
    {
      value: true,
      fieldName: 'todoTask',
      label: $t('page.profile.todoTask'),
      description: $t('page.profile.todoTaskDesc'),
    },
  ];
});
</script>
<template>
  <ProfileNotificationSetting :form-schema="formSchema" />
</template>
