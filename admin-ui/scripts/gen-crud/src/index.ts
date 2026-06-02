#!/usr/bin/env node

/**
 * DevingGo CRUD Page Generator
 *
 * Usage:
 *   node scripts/gen-crud/src/index.ts --name=user --module=system
 *   node scripts/gen-crud/src/index.ts --name=category --module=system --fields=id,name,code,status,sort,remark,created_at
 *   node scripts/gen-crud/src/index.ts --name=product --module=system --cn-name=产品 --fields=id,name,code,status,sort,price,remark
 *
 * Generates a complete CRUD page structure:
 *   - api/{module}/{name}.ts
 *   - views/{module}/{name}/index.vue
 *   - views/{module}/{name}/model.ts
 *   - views/{module}/{name}/schemas.ts
 *   - views/{module}/{name}/components/{name}-modal.vue
 *   - views/{module}/{name}/use-{name}-crud.ts
 */

import cac from 'cac';
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

interface Field {
  name: string;
  type: string;
  label: string;
  required?: boolean;
  options?: { label: string; value: string | number }[];
  multiple?: boolean;
  dictType?: string;
}

interface CliOptions {
  name: string;
  module: string;
  fields?: string;
  cnName?: string;
  permission?: string;
  tableName?: string;
}

interface TemplateData extends CliOptions {
  upperName: string;
  camelName: string;
  pascalName: string;
  apiPath: string;
  tableNameResolved: string;
  permissionResolved: string;
  cnNameResolved: string;
  fields: Field[];
  searchFields: Field[];
  formFields: Field[];
  hasSort: boolean;
  hasStatus: boolean;
}

const cli = cac('gen-crud');

/**
 * Supported field types with their display labels and default configurations
 */
const fieldTypes: Record<string, { label: string; component: string; defaultOptions?: { label: string; value: string | number }[] }> = {
  // Text inputs
  'string': { label: '文本框' },
  'text': { label: '文本' },
  'input': { label: '输入框' },

  // Number inputs
  'number': { label: '数字' },
  'inputNumber': { label: '数字输入' },
  'integer': { label: '整数' },

  // Textarea
  'textarea': { label: '文本域' },

  // Select/Dropdown
  'select': {
    label: '下拉选择',
    defaultOptions: [
      { label: '正常', value: 1 },
      { label: '停用', value: 2 },
    ],
  },

  // Radio
  'radio': {
    label: '单选按钮',
    defaultOptions: [
      { label: '是', value: 1 },
      { label: '否', value: 0 },
    ],
  },

  // Checkbox
  'checkbox': { label: '多选框' },

  // Switch
  'switch': {
    label: '开关',
    defaultOptions: [
      { label: '开', value: 1 },
      { label: '关', value: 0 },
    ],
  },

  // Date/Time
  'date': { label: '日期' },
  'datePicker': { label: '日期选择' },
  'dateRange': { label: '日期范围' },
  'dateTime': { label: '日期时间' },
  'dateTimeRange': { label: '日期时间范围' },
  'time': { label: '时间' },
  'timePicker': { label: '时间选择' },

  // Tree Select
  'treeSelect': { label: '树形选择' },
  'tree': { label: '树形选择' },

  // Upload
  'upload': { label: '上传' },
  'image': { label: '图片' },
  'images': { label: '图片组' },
  'file': { label: '文件' },
  'files': { label: '文件组' },
  'avatar': { label: '头像' },

  // Password
  'password': { label: '密码' },

  // Email
  'email': { label: '邮箱' },

  // Phone/Mobile
  'phone': { label: '手机号' },
  'mobile': { label: '手机号' },

  // URL
  'url': { label: '网址' },

  // Color
  'color': { label: '颜色' },

  // Slider
  'slider': { label: '滑块' },

  // Rate
  'rate': { label: '评分' },

  // Cascader
  'cascader': { label: '级联选择' },

  // AutoComplete
  'autocomplete': { label: '自动完成' },

  // InputGroup (组合输入)
  'inputGroup': { label: '组合输入' },

  // Divider
  'divider': { label: '分割线' },
};

// Default fields if not specified
const defaultFields: Field[] = [
  { name: 'id', type: 'number', label: 'ID' },
  { name: 'name', type: 'string', label: '名称', required: true },
  { name: 'code', type: 'string', label: '编码' },
  { name: 'sort', type: 'number', label: '排序' },
  { name: 'status', type: 'select', label: '状态', options: [{ label: '正常', value: 1 }, { label: '停用', value: 2 }] },
  { name: 'remark', type: 'textarea', label: '备注' },
  { name: 'created_at', type: 'dateRange', label: '创建时间' },
];

cli
  .option('--name <name>', 'Business name (e.g., user, post, category)', { demandOption: true })
  .option('--module <module>', 'Module name (e.g., system)', { demandOption: true })
  .option('--fields <fields>', 'Comma-separated field list (default: name,code,status,sort,remark,created_at)')
  .option('--cn-name <cnName>', 'Chinese name for display (e.g., 用户)')
  .option('--permission <permission>', 'Permission prefix (e.g., system:user)')
  .option('--table-name <tableName>', 'Database table name (e.g., system_user)');

cli.help();
cli.parse();

function toCamelCase(str: string): string {
  return str.replace(/-(\w)/g, (_, c) => (c ? c.toUpperCase() : ''));
}

function toPascalCase(str: string): string {
  const camel = toCamelCase(str);
  return camel.charAt(0).toUpperCase() + camel.slice(1);
}

