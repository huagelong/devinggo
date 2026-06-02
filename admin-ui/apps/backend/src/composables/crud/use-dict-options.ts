import type { IdType, OptionItem } from '#/types/common';

import { getDictList } from '#/api/system/dict';
import { logger } from '#/utils/logger';

export interface DictItem<TKey extends IdType = IdType> {
  [key: string]: unknown;
  key: TKey;
  title: string;
}

export type DictOption<TValue extends IdType = IdType> = OptionItem<TValue>;

interface DictOptionConfig<TItem extends DictItem = DictItem> {
  labelKey?: Extract<keyof TItem, string>;
  valueKey?: Extract<keyof TItem, string>;
}

const DEFAULT_CACHE_TTL = 5 * 60 * 1000;

interface CacheEntry<T> {
  data: T;
  timestamp: number;
}

const dictCache = new Map<string, CacheEntry<DictItem[]>>();
const dictPromiseCache = new Map<string, Promise<DictItem[]>>();

function normalizeOptionValue(value: unknown): IdType {
  if (typeof value === 'number' || typeof value === 'string') {
    return value;
  }
  return String(value ?? '');
}

function toOptions<TItem extends DictItem>(
  list: TItem[],
  config?: DictOptionConfig<TItem>,
): DictOption[] {
  const labelKey = config?.labelKey ?? 'title';
  const valueKey = config?.valueKey ?? 'key';

  return (list || []).map((item) => ({
    label: String(item?.[labelKey] ?? ''),
    value: normalizeOptionValue(item?.[valueKey]),
  }));
}

async function fetchDictList(code: string, forceRefresh = false): Promise<DictItem[]> {
  if (!forceRefresh) {
    const cached = dictCache.get(code);
    if (cached && Date.now() - cached.timestamp < DEFAULT_CACHE_TTL) {
      return cached.data;
    }
    dictCache.delete(code);
  }

  if (!forceRefresh && dictPromiseCache.has(code)) {
    return dictPromiseCache.get(code)!;
  }

  const requestPromise = getDictList(code)
    .then((response) => {
      const list = Array.isArray(response) ? response : [];
      dictCache.set(code, { data: list, timestamp: Date.now() });
      return list;
    })
    .catch((error) => {
      logger.error(`Failed to load dict list for code "${code}":`, error);
      return [] as DictItem[];
    })
    .finally(() => {
      dictPromiseCache.delete(code);
    });

  dictPromiseCache.set(code, requestPromise);
  return requestPromise;
}

export function useDictOptions() {
  async function getDictOptions(
    code: string,
    config?: DictOptionConfig & { forceRefresh?: boolean },
  ) {
    const list = await fetchDictList(code, config?.forceRefresh);
    return toOptions(list, config);
  }

  async function getMultipleDictOptions(
    codes: string[],
    config?: DictOptionConfig & { forceRefresh?: boolean },
  ) {
    const entries = await Promise.all(
      codes.map(async (code) => [code, await getDictOptions(code, config)] as const),
    );
    return Object.fromEntries(entries) as Record<string, DictOption[]>;
  }

  function clearDictCache(code?: string) {
    if (code) {
      dictCache.delete(code);
      dictPromiseCache.delete(code);
      return;
    }

    dictCache.clear();
    dictPromiseCache.clear();
  }

  return {
    clearDictCache,
    getDictOptions,
    getMultipleDictOptions,
  };
}
