import type { AppGroupApi } from '#/api/system/app-group';
import type { OptionItem } from '#/types/common';

export type AppGroupListItem = AppGroupApi.ListItem;

export type AppGroupFormModel = AppGroupApi.SubmitPayload;

export interface AppGroupSearchFormModel {
  created_at: string[];
  name: string;
  status?: number;
}

export type AppGroupColumnOptionItem = OptionItem<string>;

export interface AppGroupTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
