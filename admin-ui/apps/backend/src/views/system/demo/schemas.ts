import type {
  DemoColumnOptionItem,
  DemoFormModel,
  DemoSearchFormModel,
  DemoTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createDemoSearchForm(): DemoSearchFormModel {
  return {
    name: '',
    code: '',
    status: undefined,
    birthday: '',
    created_at: [],
  };
}

export function createDemoFormDefaultValues(): DemoFormModel {
  return {
    name: '',
    code: '',
    status: 1,
    sort: 1,
    price: 0,
    cover: '',
    email: '',
    phone: '',
    birthday: '',
    remark: '',
  };
}

export function createDemoTableColumns(): DemoTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { align: 'center', colKey: 'id', title: 'ID', width: 80 },
    { align: 'center', colKey: 'name', minWidth: 120, title: $t('common.name') },
    { align: 'center', colKey: 'code', minWidth: 120, title: $t('system.demo.code') },
    { align: 'center', colKey: 'status', minWidth: 120, title: $t('common.status') },
    { align: 'center', colKey: 'sort', minWidth: 140, title: $t('common.sort') },
    { align: 'center', colKey: 'price', minWidth: 120, title: $t('system.demo.price') },
    { align: 'center', colKey: 'email', minWidth: 120, title: $t('system.demo.email') },
    { align: 'center', colKey: 'phone', minWidth: 120, title: $t('system.demo.phone') },
    { align: 'center', colKey: 'birthday', minWidth: 120, title: $t('system.demo.birthday') },
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

export function createDemoColumnOptions(
  columns: DemoTableColumn[],
): DemoColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
