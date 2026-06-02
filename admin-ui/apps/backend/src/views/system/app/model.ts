import type { AppApi } from '#/api/system/app';
import type { OptionItem } from '#/types/common';

export type AppListItem = AppApi.ListItem;

export type AppFormModel = AppApi.SubmitPayload;

export interface AppSearchFormModel {
  created_at: string[];
  name: string;
  status?: number;
}

export type AppColumnOptionItem = OptionItem<string>;

export interface AppTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
