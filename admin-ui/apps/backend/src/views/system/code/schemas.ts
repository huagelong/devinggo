import type {
  CodeColumnOptionItem,
  CodeSearchFormModel,
  CodeTableColumn,
} from './model';

import { $t } from '@vben/locales';

export const generateTypeOptions = [
  { label: $t('system.code.singleCrud'), value: 'single' },
  { label: $t('system.code.treeCrud'), value: 'tree' },
];

export const componentTypeOptions = [
  { label: $t('system.code.componentModal'), value: 1 },
  { label: $t('system.code.componentDrawer'), value: 2 },
  { label: $t('system.code.tagConfig'), value: 3 },
];

export const tplTypeOptions = [
  { label: 'default', value: 'default' },
  { label: 'ruoyi', value: 'ruoyi' },
];

export const queryTypeOptions = [
  { label: '=', value: 'eq' },
  { label: '!=', value: 'neq' },
  { label: '>', value: 'gt' },
  { label: '>=', value: 'gte' },
  { label: '<', value: 'lt' },
  { label: '<=', value: 'lte' },
  { label: 'LIKE', value: 'like' },
  { label: 'IN', value: 'in' },
  { label: 'NOT IN', value: 'notin' },
  { label: 'BETWEEN', value: 'between' },
];

export const viewTypeOptions = [
  { label: $t('system.code.viewTypes.text'), value: 'text' },
  { label: $t('system.code.viewTypes.password'), value: 'password' },
  { label: $t('system.code.viewTypes.textarea'), value: 'textarea' },
  { label: $t('system.code.viewTypes.inputNumber'), value: 'inputNumber' },
  { label: $t('system.code.viewTypes.switch'), value: 'switch' },
  { label: $t('system.code.viewTypes.slider'), value: 'slider' },
  { label: $t('system.code.viewTypes.select'), value: 'select' },
  { label: $t('system.code.viewTypes.treeSelect'), value: 'treeSelect' },
  { label: $t('system.code.viewTypes.radio'), value: 'radio' },
  { label: $t('system.code.viewTypes.checkbox'), value: 'checkbox' },
  { label: $t('system.code.viewTypes.date'), value: 'date' },
  { label: $t('system.code.viewTypes.time'), value: 'time' },
  { label: $t('system.code.viewTypes.rate'), value: 'rate' },
  { label: $t('system.code.viewTypes.cascader'), value: 'cascader' },
  { label: $t('system.code.viewTypes.transfer'), value: 'transfer' },
  { label: $t('system.code.viewTypes.selectUser'), value: 'selectUser' },
  { label: $t('system.code.viewTypes.cityLinkage'), value: 'cityLinkage' },
  { label: $t('system.code.viewTypes.upload'), value: 'upload' },
  { label: $t('system.code.viewTypes.editor'), value: 'editor' },
  { label: $t('system.code.viewTypes.codeEditor'), value: 'codeEditor' },
];

export const menuButtonOptions = [
  { label: $t('system.code.buttons.save'), value: 'save' },
  { label: $t('system.code.buttons.update'), value: 'update' },
  { label: $t('system.code.buttons.read'), value: 'read' },
  { label: $t('system.code.buttons.delete'), value: 'delete' },
  { label: $t('system.code.buttons.recycle'), value: 'recycle' },
  { label: $t('system.code.buttons.changeStatus'), value: 'changeStatus' },
  { label: $t('system.code.buttons.numberOperation'), value: 'numberOperation' },
  { label: $t('system.code.buttons.import'), value: 'import' },
  { label: $t('system.code.buttons.export'), value: 'export' },
];

export function createCodeSearchForm(): CodeSearchFormModel {
  return {
    table_name: '',
    type: undefined,
  };
}

export function createCodeTableColumns(): CodeTableColumn[] {
  return [
    { colKey: 'row-select', title: '', width: 52, fixed: 'left' },
    { colKey: 'id', title: 'ID', width: 80 },
    { colKey: 'table_name', title: $t('system.code.tableName'), minWidth: 200 },
    { colKey: 'table_comment', title: $t('system.code.tableComment'), minWidth: 200 },
    { colKey: 'type', title: $t('system.code.genType'), width: 120 },
    { colKey: 'module_name', title: $t('system.code.moduleName'), width: 150 },
    { colKey: 'menu_name', title: $t('system.code.menuName'), width: 150 },
    { colKey: 'created_at', title: $t('common.createTime'), width: 180 },
    { colKey: 'action', title: $t('common.action'), width: 320, fixed: 'right' },
  ];
}

export function createCodeColumnOptions(
  columns: CodeTableColumn[],
): CodeColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
