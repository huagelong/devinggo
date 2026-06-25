import type { DataMaintainApi } from '#/api/system/data-maintain';
import type { OptionItem } from '#/types/common';

export type DataMaintainListItem = DataMaintainApi.ListItem;

export interface DataMaintainSearchFormModel {
  group_name: string;
  name: string;
}

export type DataMaintainColumnOptionItem = OptionItem<string>;

export interface DataMaintainTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
