import type {
  CrontabColumnOptionItem,
  CrontabFormModel,
  CrontabSearchFormModel,
  CrontabTableColumn,
} from './model';

import { $t } from '@vben/locales';

export const crontabTypeOptions = [
  { label: $t('system.crontab.typeInterval'), value: 1 },
  { label: $t('system.crontab.typeCron'), value: 2 },
];

export const crontabFinallyOptions = [
  { label: $t('common.yes'), value: 1 },
  { label: $t('common.no'), value: 2 },
];

export function createCrontabSearchForm(): CrontabSearchFormModel {
  return {
    created_at: [],
    is_finally: undefined,
    name: '',
    type: undefined,
  };
}

export function createCrontabFormDefaultValues(): CrontabFormModel {
  return {
    is_finally: 2,
    name: '',
    remark: '',
    rule: '',
    target: '',
    type: 1,
  };
}

export function createCrontabTableColumns(): CrontabTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'id', title: 'ID', width: 80 },
    { colKey: 'name', title: $t('system.crontab.name'), minWidth: 160 },
    { colKey: 'type', title: $t('system.crontab.taskType'), width: 120 },
    { colKey: 'rule', title: $t('system.crontab.rule'), minWidth: 180 },
    { colKey: 'is_finally', title: $t('system.crontab.isFinally'), width: 100 },
    { colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 360,
    },
  ];
}

export function createCrontabColumnOptions(
  columns: CrontabTableColumn[],
): CrontabColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
