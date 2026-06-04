<script lang="ts" setup>
import { $t } from '@vben/locales';

import { Button, Form, FormItem, Input, Select } from 'tdesign-vue-next';

interface Props {
  formData: {
    username: string;
    nickname: string;
    phone: string;
    email: string;
    dept_id: number | undefined;
    role_id: number | undefined;
    post_id: number | undefined;
  };
  deptOptions: Array<{ label: string; value: number }>;
  roleOptions: Array<{ label: string; value: number }>;
  postOptions: Array<{ label: string; value: number }>;
}

interface Emits {
  (e: 'search'): void;
  (e: 'reset'): void;
  (e: 'update:formData', value: Props['formData']): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

function handleSearch() {
  emit('search');
}

function handleReset() {
  emit('reset');
}

function updateField(field: keyof Props['formData'], value: string | number | undefined) {
  emit('update:formData', { ...props.formData, [field]: value });
}
</script>

<template>
  <Form :data="formData" label-width="90px" colon>
    <div class="grid grid-cols-4 gap-x-4 gap-y-3 items-end">
      <FormItem :label="$t('system.user.username')" name="username">
        <Input
          :value="formData.username"
          :placeholder="$t('ui.placeholder.input', [$t('system.user.username')])"
          clearable
          @change="(val: any) => updateField('username', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.nickname')" name="nickname">
        <Input
          :value="formData.nickname"
          :placeholder="$t('ui.placeholder.input', [$t('system.user.nickname')])"
          clearable

          @change="(val: any) => updateField('nickname', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.phone')" name="phone">
        <Input
          :value="formData.phone"
          :placeholder="$t('ui.placeholder.input', [$t('system.user.phone')])"
          clearable

          @change="(val: any) => updateField('phone', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.email')" name="email">
        <Input
          :value="formData.email"
          :placeholder="$t('ui.placeholder.input', [$t('system.user.email')])"
          clearable

          @change="(val: any) => updateField('email', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.dept')" name="dept_id">
        <Select
          :value="formData.dept_id"
          :options="deptOptions"
          clearable

          :placeholder="$t('ui.placeholder.select', [$t('system.user.dept')])"
          @change="(val: any) => updateField('dept_id', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.role')" name="role_id">
        <Select
          :value="formData.role_id"
          :options="roleOptions"
          clearable

          :placeholder="$t('ui.placeholder.select', [$t('system.user.role')])"
          @change="(val: any) => updateField('role_id', val)"
        />
      </FormItem>
      <FormItem :label="$t('system.user.post')" name="post_id">
        <Select
          :value="formData.post_id"
          :options="postOptions"
          clearable

          :placeholder="$t('ui.placeholder.select', [$t('system.user.post')])"
          @change="(val: any) => updateField('post_id', val)"
        />
      </FormItem>
      <div class="flex justify-end gap-2 pb-[7px]">
        <Button theme="default" @click="handleReset">
          {{ $t('common.reset') }}
        </Button>
        <Button theme="primary" @click="handleSearch">
          {{ $t('common.search') }}
        </Button>
      </div>
    </div>
  </Form>
</template>