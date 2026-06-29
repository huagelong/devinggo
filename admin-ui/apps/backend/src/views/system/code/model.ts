import type { GenerateApi } from '#/api/system/generate';
import type { OptionItem } from '#/types/common';

export type CodeListItem = GenerateApi.ListItem;

export interface CodeSearchFormModel {
  table_name: string;
  type?: string;
}

export type CodeColumnOptionItem = OptionItem<string>;

export interface CodeTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  width?: number;
}

// 字段配置行
export interface FieldConfigRow extends GenerateApi.FieldConfig {
  _checked?: boolean;
}

// 预览代码项
export interface PreviewCodeRow {
  name: string;
  tab_name: string;
  code: string;
  lang: string;
}

// 装载数据表行
export interface LoadTableRow {
  name: string;
  comment: string;
  sourceName: string;
  selected?: boolean;
}
