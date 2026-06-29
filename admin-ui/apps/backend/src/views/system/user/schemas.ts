import type {
  ColumnOptionItem,
  UserActionDropdownItem,
  UserSearchFormModel,
  UserTableColumn,
} from './model';

import { $t } from '#/locales';

export function createUserSearchForm(): UserSearchFormModel {
  return {
    created_at: [],
    dept_ids: [],
    email: '',
    phone: '',
    post_id: undefined,
    role_id: undefined,
    status: undefined,
    user_type: undefined,
    username: '',
  };
}

export function createUserTableColumns(): UserTableColumn[] {
  return [
    {
      align: 'center',
      colKey: 'row-select',
      type: 'multiple',
      width: 50,
    },
    { align: 'center', colKey: 'avatar', title: $t('system.user.avatar'), width: 80 },
    { align: 'center', colKey: 'username', minWidth: 100, title: $t('system.user.username') },
    { align: 'center', colKey: 'dept_name', minWidth: 100, title: $t('system.user.dept') },
    { align: 'center', colKey: 'nickname', minWidth: 100, title: $t('system.user.nickname') },
    { align: 'center', colKey: 'role_name', minWidth: 100, title: $t('system.user.role') },
    { align: 'center', colKey: 'phone', minWidth: 120, title: $t('system.user.phone') },
    { align: 'center', colKey: 'post_name', minWidth: 100, title: $t('system.user.post') },
    { align: 'center', colKey: 'email', minWidth: 150, title: $t('system.user.email') },
    { align: 'center', colKey: 'status', title: $t('common.status'), width: 100 },
    { align: 'center', colKey: 'user_type', title: $t('system.user.userType'), width: 100 },
    { align: 'center', colKey: 'created_at', title: $t('system.user.registerTime'), width: 160 },
    {
      align: 'center',
      colKey: 'action',
      fixed: 'right',
      title: $t('common.action'),
      width: 220,
    },
  ];
}

export function createUserColumnOptions(
  columns: UserTableColumn[],
): ColumnOptionItem[] {
  return columns
    .filter((column) => column.colKey !== 'row-select' && column.title)
    .map((column) => ({
      label: String(column.title),
      value: String(column.colKey),
    }));
}

export const userActionDropdownOptions: UserActionDropdownItem[] = [
  { content: $t('system.user.resetPassword'), value: 'reset_password' },
  { content: $t('system.user.clearCache'), value: 'clear_cache' },
  { content: $t('system.user.setHome'), value: 'set_homepage' },
];
