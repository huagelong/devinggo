import type { Ref } from 'vue';
import type { IdType } from '#/types/common';

import { ref } from 'vue';

import { $t } from '@vben/locales';

import { logger } from '#/utils/logger';

import {
  changeUserStatus,
  clearUserCache,
  deleteUser,
  downloadUserImportTemplate,
  exportUserList,
  importUserFile,
  realDeleteUser,
  recoveryUser,
  resetPassword,
  setHomePage,
} from '#/api/system/user';
import { dialog, message } from '#/adapter/tdesign';
import { downloadResponseBlob } from '#/utils/download';

import type { UserActionDropdownItem, UserListItem } from './model';

type UserActionRow = Pick<UserListItem, 'id'> & Partial<UserListItem>;

interface UseUserActionsOptions {
  buildRequestParams: (includePagination?: boolean) => Record<string, unknown>;
  clearSelectedRowKeys: () => void;
  fetchTableData: () => void;
  isRecycleBin: Ref<boolean>;
  selectedRowKeys: Ref<IdType[]>;
}

export function useUserActions(options: UseUserActionsOptions) {
  const importInputRef = ref<HTMLInputElement>();
  const importLoading = ref(false);
  const exportLoading = ref(false);
  const templateLoading = ref(false);

  const setHomePageVisible = ref(false);
  const setHomePageLoading = ref(false);
  const selectedHomePage = ref('');
  const selectedHomePageUserId = ref<null | number>(null);

  function normalizeDashboardValue(value?: string) {
    if (!value) {
      return '';
    }

    if (value === '/analytics') {
      return 'statistics';
    }

    if (value === '/workspace') {
      return 'work';
    }

    return value;
  }

  function isSuperAdmin(row: UserActionRow) {
    return Number(row?.id) === 1;
  }

  function toIds(keys: IdType[]) {
    return keys.map((key) => Number(key));
  }

  async function handleDelete(row: UserActionRow) {
    if (isSuperAdmin(row)) {
      message.warning($t('common.superAdminCannotDelete2'));
      return;
    }

    try {
      await (options.isRecycleBin.value
        ? realDeleteUser([row.id])
        : deleteUser([row.id]));
      message.success($t('common.deleteSuccess'));
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.deleteFailed'));
    }
  }

  async function handleBatchDelete() {
    if (options.selectedRowKeys.value.length === 0) {
      message.warning($t('common.selectDataFirst'));
      return;
    }

    const ids = toIds(options.selectedRowKeys.value);
    if (ids.some((id) => id === 1)) {
      message.warning($t('common.superAdminCannotDelete2'));
      return;
    }

    try {
      await (options.isRecycleBin.value ? realDeleteUser(ids) : deleteUser(ids));
      message.success($t('common.operationSuccess'));
      options.clearSelectedRowKeys();
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.batchDeleteFailed'));
    }
  }

  async function handleRecovery(row: UserActionRow) {
    try {
      await recoveryUser([row.id]);
      message.success($t('common.recoverySuccess'));
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.recoveryFailed'));
    }
  }

  async function handleBatchRecovery() {
    if (options.selectedRowKeys.value.length === 0) {
      message.warning($t('common.selectDataFirst'));
      return;
    }

    const ids = toIds(options.selectedRowKeys.value);
    if (ids.some((id) => id === 1)) {
      message.warning($t('common.superAdminCannotRecover'));
      return;
    }

    try {
      await recoveryUser(ids);
      message.success($t('common.recoverySuccess'));
      options.clearSelectedRowKeys();
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.batchRecoveryFailed'));
    }
  }

  async function handleStatusChange(row: UserActionRow, checked: boolean) {
    if (isSuperAdmin(row)) {
      message.warning($t('common.superAdminCannotDisable'));
      return;
    }

    const status = checked ? 1 : 2;
    try {
      await changeUserStatus({ id: row.id, status });
      message.success($t('common.updateStatusSuccess'));
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.updateStatusFailed'));
    }
  }

  async function handleResetPassword(row: UserActionRow) {
    try {
      await resetPassword({ id: row.id });
      message.success($t('common.passwordResetSuccess'));
    } catch (error) {
      logger.error(error);
      message.error($t('common.passwordResetFailed'));
    }
  }

  async function handleClearCache(row: UserActionRow) {
    try {
      await clearUserCache({ id: row.id });
      message.success($t('common.clearCacheSuccess2'));
    } catch (error) {
      logger.error(error);
      message.error($t('common.clearCacheFailed2'));
    }
  }

  function triggerImport() {
    importInputRef.value?.click();
  }

  async function handleImportChange(event: Event) {
    const input = event.target as HTMLInputElement;
    const file = input.files?.[0];
    if (!file) return false;

    importLoading.value = true;
    try {
      await importUserFile(file);
      message.success($t('common.importSuccess'));
      options.fetchTableData();
      return true;
    } catch (error) {
      logger.error(error);
      message.error($t('common.importFailed'));
      return false;
    } finally {
      importLoading.value = false;
      input.value = '';
    }
  }

  async function handleExport() {
    exportLoading.value = true;
    try {
      const response = await exportUserList(options.buildRequestParams(false));
      downloadResponseBlob(response, `${$t('system.user.systemUser')}.xlsx`);
      message.success($t('common.exportSuccess'));
    } catch (error) {
      logger.error(error);
      message.error($t('common.exportFailed'));
    } finally {
      exportLoading.value = false;
    }
  }

  async function handleDownloadTemplate() {
    templateLoading.value = true;
    try {
      const response = await downloadUserImportTemplate();
      downloadResponseBlob(response, `${$t('common.importTemplate')}.xlsx`);
      message.success($t('common.templateDownloadSuccess'));
    } catch (error) {
      logger.error(error);
      message.error($t('common.templateDownloadFailed'));
    } finally {
      templateLoading.value = false;
    }
  }

  function openSetHomePageDialog(row: UserActionRow) {
    if (isSuperAdmin(row)) {
      message.warning($t('common.superAdminCannotSetHome'));
      return;
    }

    selectedHomePageUserId.value = Number(row.id);
    selectedHomePage.value = normalizeDashboardValue(row.dashboard);
    setHomePageVisible.value = true;
  }

  async function handleSetHomePage() {
    if (!selectedHomePageUserId.value) {
      message.warning($t('common.userInfoInvalid'));
      return;
    }

    const dashboard = normalizeDashboardValue(selectedHomePage.value);
    if (!dashboard) {
      message.warning($t('common.selectUserHome'));
      return;
    }

    setHomePageLoading.value = true;
    try {
      await setHomePage({
        dashboard,
        id: selectedHomePageUserId.value,
      });
      message.success($t('common.setHomeSuccess'));
      setHomePageVisible.value = false;
      options.fetchTableData();
    } catch (error) {
      logger.error(error);
      message.error($t('common.setHomeFailed'));
    } finally {
      setHomePageLoading.value = false;
    }
  }

  function handleActionDropdownClick(
    data: Pick<UserActionDropdownItem, 'value'>,
    row: UserActionRow,
  ) {
    if (data.value === 'reset_password') {
      const dialogInstance = dialog.confirm({
        body: $t('common.confirmResetPassword'),
        header: $t('common.prompt'),
        onClose: () => dialogInstance.hide(),
        onConfirm: () => {
          void handleResetPassword(row);
          dialogInstance.hide();
        },
      });
      return;
    }

    if (data.value === 'clear_cache') {
      const dialogInstance = dialog.confirm({
        body: $t('common.confirmUpdateCache'),
        header: $t('common.prompt'),
        onClose: () => dialogInstance.hide(),
        onConfirm: () => {
          void handleClearCache(row);
          dialogInstance.hide();
        },
      });
      return;
    }

    if (data.value === 'set_homepage') {
      openSetHomePageDialog(row);
    }
  }

  return {
    exportLoading,
    handleActionDropdownClick,
    handleBatchDelete,
    handleBatchRecovery,
    handleClearCache,
    handleDelete,
    handleDownloadTemplate,
    handleExport,
    handleImportChange,
    handleRecovery,
    handleSetHomePage,
    handleStatusChange,
    importInputRef,
    importLoading,
    isSuperAdmin,
    selectedHomePage,
    setHomePageLoading,
    setHomePageVisible,
    templateLoading,
    triggerImport,
  };
}
