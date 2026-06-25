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

import type { OptionItem, IdType } from '#/types/common';
import type { DictOption } from '#/composables/crud/use-dict-options';

interface Props {
  formData: {
    group_id?: IdType | undefined;
    name: string;
    access_name: string;
    request_mode?: string | number | undefined;
    status?: number | undefined;
    created_at: string[] | undefined;
  };
  groupOptions: OptionItem<IdType>[];
  requestModeOptions: DictOption[];
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
  <div class="rounded-md bg-card p-4">
    <Form :data="formData" label-width="90px" layout="inline" colon>
      <div class="w-full">
        <div class="grid grid-cols-4 gap-x-4 gap-y-3 w-full">
          <FormItem :label="$t('system.api.group')" name="group_id">
            <Select
              :value="formData.group_id"
              :options="groupOptions"
              :placeholder="$t('ui.placeholder.select')"
              clearable
              class="w-full"
              @change="(val: any) => updateField('group_id', val)"
            />
          </FormItem>
          <FormItem :label="$t('system.api.name')" name="name">
            <Input
              :value="formData.name"
              :placeholder="$t('ui.placeholder.input')"
              clearable
              @change="(val: any) => updateField('name', val)"
            />
          </FormItem>
          <FormItem :label="$t('system.api.code')" name="access_name">
            <Input
              :value="formData.access_name"
              :placeholder="$t('ui.placeholder.input')"
              clearable
              @change="(val: any) => updateField('access_name', val)"
            />
          </FormItem>
          <FormItem :label="$t('system.api.requestMode')" name="request_mode">
            <Select
              :value="formData.request_mode"
              :options="requestModeOptions"
              :placeholder="$t('ui.placeholder.select')"
              clearable
              class="w-full"
              @change="(val: any) => updateField('request_mode', val)"
            />
          </FormItem>
          <FormItem :label="$t('common.status')" name="status">
            <Select
              :value="formData.status"
              :options="statusOptions"
              :placeholder="$t('ui.placeholder.select')"
              clearable
              class="w-full"
              @change="(val: any) => updateField('status', val)"
            />
          </FormItem>
          <FormItem :label="$t('common.createTime')" name="created_at" class="col-span-2">
            <DateRangePicker
              :value="formData.created_at"
              :placeholder="[$t('common.startTime'), $t('common.endTime')]"
              clearable
              class="w-full"
              @change="(val: any) => updateField('created_at', val)"
            />
          </FormItem>
        </div>
        <div class="flex justify-start gap-2 pt-4">
          <Button theme="default" @click="handleReset">{{ $t('common.reset') }}</Button>
          <Button theme="primary" @click="handleSearch">
            <template #icon><SearchIcon /></template>
            {{ $t('common.query') }}
          </Button>
        </div>
      </div>
    </Form>
  </div>
</template>