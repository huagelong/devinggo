import type {
  MenuColumnOptionItem,
  MenuFormModel,
  MenuSearchFormModel,
  MenuTableColumn,
  MenuTypeValue,
} from './model';

import { $t } from '@vben/locales';

export const menuTypeOptions: Array<{ label: string; value: MenuTypeValue }> = [
  { label: $t('system.menu.typeDir'), value: 'M' },
  { label: $t('system.menu.typeButton'), value: 'B' },
  { label: $t('system.menu.externalLink'), value: 'L' },
  { label: $t('system.menu.iframe'), value: 'I' },
];

export const menuHiddenOptions = [
  { label: $t('common.yes'), value: 1 },
  { label: $t('common.no'), value: 2 },
];

export const restfulOptions = [
  { label: $t('common.yes'), value: '1' },
  { label: $t('common.no'), value: '2' },
];

export const menuTypeTagMap: Record<string, { label: string; theme: 'default' | 'primary' | 'success' | 'warning' | 'danger' }> =
  {
    M: { label: $t('system.menu.typeDir'), theme: 'primary' },
    B: { label: $t('system.menu.typeButton'), theme: 'warning' },
    L: { label: $t('system.menu.externalLink'), theme: 'success' },
    I: { label: $t('system.menu.iframe'), theme: 'default' },
  };

export function createMenuSearchForm(): MenuSearchFormModel {
  return {
    code: '',
    created_at: [],
    level: '',
    name: '',
    status: undefined,
  };
}

export function createMenuFormDefaultValues(): MenuFormModel {
  return {
    code: '',
    component: '',
    icon: '',
    is_hidden: 2,
    level: '',
    name: '',
    parent_id: 0,
    redirect: '',
    remark: '',
    restful: '2',
    route: '',
    sort: 1,
    status: 1,
    type: 'M',
  };
}

export function createMenuTableColumns(): MenuTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'name', minWidth: 200, title: $t('system.menu.title') },
    { align: 'center', colKey: 'type', title: $t('system.menu.type'), width: 120 },
    { colKey: 'code', minWidth: 180, title: $t('system.menu.code') },
    { colKey: 'icon', minWidth: 120, title: $t('system.menu.icon') },
    { colKey: 'route', minWidth: 180, title: $t('system.menu.router') },
    { colKey: 'component', minWidth: 200, title: $t('system.menu.component') },
    { align: 'center', colKey: 'sort', title: $t('common.sort'), width: 120 },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { colKey: 'created_at', minWidth: 180, title: $t('common.createTime') },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 360,
    },
  ];
}

export function createMenuColumnOptions(columns: MenuTableColumn[]): MenuColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
