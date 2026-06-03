<script lang="ts" setup>
import { $t } from '@vben/locales';

import { Button, Form, FormItem, Input, Select } from 'tdesign-vue-next';

interface Props {
  formData: {
    username: string;
    nickname: string;
    status: number | undefined;
  };
  statusOptions: Array<{ label: string; value: number }>;
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
  <div class="rounded-md border border-gray-100 bg-white p-4">
    <Form :data="formData" label-width="90px" layout="inline" colon size="small">
      <div class="grid grid-cols-3 gap-x-4 gap-y-3">
        <FormItem :label="$t('system.dept.username')" name="username">
          <Input
            :value="formData.username"
            :placeholder="$t('ui.placeholder.input', [$t('system.dept.username')])"
            clearable
            size="small"
            @change="(val: any) => updateField('username', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.dept.nickname')" name="nickname">
          <Input
            :value="formData.nickname"
            :placeholder="$t('ui.placeholder.input', [$t('system.dept.nickname')])"
            clearable
            size="small"
            @change="(val: any) => updateField('nickname', val)"
          />
        </FormItem>
        <FormItem :label="$t('common.status')" name="status">
          <Select
            :value="formData.status"
            :options="statusOptions"
            :placeholder="$t('ui.placeholder.select', [$t('common.status')])"
            clearable
            size="small"
            @change="(val: any) => updateField('status', val)"
          />
        </FormItem>
      </div>

      <div class="flex justify-end gap-2 pt-3">
        <Button size="small" theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
        <Button size="small" theme="primary" @click="handleSearch">
          {{ $t('common.search') }}
        </Button>
      </div>
    </Form>
  </div>
</template>