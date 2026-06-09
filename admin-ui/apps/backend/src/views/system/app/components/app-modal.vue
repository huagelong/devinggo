<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { AppApi } from '#/api/system/app';

import { markRaw, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, Input, MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getAppGroupPageList } from '#/api/system/app-group';
import { getAppId, getAppSecret, saveApp, updateApp } from '#/api/system/app';

import ConfigRichTextEditor from '#/views/system/config/components/config-rich-text-editor.vue';
import { createAppFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const baseValues = ref<AppApi.SubmitPayload>(createAppFormDefaultValues());
const groupOptions = ref<{ label: string; value: number }[]>([]);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  layout: 'vertical',
  commonConfig: {
    labelWidth: 90,
  },
  wrapperClass: 'grid-cols-1 gap-y-6 md:grid-cols-2 md:gap-x-6 md:gap-y-8',
  schema: [
    {
      component: 'Input',
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: 'Select',
      componentProps: {
        options: groupOptions.value,
        placeholder: '请选择所属组',
      },
      fieldName: 'group_id',
      label: '所属组',
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: '请输入应用名称',
      },
      fieldName: 'app_name',
      label: '应用名称',
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: 'e26264a666',
      },
      fieldName: 'app_id',
      label: 'APP ID',
      rules: 'required',
      formItemClass: 'md:col-span-2',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: '请输入APP SECRET',
      },
      fieldName: 'app_secret',
      label: 'APP SECRET',
      rules: 'required',
      formItemClass: 'md:col-span-2',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: '正常', value: 1 },
          { label: '停用', value: 2 },
        ],
      },
      defaultValue: 1,
      fieldName: 'status',
      label: '状态',
      formItemClass: 'md:col-span-2',
    },
    {
      component: markRaw(ConfigRichTextEditor),
      fieldName: 'description',
      label: '应用介绍',
      formItemClass: 'md:col-span-2',
    },
    {
      component: 'Textarea',
      componentProps: {
        placeholder: '请输入备注',
      },
      fieldName: 'remark',
      label: '备注',
      formItemClass: 'md:col-span-2',
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    let isEdit = false;
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<Partial<AppApi.SubmitPayload>>();
      const payload: AppApi.SubmitPayload = {
        ...baseValues.value,
        ...values,
      };
      isEdit = !!payload.id;

      modalApi.setState({ confirmLoading: true });

      if (payload.id) {
        await updateApp(Number(payload.id), payload);
      } else {
        await saveApp(payload);
      }

      MessagePlugin.success(isEdit ? $t('common.updateSuccess') : $t('common.createSuccess'));
      emit('success');
      modalApi.close();
    } catch (error) {
      logger.error(error);
      MessagePlugin.error(isEdit ? $t('common.updateFailed') : $t('common.createFailed'));
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  cancelText: '关闭',
  confirmText: '保存',
  class: 'w-[900px] max-w-[94vw] !p-8',
});

async function fetchGroupOptions() {
  try {
    const res = await getAppGroupPageList({ pageSize: 999 });
    const items = (res as any)?.items || [];
    groupOptions.value = items.map((item: any) => ({
      label: item.name,
      value: item.id,
    }));
    formApi.updateSchema([
      {
        fieldName: 'group_id',
        componentProps: {
          options: groupOptions.value,
          placeholder: '请选择所属组',
        },
      },
    ]);
  } catch (error) {
    logger.error(error);
  }
}

async function handleRefreshAppId() {
  try {
    const res = await getAppId();
    const newId = (res as any)?.app_id || '';
    formApi.setFieldValue('app_id', newId);
  } catch (error) {
    logger.error(error);
  }
}

async function handleRefreshAppSecret() {
  try {
    const res = await getAppSecret();
    const newSecret = (res as any)?.app_secret || '';
    formApi.setFieldValue('app_secret', newSecret);
  } catch (error) {
    logger.error(error);
  }
}

async function open(data?: Partial<AppApi.SubmitPayload>) {
  await fetchGroupOptions();
  const defaultValues = createAppFormDefaultValues();
  baseValues.value = {
    ...defaultValues,
    ...data,
  };

  if (!data?.id) {
    try {
      const [idRes, secretRes] = await Promise.all([
        getAppId(),
        getAppSecret(),
      ]);
      baseValues.value.app_id = (idRes as any)?.app_id || '';
      baseValues.value.app_secret = (secretRes as any)?.app_secret || '';
    } catch (error) {
      logger.error(error);
    }
  }

  modalApi.setState({
    title: data?.id ? '编辑' : '新增',
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(baseValues.value);
  await nextTick();
  await formApi.resetValidate();
}

defineExpose({
  open,
});
</script>

<template>
  <Modal>
    <Form>
      <template #app_id="slotProps">
        <div class="flex w-full items-center gap-2">
          <Input v-bind="slotProps" style="width: 100%" />
          <Button theme="primary" @click="handleRefreshAppId">刷新APP ID</Button>
        </div>
      </template>
      <template #app_secret="slotProps">
        <div class="flex w-full items-center gap-2">
          <Input v-bind="slotProps" style="width: 100%" />
          <Button theme="primary" @click="handleRefreshAppSecret">刷新APP SECRET</Button>
        </div>
      </template>
    </Form>
  </Modal>
</template>
