import type { PostListItem } from './model';
import type { PostApi } from '#/api/system/post';

import { getPostPageList, getRecyclePostList } from '#/api/system/post';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { createPostSearchForm } from './schemas';

export function usePostCrud() {
  return useCrudPage<PostListItem, ReturnType<typeof createPostSearchForm>>({
    defaultSearchForm: createPostSearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecyclePostList(params) : getPostPageList(params),
    buildParams: (form) => {
      const params: Partial<PostApi.ListQuery> = {};
      if (form.name) params.name = form.name;
      if (form.code) params.code = form.code;
      if (form.status !== undefined) params.status = form.status;
      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
