import type { NoticeApi } from '#/api/system/notice';
import type { OptionItem } from '#/types/common';

export interface NoticeSearchFormModel {
  created_at: string[];
  title: string;
  type?: number | string;
}

export type NoticeListItem = NoticeApi.ListItem;

export interface NoticeFormModel extends NoticeApi.SubmitPayload {}

export type NoticeColumnOptionItem = OptionItem<string>;

export interface NoticeTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
  fixed?: 'left' | 'right';
}
