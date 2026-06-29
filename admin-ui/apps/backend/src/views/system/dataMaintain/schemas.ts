import type {
  DataMaintainColumnOptionItem,
  DataMaintainSearchFormModel,
  DataMaintainTableColumn,
} from './model';

import { $t } from '@vben/locales';

export function createDataMaintainSearchForm(): DataMaintainSearchFormModel {
  return {
    group_name: 'default',
    name: '',
  };
}

export function createDataMaintainTableColumns(): DataMaintainTableColumn[] {
  return [
    { colKey: 'name', title: $t('system.dataMaintain.tableName'), minWidth: 220 },
    { colKey: 'comment', title: $t('system.dataMaintain.tableComment'), minWidth: 220 },
    { colKey: 'engine', title: $t('system.dataMaintain.engine'), width: 140 },
    { colKey: 'collation', title: $t('system.dataMaintain.collation'), width: 160 },
    { colKey: 'rows', title: $t('system.dataMaintain.rows'), width: 120 },
    { colKey: 'create_time', title: $t('common.createTime'), minWidth: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 260,
    },
  ];
}

export function createDataMaintainColumnOptions(
  columns: DataMaintainTableColumn[],
): DataMaintainColumnOptionItem[] {
  return columns
    .filter((column) => column.title && column.colKey !== 'action')
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
