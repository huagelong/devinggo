<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { NoticeApi } from '#/api/system/notice';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { markRaw, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin, Select } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { getUserInfoByIds } from '#/api/system/common';
import { getUserList } from '#/api/system/user';
import { saveNotice, updateNotice } from '#/api/system/notice';
import { useDictOptions } from '#/composables/crud/use-dict-options';

type UserSelectOption = { label: string; value: number };

const emit = defineEmits(['success']);

const fallbackNoticeTypeOptions: DictOption[] = [
  { label: $t('system.notice.noticeType'), value: 1 },
  { label: $t('system.notice.announcementType'), value: 2 },
];

function normalizeNoticeTypeOptions(options: DictOption[]) {
  return (options || []).map((item) => {
    const numericValue = Number(item.value);
    return Number.isNaN(numericValue) ? { ...item } : { ...item, value: numericValue };
  });
}

const noticeTypeOptions = ref<DictOption[]>([]);
const userOptions = ref<UserSelectOption[]>([]);
const userLoading = ref(false);
const isEdit = ref(false);

const { getDictOptions } = useDictOptions();

const handleUserSearch = (value: string) => {
  void fetchUserOptions(value);
};

function createNoticeTypeProps(disabled?: boolean) {
  return {
    options: noticeTypeOptions.value,
    disabled: typeof disabled === 'boolean' ? disabled : isEdit.value,
  };
}

function createUserSelectProps(disabled?: boolean) {
  return {
    multiple: true,
    filterable: true,
    placeholder: $t('system.notice.selectReceiveUser'),
    loading: userLoading.value,
    options: userOptions.value,
    onSearch: handleUserSearch,
    disabled: typeof disabled === 'boolean' ? disabled : isEdit.value,
    minCollapsedNum: 3,
  };
}

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
      fieldName: 'title',
      label: $t('system.notice.title'),
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      componentProps: createNoticeTypeProps(),
      defaultValue: 1,
      fieldName: 'type',
      label: $t('system.notice.type'),
      rules: 'required',
    },
    {
      component: markRaw(Select),
      componentProps: createUserSelectProps(),
      defaultValue: [],
      fieldName: 'users',
      label: $t('system.notice.receiveUser'),
    },
    {
      component: 'Textarea',
      componentProps: {
        autosize: { minRows: 8, maxRows: 20 },
        placeholder: $t('system.notice.contentPlaceholder'),
      },
      fieldName: 'content',
      formItemClass: 'md:col-span-2',
      label: $t('system.notice.content'),
      rules: 'required',
      description: $t('system.notice.contentDescription'),
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
      const values = await formApi.getValues<NoticeApi.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await updateNotice(Number(values.id), values);
      } else {
        await saveNotice(values);
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
  class: 'w-[1000px] max-w-[94vw]',
});

function updateNoticeTypeSchema() {
  formApi.updateSchema([
    {
      fieldName: 'type',
      componentProps: createNoticeTypeProps(),
    },
  ]);
}

function updateUserSchema() {
  formApi.updateSchema([
    {
      fieldName: 'users',
      componentProps: createUserSelectProps(),
    },
  ]);
}

async function fetchNoticeTypeOptions() {
  const options = await getDictOptions('backend_notice_type');
  noticeTypeOptions.value = normalizeNoticeTypeOptions(
    options.length > 0 ? options : fallbackNoticeTypeOptions,
  );
  updateNoticeTypeSchema();
}

function normalizeUserOption(item: {
  id: number | string;
  nickname?: string;
  username: string;
}): UserSelectOption | null {
  const value = Number(item.id);
  if (Number.isNaN(value)) {
    return null;
  }
  return {
    label: item.nickname ? `${item.nickname} (${item.username})` : item.username,
    value,
  };
}

async function fetchUserOptions(keyword = '') {
  userLoading.value = true;
  updateUserSchema();
  try {
    const response = await getUserList({
      page: 1,
      pageSize: 50,
      username: keyword || undefined,
    });
    const list = response.items ?? [];
    userOptions.value = list
      .map(normalizeUserOption)
      .filter((item): item is UserSelectOption => Boolean(item));
  } catch (error) {
    logger.error(error);
  } finally {
    userLoading.value = false;
    updateUserSchema();
  }
}

async function ensureSelectedUsers(userIds?: number[]) {
  if (!Array.isArray(userIds) || userIds.length === 0) {
    return;
  }
  try {
    const response = await getUserInfoByIds({ ids: userIds });
    if (!Array.isArray(response)) return;
    const existingIds = new Set(userOptions.value.map((item) => item.value));
    const extraOptions = response
      .map(normalizeUserOption)
      .filter((item): item is UserSelectOption => {
        if (!item) return false;
        return !existingIds.has(item.value);
      });
    if (extraOptions.length > 0) {
      userOptions.value = [...userOptions.value, ...extraOptions];
      updateUserSchema();
    }
  } catch (error) {
    logger.error(error);
  }
}

async function open(data?: NoticeApi.SubmitPayload) {
  isEdit.value = Boolean(data?.id);
  updateNoticeTypeSchema();
  updateUserSchema();
  modalApi.setState({
    title: isEdit.value ? $t('system.notice.editTitle') : $t('system.notice.createTitle'),
  });
  modalApi.open();
  await Promise.all([fetchNoticeTypeOptions(), fetchUserOptions()]);
  await formApi.resetForm();
  if (data) {
    await ensureSelectedUsers(data.users);
    formApi.setValues({
      ...data,
      users: data.users ?? [],
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
