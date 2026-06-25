<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { ApiFormModel } from '../model';
import type { OptionItem, IdType } from '#/types/common';
import type { DictOption } from '#/composables/crud/use-dict-options';
import type { ApiManageApi } from '#/api/system/api';

import { markRaw, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin, Select } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveApi, updateApi } from '#/api/system/api';
import { getApiGroupList } from '#/api/system/api-group';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import { createApiFormDefaultValues } from '../schemas';

type SelectOption = OptionItem<IdType>;

const emit = defineEmits(['success']);

const groupOptions = ref<SelectOption[]>([]);
const requestModeOptions = ref<DictOption[]>([]);
const statusOptions = ref<DictOption[]>([]);

const fallbackRequestModes: DictOption[] = [
  { label: 'GET', value: 'GET' },
  { label: 'POST', value: 'POST' },
  { label: 'PUT', value: 'PUT' },
  { label: 'DELETE', value: 'DELETE' },
];

const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const authModeOptions: DictOption[] = [
  { label: $t('system.api.simpleMode'), value: 1 },
  { label: $t('system.api.complexMode'), value: 2 },
];

const { getDictOptions } = useDictOptions();

function createSelectProps(options: DictOption[] | SelectOption[], placeholder: string) {
  return {
    options,
    placeholder,
    clearable: true,
  };
}

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 100,
  },
  wrapperClass: 'grid-cols-1 md:grid-cols-2',
  schema: [
    {
      component: 'Input',
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: markRaw(Select),
      componentProps: createSelectProps(groupOptions.value, $t('ui.placeholder.select')),
      fieldName: 'group_id',
      label: $t('system.api.group'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'name',
      label: $t('system.api.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.api.codePlaceholder') },
      fieldName: 'access_name',
      label: $t('system.api.code'),
      rules: 'required',
    },
    {
      component: markRaw(Select),
      componentProps: createSelectProps(requestModeOptions.value, $t('ui.placeholder.select')),
      fieldName: 'request_mode',
      label: $t('system.api.requestMode'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: { options: authModeOptions },
      defaultValue: 1,
      fieldName: 'auth_mode',
      label: $t('system.api.authMode'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: { options: statusOptions.value },
      defaultValue: 1,
      fieldName: 'status',
      label: $t('common.status'),
      rules: 'required',
    },
    {
      component: 'Textarea',
      componentProps: { placeholder: $t('ui.placeholder.input') },
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
      const values = await formApi.getValues<ApiFormModel>();
      modalApi.setState({ confirmLoading: true });
      const payload: ApiManageApi.SubmitPayload = {
        ...values,
        group_id: values.group_id as IdType,
        request_mode: values.request_mode as number | string,
      };
      if (values.id) {
        await updateApi(Number(values.id), payload);
      } else {
        await saveApi(payload);
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
  class: 'w-[960px] max-w-[94vw]',
});

function updateFormSchemas() {
  formApi.updateSchema([
    {
      fieldName: 'group_id',
      componentProps: createSelectProps(groupOptions.value, $t('ui.placeholder.select')),
    },
    {
      fieldName: 'request_mode',
      componentProps: createSelectProps(requestModeOptions.value, $t('ui.placeholder.select')),
    },
    {
      fieldName: 'status',
      componentProps: { options: statusOptions.value },
    },
  ]);
}

async function fetchFormOptions() {
  try {
    const [groupList, requestModes, statuses] = await Promise.all([
      getApiGroupList().catch(() => []),
      getDictOptions('request_mode'),
      getDictOptions('data_status'),
    ]);
    groupOptions.value = (groupList || []).map((item) => ({
      label: item.name,
      value: item.id,
    }));
    requestModeOptions.value =
      (requestModes && requestModes.length > 0 ? requestModes : fallbackRequestModes).map(
        (item) => ({
          ...item,
          value: item.value ?? item.label,
        }),
      );
    statusOptions.value =
      statuses && statuses.length > 0 ? statuses : fallbackStatusOptions;
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.selectOptionsLoadFailed'));
    groupOptions.value = [];
    requestModeOptions.value = fallbackRequestModes;
    statusOptions.value = fallbackStatusOptions;
  } finally {
    updateFormSchemas();
  }
}

async function open(data?: ApiFormModel) {
  modalApi.setState({
    title: data?.id ? $t('system.api.editTitle') : $t('system.api.createTitle'),
  });
  modalApi.open();
  await fetchFormOptions();
  await formApi.resetForm();
  const defaultValues = createApiFormDefaultValues();
  if (!defaultValues.request_mode && requestModeOptions.value.length > 0) {
    defaultValues.request_mode = requestModeOptions.value[0]?.value;
  }
  formApi.setValues(defaultValues);
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
