<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { PostApi } from '#/api/system/post';

import { nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { savePost, updatePost } from '#/api/system/post';

import { createPostFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 90,
  },
  wrapperClass: 'grid-cols-1 md:grid-cols-2',
  schema: [
    {
      component: 'Input',
      dependencies: {
        show: false,
        triggerFields: [''],
      },
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'name',
      label: $t('system.post.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'code',
      label: $t('system.post.code'),
      rules: 'required',
    },
    {
      component: 'InputNumber',
      componentProps: {
        max: 1000,
        min: 0,
      },
      defaultValue: 1,
      fieldName: 'sort',
      label: $t('common.sort'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: [
          { label: $t('common.statusEnabled'), value: 1 },
          { label: $t('common.statusDisabled'), value: 2 },
        ],
      },
      defaultValue: 1,
      fieldName: 'status',
      label: $t('common.status'),
      rules: 'required',
    },
    {
      component: 'Textarea',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'remark',
      formItemClass: 'md:col-span-2',
      label: $t('common.remark'),
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<PostApi.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await updatePost(Number(values.id), values);
      } else {
        await savePost(values);
      }

      MessagePlugin.success(values.id ? $t('common.updateSuccess') : $t('common.createSuccess'));
      emit('success');
      modalApi.close();
    } catch (error) {
      logger.error(error);
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  class: 'w-[840px] max-w-[94vw]',
});

async function open(data?: Partial<PostApi.SubmitPayload>) {
  modalApi.setState({
    title: data?.id ? $t('system.post.editTitle') : $t('system.post.createTitle'),
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(createPostFormDefaultValues());
  if (data) {
    formApi.setValues(data);
  }
  await nextTick();
  await formApi.resetValidate();
}

defineExpose({
  open,
});
</script>

<template>
  <Modal>
    <Form />
  </Modal>
</template>
