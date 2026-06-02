import type { SystemModulesApi } from '#/api/system/system-modules';
import type { OptionItem } from '#/types/common';

export type SystemModulesListItem = SystemModulesApi.ListItem;

export type SystemModulesFormModel = SystemModulesApi.SubmitPayload;

export interface SystemModulesSearchFormModel {
  created_at: string[];
  name: string;
  status?: number;
}

export type SystemModulesColumnOptionItem = OptionItem<string>;

export interface SystemModulesTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
