<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { <%.EntityName%>Api } from '#/api/<%.ModuleName%>/<%.VarName%>';

import { nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { save<%.EntityName%>, update<%.EntityName%> } from '#/api/<%.ModuleName%>/<%.VarName%>';

import { create<%.EntityName%>FormDefaultValues } from '../schemas';

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
<%range .EditableFields%>
    {
      component: '<%.Component%>',
      fieldName: '<%.Name%>',
      label: $t('<%$.ModuleName%>.<%$.VarName%>.<%.Name%>'),
<%if eq .Component "RadioGroup"%>
      componentProps: {
        options: [
          { label: '已上架', value: 1 },
          { label: '未上架', value: 2 },
        ],
      },
<%else if eq .Component "InputNumber"%>
      componentProps: {
        min: 0,
        max: 1000,
      },
<%else%>
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
<%end%>
<%if .DefaultValue%>
      defaultValue: <%.DefaultValue%>,
<%end%>
<%if .Rules%>
      rules: '<%.Rules%>',
<%end%>
<%if .FormItemClass%>
      formItemClass: '<%.FormItemClass%>',
<%end%>
    },
<%end%>
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<<%.EntityName%>Api.SubmitPayload & { id?: number }>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await update<%.EntityName%>(Number(values.id), values);
      } else {
        await save<%.EntityName%>(values);
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

async function open(data?: Partial<<%.EntityName%>Api.SubmitPayload & { id?: number }>) {
  modalApi.setState({
    title: data?.id ? $t('<%.ModuleName%>.<%.VarName%>.editTitle') : $t('<%.ModuleName%>.<%.VarName%>.createTitle'),
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(create<%.EntityName%>FormDefaultValues());
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
