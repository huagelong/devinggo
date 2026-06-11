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
  TreeSelect,
} from 'tdesign-vue-next';

import type { DeptApi } from '#/api/system/dept';
import type { PostApi } from '#/api/system/post';
import type { RoleApi } from '#/api/system/role';
import type { IdType } from '#/types/common';
import type { DictOption } from '#/composables/crud/use-dict-options';

interface Props {
  formData: {
    username: string;
    dept_ids?: IdType[] | undefined;
    role_id?: number | undefined;
    phone: string;
    post_id?: number | undefined;
    email: string;
    status?: number | undefined;
    user_type?: string | undefined;
    created_at?: string[] | undefined;
  };
  roleOptions: RoleApi.ListItem[];
  postOptions: PostApi.ListItem[];
  deptTreeData: DeptApi.TreeNode[];
  statusOptions: DictOption[];
  userTypeOptions: DictOption[];
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
    <Form :data="formData" label-width="80px" layout="inline" colon>
      <div class="grid grid-cols-3 gap-x-4 gap-y-3">
        <FormItem :label="$t('system.user.username')" name="username">
          <Input
            :value="formData.username"
            :placeholder="$t('ui.placeholder.input', [$t('system.user.username')])"
            clearable
            @change="(val: any) => updateField('username', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.user.dept')" name="dept_ids">
          <TreeSelect
            :value="formData.dept_ids"
            :data="deptTreeData"
            :keys="{ value: 'id', label: 'label', children: 'children' }"
            :multiple="true"
            :tree-props="{ checkStrictly: true }"
            :placeholder="$t('ui.placeholder.select', [$t('system.user.dept')])"
            clearable
            class="w-full"
            @change="(val: any) => updateField('dept_ids', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.user.role')" name="role_id">
          <Select
            :value="formData.role_id"
            :options="roleOptions"
            :keys="{ value: 'id', label: 'name' }"
            :placeholder="$t('ui.placeholder.select', [$t('system.user.role')])"
            clearable
            class="w-full"
            @change="(val: any) => updateField('role_id', val)"
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
        <FormItem :label="$t('system.user.post')" name="post_id">
          <Select
            :value="formData.post_id"
            :options="postOptions"
            :keys="{ value: 'id', label: 'name' }"
            :placeholder="$t('ui.placeholder.select', [$t('system.user.post')])"
            clearable
            class="w-full"
            @change="(val: any) => updateField('post_id', val)"
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
        <FormItem :label="$t('common.status')" name="status">
          <Select
            :value="formData.status"
            :options="statusOptions"
            :placeholder="$t('ui.placeholder.select', [$t('common.status')])"
            clearable
            class="w-full"
            @change="(val: any) => updateField('status', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.user.userType')" name="user_type">
          <Select
            :value="formData.user_type"
            :options="userTypeOptions"
            :placeholder="$t('ui.placeholder.select', [$t('system.user.userType')])"
            clearable
            class="w-full"
            @change="(val: any) => updateField('user_type', val)"
          />
        </FormItem>
        <FormItem :label="$t('system.user.registerTime')" name="created_at" class="col-span-1">
          <DateRangePicker
            :value="formData.created_at"
            :placeholder="[$t('common.startTime'), $t('common.endTime')]"
            clearable
            class="w-full"
            @change="(val: any) => updateField('created_at', val)"
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