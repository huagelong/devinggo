import type {
  DictColumnOptionItem,
  DictDataFormModel,
  DictDataSearchFormModel,
  DictTableColumn,
  DictTypeFormModel,
  DictTypeSearchFormModel,
} from './model';

import { $t } from '@vben/locales';

export function createDictTypeSearchForm(): DictTypeSearchFormModel {
  return {
    code: '',
    created_at: [],
    name: '',
    status: undefined,
  };
}

export function createDictTypeFormDefaultValues(): DictTypeFormModel {
  return {
    code: '',
    name: '',
    remark: '',
    status: 1,
  };
}

export function createDictTypeTableColumns(): DictTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'name', minWidth: 160, title: $t('system.dict.name') },
    { colKey: 'code', minWidth: 200, title: $t('system.dict.code') },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { colKey: 'remark', minWidth: 180, title: $t('common.remark') },
    { colKey: 'created_at', minWidth: 180, title: $t('common.createTime') },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 320,
    },
  ];
}

export function createDictTypeColumnOptions(columns: DictTableColumn[]): DictColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}

export function createDictDataSearchForm(): DictDataSearchFormModel {
  return {
    code: '',
    created_at: [],
    label: '',
    status: undefined,
    type_id: undefined,
    value: '',
  };
}

export function createDictDataFormDefaultValues(typeId?: number, code?: string): DictDataFormModel {
  return {
    code: code ?? '',
    label: '',
    remark: '',
    sort: 1,
    status: 1,
    type_id: typeId,
    value: '',
  };
}

export function createDictDataTableColumns(): DictTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'label', minWidth: 160, title: $t('system.dict.label') },
    { colKey: 'value', minWidth: 140, title: $t('system.dict.value') },
    { align: 'center', colKey: 'sort', title: $t('common.sort'), width: 120 },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 120 },
    { colKey: 'remark', minWidth: 180, title: $t('common.remark') },
    { colKey: 'created_at', minWidth: 180, title: $t('common.createTime') },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 240,
    },
  ];
}
