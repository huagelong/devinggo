<script lang="ts" setup>
import type { CrontabApi } from '#/api/system/crontab';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getCrontabTarget, saveCrontab, updateCrontab } from '#/api/system/crontab';
import { logger } from '#/utils/logger';

import {
  createCrontabFormDefaultValues,
  crontabSingletonOptions,
  crontabStatusOptions,
  crontabTypeOptions,
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
  wrapperClass: 'grid-cols-1',
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
        placeholder: $t('ui.placeholder.input', [$t('system.crontab.name')]),
      },
      fieldName: 'name',
      label: $t('system.crontab.name'),
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        options: crontabTypeOptions,
        placeholder: $t('ui.placeholder.select', [$t('system.crontab.taskType')]),
        onChange: (value: number) => {
          formApi.updateSchema([
            {
              fieldName: 'target',
              componentProps: {
                api: getCrontabTarget,
                params: { type: value },
                resultField: 'data',
                placeholder: $t('ui.placeholder.select', [$t('system.crontab.target')]),
              },
            },
          ]);
          formApi.setFieldValue('target', undefined);
        },
      },
      defaultValue: 1,
      fieldName: 'type',
      label: $t('system.crontab.taskType'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input', [$t('system.crontab.rule')]),
      },
      description: $t('system.crontab.expression'),
      fieldName: 'rule',
      label: $t('system.crontab.rule'),
      rules: 'required',
    },
    {
      component: 'ApiSelect',
      componentProps: {
        api: getCrontabTarget,
        params: { type: 1 },
        resultField: 'data',
        placeholder: $t('ui.placeholder.select', [$t('system.crontab.target')]),
      },
      fieldName: 'target',
      label: $t('system.crontab.target'),
      rules: 'required',
    },
    {
      component: 'CodeEditor',
      componentProps: {
        placeholder: $t('ui.placeholder.input', [$t('system.crontab.parameter')]),
      },
      description:
        '必须json格式,例子: {"name":"name1", "value":"shuju1"} URL任务参数例子: {"url":"http://www.example.com","method":"get", "header":{},"params":{}}',
      fieldName: 'parameter',
      label: $t('system.crontab.parameter'),
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: crontabSingletonOptions,
      },
      defaultValue: 2,
      description: $t('system.crontab.singletonTip'),
      fieldName: 'singleton',
      label: $t('system.crontab.singleton'),
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: crontabStatusOptions,
      },
      defaultValue: 1,
      fieldName: 'status',
      label: $t('common.status'),
    },
    {
      component: 'Textarea',
      componentProps: {
        placeholder: $t('ui.placeholder.input', [$t('common.remark')]),
        autosize: { minRows: 2, maxRows: 4 },
      },
      fieldName: 'remark',
      label: $t('common.remark'),
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onCancel: () => {
    modalApi.close();
  },
  onOpened: () => {
    // 弹窗完全展开后通知表单重新布局
    nextTick(() => {
      formApi.validate();
    });
  },
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

      await (payload.id ? updateCrontab(Number(payload.id), payload) : saveCrontab(payload));

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
  cancelText: $t('common.cancel'),
  confirmText: $t('common.save'),
  class: 'w-[800px] max-w-[94vw]',
});

async function open(data?: Partial<CrontabApi.SubmitPayload>) {
  const defaultValues = createCrontabFormDefaultValues();
  const mergedData = { ...data };
  if (mergedData.parameter && typeof mergedData.parameter === 'object') {
    mergedData.parameter = JSON.stringify(mergedData.parameter, null, 2);
  }
  baseValues.value = {
    ...defaultValues,
    ...mergedData,
  };

  modalApi.setState({
    title: data?.id ? $t('system.crontab.editTitle') : $t('system.crontab.createTitle'),
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(baseValues.value);
  await nextTick();

  const currentType = baseValues.value.type || 1;
  formApi.updateSchema([
    {
      fieldName: 'target',
      componentProps: {
        api: getCrontabTarget,
        params: { type: currentType },
        resultField: 'data',
        placeholder: $t('ui.placeholder.select', [$t('system.crontab.target')]),
      },
    },
  ]);

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
