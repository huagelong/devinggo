import type {
  RoleColumnOptionItem,
  RoleFormModel,
  RoleSearchFormModel,
  RoleTableColumn,
} from './model';

import { $t } from '#/locales';

export const roleDataScopeOptions = [
  { label: $t('system.role.allDataScope'), value: 1 },
  { label: $t('system.role.customDataScope'), value: 2 },
  { label: $t('system.role.deptDataScope'), value: 3 },
  { label: $t('system.role.deptAndBelowDataScope'), value: 4 },
  { label: $t('system.role.selfDataScope'), value: 5 },
  { label: $t('system.role.deptFilterScope'), value: 6 },
];

export function createRoleSearchForm(): RoleSearchFormModel {
  return {
    code: '',
    created_at: [],
    name: '',
    status: undefined,
  };
}

export function createRoleFormDefaultValues(): RoleFormModel {
  return {
    code: '',
    data_scope: 1,
    dept_ids: [],
    menu_ids: [],
    name: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createRoleTableColumns(): RoleTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { align: 'center', colKey: 'id', title: 'ID', width: 80 },
    { align: 'center', colKey: 'name', minWidth: 140, title: $t('system.role.name') },
    { align: 'center', colKey: 'code', minWidth: 160, title: $t('system.role.code') },
    { align: 'center', colKey: 'data_scope', minWidth: 160, title: $t('system.role.dataScope') },
    { align: 'center', colKey: 'sort', title: $t('common.sort'), width: 140 },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { align: 'center', colKey: 'remark', minWidth: 180, title: $t('common.remark') },
    { align: 'center', colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 420,
    },
  ];
}

export function createRoleColumnOptions(
  columns: RoleTableColumn[],
): RoleColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}

export function getRoleDataScopeLabel(value?: number | string) {
  const normalizedValue = Number(value);
  return (
    roleDataScopeOptions.find((item) => Number(item.value) === normalizedValue)
      ?.label ?? '-'
  );
}
