<script lang="ts" setup>
import { $t } from '@vben/locales';

import { InfoCircleIcon } from 'tdesign-icons-vue-next';
import { Button, Form, FormItem, Tabs, TabPanel } from 'tdesign-vue-next';

import type { ConfigFormModel, ConfigGroup, ConfigFieldMeta } from '../model';
import ConfigFieldRenderer from './config-field-renderer.vue';

interface Props {
  groupList: ConfigGroup[];
  activeGroupKey: number | undefined;
  groupLoading: boolean;
  configFieldsMap: Record<number, ConfigFieldMeta[]>;
  configFormMap: Record<number, ConfigFormModel>;
  hasGroups: boolean;
}

interface Emits {
  (e: 'tab-change', value: string | number): void;
  (e: 'submit', groupId: number): void;
  (e: 'update:configFormMap', value: Record<number, ConfigFormModel>): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

function handleTabChange(value: string | number) {
  emit('tab-change', value);
}

function handleSubmit(groupId: number) {
  emit('submit', groupId);
}

function updateFormValue(groupId: number, key: string, value: unknown) {
  const newFormMap = { ...props.configFormMap };
  if (!newFormMap[groupId]) {
    newFormMap[groupId] = {};
  }
    newFormMap[groupId] = { ...newFormMap[groupId]!, [key]: value } as ConfigFormModel;
  emit('update:configFormMap', newFormMap);
}
</script>

<template>
  <div v-if="hasGroups" class="config-content">
    <Tabs
      :value="activeGroupKey"
      :loading="groupLoading"
      class="config-tabs"
      @change="handleTabChange"
    >
      <TabPanel
        v-for="group in groupList"
        :key="group.id"
        :value="group.id"
        :label="group.name"
      >
        <Form
          v-if="configFormMap[group.id]"
          :data="configFormMap[group.id]!"
          class="config-edit-form"
          label-width="auto" colon
        >
          <div class="config-form-list">
            <FormItem
              v-for="field in configFieldsMap[group.id] ?? []"
              :key="field.id"
              :label="field.label"
              :name="field.key"
              class="config-field-item"
            >
              <div class="field-content">
                <div style="width: 100%">
                  <ConfigFieldRenderer
                    :model-value="configFormMap[group.id]![field.key]"
                    :field="field"
                    @update:model-value="(val: unknown) => updateFormValue(group.id, field.key, val)"
                  />
                </div>
                <div
                  v-if="field.remark"
                  class="config-field-remark"
                >
                  {{ field.remark }}
                </div>
              </div>
            </FormItem>
          </div>
          <div class="config-save-bar">
            <Button theme="primary" @click="handleSubmit(group.id)">
              {{ $t('system.config.saveConfig') }}
            </Button>
          </div>
        </Form>
        <div
          v-else
          class="flex min-h-[200px] items-center justify-center text-gray-500"
        >
          <InfoCircleIcon class="mr-2" /> {{ $t('system.config.configLoading') }}
        </div>
      </TabPanel>
    </Tabs>
  </div>
  <div
    v-else
    class="flex min-h-[200px] items-center justify-center text-gray-500"
  >
    {{ $t('system.config.noConfigGroup') }}
  </div>
</template>

<style scoped>
.config-content {
  min-height: 420px;
}

.config-form-list {
  display: flex;
  flex-direction: column;
}

.config-edit-form {
  display: block;
}

.config-field-item {
  margin-bottom: 0 !important;
  padding: 18px 16px;
  border-bottom: 1px solid #f0f2f5;
  transition: background-color 0.2s ease;
}

.config-field-item:hover {
  background-color: #fafbfc;
}

.config-field-remark {
  margin-top: 6px;
  margin-left: 0;
  font-size: 13px;
  color: #9ca3af;
  line-height: 1.5;
}

.config-field-item:last-child {
  border-bottom: none;
}

.config-edit-form :deep(.t-form__item) {
  width: 100%;
  margin-right: 0;
  display: flex !important;
  flex-direction: row !important;
  align-items: flex-start;
}

.config-edit-form :deep(.t-form__label) {
  min-width: 160px;
  flex-shrink: 0;
  padding-right: 16px;
  color: var(--td-text-color-primary, #374151);
  font-weight: 500;
  line-height: 40px;
  font-size: 14px;
}

.config-edit-form :deep(.t-form__controls) {
  flex: 1;
  min-width: 0;
  max-width: 100%;
}

.config-edit-form :deep(.t-form__controls-content) {
  display: block !important;
  width: 100%;
}

.field-content {
  display: block;
  width: 100%;
}

.config-edit-form :deep(.t-input),
.config-edit-form :deep(.t-textarea),
.config-edit-form :deep(.t-textarea__inner),
.config-edit-form :deep(.t-select),
.config-edit-form :deep(.t-select__wrap) {
  border-radius: 8px;
  width: 100% !important;
  max-width: 100%;
}

.config-edit-form :deep(.t-radio-group),
.config-edit-form :deep(.t-checkbox-group) {
  width: 100%;
}

.config-save-bar {
  margin-top: 0;
  padding: 16px 20px;
  display: flex;
  justify-content: flex-end;
  border-top: 1px solid #f0f2f5;
  background: linear-gradient(to bottom, #fafbfc, #ffffff);
  border-radius: 0 0 10px 10px;
}

.config-tabs :deep(.t-tabs__content) {
  margin-top: 4px;
}

@media (max-width: 1023px) {
  .config-content {
    min-height: auto;
  }

  .config-field-item {
    padding: 14px 12px;
  }

  .config-save-bar {
    padding: 12px 16px;
  }
}
</style>