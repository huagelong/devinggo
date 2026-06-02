import type { ApiGroupApi } from '#/api/system/api-group';
import type { OptionItem } from '#/types/common';

export interface ApiGroupSearchFormModel {
  name: string;
  status?: number;
  created_at: string[];
}

export type ApiGroupListItem = ApiGroupApi.ListItem;

export interface ApiGroupFormModel extends ApiGroupApi.SubmitPayload {}

export type ApiGroupColumnOptionItem = OptionItem<string>;

export interface ApiGroupTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
