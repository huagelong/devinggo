export type CrontabParameterMode = 'common' | 'url';

export interface KeyValueRow {
  key: string;
  value: string;
}

export interface UrlParameterForm {
  headers: KeyValueRow[];
  method: string;
  params: KeyValueRow[];
  url: string;
}

export interface CrontabParameterFormValue {
  commonParams: KeyValueRow[];
  mode: CrontabParameterMode;
  rawJson: string;
  url: UrlParameterForm;
  useRawJson: boolean;
}

const EMPTY_ROW: KeyValueRow = { key: '', value: '' };

export function createEmptyCrontabParameterValue(): CrontabParameterFormValue {
  return {
    commonParams: [{ ...EMPTY_ROW }],
    mode: 'common',
    rawJson: '',
    url: {
      headers: [{ ...EMPTY_ROW }],
      method: 'get',
      params: [{ ...EMPTY_ROW }],
      url: '',
    },
    useRawJson: false,
  };
}

export function buildCrontabParameter(value: CrontabParameterFormValue): string {
  if (value.useRawJson) {
    return value.rawJson.trim();
  }

  if (value.mode === 'url') {
    return JSON.stringify(
      {
        url: value.url.url,
        method: value.url.method,
        header: rowsToObject(value.url.headers),
        params: rowsToObject(value.url.params),
      },
      null,
      2,
    );
  }

  return JSON.stringify(rowsToObject(value.commonParams), null, 2);
}

export function parseCrontabParameter(
  parameter: unknown,
  mode: CrontabParameterMode,
): CrontabParameterFormValue {
  const emptyValue = createEmptyCrontabParameterValue();
  emptyValue.mode = mode;
  const rawText =
    typeof parameter === 'string'
      ? parameter.trim()
      : parameter
        ? JSON.stringify(parameter, null, 2)
        : '';

  if (!rawText) {
    return emptyValue;
  }

  try {
    const parsed = JSON.parse(rawText);
    if (!isPlainObject(parsed)) {
      return {
        ...emptyValue,
        rawJson: rawText,
        useRawJson: true,
      };
    }

    if (mode === 'url') {
      return {
        ...emptyValue,
        url: {
          headers: objectToRows(parsed.header),
          method: String(parsed.method || 'get'),
          params: objectToRows(parsed.params),
          url: String(parsed.url || ''),
        },
      };
    }

    return {
      ...emptyValue,
      commonParams: objectToRows(parsed),
    };
  } catch {
    return {
      ...emptyValue,
      rawJson: rawText,
      useRawJson: true,
    };
  }
}

function rowsToObject(rows: KeyValueRow[]) {
  return rows.reduce<Record<string, string>>((result, item) => {
    const key = item.key.trim();
    if (key) {
      result[key] = item.value;
    }
    return result;
  }, {});
}

function objectToRows(value: unknown): KeyValueRow[] {
  if (!isPlainObject(value)) {
    return [{ ...EMPTY_ROW }];
  }

  const rows = Object.entries(value).map(([key, rowValue]) => ({
    key,
    value: formatRowValue(rowValue),
  }));

  return rows.length > 0 ? rows : [{ ...EMPTY_ROW }];
}

function formatRowValue(value: unknown): string {
  if (typeof value === 'string') {
    return value;
  }

  if (value === null || value === undefined) {
    return '';
  }

  if (typeof value === 'object') {
    return JSON.stringify(value);
  }

  return String(value);
}

function isPlainObject(value: unknown): value is Record<string, unknown> {
  return Object.prototype.toString.call(value) === '[object Object]';
}
