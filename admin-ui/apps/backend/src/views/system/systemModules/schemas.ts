import type {
  SystemModulesColumnOptionItem,
  SystemModulesFormModel,
  SystemModulesSearchFormModel,
  SystemModulesTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createSystemModulesSearchForm(): SystemModulesSearchFormModel {
  return {
    created_at: [],
    id: undefined,
    installed: undefined,
    label: '',
    name: '',
    status: undefined,
  };
}

export function createSystemModulesFormDefaultValues(): SystemModulesFormModel {
  return {
    description: '',
    installed: 1,
    label: '',
    name: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createSystemModulesTableColumns(): SystemModulesTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      disabled: ({ row }: { row: { name: string } }) => row.name === 'system',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'id', title: 'ID', minWidth: 80, sorter: true },
    { colKey: 'name', title: $t('system.systemModules.name'), minWidth: 140 },
    { colKey: 'label', title: $t('system.systemModules.label'), minWidth: 140 },
    {
      colKey: 'installed',
      title: $t('system.systemModules.installed'),
      minWidth: 100,
      sorter: true,
    },
    { colKey: 'status', title: $t('common.status'), minWidth: 100, sorter: true },
    { colKey: 'created_at', title: $t('common.createTime'), minWidth: 180 },
    { colKey: 'description', title: $t('system.systemModules.description'), minWidth: 200 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 240,
    },
  ];
}

export function createSystemModulesColumnOptions(
  columns: SystemModulesTableColumn[],
): SystemModulesColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
