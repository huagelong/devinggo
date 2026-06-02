import type { AttachmentApi } from '#/api/system/attachment';
import type { OptionItem } from '#/types/common';

export type AttachmentListItem = AttachmentApi.ListItem;

export interface AttachmentSearchFormModel {
  created_at: string[];
  mime_type?: string;
  origin_name: string;
  storage_mode?: number;
}

export type AttachmentColumnOptionItem = OptionItem<string>;

export interface AttachmentTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}

export interface AttachmentTreeItem {
  title: string;
  key: string;
  children?: AttachmentTreeItem[];
}
