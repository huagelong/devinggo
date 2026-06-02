<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DeptApi } from '#/api/system/dept';

import { nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getDeptTree, saveDept, updateDept } from '#/api/system/dept';

import { createDeptFormDefaultValues } from '../schemas';

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
      component: 'TreeSelect',
      componentProps: {
        data: [],
        keys: { label: 'label', value: 'value', children: 'children' },
        placeholder: $t('ui.placeholder.select'),
      },
      fieldName: 'parent_id',
      label: $t('system.dept.parentDept'),
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'name',
      label: $t('system.dept.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'leader',
      label: $t('system.dept.leader'),
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'phone',
      label: $t('system.dept.phone'),
      rules: 'pattern:^1[3-9]\\d{9}$#' + $t('system.dept.invalidPhone'),
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

      const values = await formApi.getValues<DeptApi.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await updateDept(Number(values.id), values);
      } else {
        await saveDept(values);
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
  class: 'w-[940px] max-w-[94vw]',
});

async function open(data?: Partial<DeptApi.SubmitPayload>) {
  modalApi.setState({
    title: data?.id ? $t('system.dept.editTitle') : $t('system.dept.createTitle'),
  });
  modalApi.open();

  try {
    const deptTree = await getDeptTree().catch(() => [] as DeptApi.TreeNode[]);

    formApi.updateSchema([
      {
        fieldName: 'parent_id',
        componentProps: {
          data: deptTree,
        },
      },
    ]);

    await formApi.resetForm();
    formApi.setValues(createDeptFormDefaultValues());
    if (data) {
      formApi.setValues({
        ...data,
        status:
          data.status !== undefined ? Number(data.status) : createDeptFormDefaultValues().status,
      });
    }
    await nextTick();
    await formApi.resetValidate();
  } catch (error) {
    logger.error($t('common.deptFormLoadFailed'), error);
  }
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
