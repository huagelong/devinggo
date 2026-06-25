import type { UploadApi } from '#/api/system/upload';

export type UploadListItem = UploadApi.FileInfoResponse;

export interface UploadTreeItem {
  label: string;
  value: string;
  children?: UploadTreeItem[];
}

export interface UploadSearchFormModel {
  origin_name?: string;
  mime_type?: string;
  storage_mode?: number;
  created_at?: string[];
}

export interface UploadTableColumn {
  colKey: string;
  title?: string;
  width?: number;
  align?: 'left' | 'center' | 'right';
  ellipsis?: boolean;
  fixed?: 'left' | 'right';
  minWidth?: number;
  type?: 'multiple' | 'single';
}
