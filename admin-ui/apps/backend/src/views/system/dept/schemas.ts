import type {
  DeptColumnOptionItem,
  DeptFormModel,
  DeptSearchFormModel,
  DeptTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createDeptSearchForm(): DeptSearchFormModel {
  return {
    created_at: [],
    leader: '',
    level: '',
    name: '',
    phone: '',
    status: undefined,
  };
}

export function createDeptFormDefaultValues(): DeptFormModel {
  return {
    leader: '',
    level: '',
    name: '',
    parent_id: 0,
    phone: '',
    remark: '',
    sort: 1,
    status: 1,
  };
}

export function createDeptTableColumns(): DeptTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'name', minWidth: 180, title: $t('system.dept.name') },
    { align: 'center', colKey: 'leader', minWidth: 120, title: $t('system.dept.leader') },
    { align: 'center', colKey: 'phone', minWidth: 150, title: $t('system.dept.phone') },
    { align: 'center', colKey: 'sort', title: $t('common.sort'), width: 140 },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { align: 'center', colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 320,
    },
  ];
}

export function createDeptColumnOptions(
  columns: DeptTableColumn[],
): DeptColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
