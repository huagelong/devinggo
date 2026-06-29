import type { OptionItem, StatusValue } from '#/types/common';

export type RoleDataScopeValue = 1 | 2 | 3 | 4 | 5 | number;

export interface RoleSearchFormModel {
  code: string;
  created_at: string[];
  name: string;
  status?: number;
}

export interface RoleListItem {
  code: string;
  created_at?: string;
  data_scope?: RoleDataScopeValue;
  id: number;
  name: string;
  remark?: string;
  sort?: number;
  status?: StatusValue;
}

export interface RoleFormModel {
  code: string;
  data_scope: RoleDataScopeValue;
  dept_ids: number[];
  id?: number;
  menu_ids: number[];
  name: string;
  remark: string;
  sort: number;
  status: number;
}

export interface RolePermissionRelationItem {
  id: number;
  pivot: {
    dept_id?: number;
    menu_id?: number;
    role_id: number;
  };
}

export interface RolePermissionSelection {
  depts?: RolePermissionRelationItem[];
  id: number;
  menus?: RolePermissionRelationItem[];
}

export type RoleColumnOptionItem = OptionItem<string>;

export interface RoleTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
