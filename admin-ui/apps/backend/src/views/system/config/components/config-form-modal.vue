<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { ConfigApi } from '#/api/system/config';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getConfigGroupList, saveConfig, updateConfig } from '#/api/system/config';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import { createConfigFormDefaultValues } from '../schemas';
import { inputComponentOptions } from '../model';

const emit = defineEmits(['success']);

const groupOptions = ref<DictOption<number>[]>([]);
const statusOptions = ref<DictOption[]>([]);

const { getDictOptions } = useDictOptions();

const fallbackStatusOptions: DictOption[] = [
  { label: $t('common.statusEnabled'), value: 1 },
  { label: $t('common.statusDisabled'), value: 2 },
];

function normalizeStatusOptionValue(value: unknown) {
  if (value === '1' || value === 1) return 1;
  if (value === '2' || value === 2) return 2;
  return value as string | number;
}

function normalizeStatusOptions(options: DictOption[]) {
  const fallbackLabelMap = new Map([
    ['1', String(fallbackStatusOptions[0]?.label ?? '启用')],
    ['2', String(fallbackStatusOptions[1]?.label ?? '禁用')],
  ]);
  return options.map((option) => {
    const normalizedValue = normalizeStatusOptionValue(option.value);
    const valueText = String(normalizedValue ?? '').trim();
    const labelText = String(option.label ?? '').trim();
    const normalizedLabel =
      !labelText || labelText === valueText
        ? (fallbackLabelMap.get(valueText) ?? labelText)
        : labelText;
    return {
      ...option,
      label: normalizedLabel,
      value: normalizedValue,
    };
  });
}

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: { labelWidth: 100 },
  wrapperClass: 'grid-cols-1 gap-x-4 md:grid-cols-2',
  schema: [
    {
      component: 'Input',
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: 'Select',
      componentProps: {
        options: groupOptions.value,
        placeholder: $t('ui.placeholder.select'),
      },
      fieldName: 'group_id',
      label: $t('system.config.group'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.config.placeholder.enterConfigName') },
      fieldName: 'name',
      label: $t('system.config.name'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.config.placeholder.enterConfigCode') },
      fieldName: 'key',
      label: $t('system.config.code'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.config.placeholder.enterConfigValue') },
      fieldName: 'value',
      label: $t('system.config.value'),
    },
    {
      component: 'InputNumber',
      componentProps: { min: 0, max: 999 },
      defaultValue: 0,
      fieldName: 'sort',
      label: $t('common.sort'),
    },
    {
      component: 'Select',
      componentProps: {
        options: inputComponentOptions,
        placeholder: $t('ui.placeholder.select'),
      },
      defaultValue: 'input',
      fieldName: 'input_type',
      label: $t('system.config.inputComponent'),
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        options: statusOptions.value,
        placeholder: $t('ui.placeholder.select'),
      },
      defaultValue: 1,
      fieldName: 'status',
      label: $t('common.status'),
      rules: 'required',
    },
    {
      component: 'Textarea',
      componentProps: { placeholder: $t('system.config.placeholder.enterRemark') },
      fieldName: 'remark',
      formItemClass: 'md:col-span-2',
      label: $t('common.remark'),
    },
    {
      component: 'Textarea',
      componentProps: {
        autosize: { minRows: 3, maxRows: 6 },
        placeholder: $t('system.config.optionsPlaceholder'),
      },
      dependencies: {
        show: (values) =>
          ['select', 'radio', 'checkbox'].includes(values?.input_type),
        triggerFields: ['input_type'],
      },
      fieldName: 'config_select_data',
      formItemClass: 'md:col-span-2',
      label: $t('system.config.options'),
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<ConfigApi.ConfigSubmitPayload & {
        config_select_data?: string;
      }>();
      const payload: ConfigApi.ConfigSubmitPayload = {
        ...values,
        group_id: Number(values.group_id),
        sort: Number(values.sort ?? 0),
        status: Number(values.status ?? 1),
      };

      if (values.config_select_data?.trim()) {
        try {
          const parsed = JSON.parse(values.config_select_data);
          if (!Array.isArray(parsed)) {
            MessagePlugin.error($t('common.jsonArrayRequired'));
            return;
          }
          payload.config_select_data = values.config_select_data;
        } catch {
          MessagePlugin.error($t('common.jsonArrayRequired'));
          return;
        }
      } else {
        payload.config_select_data = undefined;
      }

      modalApi.setState({ confirmLoading: true });
      if (payload.id) {
        await updateConfig(payload);
      } else {
        await saveConfig(payload);
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
  class: 'w-[980px] max-w-[94vw]',
});

interface OpenOptions {
  data?: ConfigApi.ConfigSubmitPayload;
}

async function fetchGroupOptions() {
  const list = await getConfigGroupList();
  groupOptions.value = list.map((item) => ({
    label: item.name,
    value: item.id,
  }));
  formApi.updateSchema([
    {
      fieldName: 'group_id',
      componentProps: { options: groupOptions.value },
    },
  ]);
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  statusOptions.value =
    options && options.length > 0
      ? normalizeStatusOptions(options)
      : fallbackStatusOptions;
  formApi.updateSchema([
    {
      fieldName: 'status',
      componentProps: { options: statusOptions.value },
    },
  ]);
}

async function open(options?: OpenOptions) {
  modalApi.setState({
    title: options?.data?.id ? $t('system.config.editTitle') : $t('system.config.addConfigTitle'),
  });
  modalApi.open();

  await Promise.all([fetchGroupOptions(), fetchStatusOptions()]);
  await formApi.resetForm();
  const defaults = createConfigFormDefaultValues();
  if (options?.data) {
    const configData = { ...options.data } as ConfigApi.ConfigSubmitPayload & {
      config_select_data?: string;
    };
    if (Array.isArray(configData.config_select_data)) {
      configData.config_select_data = JSON.stringify(
        configData.config_select_data,
      );
    }
    formApi.setValues({
      ...defaults,
      ...configData,
      group_id: configData.group_id ?? defaults.group_id,
    });
  } else {
    formApi.setValues(defaults);
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
    <div class="config-form-modal">
      <Form />
    </div>
  </Modal>
</template>

<style scoped>
.config-form-modal {
  padding: 2px;
}

.config-form-modal :deep(.t-form__item) {
  margin-bottom: 14px;
}

.config-form-modal :deep(.t-form__label) {
  color: var(--td-text-color-secondary, #6b7280);
  font-weight: 500;
}

.config-form-modal :deep(.t-input),
.config-form-modal :deep(.t-textarea__inner),
.config-form-modal :deep(.t-select),
.config-form-modal :deep(.t-select__wrap) {
  border-radius: 8px;
}
</style>
