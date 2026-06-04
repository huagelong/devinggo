<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { AppApi } from '#/api/system/app';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getAppGroupPageList } from '#/api/system/app-group';
import { getAppId, getAppSecret, saveApp, updateApp } from '#/api/system/app';

import { createAppFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const baseValues = ref<AppApi.SubmitPayload>(createAppFormDefaultValues());
const groupOptions = ref<{ label: string; value: number }[]>([]);

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
      fieldName: 'app_name',
      label: $t('system.app.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'app_id',
      label: $t('system.app.code'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'app_secret',
      label: $t('system.app.secret'),
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        options: groupOptions.value,
        placeholder: $t('ui.placeholder.select'),
      },
      fieldName: 'group_id',
      label: $t('system.app.group'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
      },
      fieldName: 'description',
      label: $t('system.app.description'),
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
  class: 'w-[900px] max-w-[94vw]',
});

async function fetchGroupOptions() {
  try {
    const res = await getAppGroupPageList({ pageSize: 999 });
    const items = (res as any)?.items || [];
    groupOptions.value = items.map((item: any) => ({
      label: item.name,
      value: item.id,
    }));
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
    title: data?.id ? $t('system.app.editTitle') : $t('system.app.createTitle'),
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
