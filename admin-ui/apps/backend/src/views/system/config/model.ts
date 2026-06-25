import type { ConfigApi } from '#/api/system/config';
import type { OptionItem } from '#/types/common';

import { $t } from '@vben/locales';

export interface ConfigGroup extends ConfigApi.ConfigGroupItem {}

export type ConfigListItem = ConfigApi.ConfigItem;

export interface ConfigFieldMeta {
  config_select_data?: ConfigApi.ConfigSelectOption[];
  id: number;
  input_type: ConfigApi.InputType;
  key: string;
  label: string;
  remark?: string;
  sort?: number;
  switchValues?: {
    checked: string | number | boolean;
    unchecked: string | number | boolean;
  };
}

export type ConfigKeyValueItem = { key: string; value: string };

export interface ConfigFormModel {
  [key: string]:
    | string
    | number
    | boolean
    | string[]
    | Record<string, unknown>
    | ConfigKeyValueItem[]
    | undefined;
}

export type ConfigGroupOption = OptionItem<number>;

export const inputComponentOptions: ConfigApi.ConfigSelectOption[] = [
  { label: $t('system.config.inputTypeInput'), value: 'input' },
  { label: $t('system.config.inputTypeTextarea'), value: 'textarea' },
  { label: $t('system.config.inputTypeSelect'), value: 'select' },
  { label: $t('system.config.inputTypeRadio'), value: 'radio' },
  { label: $t('system.config.inputTypeCheckbox'), value: 'checkbox' },
  { label: $t('system.config.inputTypeSwitch'), value: 'switch' },
  { label: $t('system.config.inputTypeUpload'), value: 'upload' },
  { label: $t('system.config.inputTypeKeyValue'), value: 'key-value' },
  { label: $t('system.config.inputTypeEditor'), value: 'editor' },
];
