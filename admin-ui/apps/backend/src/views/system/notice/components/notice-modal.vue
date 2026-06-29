<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { NoticeApi } from '#/api/system/notice';
import type { DictOption } from '#/composables/crud/use-dict-options';

import { computed, markRaw, nextTick, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { Button, MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { saveNotice, updateNotice } from '#/api/system/notice';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import ConfigRichTextEditor from '#/views/system/config/components/config-rich-text-editor.vue';
import UserSelectModal from './user-select-modal.vue';

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
const isEdit = ref(false);
const userSelectVisible = ref(false);
const selectedUserIds = ref<number[]>([]);

const selectedUserCount = computed(() => selectedUserIds.value.length);

const { getDictOptions } = useDictOptions();

function createNoticeTypeProps(disabled?: boolean) {
  return {
    options: noticeTypeOptions.value,
    disabled: typeof disabled === 'boolean' ? disabled : isEdit.value,
  };
}

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  layout: 'vertical',
  commonConfig: {
    labelWidth: 90,
  },
  wrapperClass: 'grid-cols-1',
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
      component: 'Input',
      defaultValue: [],
      fieldName: 'users',
      hideLabel: true,
      formItemClass: 'md:col-span-2',
      label: $t('system.notice.receiveUser'),
    },
    {
      component: markRaw(ConfigRichTextEditor),
      fieldName: 'content',
      formItemClass: 'md:col-span-2',
      label: $t('system.notice.content'),
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
  class: 'w-[1000px] max-w-[94vw] !p-4',
});

function updateNoticeTypeSchema() {
  formApi.updateSchema([
    {
      fieldName: 'type',
      componentProps: createNoticeTypeProps(),
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

function openUserSelect() {
  userSelectVisible.value = true;
}

function handleUserSelectConfirm(ids: number[]) {
  selectedUserIds.value = ids;
  formApi.setFieldValue('users', ids);
}

function clearSelectedUsers() {
  selectedUserIds.value = [];
  formApi.setFieldValue('users', []);
}

async function open(data?: NoticeApi.SubmitPayload) {
  isEdit.value = Boolean(data?.id);
  updateNoticeTypeSchema();
  modalApi.setState({
    title: isEdit.value ? $t('system.notice.editTitle') : $t('system.notice.createTitle'),
  });
  modalApi.open();
  await fetchNoticeTypeOptions();
  await formApi.resetForm();
  if (data) {
    selectedUserIds.value = data.users ?? [];
    formApi.setValues({
      ...data,
      users: data.users ?? [],
    });
  } else {
    selectedUserIds.value = [];
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
    <Form>
      <template #users>
        <div class="flex flex-col gap-1">
          <label class="flex items-center gap-1 text-sm leading-6">
            <span>{{ $t('system.notice.receiveUser') }}</span>
          </label>
          <div class="flex items-center gap-2">
            <Button theme="primary" @click="openUserSelect">
              {{ $t('system.notice.selectUser') }}
            </Button>
            <span v-if="selectedUserCount > 0" class="text-primary">
              已选择 {{ selectedUserCount }} 位
            </span>
            <span v-else class="text-muted-foreground/80">
              {{ $t('system.notice.receiveUserPlaceholder') }}
            </span>
            <Button
              v-if="selectedUserCount > 0"
              theme="default"
              variant="text"
              size="small"
              @click="clearSelectedUsers"
            >
              清空
            </Button>
          </div>
        </div>
      </template>
    </Form>
  </Modal>

  <UserSelectModal
    v-model:visible="userSelectVisible"
    :selected-ids="selectedUserIds"
    @confirm="handleUserSelectConfirm"
  />
</template>
