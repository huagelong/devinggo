import type {
  PostColumnOptionItem,
  PostFormModel,
  PostSearchFormModel,
  PostTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createPostSearchForm(): PostSearchFormModel {
  return {
    code: '',
    created_at: [],
    name: '',
    status: undefined,
  };
}

export function createPostFormDefaultValues(): PostFormModel {
  return {
    code: '',
    name: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createPostTableColumns(): PostTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { align: 'center', colKey: 'id', title: 'ID', width: 80 },
    { align: 'center', colKey: 'name', minWidth: 140, title: $t('system.post.name') },
    { align: 'center', colKey: 'code', minWidth: 140, title: $t('system.post.code') },
    { align: 'center', colKey: 'sort', title: $t('common.sort'), width: 140 },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { align: 'center', colKey: 'remark', minWidth: 160, title: $t('common.remark') },
    { align: 'center', colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 220,
    },
  ];
}

export function createPostColumnOptions(
  columns: PostTableColumn[],
): PostColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