function toSnakeCase(str: string): string {
  return str.replace(/[A-Z]/g, (c) => `_${c.toLowerCase()}`).replace(/^_/, '');
}

function parseFields(fieldsStr?: string): Field[] {
  if (!fieldsStr) {
    return defaultFields;
  }

  // Field names that should auto-map to specific types if no type specified
  const nameToTypeMap: Record<string, string> = {
    phone: 'phone',
    mobile: 'mobile',
    email: 'email',
    password: 'password',
    url: 'url',
    avatar: 'avatar',
    cover: 'image',
    image: 'image',
    images: 'images',
    file: 'file',
    files: 'files',
    content: 'textarea',
    description: 'textarea',
    remark: 'textarea',
  };

  return fieldsStr.split(',').map((f) => {
    // Format: name:type:option1;option2 or name:type:dict:dictType
    // Option format: label=value (e.g., normal=1;disabled=2)
    const parts = f.trim().split(':');
    const name = parts[0];
    let type = parts[1];

    // Auto-map type from name if not explicitly specified
    if (!type && nameToTypeMap[name]) {
      type = nameToTypeMap[name];
    }
    type = type || 'string';

    let options: { label: string; value: string | number }[] | undefined;
    let dictType: string | undefined;

    if (parts[2]) {
      if (parts[2].startsWith('dict:')) {
        dictType = parts[2].substring(5);
      } else {
        options = parts[2].split(';').map((opt) => {
          // Format: label=value (e.g., normal=1)
          const eqIndex = opt.indexOf('=');
          if (eqIndex === -1) {
            return { label: opt, value: opt };
          }
          const label = opt.substring(0, eqIndex);
          const valueStr = opt.substring(eqIndex + 1);
          const value = isNaN(Number(valueStr)) ? valueStr : Number(valueStr);
          return { label, value };
        });
      }
    }

    const defaultField = defaultFields.find((df) => df.name === name);
    const fieldType = fieldTypes[type];

    return {
      name,
      type,
      // Priority: explicit options label > defaultField label > fieldType label > name
      label: defaultField?.label || fieldType?.label || name.charAt(0).toUpperCase() + name.slice(1).replace(/([A-Z])/g, ' $1'),
      required: name === 'id' ? false : ['name', 'code'].includes(name),
      options: options || defaultField?.options || fieldType?.defaultOptions,
      dictType,
    };
  });
}

function getFieldComponentConfig(field: Field): { component: string; props: Record<string, unknown>; rules?: string } {
  const { type, name, options, required } = field;
  const rules = required ? 'required' : undefined;

  switch (type) {
    case 'number':
    case 'integer':
      return {
        component: 'InputNumber',
        props: { min: 0, max: 1000, placeholder: '请输入' },
        rules,
      };

    case 'inputNumber':
      return {
        component: 'InputNumber',
        props: { min: 0, placeholder: '请输入' },
        rules,
      };

    case 'textarea':
      return {
        component: 'Textarea',
        props: { placeholder: '请输入', autosize: { minRows: 3, maxRows: 6 } },
        rules,
      };

    case 'select':
      return {
        component: 'Select',
        props: {
          options: options || [],
          placeholder: '请选择',
          clearable: true,
          ...(field.multiple ? { multiple: true } : {}),
        },
        rules,
      };

    case 'radio':
      return {
        component: 'RadioGroup',
        props: {
          options: options || [],
        },
        rules,
      };

    case 'checkbox':
      return {
        component: 'CheckboxGroup',
        props: {
          options: options || [],
        },
        rules,
      };

    case 'switch':
      return {
        component: 'Switch',
        props: {},
        rules,
      };

    case 'date':
    case 'datePicker':
      return {
        component: 'DatePicker',
        props: { placeholder: '请选择日期', clearable: true },
        rules,
      };

    case 'dateRange':
    case 'dateRangePicker':
      return {
        component: 'DateRangePicker',
        props: { placeholder: ['开始日期', '结束日期'], clearable: true },
        rules,
      };

    case 'dateTime':
      return {
        component: 'DatePicker',
        props: { placeholder: '请选择日期时间', showTime: true, clearable: true },
        rules,
      };

    case 'dateTimeRange':
      return {
        component: 'DateRangePicker',
        props: { placeholder: ['开始时间', '结束时间'], showTime: true, clearable: true },
        rules,
      };

    case 'time':
    case 'timePicker':
      return {
        component: 'TimePicker',
        props: { placeholder: '请选择时间', clearable: true },
        rules,
      };

    case 'treeSelect':
    case 'tree':
      return {
        component: 'TreeSelect',
        props: {
          data: [],
          keys: { label: 'label', value: 'value', children: 'children' },
          placeholder: '请选择',
          clearable: true,
          ...(field.multiple ? { multiple: true } : {}),
        },
        rules,
      };

    case 'password':
      return {
        component: 'Input',
        props: { type: 'password', placeholder: '请输入' },
        rules,
      };

    case 'email':
      return {
        component: 'Input',
        props: { type: 'email', placeholder: '请输入邮箱' },
        rules: 'email',
      };

    case 'phone':
    case 'mobile':
      return {
        component: 'Input',
        props: { type: 'phone', placeholder: '请输入手机号' },
        rules: 'phone',
      };

    case 'url':
      return {
        component: 'Input',
        props: { type: 'url', placeholder: '请输入网址' },
        rules: 'url',
      };

    case 'color':
      return {
        component: 'ColorPicker',
        props: { placeholder: '请选择颜色' },
      };

    case 'slider':
      return {
        component: 'Slider',
        props: { min: 0, max: 100 },
      };

    case 'rate':
      return {
        component: 'Rate',
        props: {},
      };

    case 'cascader':
      return {
        component: 'Cascader',
        props: {
          options: [],
          placeholder: '请选择',
          clearable: true,
        },
        rules,
      };

    case 'avatar':
      return {
        component: 'Upload',
        props: {
          accept: 'image/*',
          placeholder: '请上传头像',
        },
      };

    case 'image':
      return {
        component: 'Upload',
        props: {
          accept: 'image/*',
          placeholder: '请上传图片',
        },
      };

    case 'images':
      return {
        component: 'Upload',
        props: {
          accept: 'image/*',
          multiple: true,
          placeholder: '请上传图片',
        },
      };

    case 'file':
      return {
        component: 'Upload',
        props: {
          placeholder: '请上传文件',
        },
      };

    case 'files':
      return {
        component: 'Upload',
        props: {
          multiple: true,
          placeholder: '请上传文件',
        },
      };

    case 'inputGroup':
      return {
        component: 'InputGroup',
        props: { placeholder: '请输入' },
        rules,
      };

    case 'divider':
      return {
        component: 'Divider',
        props: {},
      };

    case 'string':
    case 'text':
    case 'input':
    default:
      return {
        component: 'Input',
        props: { placeholder: '请输入' },
        rules,
      };
  }
}

