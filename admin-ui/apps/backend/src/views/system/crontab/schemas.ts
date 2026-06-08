import type {
  CrontabColumnOptionItem,
  CrontabFormModel,
  CrontabSearchFormModel,
  CrontabTableColumn,
} from './model';

import { $t } from '@vben/locales';

export const crontabTypeOptions = [
  { label: $t('system.crontab.typeCommand'), value: 1 },
  { label: $t('system.crontab.typeClass'), value: 2 },
  { label: $t('system.crontab.typeUrl'), value: 3 },
  { label: $t('system.crontab.typeEval'), value: 4 },
];

export const crontabStatusOptions = [
  { label: $t('common.enable'), value: 1 },
  { label: $t('common.disable'), value: 2 },
];

export const crontabSingletonOptions = [
  { label: $t('common.yes'), value: 1 },
  { label: $t('common.no'), value: 2 },
];

export function createCrontabSearchForm(): CrontabSearchFormModel {
  return {
    created_at: [],
    name: '',
    status: undefined,
    type: undefined,
  };
}

export function createCrontabFormDefaultValues(): CrontabFormModel {
  return {
    name: '',
    parameter: '',
    remark: '',
    rule: '',
    singleton: 2,
    status: 1,
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
    { colKey: 'name', title: $t('system.crontab.name'), minWidth: 160 },
    { colKey: 'type', title: $t('system.crontab.taskType'), width: 120 },
    { colKey: 'rule', title: $t('system.crontab.rule'), width: 160 },
    { colKey: 'target', title: $t('system.crontab.target'), minWidth: 160 },
    { colKey: 'status', title: $t('common.status'), width: 100 },
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
