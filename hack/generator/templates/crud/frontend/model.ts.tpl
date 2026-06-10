import type { OptionItem } from '#/types/common';

export interface <%.EntityName%>SearchFormModel {
<%range .SearchFields%>  <%.Name%>: <%.DefaultValue%>;
<%end%>}

export interface <%.EntityName%>ListItem {
<%range .ListFields%>  <%.Name%>?: <%.Type%>;
<%end%>  id: number;
}

export interface <%.EntityName%>FormModel {
<%range .EditableFields%>  <%.Name%>: <%.Type%>;
<%end%>  id?: number;
}

export type <%.EntityName%>ColumnOptionItem = OptionItem<string>;

export interface <%.EntityName%>TableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
