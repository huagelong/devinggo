<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { ConfigApi } from '#/api/system/config';

import { nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';

import { saveConfigGroup } from '#/api/system/config';

import { createConfigGroupFormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: { labelWidth: 90 },
  wrapperClass: 'grid-cols-1 gap-x-4 md:grid-cols-2',
  schema: [
    {
      component: 'Input',
      dependencies: { show: false, triggerFields: [''] },
      fieldName: 'id',
      label: 'ID',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.config.placeholder.enterGroupName') },
      fieldName: 'name',
      label: $t('system.config.groupName'),
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: $t('system.config.placeholder.enterGroupCode') },
      fieldName: 'code',
      label: $t('system.config.groupCode'),
      rules: 'required',
    },
    {
      component: 'Textarea',
      componentProps: { placeholder: $t('system.config.placeholder.enterRemark') },
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
      const values = await formApi.getValues<ConfigApi.ConfigGroupSubmitPayload>();
      modalApi.setState({ confirmLoading: true });
      await saveConfigGroup(values);
      MessagePlugin.success($t('common.saveSuccess'));
      emit('success');
      modalApi.close();
    } catch (error) {
      logger.error(error);
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  class: 'w-[820px] max-w-[94vw]',
});

async function open() {
  modalApi.setState({ title: $t('system.config.addGroupTitle') });
  modalApi.open();
  await formApi.resetForm();
  formApi.setValues(createConfigGroupFormDefaultValues());
  await nextTick();
  await formApi.resetValidate();
}

defineExpose({
  open,
});
</script>

<template>
  <Modal>
    <div class="config-group-modal">
      <Form />
    </div>
  </Modal>
</template>

<style scoped>
.config-group-modal {
  padding: 2px;
}

.config-group-modal :deep(.t-form__item) {
  margin-bottom: 14px;
}

.config-group-modal :deep(.t-form__label) {
  color: var(--td-text-color-secondary, #6b7280);
  font-weight: 500;
}

.config-group-modal :deep(.t-input),
.config-group-modal :deep(.t-textarea__inner) {
  border-radius: 8px;
}
</style>
