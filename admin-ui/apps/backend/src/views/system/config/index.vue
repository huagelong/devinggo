<script lang="ts" setup>
import type { ConfigApi } from '#/api/system/config';
import type { ConfigFormModel, ConfigGroup } from './model';

import { computed, onMounted, reactive, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import { MessagePlugin } from 'tdesign-vue-next';

import {
  deleteConfigGroup,
  getConfigGroupList,
  getConfigList,
  updateConfigByKeys,
} from '#/api/system/config';

import ConfigAddPanel from './components/config-add-panel.vue';
import ConfigFormModal from './components/config-form-modal.vue';
import ConfigGroupModal from './components/config-group-modal.vue';
import ConfigGroupTabs from './components/config-group-tabs.vue';
import ConfigManageModal from './components/config-manage-modal.vue';
import ConfigToolbar from './components/config-toolbar.vue';
import type { ConfigFieldMeta } from './model';

const configGroupModalRef = ref<{ open: () => void }>();
const configFormModalRef = ref<{ open: () => void }>();
const configManageModalRef = ref<{ open: (groupId: number) => void }>();

const groupList = ref<ConfigGroup[]>([]);
const activeGroupKey = ref<number>();
const groupLoading = ref(false);
const configFieldsMap = reactive<Record<number, ConfigFieldMeta[]>>({});
const configFormMap = reactive<Record<number, ConfigFormModel>>({});

const hasGroups = computed(() => groupList.value.length > 0);

function normalizeOptions(data: unknown): ConfigFieldMeta['config_select_data'] {
  if (Array.isArray(data)) {
    return data as ConfigFieldMeta['config_select_data'];
  }
  if (typeof data === 'string' && data.trim()) {
    try {
      const parsed = JSON.parse(data);
      if (Array.isArray(parsed)) {
        return parsed as ConfigFieldMeta['config_select_data'];
      }
    } catch {
      return undefined;
    }
  }
  return undefined;
}

function normalizeSwitchValues(value: unknown) {
  if (typeof value === 'boolean') {
    return { checked: true, unchecked: false };
  }
  if (typeof value === 'number') {
    return { checked: 1, unchecked: 0 };
  }
  if (typeof value === 'string') {
    const normalized = value.trim().toLowerCase();
    if (normalized === '1' || normalized === '0') {
      return { checked: '1', unchecked: '0' };
    }
    if (normalized === 'true' || normalized === 'false') {
      return { checked: 'true', unchecked: 'false' };
    }
  }
  return { checked: true, unchecked: false };
}

function parseCheckboxValue(rawValue: unknown): string[] {
  if (Array.isArray(rawValue)) {
    return rawValue.map((item) => String(item));
  }
  if (typeof rawValue === 'string') {
    try {
      const parsed = JSON.parse(rawValue);
      if (Array.isArray(parsed)) {
        return parsed.map((item) => String(item));
      }
    } catch {
      if (rawValue.includes(',')) {
        return rawValue.split(',').map((value) => value.trim());
      }
    }
  }
  return [];
}

function parseKeyValueInput(rawValue: unknown) {
  if (Array.isArray(rawValue)) {
    return rawValue;
  }
  if (typeof rawValue === 'string' && rawValue.trim()) {
    try {
      const parsed = JSON.parse(rawValue);
      if (Array.isArray(parsed)) {
        return parsed;
      }
    } catch {
      return [{ key: '', value: rawValue }];
    }
  }
  return [{ key: '', value: '' }];
}

async function fetchGroups() {
  groupLoading.value = true;
  try {
    const list = await getConfigGroupList();
    groupList.value = list;
    if (list.length > 0) {
      const firstGroup = list[0]!;
      activeGroupKey.value = firstGroup.id;
      await fetchGroupConfigs(firstGroup.id);
    } else {
      activeGroupKey.value = undefined;
    }
  } catch (error) {
    logger.error(error);
    message.error($t('common.configGroupLoadFailed'));
  } finally {
    groupLoading.value = false;
  }
}

async function fetchGroupConfigs(groupId: number) {
  try {
    const items = await getConfigList({
      group_id: groupId,
      orderBy: 'sort',
      orderType: 'asc',
    });
    const fields: ConfigFieldMeta[] = [];
    const form: ConfigFormModel = {};
    items.forEach((item) => {
      const field: ConfigFieldMeta = {
        id: item.id,
        input_type: item.input_type,
        key: item.key,
        label: item.name,
        remark: item.remark,
        sort: item.sort,
        config_select_data: normalizeOptions(item.config_select_data),
      };
      let value: ConfigFormModel[string];
      if (item.input_type === 'switch') {
        field.switchValues = normalizeSwitchValues(item.value);
      }
      if (item.input_type === 'checkbox') {
        value = parseCheckboxValue(item.value);
      } else if (item.input_type === 'key-value') {
        value = parseKeyValueInput(item.value);
      } else {
        value = item.value as ConfigFormModel[string];
      }
      form[item.key] = value;
      fields.push(field);
    });
    configFieldsMap[groupId] = fields;
    configFormMap[groupId] = form;
  } catch (error) {
    logger.error(error);
    message.error($t('common.configDataLoadFailed'));
  }
}

function handleTabChange(value: string | number) {
  const id = Number(value);
  if (Number.isNaN(id)) return;
  activeGroupKey.value = id;
  if (!configFieldsMap[id]) {
    void fetchGroupConfigs(id);
  }
}

function formatSubmitValue(
  field: ConfigFieldMeta,
  value: unknown,
): ConfigApi.UpdateByKeysPayload[string] {
  if (
    field.input_type === 'key-value' ||
    Array.isArray(value) ||
    (value !== null && typeof value === 'object')
  ) {
    try {
      return JSON.stringify(value ?? []);
    } catch {
      return JSON.stringify([]);
    }
  }
  return value as ConfigApi.UpdateByKeysPayload[string];
}

async function handleSubmit(groupId: number) {
  const currentForm = configFormMap[groupId];
  if (!currentForm) return;
  const fields = configFieldsMap[groupId] ?? [];
  const payload: ConfigApi.UpdateByKeysPayload = {};
  fields.forEach((field) => {
    payload[field.key] = formatSubmitValue(field, currentForm[field.key]);
  });

  try {
    await updateConfigByKeys(payload);
    MessagePlugin.success($t('common.configUpdateSuccess'));
    await fetchGroupConfigs(groupId);
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.configUpdateFailed'));
  }
}

async function handleDeleteGroup(groupId: number) {
  const group = groupList.value.find((item) => item.id === groupId);
  if (!group) return;
  if (groupId === 1 || groupId === 2) {
    MessagePlugin.info($t('common.defaultGroupCannotDelete'));
    return;
  }
  try {
    await deleteConfigGroup({ id: groupId });
    MessagePlugin.success($t('common.groupDeleteSuccess'));
    await fetchGroups();
  } catch (error) {
    logger.error(error);
    MessagePlugin.error($t('common.groupDeleteFailed'));
  }
}

function handleManage() {
  if (activeGroupKey.value) {
    configManageModalRef.value?.open(Number(activeGroupKey.value));
  }
}

function handleConfigFormMapUpdate(value: Record<number, ConfigFormModel>) {
  Object.assign(configFormMap, value);
}

onMounted(() => {
  void fetchGroups();
});
</script>

<template>
  <Page auto-content-height>
    <div class="flex h-full flex-col gap-5">
      <div class="flex flex-col gap-5 lg:flex-row">
        <div class="config-card flex-1 lg:w-2/3">
          <ConfigToolbar
            :active-group-key="activeGroupKey"
            @add-group="configGroupModalRef?.open()"
            @manage="handleManage"
            @delete-group="activeGroupKey && handleDeleteGroup(activeGroupKey)"
          />

          <ConfigGroupTabs
            :group-list="groupList"
            :active-group-key="activeGroupKey"
            :group-loading="groupLoading"
            :config-fields-map="configFieldsMap"
            :config-form-map="configFormMap"
            :has-groups="hasGroups"
            @tab-change="handleTabChange"
            @submit="handleSubmit"
            @update:config-form-map="handleConfigFormMapUpdate"
          />
        </div>

        <div class="config-card lg:w-1/3">
          <ConfigAddPanel @create="configFormModalRef?.open()" />
        </div>
      </div>
    </div>

    <ConfigGroupModal ref="configGroupModalRef" @success="fetchGroups" />
    <ConfigFormModal ref="configFormModalRef" @success="fetchGroups" />
    <ConfigManageModal ref="configManageModalRef" />
  </Page>
</template>

<style scoped>
.config-card {
  border: 1px solid var(--td-component-border, #e7e7e7);
  border-radius: 10px;
  background: #fff;
  padding: 20px 24px;
  box-shadow: 0 6px 18px rgb(15 23 42 / 4%);
  transition: box-shadow 0.3s ease;
}

.config-card:hover {
  box-shadow: 0 8px 24px rgb(15 23 42 / 6%);
}

@media (max-width: 1023px) {
  .config-card {
    padding: 14px;
  }
}
</style>