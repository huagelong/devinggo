import type { ConfigApi } from '#/api/system/config';

export interface ConfigFormValues {
  config_select_data?: string;
  group_id?: number;
  id?: number;
  input_type: ConfigApi.InputType;
  key: string;
  name: string;
  remark: string;
  sort: number;
  status: number;
  value: string;
}

export interface ConfigGroupFormValues {
  code: string;
  id?: number;
  name: string;
  remark: string;
}

export function createConfigFormDefaultValues(): ConfigFormValues {
  return {
    config_select_data: '',
    group_id: undefined,
    input_type: 'input',
    key: '',
    name: '',
    remark: '',
    sort: 0,
    status: 1,
    value: '',
  };
}

export function createConfigGroupFormDefaultValues(): ConfigGroupFormValues {
  return {
    code: '',
    name: '',
    remark: '',
  };
}