function getSearchFieldComponent(field: Field): string {
  const { type } = field;

  switch (type) {
    case 'number':
    case 'integer':
    case 'inputNumber':
      return `            <FormItem label="${field.label}" name="${field.name}">
              <InputNumber
                v-model="searchForm.${field.name}"
                placeholder="请输入"
                clearable
                class="w-full"
              />
            </FormItem>`;

    case 'select':
      return `            <FormItem label="${field.label}" name="${field.name}">
              <Select
                v-model="searchForm.${field.name}"
                :options="statusOptions"
                placeholder="请选择"
                clearable
                class="w-full"
              />
            </FormItem>`;

    case 'dateRange':
    case 'dateRangePicker':
      return `            <FormItem label="${field.label}" name="${field.name}" class="col-span-2">
              <DateRangePicker
                v-model="searchForm.${field.name}"
                :placeholder="['开始时间', '结束时间']"
                clearable
                class="w-full"
              />
            </FormItem>`;

    case 'date':
    case 'datePicker':
      return `            <FormItem label="${field.label}" name="${field.name}">
              <DatePicker
                v-model="searchForm.${field.name}"
                placeholder="请选择日期"
                clearable
                class="w-full"
              />
            </FormItem>`;

    case 'switch':
      return `            <FormItem label="${field.label}" name="${field.name}">
              <Select
                v-model="searchForm.${field.name}"
                :options="[{ label: '开启', value: 1 }, { label: '关闭', value: 0 }]"
                placeholder="请选择"
                clearable
                class="w-full"
              />
            </FormItem>`;

    case 'treeSelect':
    case 'tree':
      return `            <FormItem label="${field.label}" name="${field.name}">
              <TreeSelect
                v-model="searchForm.${field.name}"
                :data="[]"
                placeholder="请选择"
                clearable
                class="w-full"
              />
            </FormItem>`;

    default:
      return `            <FormItem label="${field.label}" name="${field.name}">
              <Input
                v-model="searchForm.${field.name}"
                placeholder="请输入"
                clearable
              />
            </FormItem>`;
  }
}

function generateTemplateData(opts: CliOptions): TemplateData {
  const upperName = (opts.module + ':' + opts.name).toUpperCase().replace(/:/, '_');
  const camelName = toCamelCase(opts.name);
  const pascalName = toPascalCase(opts.name);
  const fields = parseFields(opts.fields);
  const searchFields = fields.filter((f) =>
    ['name', 'code', 'status', 'created_at', 'type'].includes(f.name) ||
    ['select', 'dateRange', 'date', 'switch', 'treeSelect'].includes(f.type),
  );
  const formFields = fields.filter((f) =>
    !['id', 'created_at', 'updated_at', 'deleted_at'].includes(f.name),
  );

  return {
    ...opts,
    upperName,
    camelName,
    pascalName,
    apiPath: `/system/${opts.name.replace(/([A-Z])/g, '-$1').toLowerCase()}`,
    tableNameResolved: opts.tableName || `system_${toSnakeCase(opts.name)}`,
    permissionResolved: opts.permission || `system:${opts.name.replace(/([A-Z])/g, '-$1').toLowerCase()}`,
    cnNameResolved: opts.cnName || opts.name,
    fields,
    searchFields,
    formFields,
    hasSort: fields.some((f) => f.name === 'sort'),
    hasStatus: fields.some((f) => f.name === 'status'),
  };
}

