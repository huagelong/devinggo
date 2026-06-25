import type { UploadTableColumn, UploadTreeItem } from './model';

import { $t } from '@vben/locales';

export const defaultUploadTreeData: UploadTreeItem[] = [
  {
    label: $t('common.all'),
    value: 'all',
  },
  {
    label: $t('system.upload.imageCategory'),
    value: 'image',
    children: [
      { label: 'JPEG', value: 'image/jpeg' },
      { label: 'PNG', value: 'image/png' },
      { label: 'GIF', value: 'image/gif' },
      { label: 'WebP', value: 'image/webp' },
    ],
  },
  {
    label: $t('system.upload.documentCategory'),
    value: 'document',
    children: [
      { label: 'PDF', value: 'application/pdf' },
      { label: 'Word', value: 'application/msword' },
      { label: 'Excel', value: 'application/vnd.ms-excel' },
      { label: 'PowerPoint', value: 'application/vnd.ms-powerpoint' },
    ],
  },
  {
    label: $t('system.upload.otherCategory'),
    value: 'other',
  },
];

export function createUploadTableColumns(): UploadTableColumn[] {
  return [
    { colKey: 'row-select', width: 50, fixed: 'left', type: 'multiple' },
    { title: 'ID', colKey: 'id', width: 80 },
    { title: $t('system.upload.fileName'), colKey: 'origin_name', ellipsis: true, minWidth: 160 },
    { title: $t('system.attachment.mimeType'), colKey: 'mime_type', width: 150 },
    { title: $t('system.attachment.storagePath'), colKey: 'storage_path', ellipsis: true },
    { title: $t('system.attachment.fileSize'), colKey: 'size_info', width: 100 },
    { title: $t('system.upload.storageMode'), colKey: 'storage_mode', width: 100 },
    { title: $t('common.createTime'), colKey: 'created_at', width: 180 },
    { title: $t('common.action'), colKey: 'action', width: 200, align: 'center', fixed: 'right' },
  ];
}

export function createUploadColumnOptions(columns: UploadTableColumn[]) {
  return columns
    .filter((col) => col.colKey !== 'row-select' && col.colKey !== 'action')
    .map((col) => ({
      label: col.title || col.colKey,
      value: col.colKey,
    }));
}

export function createUploadSearchForm() {
  return {
    origin_name: '',
    mime_type: '',
    storage_mode: undefined as number | undefined,
    created_at: [] as string[],
  };
}
