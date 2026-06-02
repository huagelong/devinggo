<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { ApiGroupFormModel } from '../model';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveApiGroup, updateApiGroup } from '#/api/system/api-group';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import { createApiGroupFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const statusOptions = ref<DictOption[]>([]);
const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

const { getDictOptions } = useDictOptions();

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 90,
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
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'name',
      label: $t('system.apiGroup.name'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: { options: statusOptions.value },
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
      const values = await formApi.getValues<ApiGroupFormModel>();
      modalApi.setState({ confirmLoading: true });
      if (values.id) {
        await updateApiGroup(Number(values.id), values);
      } else {
        await saveApiGroup(values);
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

async function fetchStatusOptions() {
  try {
    const options = await getDictOptions('data_status');
    statusOptions.value =
      options && options.length > 0 ? options : fallbackStatusOptions;
  } catch (error) {
    logger.error(error);
    statusOptions.value = fallbackStatusOptions;
  } finally {
    formApi.updateSchema([
      {
        fieldName: 'status',
        componentProps: { options: statusOptions.value },
      },
    ]);
  }
}

async function open(data?: ApiGroupFormModel) {
  modalApi.setState({
    title: data?.id ? $t('system.apiGroup.editTitle') : $t('system.apiGroup.createTitle'),
  });
  modalApi.open();
  await fetchStatusOptions();
  await formApi.resetForm();
  formApi.setValues(createApiGroupFormDefaultValues());
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