function generateApi(data: TemplateData): string {
  const formFieldTypes = data.formFields.map(f => {
    const isNumber = ['number', 'integer', 'inputNumber', 'sort', 'status'].includes(f.type) ||
                     ['sort', 'status'].includes(f.name);
    return {
      ...f,
      tsType: isNumber ? 'number' : 'string',
    };
  });

  return `import { requestClient } from '#/api/request';
import type { BatchIdsPayload, StatusValue } from '#/types/common';
import type { PageQuery, PageResponse } from '#/types/paging';

export namespace ${data.pascalName}Api {
  export interface ListItem {
    id: number;
${formFieldTypes.filter(f => f.name !== 'id').map(f => `    ${f.name}${f.name === 'remark' || f.name === 'code' ? '?' : ''}: ${f.tsType};`).join('\n')}
  }

  export interface ListQuery extends Partial<PageQuery> {
${data.searchFields.map(f => {
  const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);
  const isDateRange = ['dateRange', 'dateRangePicker'].includes(f.type);
  if (isDateRange) return `    ${f.name}?: string[];`;
  if (isNumber) return `    ${f.name}?: number;`;
  return `    ${f.name}?: string;`;
}).join('\n')}
  }

  export interface SubmitPayload {
${formFieldTypes.filter(f => f.name !== 'id').map(f => {
  const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);
  return `    ${f.name}: ${isNumber ? 'number' : 'string'};`;
}).join('\n')}
  }

  export interface ChangeStatusPayload {
    id: number;
    status: number;
  }

  export type BatchPayload = BatchIdsPayload<number>;
  export type ListResponse = PageResponse<ListItem>;
  export type OptionListResponse = ListItem[] | ListResponse;
}

export function get${data.pascalName}PageList(params: ${data.pascalName}Api.ListQuery) {
  return requestClient.get<${data.pascalName}Api.ListResponse>('/system/${data.name}/index', { params });
}

export function get${data.pascalName}List(params?: ${data.pascalName}Api.ListQuery) {
  return requestClient.get<${data.pascalName}Api.OptionListResponse>('/system/${data.name}/list', { params });
}

export function getRecycle${data.pascalName}List(params: ${data.pascalName}Api.ListQuery) {
  return requestClient.get<${data.pascalName}Api.ListResponse>('/system/${data.name}/recycle', { params });
}

export function save${data.pascalName}(data: ${data.pascalName}Api.SubmitPayload) {
  return requestClient.post<void>('/system/${data.name}/save', data);
}

export function update${data.pascalName}(id: number, data: ${data.pascalName}Api.SubmitPayload) {
  return requestClient.put<void>(\`/system/${data.name}/update/\${id}\`, data);
}

export function delete${data.pascalName}(ids: number[]) {
  return requestClient.delete<void>('/system/${data.name}/delete', { data: { ids } });
}

export function realDelete${data.pascalName}(ids: number[]) {
  return requestClient.delete<void>('/system/${data.name}/realDelete', { data: { ids } });
}

export function recovery${data.pascalName}(ids: number[]) {
  return requestClient.put<void>('/system/${data.name}/recovery', { ids });
}

export function change${data.pascalName}Status(data: ${data.pascalName}Api.ChangeStatusPayload) {
  return requestClient.put<void>('/system/${data.name}/changeStatus', data);
}
`;
}

function generateModel(data: TemplateData): string {
  const searchFieldTypes = data.searchFields.map(f => {
    const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);
    const isDateRange = ['dateRange', 'dateRangePicker'].includes(f.type);
    return {
      ...f,
      tsType: isDateRange ? 'string[]' : isNumber ? 'number' : 'string',
      optional: f.name === 'status' || f.name === 'sort',
    };
  });

  const formFieldTypes = data.formFields.map(f => {
    const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);
    return {
      ...f,
      tsType: isNumber ? 'number' : 'string',
    };
  });

  return `import type { ${data.pascalName}Api } from '#/api/${data.module}/${data.name}';
import type { OptionItem } from '#/types/common';

export type ${data.pascalName}ListItem = ${data.pascalName}Api.ListItem;

export interface ${data.pascalName}SearchFormModel {
${searchFieldTypes.map(f => `  ${f.name}${f.optional ? '?' : ''}: ${f.tsType};`).join('\n')}
}

export interface ${data.pascalName}FormModel {
${formFieldTypes.map(f => {
  if (f.name === 'id') return `  ${f.name}: number;`;
  return `  ${f.name}: ${f.tsType};`;
}).join('\n')}
}

export type ${data.pascalName}ColumnOptionItem = OptionItem<string>;

export interface ${data.pascalName}TableColumn {
  align?: 'left' | 'center' | 'right';
  colKey: string;
  fixed?: 'left' | 'right';
  minWidth?: number;
  title?: string;
  type?: 'multiple' | 'single';
  width?: number;
}
`;
}

function generateSchemas(data: TemplateData): string {
  const searchFieldsInit = data.searchFields.map(f => {
    const isDateRange = ['dateRange', 'dateRangePicker'].includes(f.type);
    const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);

    if (isDateRange) return `    ${f.name}: [],`;
    if (isNumber) return f.name === 'status' ? `    ${f.name}: undefined,` : `    ${f.name}: 0,`;
    return `    ${f.name}: '',`;
  }).join('\n');

  const formFieldsInit = data.formFields.map(f => {
    const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);
    if (f.name === 'sort') return `    ${f.name}: 1,`;
    if (f.name === 'status') return `    ${f.name}: 1,`;
    if (isNumber) return `    ${f.name}: 0,`;
    return `    ${f.name}: '',`;
  }).join('\n');

  const tableColumns = data.formFields
    .filter(f => f.name !== 'id' && f.name !== 'remark' && f.type !== 'textarea' && f.type !== 'avatar' && f.type !== 'image' && f.type !== 'file')
    .map(f => {
      const width = f.name === 'sort' ? '140' : f.name === 'status' ? '120' : '120';
      return `    { align: 'center', colKey: '${f.name}', minWidth: ${width}, title: '${f.label}' },`;
    }).join('\n');

  const hasRemark = data.formFields.some(f => f.name === 'remark' || f.type === 'textarea');
  const hasImage = data.formFields.some(f => ['avatar', 'image', 'images', 'file', 'files'].includes(f.type));

  return `import type {
  ${data.pascalName}ColumnOptionItem,
  ${data.pascalName}FormModel,
  ${data.pascalName}SearchFormModel,
  ${data.pascalName}TableColumn,
} from './model';

export function create${data.pascalName}SearchForm(): ${data.pascalName}SearchFormModel {
  return {
${searchFieldsInit}
  };
}

export function create${data.pascalName}FormDefaultValues(): ${data.pascalName}FormModel {
  return {
${formFieldsInit}
  };
}

export function create${data.pascalName}TableColumns(): ${data.pascalName}TableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 52,
    },
${hasImage ? "    { align: 'center', colKey: 'id', title: 'ID', width: 80 },\n" : ''}${hasImage ? '' : "    { align: 'center', colKey: 'id', title: 'ID', width: 80 },\n"}${tableColumns}
${hasRemark ? `    { align: 'center', colKey: 'remark', minWidth: 160, title: '备注' },` : ''}
    { align: 'center', colKey: 'created_at', title: '创建时间', width: 180 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: '操作',
      width: 220,
    },
  ];
}

export function create${data.pascalName}ColumnOptions(
  columns: ${data.pascalName}TableColumn[],
): ${data.pascalName}ColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}
`;
}

