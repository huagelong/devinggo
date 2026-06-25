# Gen-CRUD

DevingGo Admin UI CRUD page generator. Quickly generates complete CRUD page structure based on the DevingGo frontend conventions.

## Features

- Generates standard CRUD page structure following DevingGo conventions
- Supports custom fields configuration with 30+ field types
- Generates API client, model, schemas, Vue components, and CRUD composable
- Supports Chinese naming for better i18n integration

## Usage

```bash
# Generate with default fields (name, code, status, sort, remark, created_at)
node scripts/gen-crud/src/index.ts --name=user --module=system --cn-name=用户

# Generate with custom fields
node scripts/gen-crud/src/index.ts --name=product --module=system --cn-name=产品 --fields=id,name,code,status,sort,price,remark

# Run from admin-ui directory
cd admin-ui
node scripts/gen-crud/src/index.ts --name=category --module=system --cn-name=分类
```

## Generated Structure

```
apps/backend/src/
├── api/system/
│   └── {name}.ts              # API client with all CRUD operations
└── views/system/{name}/
    ├── index.vue               # Main page component
    ├── model.ts                # TypeScript interfaces
    ├── schemas.ts              # Form/search/table schema factories
    ├── components/
    │   └── {name}-modal.vue   # Add/Edit modal
    └── use-{name}-crud.ts     # CRUD composable using useCrudPage
```

## Options

| Option | Description | Required |
|--------|-------------|----------|
| `--name` | Business name (e.g., user, post, category) | Yes |
| `--module` | Module name (e.g., system) | Yes |
| `--cn-name` | Chinese name for display | No (defaults to name) |
| `--fields` | Comma-separated field list | No (uses defaults) |
| `--permission` | Permission prefix | No (auto-generated) |
| `--table-name` | Database table name | No (auto-generated) |

## Field Format

Fields can be specified in multiple formats:

```bash
# Simple type
--fields=name,code,status

# With explicit type
--fields=name:string,code:string,status:select

# With options (for select, radio, checkbox)
--fields=status:select:正常=1;停用=2

# With dict type reference
--fields=status:select:dict:data_status

# Multiple words in name using underscore or camelCase
--fields=first_name:string,lastName:string
```

## Supported Field Types

### Text Inputs
| Type | Description |
|------|-------------|
| `string` | Text input (default) |
| `text` | Text |
| `input` | Input field |
| `password` | Password input |
| `email` | Email input |
| `phone` | Phone number |
| `mobile` | Mobile number |
| `url` | URL input |

### Number Inputs
| Type | Description |
|------|-------------|
| `number` | Number with default 0-1000 range |
| `integer` | Integer |
| `inputNumber` | Number input |
| `slider` | Slider control |

### Textarea
| Type | Description |
|------|-------------|
| `textarea` | Multi-line text input |

### Select/Dropdown
| Type | Description |
|------|-------------|
| `select` | Dropdown select |

### Radio/Checkbox
| Type | Description |
|------|-------------|
| `radio` | Radio buttons |
| `checkbox` | Checkbox group |
| `switch` | Toggle switch |

### Date/Time
| Type | Description |
|------|-------------|
| `date` | Single date |
| `datePicker` | Date picker |
| `dateRange` | Date range picker |
| `dateTime` | Date with time |
| `dateTimeRange` | Date/time range |
| `time` | Time |
| `timePicker` | Time picker |

### Tree/Cascader
| Type | Description |
|------|-------------|
| `treeSelect` | Tree dropdown select |
| `tree` | Tree select |
| `cascader` | Cascader select |

### Upload
| Type | Description |
|------|-------------|
| `upload` | File upload |
| `image` | Single image upload |
| `images` | Multiple image upload |
| `file` | Single file upload |
| `files` | Multiple file upload |
| `avatar` | Avatar image upload |

### Other
| Type | Description |
|------|-------------|
| `color` | Color picker |
| `rate` | Rating control |
| `autocomplete` | Auto-complete input |
| `divider` | Divider line |

## Field Options

For `select`, `radio`, `checkbox` fields, you can specify options:

```bash
# Format: value1=Label1;value2=Label2
--fields=status:select:正常=1;停用=2

# For status field, if no options specified, defaults to:
# [{ label: '正常', value: 1 }, { label: '停用', value: 2 }]
```

## Default Fields

If no fields are specified, the following defaults are used:

```typescript
[
  { name: 'id', type: 'number', label: 'ID' },
  { name: 'name', type: 'string', label: '名称', required: true },
  { name: 'code', type: 'string', label: '编码' },
  { name: 'sort', type: 'number', label: '排序' },
  { name: 'status', type: 'select', label: '状态', options: [{ label: '正常', value: 1 }, { label: '停用', value: 2 }] },
  { name: 'remark', type: 'textarea', label: '备注' },
  { name: 'created_at', type: 'dateRange', label: '创建时间' },
]
```

## Next Steps

After generation:

1. Review generated files in `apps/backend/src/views/{module}/{name}/`
2. Customize fields and API endpoints as needed
3. Add the page to router configuration
4. Add menu entry in backend

## Development

```bash
# Install dependencies
pnpm install

# Run in development mode
pnpm run dev

# Or run directly with node
node scripts/gen-crud/src/index.ts --name=user --module=system
```
