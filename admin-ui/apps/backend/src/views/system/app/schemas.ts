import type {
  AppColumnOptionItem,
  AppFormModel,
  AppSearchFormModel,
  AppTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createAppSearchForm(): AppSearchFormModel {
  return {
    created_at: [],
    name: '',
    status: undefined,
  };
}

export function createAppFormDefaultValues(): AppFormModel {
  return {
    intro: '',
    name: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createAppTableColumns(): AppTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'id', title: 'ID', width: 80 },
    { colKey: 'name', title: $t('system.app.name'), minWidth: 160 },
    { colKey: 'app_id', title: 'AppId', minWidth: 200 },
    { colKey: 'intro', title: $t('system.app.description'), minWidth: 200 },
    { colKey: 'sort', title: $t('common.sort'), width: 120 },
    { colKey: 'status', title: $t('common.status'), width: 120 },
    { colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 320,
    },
  ];
}

export function createAppColumnOptions(
  columns: AppTableColumn[],
): AppColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
