import type {
  ApiGroupColumnOptionItem,
  ApiGroupFormModel,
  ApiGroupSearchFormModel,
  ApiGroupTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createApiGroupSearchForm(): ApiGroupSearchFormModel {
  return {
    name: '',
    status: undefined,
    created_at: [],
  };
}

export function createApiGroupTableColumns(): ApiGroupTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'name', title: $t('system.apiGroup.name'), minWidth: 220 },
    { colKey: 'status', title: $t('common.status'), width: 120, align: 'center' },
    { colKey: 'remark', title: $t('common.remark'), minWidth: 200 },
    { colKey: 'created_at', title: $t('common.createTime'), minWidth: 180 },
    { colKey: 'updated_at', title: $t('common.updateTime'), minWidth: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 200,
    },
  ];
}

export function createApiGroupColumnOptions(
  columns: ApiGroupTableColumn[],
): ApiGroupColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}

export function createApiGroupFormDefaultValues(): ApiGroupFormModel {
  return {
    name: '',
    remark: '',
    status: 1,
  };
}
