<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { CrontabApi } from '#/api/system/crontab';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveCrontab, updateCrontab } from '#/api/system/crontab';

import {
  crontabFinallyOptions,
  crontabTypeOptions,
  createCrontabFormDefaultValues,
} from '../schemas';

const emit = defineEmits(['success']);

const baseValues = ref<CrontabApi.SubmitPayload>(
  createCrontabFormDefaultValues(),
);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 100,
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
      label: $t('system.crontab.name'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: crontabTypeOptions,
      },
      defaultValue: 1,
      fieldName: 'type',
      label: $t('system.crontab.taskType'),
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'rule',
      label: $t('system.crontab.rule'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'target',
      label: $t('system.crontab.target'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: crontabFinallyOptions,
      },
      defaultValue: 2,
      fieldName: 'is_finally',
      label: $t('system.crontab.isFinally'),
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

      const values = await formApi.getValues<Partial<CrontabApi.SubmitPayload>>();
      const payload: CrontabApi.SubmitPayload = {
        ...baseValues.value,
        ...values,
      };
      isEdit = !!payload.id;

      modalApi.setState({ confirmLoading: true });

      if (payload.id) {
        await updateCrontab(Number(payload.id), payload);
      } else {
        await saveCrontab(payload);
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
  class: 'w-[940px] max-w-[94vw]',
});

async function open(data?: Partial<CrontabApi.SubmitPayload>) {
  const defaultValues = createCrontabFormDefaultValues();
  baseValues.value = {
    ...defaultValues,
    ...data,
  };

  modalApi.setState({
    title: data?.id ? $t('system.crontab.editTitle') : $t('system.crontab.createTitle'),
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
