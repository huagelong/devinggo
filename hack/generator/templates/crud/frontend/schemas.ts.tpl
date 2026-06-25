import type {
  <%.EntityName%>ColumnOptionItem,
  <%.EntityName%>FormModel,
  <%.EntityName%>SearchFormModel,
  <%.EntityName%>TableColumn,
} from './model';

import { $t } from '@vben/locales';

export function create<%.EntityName%>SearchForm(): <%.EntityName%>SearchFormModel {
  return {
<%range .SearchFields%>    <%.Name%>: <%.DefaultValue%>,
<%end%>  };
}

export function create<%.EntityName%>FormDefaultValues(): <%.EntityName%>FormModel {
  return {
<%range .EditableFields%>    <%.Name%>: <%.DefaultValue%>,
<%end%>  };
}

export function create<%.EntityName%>TableColumns(): <%.EntityName%>TableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
<%range .TableColumns%>    { <%if .Align%>align: '<%.Align%>', <%end%>colKey: '<%.ColKey%>', <%if .MinWidth%>minWidth: <%.MinWidth%>, <%end%>title: <%.Title%>, <%if .Width%>width: <%.Width%><%end%> },
<%end%>    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 220,
    },
  ];
}

export function create<%.EntityName%>ColumnOptions(
  columns: <%.EntityName%>TableColumn[],
): <%.EntityName%>ColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
