import type {
  AttachmentColumnOptionItem,
  AttachmentSearchFormModel,
  AttachmentTableColumn,
  AttachmentTreeItem,
} from './model';

import { $t } from '@vben/locales';

export const storageModeOptions = [
  { label: $t('system.attachment.localStorage'), value: 1 },
  { label: $t('system.attachment.aliyunOss'), value: 2 },
  { label: $t('system.attachment.tencentCos'), value: 3 },
  { label: $t('system.attachment.qiniu'), value: 4 },
  { label: 'FTP', value: 5 },
];

export const defaultAttachmentTreeData: AttachmentTreeItem[] = [
  { title: $t('system.attachment.filterAll'), key: 'all' },
  { title: $t('system.attachment.filterImage'), key: 'image' },
  { title: $t('system.attachment.filterVideo'), key: 'video' },
  { title: $t('system.attachment.filterAudio'), key: 'audio' },
  { title: $t('system.attachment.filterDocument'), key: 'document' },
  { title: $t('system.attachment.filterArchive'), key: 'archive' },
  { title: $t('system.attachment.filterOther'), key: 'other' },
];

export function createAttachmentSearchForm(): AttachmentSearchFormModel {
  return {
    created_at: [],
    mime_type: undefined,
    origin_name: '',
    storage_mode: undefined,
  };
}

export function createAttachmentTableColumns(): AttachmentTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
    { colKey: 'id', title: 'ID', width: 80 },
    {
      colKey: 'url',
      title: $t('common.preview'),
      width: 80,
    },
    { colKey: 'object_name', title: $t('system.attachment.objectName'), minWidth: 200 },
    { colKey: 'origin_name', title: $t('system.attachment.originName'), minWidth: 150 },
    { colKey: 'storage_mode', title: $t('system.attachment.storageMode'), width: 120 },
    { colKey: 'mime_type', title: $t('system.attachment.mimeType'), minWidth: 130 },
    { colKey: 'size_info', title: $t('system.attachment.fileSize'), width: 130 },
    { colKey: 'created_at', title: $t('system.attachment.uploadTime'), width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 120,
    },
  ];
}

export function createAttachmentColumnOptions(
  columns: AttachmentTableColumn[],
): AttachmentColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
