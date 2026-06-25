import type { IdType, OptionItem, StatusValue } from '#/types/common';

export interface DeptSearchFormModel {
  created_at: string[];
  leader: string;
  level: string;
  name: string;
  phone: string;
  status?: number;
}

export interface DeptTreeItem {
  children?: DeptTreeItem[];
  created_at?: string;
  id: number;
  leader?: string;
  name: string;
  parent_id?: IdType;
  phone?: string;
  sort?: number;
  status?: StatusValue;
}

export interface DeptTreeOptionItem {
  children?: DeptTreeOptionItem[];
  id: number;
  label: string;
  leader?: string;
  parent_id?: IdType;
  value?: number;
}

export interface DeptFormModel {
  id?: number;
  leader: string;
  level: string;
  name: string;
  parent_id: number;
  phone: string;
  remark: string;
  sort: number;
  status: StatusValue;
}

export interface DeptLeaderListItem {
  avatar?: string;
  email?: string;
  id: number;
  leader_add_time?: string;
  nickname?: string;
  phone?: string;
  status?: StatusValue;
  username: string;
}

export type DeptColumnOptionItem = OptionItem<string>;

export interface DeptTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
