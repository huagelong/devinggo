import file2md5 from 'file2md5';

import { requestClient } from '#/api/request';
import type { BatchIdsPayload } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace UploadApi {
  // File types
  export type StorageMode = number;

  export interface UploadFileInfo {
    url: string;
    path: string;
    name: string;
    size: number;
    mime_type: string;
    storage_mode: StorageMode;
  }

  export interface UploadImageRequest {
    image: File;
    isChunk?: boolean;
    hash?: string;
    [key: string]: string | boolean | File | undefined;
  }

  export interface UploadChunkRequest {
    file: File;
    chunkIndex: number;
    totalChunks: number;
    hash: string;
    [key: string]: string | number | File | undefined;
  }

  export interface SaveNetworkImageRequest {
    url: string;
    [key: string]: string | undefined;
  }

  export interface FileInfoRequest {
    id: number;
  }

  export interface DownloadRequest {
    id: number;
  }

  export interface ShowFileRequest {
    hash: string;
  }

  export interface UploadResponse {
    url: string;
    path: string;
    name: string;
    size: number;
    mime_type: string;
    storage_mode: StorageMode;
  }

  export interface FileInfoResponse {
    id: number;
    object_name: string;
    origin_name: string;
    storage_mode: StorageMode;
    mime_type: string;
    storage_path: string;
    size_info: string;
    url: string;
    created_at: string;
    updated_at: string;
  }

  export interface ListQuery extends Partial<PageQuery> {
    created_at?: string[];
    mime_type?: string;
    origin_name?: string;
    storage_mode?: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<FileInfoResponse>;
}

async function fileToMd5(file: File) {
  return file2md5(file);
}

export async function buildImageUploadFormData(
  file: File,
  requestData: Record<string, string | number | boolean> = {},
) {
  const hash = await fileToMd5(file);
  const formData = new FormData();

  formData.append('image', file);
  formData.append('isChunk', 'false');
  formData.append('hash', hash);

  Object.entries(requestData).forEach(([name, value]) => {
    formData.append(name, String(value));
  });

  return formData;
}

// Upload APIs
export function uploadImageApi(data: FormData) {
  return requestClient.post<UploadApi.UploadResponse>(
    '/system/uploadImage',
    data,
  );
}

export async function uploadImageFileApi(
  file: File,
  requestData: Record<string, boolean | number | string> = {},
) {
  const formData = await buildImageUploadFormData(file, requestData);
  return uploadImageApi(formData);
}

export function uploadFileApi(data: FormData) {
  return requestClient.post<UploadApi.UploadResponse>(
    '/system/uploadFile',
    data,
  );
}

export async function uploadFileFromFileApi(file: File) {
  const formData = new FormData();
  formData.append('file', file);
  return uploadFileApi(formData);
}

export function uploadChunkApi(data: FormData) {
  return requestClient.post<UploadApi.UploadResponse>(
    '/system/chunkUpload',
    data,
  );
}

export function saveNetworkImageApi(params: UploadApi.SaveNetworkImageRequest) {
  return requestClient.post<UploadApi.UploadResponse>(
    '/system/saveNetworkImage',
    params,
  );
}

// File management APIs
export function getUploadPageList(params: UploadApi.ListQuery) {
  return requestClient.get<UploadApi.ListResponse>('/system/attachment/index', {
    params,
  });
}

export function getFileInfoApi(params: UploadApi.FileInfoRequest) {
  return requestClient.get<UploadApi.FileInfoResponse>(
    '/system/getFileInfoById',
    { params },
  );
}

export function downloadFileApi(params: UploadApi.DownloadRequest) {
  return requestClient.get<Blob>('/system/downloadById', {
    params,
    responseType: 'blob',
  });
}

export function showFileApi(params: UploadApi.ShowFileRequest) {
  return requestClient.get<string>(`/system/showFile/${params.hash}`);
}

export function deleteUploadFiles(ids: number[]) {
  return requestClient.delete<void>('/system/attachment/delete', {
    data: { ids },
  });
}
