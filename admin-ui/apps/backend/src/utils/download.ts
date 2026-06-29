import { downloadFileFromBlob } from '@vben/utils';

interface HeaderLike {
  get?: (name: string) => null | string;
  [key: string]: unknown;
}

interface DownloadResponseLike {
  data: Blob;
  headers?: HeaderLike | Record<string, string | undefined>;
}

function getHeaderValue(
  headers: DownloadResponseLike['headers'],
  name: string,
): string {
  if (!headers) {
    return '';
  }

  if (typeof headers.get === 'function') {
    return headers.get(name) ?? '';
  }

  const value = headers[name];
  return typeof value === 'string' ? value : '';
}

export function getFileNameFromDisposition(disposition?: string): string {
  if (!disposition) return '';

  const utf8Match = disposition.match(/filename\*=UTF-8''([^;]+)/i);
  if (utf8Match?.[1]) {
    return decodeURIComponent(utf8Match[1]);
  }

  const asciiMatch = disposition.match(/filename="?([^"]+)"?/i);
  return asciiMatch?.[1] ?? '';
}

export function downloadResponseBlob(
  response: DownloadResponseLike,
  fallbackFileName: string,
) {
  const disposition = getHeaderValue(response.headers, 'content-disposition');
  const fileName = getFileNameFromDisposition(disposition) || fallbackFileName;

  downloadFileFromBlob({
    fileName,
    source: response.data,
  });
}
