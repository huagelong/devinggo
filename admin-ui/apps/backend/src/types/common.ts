export type IdType = number | string;

export type StatusValue = number | string;

export interface OptionItem<TValue extends IdType = IdType> {
  disabled?: boolean;
  label: string;
  value: TValue;
}

export interface TreeOptionItem<TValue extends IdType = IdType> {
  children?: TreeOptionItem<TValue>[];
  disabled?: boolean;
  label: string;
  value: TValue;
}

export interface BatchIdsPayload<TValue extends IdType = number> {
  ids: TValue[];
}
