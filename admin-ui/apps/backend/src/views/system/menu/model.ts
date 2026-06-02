import type { IdType, OptionItem } from '#/types/common';

export type MenuTypeValue = 'B' | 'I' | 'L' | 'M' | string;

export interface MenuSearchFormModel {
  code: string;
  created_at: string[];
  level: string;
  name: string;
  status?: number;
}

export interface MenuTreeItem {
  children?: MenuTreeItem[];
  code?: string;
  component?: string;
  created_at?: string;
  icon?: string;
  id: number;
  is_hidden?: number;
  level?: string;
  name: string;
  parent_id?: IdType;
  redirect?: string;
  remark?: string;
  route?: string;
  sort?: number;
  status?: number;
  type?: MenuTypeValue;
}

export interface MenuFormModel {
  code: string;
  component: string;
  icon: string;
  id?: number;
  is_hidden: number;
  level: string;
  name: string;
  parent_id: number;
  redirect: string;
  remark: string;
  restful: string;
  route: string;
  sort: number;
  status: number;
  type: MenuTypeValue;
}

export interface MenuTreeOptionItem {
  children?: MenuTreeOptionItem[];
  id: number;
  label: string;
  parent_id?: IdType;
  value: number;
}

export type MenuColumnOptionItem = OptionItem<string>;

export interface MenuTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
