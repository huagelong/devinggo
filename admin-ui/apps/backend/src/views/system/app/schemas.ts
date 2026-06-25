import type {
  AppColumnOptionItem,
  AppFormModel,
  AppSearchFormModel,
  AppTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createAppSearchForm(): AppSearchFormModel {
  return {
    app_name: '',
    created_at: [],
    status: undefined,
  };
}

export function createAppFormDefaultValues(): AppFormModel {
  return {
    app_name: '',
    app_id: '',
    app_secret: '',
    description: '',
    group_id: undefined,
    remark: '',
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
    { colKey: 'app_name', title: $t('system.app.name'), minWidth: 160 },
    { colKey: 'app_id', title: 'AppId', minWidth: 200 },
    { colKey: 'description', title: $t('system.app.description'), minWidth: 200 },
    { colKey: 'group_name', title: $t('system.app.group'), minWidth: 120 },
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