function generateIndex(data: TemplateData): string {
  const searchFormItems = data.searchFields.map(f => getSearchFieldComponent(f)).join('\n');

  const statusSwitchTemplate = data.hasStatus ? `          <template #status="{ row }">
            <Switch
              :disabled="isRecycleBin"
              :value="row.status === 1"
              @change="(value: unknown) => handleStatusSwitchChange(row, value)"
            />
          </template>
` : '';

  const statusOptions = data.fields.find(f => f.name === 'status')?.options ||
    [{ label: '正常', value: 1 }, { label: '停用', value: 2 }];

  return `<script lang="ts" setup>
import type { ${data.pascalName}ListItem } from './model';

import { computed, onMounted, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { message } from '#/adapter/tdesign';

import {
  AddIcon,
  DeleteIcon,
  EditIcon,
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Button,
  DateRangePicker,
  Form,
  FormItem,
  Input,
  InputNumber,
  Popconfirm,
  Select,
  Space,
  Switch,
  Table,
  TreeSelect,
} from 'tdesign-vue-next';

import CrudToolbar from '#/components/crud/crud-toolbar.vue';
import {
  change${data.pascalName}Status,
  delete${data.pascalName},
  realDelete${data.pascalName},
  recovery${data.pascalName},
} from '#/api/${data.module}/${data.name}';
import type { DictOption } from '#/composables/crud/use-dict-options';
import { useDictOptions } from '#/composables/crud/use-dict-options';

import ${data.pascalName}Modal from './components/${data.name}-modal.vue';
import type { ${data.pascalName}TableColumn } from './model';
import { create${data.pascalName}ColumnOptions, create${data.pascalName}TableColumns } from './schemas';
import { use${data.pascalName}Crud } from './use-${data.name}-crud';

defineOptions({ name: 'System${data.pascalName}' });

type ${data.pascalName}ModalInstance = {
  open: (data?: Partial<${data.pascalName}ListItem>) => void;
};

const ${data.camelName}ModalRef = ref<${data.pascalName}ModalInstance>();
const statusOptions = ref<DictOption[]>(${JSON.stringify(statusOptions.map(o => ({ label: o.label, value: String(o.value) })))});

const columns: ${data.pascalName}TableColumn[] = create${data.pascalName}TableColumns();
const columnOptions = create${data.pascalName}ColumnOptions(columns);
const allColumnKeys = columnOptions.map((item) => item.value);
const visibleColumns = ref<string[]>([...allColumnKeys]);

const displayColumns = computed({
  get: () => ['row-select', ...visibleColumns.value],
  set: (value: string[]) => {
    visibleColumns.value = value.filter((item) => item !== 'row-select');
  },
});

const {
  clearSelectedRowKeys,
  fetchTableData,
  handlePageChange,
  handleReset,
  handleSearch,
  handleSelectChange,
  isRecycleBin,
  loading,
  pagination,
  searchForm,
  selectedRowKeys,
  tableData,
  toggleRecycleBin,
} = use${data.pascalName}Crud();

const { getDictOptions } = useDictOptions();

function toIds(keys: Array<number | string>) {
  return keys.map((key) => Number(key));
}

async function fetchStatusOptions() {
  const options = await getDictOptions('data_status');
  if (options.length > 0) {
    statusOptions.value = options;
  }
}

function handleAdd() {
  ${data.camelName}ModalRef.value?.open();
}

function handleEdit(row: ${data.pascalName}ListItem) {
  ${data.camelName}ModalRef.value?.open(row);
}

async function handleDelete(row: ${data.pascalName}ListItem) {
  try {
    await (isRecycleBin.value ? realDelete${data.pascalName}([row.id]) : delete${data.pascalName}([row.id]));
    message.success('操作成功');
    await fetchTableData();
  } catch (error) {
    console.error(error);
    message.error('删除失败，请稍后重试');
  }
}

async function handleBatchDelete() {
  if (selectedRowKeys.value.length === 0) {
    message.warning('请选择需要操作的数据');
    return;
  }

  const ids = toIds(selectedRowKeys.value);
  try {
    await (isRecycleBin.value ? realDelete${data.pascalName}(ids) : delete${data.pascalName}(ids));
    message.success('操作成功');
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    console.error(error);
    message.error('批量删除失败，请稍后重试');
  }
}

async function handleRecovery(row: ${data.pascalName}ListItem) {
  try {
    await recovery${data.pascalName}([row.id]);
    message.success('恢复成功');
    await fetchTableData();
  } catch (error) {
    console.error(error);
    message.error('恢复失败，请稍后重试');
  }
}

async function handleBatchRecovery() {
  if (selectedRowKeys.value.length === 0) {
    message.warning('请选择需要操作的数据');
    return;
  }

  const ids = toIds(selectedRowKeys.value);
  try {
    await recovery${data.pascalName}(ids);
    message.success('恢复成功');
    clearSelectedRowKeys();
    await fetchTableData();
  } catch (error) {
    console.error(error);
    message.error('批量恢复失败，请稍后重试');
  }
}

async function handleStatusChange(row: ${data.pascalName}ListItem, checked: boolean) {
  const status = checked ? 1 : 2;
  try {
    await change${data.pascalName}Status({ id: row.id, status });
    message.success('状态更新成功');
    await fetchTableData();
  } catch (error) {
    console.error(error);
    message.error('状态更新失败，请稍后重试');
  }
}

function handleSuccess() {
  void fetchTableData();
}

function handleTableSelectChange(keys: Array<number | string>) {
  handleSelectChange(keys);
}

function handleStatusSwitchChange(row: ${data.pascalName}ListItem, value: unknown) {
  void handleStatusChange(row, Boolean(value));
}

onMounted(() => {
  void fetchStatusOptions();
  void fetchTableData();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-3">
      <div class="rounded-md bg-white p-4">
        <Form :data="searchForm" label-width="80px" colon>
          <div class="grid grid-cols-4 gap-x-4">
${searchFormItems}
          </div>
          <div class="flex justify-end gap-2 pt-2">
            <Button theme="default" @click="handleReset">重置</Button>
            <Button theme="primary" @click="handleSearch">
              <template #icon><SearchIcon /></template>
              查询
            </Button>
          </div>
        </Form>
      </div>

      <div class="flex min-h-0 flex-1 flex-col rounded-md bg-white p-4">
        <div class="mb-3 flex items-center justify-between">
          <Space>
            <template v-if="!isRecycleBin">
              <Button theme="primary" @click="handleAdd">
                <template #icon><AddIcon /></template>
                新增
              </Button>
              <Button theme="danger" variant="outline" @click="handleBatchDelete">
                <template #icon><DeleteIcon /></template>
                删除
              </Button>
            </template>
            <template v-else>
              <Button theme="success" @click="handleBatchRecovery">恢复</Button>
              <Button theme="danger" @click="handleBatchDelete">彻底删除</Button>
            </template>
          </Space>

          <CrudToolbar
            v-model="visibleColumns"
            :column-options="columnOptions"
            :is-recycle-bin="isRecycleBin"
            @refresh="fetchTableData"
            @toggle-recycle="toggleRecycleBin"
          />
        </div>

        <Table
          v-model:display-columns="displayColumns"
          :columns="columns"
          :data="tableData"
          :loading="loading"
          :pagination="pagination"
          :selected-row-keys="selectedRowKeys"
          row-key="id"
          hover
          stripe
          @page-change="handlePageChange"
          @select-change="handleTableSelectChange"
        >
${statusSwitchTemplate}
          <template #action="{ row }">
            <div class="flex items-center justify-center gap-1">
              <template v-if="!isRecycleBin">
                <Button
                  size="small"
                  theme="primary"
                  variant="outline"
                  @click="handleEdit(row)"
                >
                  <template #icon><EditIcon /></template>
                  编辑
                </Button>
                <Popconfirm
                  content="确认删除吗？"
                  @confirm="handleDelete(row)"
                >
                  <Button size="small" theme="danger" variant="outline">
                    <template #icon><DeleteIcon /></template>
                    删除
                  </Button>
                </Popconfirm>
              </template>

              <template v-else>
                <Popconfirm
                  content="确认恢复吗？"
                  @confirm="handleRecovery(row)"
                >
                  <Button size="small" theme="primary" variant="outline">
                    恢复
                  </Button>
                </Popconfirm>
                <Popconfirm
                  content="确认彻底删除吗？"
                  @confirm="handleDelete(row)"
                >
                  <Button size="small" theme="danger" variant="outline">
                    彻底删除
                  </Button>
                </Popconfirm>
              </template>
            </div>
          </template>
        </Table>
      </div>
    </div>

    <${data.pascalName}Modal ref="${data.camelName}ModalRef" @success="handleSuccess" />
  </Page>
</template>
`;
}

