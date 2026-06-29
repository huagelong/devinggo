import type { DictApi } from '#/api/system/dict';
import type { OptionItem } from '#/types/common';

export type DictPanelTab = 'data' | 'type';

export interface DictTypeSearchFormModel {
  code: string;
  created_at: string[];
  name: string;
  status?: number;
}

export type DictTypeListItem = DictApi.DictTypeItem;

export type DictTypeFormModel = DictApi.DictTypeSubmitPayload;

export interface DictDataSearchFormModel {
  code: string;
  created_at: string[];
  label: string;
  status?: number;
  type_id?: number;
  value: string;
}

export type DictDataListItem = DictApi.DictDataItem;

export type DictDataFormModel = DictApi.DictDataSubmitPayload;

export interface DictOptionItem {
  code?: string;
  id?: number;
  key: string | number;
  title: string;
}

export type DictColumnOptionItem = OptionItem<string>;

export interface DictTableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
