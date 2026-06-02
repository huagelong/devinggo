import type { DemoApi } from '#/api/system/demo';
import type { OptionItem } from '#/types/common';

export type DemoListItem = DemoApi.ListItem;

export interface DemoSearchFormModel {
  name: string;
  code: string;
  status?: number;
  birthday: string;
  created_at: string[];
}

export interface DemoFormModel {
  name: string;
  code: string;
  status: number;
  sort: number;
  price: number;
  cover: string;
  email: string;
  phone: string;
  birthday: string;
  remark: string;
}

export type DemoColumnOptionItem = OptionItem<string>;

export interface DemoTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
