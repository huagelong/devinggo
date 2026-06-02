<script lang="ts" setup>
import { logger } from '#/utils/logger';
import type { DeptApi } from '#/api/system/dept';
import type { PostApi } from '#/api/system/post';
import type { RoleApi } from '#/api/system/role';
import type { UserApi } from '#/api/system/user';

import { defineComponent, h, markRaw, nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { uploadImageFileApi } from '#/api/system/upload';
import { getDeptTree } from '#/api/system/dept';
import { getPostList } from '#/api/system/post';
import { getRoleList } from '#/api/system/role';
import { getUserDetail, saveUser, updateUser } from '#/api/system/user';

const emit = defineEmits(['success']);

interface AvatarUploadResponse {
  url?: string;
}

type UserModalOpenData = Partial<UserApi.Detail>;

function extractEntityIds(list?: UserApi.RelatedEntity[]) {
  return list?.map((item) => item.id) ?? [];
}

function normalizeListData<T>(data: T[] | { items?: T[] } | null | undefined): T[] {
  if (Array.isArray(data)) {
    return data;
  }
  return Array.isArray(data?.items) ? data.items : [];
}

const AvatarUpload = defineComponent({
  props: {
    modelValue: { type: String, default: '' },
    value: { type: String, default: '' },
  },
  emits: ['update:modelValue', 'update:value', 'change'],
  setup(props, { emit: emitInner }) {
    const getAvatarValue = () => props.modelValue || props.value;

    function handleClick() {
      const input = document.createElement('input');
      input.type = 'file';
      input.accept = 'image/*';
      input.onchange = async (event: Event) => {
        const file = (event.target as HTMLInputElement).files?.[0];
        if (!file) return;

        try {
          const res = (await uploadImageFileApi(file)) as AvatarUploadResponse;
          if (res?.url) {
            emitInner('update:modelValue', res.url);
            emitInner('update:value', res.url);
            emitInner('change', res.url);
          }
        } catch {
          MessagePlugin.error($t('common.avatarUploadFailed2'));
        }
      };
      input.click();
    }

    return () =>
      h('div', { class: 'flex items-center gap-3' }, [
        h(
          'div',
          {
            onClick: handleClick,
            class:
              'relative flex h-16 w-16 cursor-pointer items-center justify-center overflow-hidden rounded-full border-2 border-dashed border-gray-300 bg-gray-50 hover:border-blue-400',
          },
          [
            getAvatarValue()
              ? h('img', {
                  src: getAvatarValue(),
                  class: 'h-full w-full rounded-full object-cover',
                  alt: $t('system.user.avatar'),
                })
              : h(
                  'span',
                  { class: 'text-3xl leading-none text-gray-400' },
                  '+',
                ),
          ],
        ),
        h('div', { class: 'flex flex-col gap-1' }, [
          h('span', { class: 'text-sm text-gray-500' }, $t('system.user.clickUploadAvatar')),
          h(
            'span',
            { class: 'text-xs text-gray-400' },
            $t('system.user.supportImageFormat'),
          ),
        ]),
      ]);
  },
});

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 80,
  },
  wrapperClass: 'grid-cols-1 md:grid-cols-2',
  schema: [
    {
      fieldName: 'id',
      label: 'ID',
      component: 'Input',
      dependencies: {
        show: false,
        triggerFields: [''],
      },
    },
    {
      fieldName: 'avatar',
      label: $t('system.user.avatar'),
      component: markRaw(AvatarUpload),
      formItemClass: 'md:col-span-2',
    },
    {
      fieldName: 'username',
      label: $t('system.user.username'),
      component: 'Input',
      rules: 'required',
      componentProps: { placeholder: $t('ui.placeholder.input', [$t('system.user.username')]) },
    },
    {
      fieldName: 'dept_ids',
      label: $t('system.user.dept'),
      component: 'TreeSelect',
      componentProps: {
        data: [],
        keys: { label: 'label', value: 'value', children: 'children' },
        multiple: true,
        placeholder: $t('ui.placeholder.select', [$t('system.user.dept')]),
      },
      rules: 'required',
    },
    {
      fieldName: 'password',
      label: $t('system.user.password'),
      component: 'Input',
      dependencies: {
        show: (values) => !values?.id,
        triggerFields: ['id'],
      },
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input', [$t('system.user.password')]),
        type: 'password',
      },
    },
    {
      fieldName: 'nickname',
      label: $t('system.user.nickname'),
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input', [$t('system.user.nickname')]) },
    },
    {
      fieldName: 'role_ids',
      label: $t('system.user.role'),
      component: 'Select',
      componentProps: {
        options: [],
        keys: { label: 'name', value: 'id' },
        multiple: true,
        placeholder: $t('ui.placeholder.select', [$t('system.user.role')]),
      },
      rules: 'required',
    },
    {
      fieldName: 'phone',
      label: $t('system.user.phone'),
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input', [$t('system.user.phone')]) },
    },
    {
      fieldName: 'post_ids',
      label: $t('system.user.post'),
      component: 'Select',
      componentProps: {
        options: [],
        keys: { label: 'name', value: 'id' },
        multiple: true,
        placeholder: $t('ui.placeholder.select', [$t('system.user.post')]),
      },
    },
    {
      fieldName: 'email',
      label: $t('system.user.email'),
      component: 'Input',
      componentProps: { placeholder: $t('ui.placeholder.input', [$t('system.user.email')]) },
    },
    {
      fieldName: 'status',
      label: $t('common.status'),
      component: 'RadioGroup',
      defaultValue: 1,
      componentProps: {
        options: [
          { label: $t('common.statusEnabled'), value: 1 },
          { label: $t('common.statusDisabled'), value: 2 },
        ],
      },
      formItemClass: 'md:col-span-2',
    },
    {
      fieldName: 'user_type',
      label: $t('system.user.userType'),
      component: 'Select',
      defaultValue: '100',
      componentProps: {
        options: [{ label: $t('system.user.systemUser'), value: '100' }],
        placeholder: $t('ui.placeholder.select', [$t('system.user.userType')]),
      },
      rules: 'required',
      formItemClass: 'md:col-span-2',
    },
    {
      fieldName: 'remark',
      label: $t('common.remark'),
      component: 'Textarea',
      formItemClass: 'md:col-span-2',
      componentProps: { placeholder: $t('ui.placeholder.input', [$t('common.remark')]) },
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<UserApi.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });

      values.id ? await updateUser(values.id, values) : await saveUser(values);

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

