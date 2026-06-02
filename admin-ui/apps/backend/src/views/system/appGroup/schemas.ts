import type {
  AppGroupColumnOptionItem,
  AppGroupFormModel,
  AppGroupSearchFormModel,
  AppGroupTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createAppGroupSearchForm(): AppGroupSearchFormModel {
  return {
    created_at: [],
    name: '',
    status: undefined,
  };
}

export function createAppGroupFormDefaultValues(): AppGroupFormModel {
  return {
    name: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createAppGroupTableColumns(): AppGroupTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'id', title: 'ID', width: 80 },
    { colKey: 'name', title: $t('system.appGroup.name'), minWidth: 160 },
    { colKey: 'sort', title: $t('common.sort'), width: 120 },
    { colKey: 'status', title: $t('common.status'), width: 120 },
    { colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 240,
    },
  ];
}

export function createAppGroupColumnOptions(
  columns: AppGroupTableColumn[],
): AppGroupColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
