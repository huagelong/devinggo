import type { CrontabApi } from '#/api/system/crontab';
import type { OptionItem } from '#/types/common';

export type CrontabListItem = CrontabApi.ListItem;

export type CrontabFormModel = CrontabApi.SubmitPayload;

export interface CrontabSearchFormModel {
  created_at: string[];
  is_finally?: number;
  name: string;
  type?: number;
}

export type CrontabLogItem = CrontabApi.LogItem;

export interface CrontabLogQuery extends CrontabApi.LogQuery {
  created_at: string[];
}

export type CrontabColumnOptionItem = OptionItem<string>;

export interface CrontabTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
