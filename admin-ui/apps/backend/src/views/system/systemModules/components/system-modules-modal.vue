<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { SystemModulesApi } from '#/api/system/system-modules';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import {
  saveSystemModules,
  updateSystemModules,
} from '#/api/system/system-modules';

import { createSystemModulesFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const baseValues = ref<SystemModulesApi.SubmitPayload>(
  createSystemModulesFormDefaultValues(),
);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 80,
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
      label: $t('system.systemModules.name'),
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
    let isEdit = false;
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<Partial<SystemModulesApi.SubmitPayload>>();
      const payload: SystemModulesApi.SubmitPayload = {
        ...baseValues.value,
        ...values,
      };
      isEdit = !!payload.id;

      modalApi.setState({ confirmLoading: true });

      if (payload.id) {
        await updateSystemModules(Number(payload.id), payload);
      } else {
        await saveSystemModules(payload);
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
  class: 'w-[840px] max-w-[94vw]',
});

async function open(data?: Partial<SystemModulesApi.SubmitPayload>) {
  const defaultValues = createSystemModulesFormDefaultValues();
  baseValues.value = {
    ...defaultValues,
    ...data,
  };

  modalApi.setState({
    title: data?.id ? $t('system.systemModules.editTitle') : $t('system.systemModules.createTitle'),
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
    <Form />
  </Modal>
</template>