async function open(data?: UserModalOpenData) {
  modalApi.setState({ title: data?.id ? $t('system.user.editTitle') : $t('system.user.createTitle') });
  modalApi.open();

  const [roleRes, postRes, deptRes] = await Promise.all([
    getRoleList().catch(() => null),
    getPostList().catch(() => null),
    getDeptTree().catch(() => null),
  ]);

  const roleOptions = normalizeListData<RoleApi.ListItem>(roleRes);
  const postOptions = normalizeListData<PostApi.ListItem>(postRes);
  const deptData = Array.isArray(deptRes) ? (deptRes as DeptApi.TreeNode[]) : [];

  formApi.updateSchema([
    {
      fieldName: 'role_ids',
      componentProps: { options: roleOptions },
    },
    {
      fieldName: 'post_ids',
      componentProps: { options: postOptions },
    },
    {
      fieldName: 'dept_ids',
      componentProps: { data: deptData },
    },
  ]);

  await formApi.resetForm();

  if (data?.id) {
    const detail = await getUserDetail(data.id).catch(() => null);
    const detailValues = detail
      ? {
          ...detail,
          dept_ids:
            detail.dept_ids ??
            extractEntityIds(detail.deptList) ??
            data.dept_ids ??
            [],
          role_ids:
            detail.role_ids ??
            extractEntityIds(detail.roleList) ??
            data.role_ids ??
            [],
          post_ids:
            detail.post_ids ??
            extractEntityIds(detail.postList) ??
            data.post_ids ??
            [],
        }
      : data;
    formApi.setValues(detailValues);
  } else if (data) {
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
