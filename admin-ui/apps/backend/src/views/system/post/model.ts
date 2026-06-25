import type { OptionItem } from '#/types/common';

export interface PostSearchFormModel {
  code: string;
  created_at: string[];
  name: string;
  status?: number;
}

export interface PostListItem {
  code: string;
  created_at?: string;
  id: number;
  name: string;
  remark?: string;
  sort?: number;
  status?: number;
}

export interface PostFormModel {
  code: string;
  id?: number;
  name: string;
  remark: string;
  sort: number;
  status: number;
}

export type PostColumnOptionItem = OptionItem<string>;

export interface PostTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
