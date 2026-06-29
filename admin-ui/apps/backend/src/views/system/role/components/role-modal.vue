<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { RoleApi } from '#/api/system/role';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveRole, updateRole } from '#/api/system/role';

import { createRoleFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const baseValues = ref<RoleApi.SubmitPayload>(createRoleFormDefaultValues());

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
        placeholder: $t('ui.placeholder.input', [$t('system.role.name')]),
      },
      fieldName: 'name',
      label: $t('system.role.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input', [$t('system.role.code')]),
      },
      fieldName: 'code',
      label: $t('system.role.code'),
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
        placeholder: $t('ui.placeholder.input', [$t('common.remark')]),
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

      const values = await formApi.getValues<Partial<RoleApi.SubmitPayload>>();
      const payload: RoleApi.SubmitPayload = {
        ...baseValues.value,
        ...values,
      };

      modalApi.setState({ confirmLoading: true });

      if (payload.id) {
        await updateRole(Number(payload.id), payload);
      } else {
        await saveRole(payload);
      }

      MessagePlugin.success(payload.id ? $t('common.updateSuccess') : $t('common.createSuccess'));
      emit('success');
      modalApi.close();
    } catch (error) {
      logger.error(error);
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  class: 'w-[900px] max-w-[94vw]',
});

async function open(data?: Partial<RoleApi.SubmitPayload>) {
  const defaultValues = createRoleFormDefaultValues();
  baseValues.value = {
    ...defaultValues,
    ...data,
  };

  modalApi.setState({
    title: data?.id ? $t('system.role.editTitle') : $t('system.role.createTitle'),
  });
  modalApi.open();

  try {
    await formApi.resetForm();
    formApi.setValues(baseValues.value);
    await nextTick();
    await formApi.resetValidate();
  } catch (error) {
    logger.error($t('common.formLoadFailed'), error);
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
