<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { MenuApi } from '#/api/system/menu';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { markRaw, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getMenuTreeOptions, saveMenu, updateMenu } from '#/api/system/menu';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import IconSelect from './icon-select.vue';

import {
  createMenuFormDefaultValues,
  menuHiddenOptions,
  menuTypeOptions,
  restfulOptions,
} from '../schemas';

const emit = defineEmits(['success']);

const parentMenuOptions = ref<MenuApi.TreeOptionItem[]>([]);
const statusOptions = ref<DictOption[]>([]);

const { getDictOptions } = useDictOptions();

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
      component: 'TreeSelect',
      componentProps: {
        data: parentMenuOptions.value,
        keys: { children: 'children', label: 'label', value: 'value' },
        placeholder: $t('ui.placeholder.select'),
      },
      fieldName: 'parent_id',
      label: $t('system.menu.parentMenu'),
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'name',
      label: $t('system.menu.title'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: { options: menuTypeOptions },
      defaultValue: 'M',
      fieldName: 'type',
      label: $t('system.menu.type'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      fieldName: 'code',
      label: $t('system.menu.code'),
      rules: 'required',
    },
    {
      component: markRaw(IconSelect),
      componentProps: { placeholder: $t('ui.placeholder.input') },
      dependencies: {
        show: (values) => isFieldVisible('icon', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'icon',
      label: $t('system.menu.icon'),
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      dependencies: {
        show: (values) => isFieldVisible('route', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'route',
      label: $t('system.menu.router'),
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      dependencies: {
        show: (values) => isFieldVisible('component', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'component',
      label: $t('system.menu.component'),
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input') },
      dependencies: {
        show: (values) => isFieldVisible('redirect', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'redirect',
      label: $t('system.menu.redirect'),
    },
    {
      component: 'InputNumber',
      componentProps: { min: 0, max: 1000 },
      defaultValue: 1,
      dependencies: {
        show: (values) => isFieldVisible('sort', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'sort',
      label: $t('common.sort'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: { options: menuHiddenOptions },
      defaultValue: 2,
      dependencies: {
        show: (values) => isFieldVisible('is_hidden', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'is_hidden',
      label: $t('system.menu.isHidden'),
    },
    {
      component: 'RadioGroup',
      componentProps: { options: restfulOptions },
      defaultValue: '2',
      dependencies: {
        show: (values) => isFieldVisible('restful', values?.type),
        triggerFields: ['type'],
      },
      fieldName: 'restful',
      label: $t('system.menu.generateButton'),
    },
    {
      component: 'RadioGroup',
      componentProps: {
        options: statusOptions.value,
      },
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

      const values = await formApi.getValues<MenuApi.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });
      const payload: MenuApi.SubmitPayload = {
        ...values,
        is_hidden: Number(values.is_hidden),
        parent_id: Number(values.parent_id ?? 0),
        sort: Number(values.sort ?? 1),
        status: Number(values.status ?? 1),
      };

      if (payload.id) {
        await updateMenu(Number(payload.id), payload);
      } else {
        await saveMenu(payload);
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

function isFieldVisible(field: string, type?: string) {
  if (type === 'B') {
    return false;
  }
  if (type === 'I' || type === 'L') {
    return ['icon', 'route', 'sort', 'is_hidden'].includes(field);
  }
  return true;
}

async function fetchParentOptions() {
  try {
    const tree = await getMenuTreeOptions({ onlyMenu: true }).catch(() => []);
    parentMenuOptions.value = [
      {
        id: 0,
        label: $t('system.menu.topMenu'),
        value: 0,
        children: tree as MenuApi.TreeOptionItem[],
      },
    ];
  } catch (error) {
    logger.error(error);
    parentMenuOptions.value = [{ id: 0, label: $t('system.menu.topMenu'), value: 0 }];
  }

  formApi.updateSchema([
    {
      fieldName: 'parent_id',
      componentProps: {
        data: parentMenuOptions.value,
      },
    },
  ]);
}

async function open(data?: Partial<MenuApi.SubmitPayload>) {
  modalApi.setState({
    title: data?.id ? $t('system.menu.editTitle') : $t('system.menu.createTitle'),
  });
  modalApi.open();

  try {
    statusOptions.value = (await getDictOptions('data_status')) || [
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

    await fetchParentOptions();
    await formApi.resetForm();
    formApi.setValues(createMenuFormDefaultValues());

    if (data) {
      formApi.setValues({
        ...data,
        parent_id: data.parent_id ?? 0,
      });
    }
    await nextTick();
    await formApi.resetValidate();
  } catch (error) {
    logger.error($t('system.menu.formLoadFailed'), error);
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
