<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DemoApi } from '#/api/system/demo';

import { nextTick } from 'vue';

import { $t } from '@vben/locales';
import { useVbenModal } from '@vben/common-ui';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveDemo, updateDemo } from '#/api/system/demo';

import { createDemoFormDefaultValues } from '../schemas';

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
      fieldName: 'name',
      label: $t('common.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input')
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: $t('system.demo.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input')
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      label: $t('common.status'),
      defaultValue: 1,
      componentProps: {
        options: [{label: $t('common.statusEnabled'), value: 1}, {label: $t('common.disabled'), value: 2}],
      },
      rules: 'required',
    },
    {
      component: 'InputNumber',
      fieldName: 'sort',
      label: $t('common.sort'),
      defaultValue: 1,
      componentProps: {
        min: 0,
        max: 1000,
      },
      rules: 'required',
    },
    {
      component: 'InputNumber',
      fieldName: 'price',
      label: $t('system.demo.price'),
      componentProps: {
        min: 0,
        max: 1000,
        placeholder: $t('ui.placeholder.input')
      },
    },
    {
      component: 'Upload',
      fieldName: 'cover',
      label: $t('system.demo.cover'),
      componentProps: {
        accept: 'image/*',
        placeholder: $t('system.demo.uploadImage'),
      },
      formItemClass: 'md:col-span-2',
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: $t('system.demo.email'),
      componentProps: {
        type: 'email',
        placeholder: $t('system.demo.enterEmail'),
      },
      rules: 'email',
    },
    {
      component: 'Input',
      fieldName: 'phone',
      label: $t('system.demo.phone'),
      componentProps: {
        placeholder: $t('system.demo.enterPhone'),
      },
    },
    {
      component: 'DatePicker',
      fieldName: 'birthday',
      label: $t('system.demo.birthday'),
      componentProps: {
        placeholder: $t('system.demo.selectDate'),
        clearable: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: $t('common.remark'),
      componentProps: {
        placeholder: $t('system.demo.enterRemark'),
        autosize: { minRows: 3, maxRows: 6 },
      },
      formItemClass: 'md:col-span-2',
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<DemoApi.SubmitPayload & { id?: number }>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await updateDemo(Number(values.id), values);
      } else {
        await saveDemo(values);
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
  class: 'w-[980px] max-w-[94vw]',
});

async function open(data?: Partial<DemoApi.SubmitPayload & { id?: number }>) {
  modalApi.setState({
    title: data?.id ? $t('system.demo.editTitle') : $t('system.demo.createTitle'),
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(createDemoFormDefaultValues());
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
