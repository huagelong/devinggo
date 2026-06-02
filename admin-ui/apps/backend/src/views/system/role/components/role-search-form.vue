<script lang="ts" setup>
import { $t } from '@vben/locales';

import { SearchIcon } from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  Select,
} from 'tdesign-vue-next';

import type { DictOption } from '#/composables/crud/use-dict-options';

interface Props {
  formData: {
    name: string;
    code: string;
    status: number | undefined;
    created_at: string[] | undefined;
  };
  statusOptions: DictOption[];
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

function updateField(field: keyof Props['formData'], value: unknown) {
  emit('update:formData', { ...props.formData, [field]: value });
}
</script>

<template>
  <div class="rounded-md bg-white p-4">
    <Form :data="formData" label-width="80px" layout="inline" colon>
      <div class="grid grid-cols-4 gap-x-4 gap-y-3">
        <FormItem :label="$t('system.role.name')" name="name">
          <Input
            :value="formData.name"
            :placeholder="$t('ui.placeholder.input', [$t('system.role.name')])"
            clearable
            @change="(val: string) => updateField('name', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.role.code')" name="code">
          <Input
            :value="formData.code"
            :placeholder="$t('ui.placeholder.input', [$t('system.role.code')])"
            clearable
            @change="(val: string) => updateField('code', val)"
          />
        </FormItem>
        <FormItem :label="$t('common.status')" name="status">
          <Select
            :value="formData.status"
            :options="statusOptions"
            :placeholder="$t('ui.placeholder.select', [$t('common.status')])"
            clearable
            class="w-full"
            @change="(val: number | undefined) => updateField('status', val)"
          />
        </FormItem>
        <FormItem :label="$t('common.createTime')" name="created_at">
          <DateRangePicker
            :value="formData.created_at"
            :placeholder="[$t('common.startTime'), $t('common.endTime')]"
            clearable
            class="w-full"
            @change="(val: string[] | undefined) => updateField('created_at', val)"
          />
        </FormItem>
      </div>
      <div class="flex justify-end gap-2 pt-2">
        <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
        <Button theme="primary" @click="handleSearch">
          <template #icon><SearchIcon /></template>
          {{ $t('common.query') }}
        </Button>
      </div>
    </Form>
  </div>
</template>