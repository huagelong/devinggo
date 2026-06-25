<script setup lang="ts">
import type { VbenFormSchema } from '#/adapter/form';

import { computed } from 'vue';

import { $t } from '@vben/locales';
import { ProfilePasswordSetting, z } from '@vben/common-ui';

import { message } from '#/adapter/tdesign';

const formSchema = computed((): VbenFormSchema[] => {
  return [
    {
      fieldName: 'oldPassword',
      label: $t('page.profile.oldPassword'),
      component: 'VbenInputPassword',
      componentProps: {
        placeholder: $t('page.profile.placeholder.oldPassword'),
      },
    },
    {
      fieldName: 'newPassword',
      label: $t('page.profile.newPassword'),
      component: 'VbenInputPassword',
      componentProps: {
        passwordStrength: true,
        placeholder: $t('page.profile.placeholder.newPassword'),
      },
    },
    {
      fieldName: 'confirmPassword',
      label: $t('page.profile.confirmPassword'),
      component: 'VbenInputPassword',
      componentProps: {
        passwordStrength: true,
        placeholder: $t('ui.placeholder.input'),
      },
      dependencies: {
        rules(values) {
          const { newPassword } = values;
          return z
            .string({ required_error: $t('ui.placeholder.input') })
            .min(1, { message: $t('ui.placeholder.input') })
            .refine((value) => value === newPassword, {
              message: $t('common.passwordMismatch'),
            });
        },
        triggerFields: ['newPassword'],
      },
    },
  ];
});

function handleSubmit() {
  message.success($t('common.passwordChangeSuccess'));
}
</script>
<template>
  <ProfilePasswordSetting
    class="w-1/3"
    :form-schema="formSchema"
    @submit="handleSubmit"
  />
</template>
