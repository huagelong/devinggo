import type { IdType, OptionItem } from '#/types/common';

export interface UserSearchFormModel {
  created_at: string[];
  dept_ids: IdType[];
  email: string;
  phone: string;
  post_id?: number;
  role_id?: number;
  status?: number;
  user_type?: string;
  username: string;
}

export interface UserListItem {
  avatar?: string;
  created_at?: string;
  dashboard?: string;
  dept_id?: IdType;
  dept_name?: string;
  email?: string;
  id: number;
  nickname?: string;
  phone?: string;
  post_name?: string;
  role_name?: string;
  status?: number;
  user_type?: string;
  username: string;
}

export type ColumnOptionItem = OptionItem<string>;

export interface UserTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}

export type UserActionDropdownValue =
  | 'clear_cache'
  | 'reset_password'
  | 'set_homepage';

export interface UserActionDropdownItem {
  content: string;
  value: UserActionDropdownValue;
}