function generateModal(data: TemplateData): string {
  // Build schema as object array first, then stringify
  const schemaItems: string[] = [
    `    {
      component: 'Input',
      dependencies: {
        show: false,
        triggerFields: [''],
      },
      fieldName: 'id',
      label: 'ID',
    },`,
  ];

  for (const f of data.formFields) {
    if (f.name === 'id') continue;

    const config = getFieldComponentConfig(f);
    const { component, props, rules } = config;

    // Build props object string
    let propsEntries: string[] = [];
    if (Object.keys(props).length > 0) {
      propsEntries = Object.entries(props).map(([k, v], idx, arr) => {
        const isLast = idx === arr.length - 1;
        let valueStr: string;
        if (typeof v === 'string') valueStr = `'${v}'`;
        else if (Array.isArray(v)) valueStr = JSON.stringify(v);
        else if (typeof v === 'object') valueStr = JSON.stringify(v);
        else valueStr = String(v);
        const comma = isLast ? '' : ',';
        return `        ${k}: ${valueStr}${comma}`;
      });
    }

    // Determine label
    let label = f.label;
    if (f.name === 'name') label = '名称';
    else if (f.name === 'code') label = '编码';
    else if (f.name === 'remark' || f.type === 'textarea') label = '备注';

    // Build this schema item
    let itemLines = [`    {`];

    // Special handling for known types
    if (f.name === 'name') {
      itemLines.push(`      component: '${component}',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      if (propsEntries.length > 0) {
        itemLines.push(`      props: {`);
        itemLines.push(...propsEntries.map(p => p.replace(/^        /, '        ')));
        itemLines.push(`      },`);
      }
      itemLines.push(`      rules: 'required',`);
    } else if (f.name === 'code') {
      itemLines.push(`      component: '${component}',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      if (propsEntries.length > 0) {
        itemLines.push(`      props: {`);
        itemLines.push(...propsEntries);
        itemLines.push(`      },`);
      }
      if (rules) itemLines.push(`      rules: '${rules}',`);
    } else if (f.name === 'sort') {
      itemLines.push(`      component: 'InputNumber',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '排序',`);
      itemLines.push(`      defaultValue: 1,`);
      itemLines.push(`      props: {`);
      itemLines.push(`        min: 0,`);
      itemLines.push(`        max: 1000,`);
      itemLines.push(`      },`);
      itemLines.push(`      rules: 'required',`);
    } else if (f.name === 'status') {
      const options = f.options || [{ label: '正常', value: 1 }, { label: '停用', value: 2 }];
      itemLines.push(`      component: 'RadioGroup',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '状态',`);
      itemLines.push(`      defaultValue: 1,`);
      itemLines.push(`      props: {`);
      itemLines.push(`        options: ${JSON.stringify(options)},`);
      itemLines.push(`      },`);
      itemLines.push(`      rules: 'required',`);
    } else if (f.name === 'remark' || f.type === 'textarea') {
      itemLines.push(`      component: 'Textarea',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '备注',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        placeholder: '请输入备注',`);
      itemLines.push(`        autosize: { minRows: 3, maxRows: 6 },`);
      itemLines.push(`      },`);
      itemLines.push(`      formItemClass: 'col-span-2',`);
    } else if (f.type === 'password') {
      itemLines.push(`      component: 'Input',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        type: 'password',`);
      itemLines.push(`        placeholder: '请输入密码',`);
      itemLines.push(`      },`);
      itemLines.push(`      rules: 'required',`);
    } else if (f.type === 'email') {
      itemLines.push(`      component: 'Input',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        type: 'email',`);
      itemLines.push(`        placeholder: '请输入邮箱',`);
      itemLines.push(`      },`);
      itemLines.push(`      rules: 'email',`);
    } else if (f.type === 'phone' || f.type === 'mobile') {
      itemLines.push(`      component: 'Input',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '手机号',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        placeholder: '请输入手机号',`);
      itemLines.push(`      },`);
    } else if (f.type === 'treeSelect' || f.type === 'tree') {
      itemLines.push(`      component: 'TreeSelect',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        data: [],`);
      itemLines.push(`        keys: { label: 'label', value: 'value', children: 'children' },`);
      itemLines.push(`        placeholder: '请选择',`);
      itemLines.push(`        clearable: true,`);
      itemLines.push(`      },`);
    } else if (f.type === 'date' || f.type === 'datePicker') {
      itemLines.push(`      component: 'DatePicker',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        placeholder: '请选择日期',`);
      itemLines.push(`        clearable: true,`);
      itemLines.push(`      },`);
    } else if (f.type === 'dateRange' || f.type === 'dateRangePicker') {
      itemLines.push(`      component: 'DateRangePicker',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        placeholder: ['开始日期', '结束日期'],`);
      itemLines.push(`        clearable: true,`);
      itemLines.push(`      },`);
    } else if (f.type === 'avatar' || f.type === 'image') {
      itemLines.push(`      component: 'Upload',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        accept: 'image/*',`);
      itemLines.push(`        placeholder: '请上传图片',`);
      itemLines.push(`      },`);
      itemLines.push(`      formItemClass: 'col-span-2',`);
    } else if (f.type === 'select' && f.options) {
      itemLines.push(`      component: 'Select',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        options: ${JSON.stringify(f.options)},`);
      itemLines.push(`        placeholder: '请选择',`);
      itemLines.push(`        clearable: true,`);
      itemLines.push(`      },`);
    } else if (f.type === 'radio' && f.options) {
      itemLines.push(`      component: 'RadioGroup',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      itemLines.push(`      props: {`);
      itemLines.push(`        options: ${JSON.stringify(f.options)},`);
      itemLines.push(`      },`);
    } else {
      // Default case
      itemLines.push(`      component: '${component}',`);
      itemLines.push(`      fieldName: '${f.name}',`);
      itemLines.push(`      label: '${label}',`);
      if (propsEntries.length > 0) {
        itemLines.push(`      props: {`);
        itemLines.push(...propsEntries);
        itemLines.push(`      },`);
      }
      if (rules) itemLines.push(`      rules: '${rules}',`);
    }

    itemLines.push(`    },`);
    schemaItems.push(itemLines.join('\n'));
  }

  const schemaStr = schemaItems.join('\n');

  return `<script lang="ts" setup>
import type { ${data.pascalName}Api } from '#/api/${data.module}/${data.name}';

import { nextTick } from 'vue';

import { useVbenModal } from '@vben/common-ui';

import { MessagePlugin } from 'tdesign-vue-next';

import { useVbenForm } from '#/adapter/form';
import { save${data.pascalName}, update${data.pascalName} } from '#/api/${data.module}/${data.name}';

import { create${data.pascalName}FormDefaultValues } from '../schemas';

const emit = defineEmits(['success']);

const [Form, formApi] = useVbenForm({
  showDefaultActions: false,
  commonConfig: {
    labelWidth: 90,
  },
  schema: [
${schemaStr}
  ],
});

const [Modal, modalApi] = useVbenModal({
  onConfirm: async () => {
    try {
      const { valid } = await formApi.validate();
      if (!valid) return;

      const values = await formApi.getValues<${data.pascalName}Api.SubmitPayload>();
      modalApi.setState({ confirmLoading: true });

      if (values.id) {
        await update${data.pascalName}(Number(values.id), values);
      } else {
        await save${data.pascalName}(values);
      }

      MessagePlugin.success(values.id ? '更新成功' : '新增成功');
      emit('success');
      modalApi.close();
    } catch (error) {
      console.error(error);
    } finally {
      modalApi.setState({ confirmLoading: false });
    }
  },
  class: 'w-[560px]',
});

async function open(data?: Partial<${data.pascalName}Api.SubmitPayload>) {
  modalApi.setState({
    title: data?.id ? '编辑${data.cnNameResolved}' : '新增${data.cnNameResolved}',
  });
  modalApi.open();

  await formApi.resetForm();
  formApi.setValues(create${data.pascalName}FormDefaultValues());
  if (data) {
    formApi.setValues(data);
  }
  await nextTick();
  await formApi.resetValidate();
}

defineExpose({
  open,
});
</script>

<template>
  <Modal>
    <Form />
  </Modal>
</template>
`;
}

function generateUseCrud(data: TemplateData): string {
  const buildParams = data.searchFields.map(f => {
    const isDateRange = ['dateRange', 'dateRangePicker'].includes(f.type);
    const isNumber = ['number', 'integer', 'inputNumber'].includes(f.type) || ['sort', 'status'].includes(f.name);

    if (isDateRange) return `      if (form.${f.name}?.length === 2 && form.${f.name}[0]) {
        params.${f.name} = form.${f.name};
      }`;
    if (isNumber) {
      if (f.name === 'status') return `      if (form.${f.name} !== undefined) params.${f.name} = form.${f.name};`;
      return `      if (form.${f.name}) params.${f.name} = form.${f.name};`;
    }
    return `      if (form.${f.name}) params.${f.name} = form.${f.name};`;
  }).join('\n');

  return `import type { ${data.pascalName}ListItem } from './model';
import type { ${data.pascalName}Api } from '#/api/${data.module}/${data.name}';

import { get${data.pascalName}PageList, getRecycle${data.pascalName}List } from '#/api/${data.module}/${data.name}';
import { useCrudPage } from '#/composables/crud/use-crud-page';

import { create${data.pascalName}SearchForm } from './schemas';

export function use${data.pascalName}Crud() {
  return useCrudPage<${data.pascalName}ListItem, ReturnType<typeof create${data.pascalName}SearchForm>>({
    defaultSearchForm: create${data.pascalName}SearchForm,
    fetchList: (params, context) =>
      context.isRecycleBin ? getRecycle${data.pascalName}List(params) : get${data.pascalName}PageList(params),
    buildParams: (form) => {
      const params: Partial<${data.pascalName}Api.ListQuery> = {};
${buildParams}
      return params;
    },
    resolveTotal: (response) =>
      Number(response?.pageInfo?.total || response?.total || 0),
  });
}
`;
}

async function main() {
  const opts = cli.options as CliOptions;

  if (!opts.name || !opts.module) {
    cli.help();
    process.exit(1);
  }

  console.log(`\n🚀 Generating CRUD page for "${opts.name}" in module "${opts.module}"...\n`);

  const data = generateTemplateData(opts);
  // Find project root by looking for apps/backend directory
  let baseDir = __dirname;
  while (baseDir && !fs.existsSync(path.join(baseDir, 'apps', 'backend'))) {
    baseDir = path.dirname(baseDir);
  }
  if (!baseDir) {
    console.error('❌ Could not find project root (apps/backend not found)');
    process.exit(1);
  }

  // Ensure directories exist
  const apiDir = path.join(baseDir, 'apps', 'backend', 'src', 'api', data.module);
  const viewDir = path.join(baseDir, 'apps', 'backend', 'src', 'views', data.module, data.name);
  const componentsDir = path.join(viewDir, 'components');

  fs.mkdirSync(apiDir, { recursive: true });
  fs.mkdirSync(componentsDir, { recursive: true });

  // Generate files
  const files = [
    [`api/${data.module}/${data.name}.ts`, generateApi(data)],
    [`views/${data.module}/${data.name}/index.vue`, generateIndex(data)],
    [`views/${data.module}/${data.name}/model.ts`, generateModel(data)],
    [`views/${data.module}/${data.name}/schemas.ts`, generateSchemas(data)],
    [`views/${data.module}/${data.name}/components/${data.name}-modal.vue`, generateModal(data)],
    [`views/${data.module}/${data.name}/use-${data.name}-crud.ts`, generateUseCrud(data)],
  ];

  for (const [relativePath, content] of files) {
    const outputPath = path.join(baseDir, 'apps', 'backend', 'src', relativePath);
    fs.writeFileSync(outputPath, content);
    console.log(`✅ Generated: ${relativePath}`);
  }

  console.log('\n✅ CRUD page generated successfully!\n');
  console.log('📝 Next steps:');
  console.log(`   1. Review generated files in apps/backend/src/views/${data.module}/${data.name}/`);
  console.log('   2. Customize the fields and API endpoints as needed');
  console.log('   3. Add the page to router and menu\n');
}

main().catch(console.error);
