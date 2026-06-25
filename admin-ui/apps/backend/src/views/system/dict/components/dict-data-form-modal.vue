<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DictApi } from '#/api/system/dict';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveDictData, updateDictData } from '#/api/system/dict';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import { createDictDataFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const statusOptions = ref<DictOption[]>([]);
const currentTypeInfo = ref<{ id: number; code: string; name?: string }>();
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
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'type_id',
      label: 'type_id',
    },
    {
      component: 'Input',
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'code',
      label: 'code',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'label',
      label: $t('system.dict.label'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'value',
      label: $t('system.dict.value'),
      rules: 'required',
    },
    {
      component: 'InputNumber',
      componentProps: { min: 0, max: 1000 },
      defaultValue: 1,
      fieldName: 'sort',
      label: $t('common.sort'),
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

      const values = await formApi.getValues<DictApi.DictDataSubmitPayload>();
      if (!values.type_id) {
        MessagePlugin.warning($t('common.missingDictTypeInfo'));
        return;
      }
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await updateDictData(Number(values.id), values);
      } else {
        await saveDictData(values);
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
  class: 'w-[860px] max-w-[94vw]',
});

interface OpenOptions {
  data?: DictApi.DictDataSubmitPayload;
  typeInfo: { id: number; code: string; name?: string };
}

async function open(options: OpenOptions) {
  currentTypeInfo.value = options.typeInfo;
  modalApi.setState({
    title: options.data?.id ? $t('system.dict.editDictData') : $t('system.dict.createDictData', [options.typeInfo.name ?? '']),
  });
  modalApi.open();

  statusOptions.value =
    (await getDictOptions('data_status')) || [
      { label: $t('common.statusEnabled'), value: 1 },
      { label: $t('common.statusDisabled'), value: 2 },
    ];

  formApi.updateSchema([
    {
      fieldName: 'status',
      componentProps: {
        options: statusOptions.value,
      },
    },
  ]);

  await formApi.resetForm();
  formApi.setValues(
    createDictDataFormDefaultValues(options.typeInfo.id, options.typeInfo.code),
  );
  if (options.data) {
    formApi.setValues({
      ...options.data,
      type_id: options.typeInfo.id,
      code: options.typeInfo.code,
    });
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
