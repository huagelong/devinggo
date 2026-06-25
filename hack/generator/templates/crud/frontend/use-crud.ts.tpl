import type { <%.EntityName%>ListItem } from './model';
import type { <%.EntityName%>Api } from '#/api/<%.ModuleName%>/<%.VarName%>';

import { get<%.EntityName%>PageList, getRecycle<%.EntityName%>List } from '#/api/<%.ModuleName%>/<%.VarName%>';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { create<%.EntityName%>SearchForm } from './schemas';

export function use<%.EntityName%>Crud() {
  return useCrudPage<<%.EntityName%>ListItem, ReturnType<typeof create<%.EntityName%>SearchForm>>({
    defaultSearchForm: create<%.EntityName%>SearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycle<%.EntityName%>List(params) : get<%.EntityName%>PageList(params),
    buildParams: (form) => {
      const params: Partial<<%.EntityName%>Api.ListQuery> = {};
<%range .SearchFields%>      if (form.<%.Name%> !== undefined && form.<%.Name%> !== '') params.<%.Name%> = form.<%.Name%>;
<%end%>      if (form.created_at?.length === 2 && form.created_at[0]) {
        params.created_at = form.created_at;
      }
      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
