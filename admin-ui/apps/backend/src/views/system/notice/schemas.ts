import type {
  NoticeColumnOptionItem,
  NoticeSearchFormModel,
  NoticeTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createNoticeSearchForm(): NoticeSearchFormModel {
  return {
    created_at: [],
    title: '',
    type: undefined,
  };
}

export function createNoticeTableColumns(): NoticeTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'title', title: $t('system.notice.title'), minWidth: 200 },
    { align: 'center', colKey: 'type', title: $t('system.notice.type'), width: 100 },
    { colKey: 'content', title: $t('system.notice.content'), minWidth: 260 },
    { colKey: 'remark', title: $t('common.remark'), minWidth: 160 },
    { colKey: 'created_at', title: $t('common.createTime'), minWidth: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 220,
    },
  ];
}

export function createNoticeColumnOptions(
  columns: NoticeTableColumn[],
): NoticeColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
