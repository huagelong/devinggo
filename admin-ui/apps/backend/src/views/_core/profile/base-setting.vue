<script setup lang="ts">
import type { BasicOption } from '@vben/types';

import type { VbenFormSchema } from '#/adapter/form';

import { computed, onMounted, ref } from 'vue';

import { $t } from '@vben/locales';
import { ProfileBaseSetting } from '@vben/common-ui';

import { getUserInfoApi } from '#/api';

const profileBaseSettingRef = ref();

const MOCK_ROLES_OPTIONS: BasicOption[] = [
  {
    label: $t('page.profile.roles.admin'),
    value: 'super',
  },
  {
    label: $t('page.profile.roles.user'),
    value: 'user',
  },
  {
    label: $t('page.profile.roles.test'),
    value: 'test',
  },
];

const formSchema = computed((): VbenFormSchema[] => {
  return [
    {
      fieldName: 'realName',
      component: 'Input',
      label: $t('page.profile.realName'),
    },
    {
      fieldName: 'username',
      component: 'Input',
      label: $t('page.profile.username'),
    },
    {
      fieldName: 'roles',
      component: 'Select',
      componentProps: {
        mode: 'tags',
        options: MOCK_ROLES_OPTIONS,
      },
      label: $t('page.profile.role'),
    },
    {
      fieldName: 'introduction',
      component: 'Textarea',
      label: $t('page.profile.bio'),
    },
  ];
});

onMounted(async () => {
  const data = await getUserInfoApi();
  profileBaseSettingRef.value.getFormApi().setValues(data);
});
</script>
<template>
  <ProfileBaseSetting ref="profileBaseSettingRef" :form-schema="formSchema" />
</template>
