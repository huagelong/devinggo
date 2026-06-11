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
  <div class="rounded-md border border-border bg-card p-4">
    <Form :data="formData" label-width="90px" colon>
      <div class="grid grid-cols-4 gap-x-4 gap-y-3 items-end">
        <FormItem :label="$t('system.dept.username')" name="username">
          <Input
            :value="formData.username"
            :placeholder="$t('ui.placeholder.input', [$t('system.dept.username')])"
            clearable
            @change="(val: any) => updateField('username', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.dept.nickname')" name="nickname">
          <Input
            :value="formData.nickname"
            :placeholder="$t('ui.placeholder.input', [$t('system.dept.nickname')])"
            clearable
            @change="(val: any) => updateField('nickname', val)"
          />
        </FormItem>
        <FormItem :label="$t('common.status')" name="status">
          <Select
            :value="formData.status"
            :options="statusOptions"
            :placeholder="$t('ui.placeholder.select', [$t('common.status')])"
            clearable
            @change="(val: any) => updateField('status', val)"
          />
        </FormItem>
        <div class="flex justify-end gap-2 pb-[7px]">
          <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
          <Button theme="primary" @click="handleSearch">
            {{ $t('common.search') }}
          </Button>
        </div>
      </div>
    </Form>
  </div>
</template>
