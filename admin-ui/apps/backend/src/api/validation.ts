export function validateResponse<T extends Record<string, unknown>>(
  data: unknown,
  requiredFields: (keyof T)[],
): data is T {
  if (!data || typeof data !== 'object') {
    return false;
  }

  const dataObj = data as Record<string, unknown>;

  return requiredFields.every((field) => {
    const fieldStr = String(field);
    return fieldStr in dataObj && dataObj[fieldStr] !== undefined && dataObj[fieldStr] !== null;
  });
}

export function validatePageResponse<T>(data: unknown): data is {
  items?: T[];
  pageInfo?: { total: number };
  total?: number;
} {
  if (!data || typeof data !== 'object') {
    return false;
  }

  const dataObj = data as Record<string, unknown>;

  const hasItems = 'items' in dataObj && (Array.isArray(dataObj.items) || dataObj.items === undefined);
  const hasTotal = 'total' in dataObj && (typeof dataObj.total === 'number' || dataObj.total === undefined);
  const hasPageInfo = 'pageInfo' in dataObj && (typeof dataObj.pageInfo === 'object' || dataObj.pageInfo === undefined);

  return hasItems && hasTotal && hasPageInfo;
}

export function validateApiResponse(response: unknown): response is {
  code: number;
  data: unknown;
  message: string;
} {
  if (!response || typeof response !== 'object') {
    return false;
  }

  const responseObj = response as Record<string, unknown>;

  return (
    typeof responseObj.code === 'number' &&
    'data' in responseObj &&
    typeof responseObj.message === 'string'
  );
}

export function validateUserListItem(item: unknown): item is {
  id: number;
  username: string;
} {
  return validateResponse(item, ['id', 'username']);
}

export function validateRoleListItem(item: unknown): item is {
  id: number;
  name: string;
  code: string;
} {
  return validateResponse(item, ['id', 'name', 'code']);
}

export function validateUserDetail(data: unknown): data is {
  username: string;
} {
  return validateResponse(data, ['username']);
}

export function validateRoleDetail(data: unknown): data is {
  name: string;
  code: string;
} {
  return validateResponse(data, ['name', 'code']);
}

export function logValidationWarning(context: string, data: unknown): void {
  console.warn(`[API Validation Warning] ${context}:`, data);
}
