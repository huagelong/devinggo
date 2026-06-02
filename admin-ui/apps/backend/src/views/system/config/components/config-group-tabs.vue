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
  newFormMap[groupId] = { ...newFormMap[groupId], [key]: value };
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
          label-width="200px" layout="inline" colon
        >
          <div class="config-form-list">
            <FormItem
              v-for="field in configFieldsMap[group.id] ?? []"
              :key="field.id"
              :label="field.label"
              :name="field.key"
              class="config-field-item"
            >
              <ConfigFieldRenderer
                :model-value="configFormMap[group.id]![field.key]"
                :field="field"
                @update:model-value="(val: unknown) => updateFormValue(group.id, field.key, val)"
              />
              <div
                v-if="field.remark"
                class="config-field-remark"
              >
                {{ field.remark }}
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
  gap: 6px;
}

.config-edit-form {
  display: block;
}

.config-field-item {
  margin-bottom: 0 !important;
  padding: 12px 8px;
  border-bottom: 1px dashed #f0f2f5;
}

.config-field-remark {
  margin-left: 10px;
  font-size: 13px;
  color: #8a94a6;
  line-height: 1.5;
}

.config-field-item:last-child {
  border-bottom: none;
}

.config-edit-form :deep(.t-form__item) {
  width: 100%;
  margin-right: 0;
}

.config-edit-form :deep(.t-form__label) {
  padding-right: 12px;
  color: var(--td-text-color-secondary, #6b7280);
  font-weight: 500;
  line-height: 40px;
}

.config-edit-form :deep(.t-form__controls) {
  flex: 1;
  min-width: 0;
}

.config-edit-form :deep(.t-input),
.config-edit-form :deep(.t-textarea__inner),
.config-edit-form :deep(.t-select),
.config-edit-form :deep(.t-select__wrap) {
  border-radius: 8px;
}

.config-save-bar {
  margin-top: 16px;
  display: flex;
  justify-content: flex-end;
}

@media (max-width: 1023px) {
  .config-content {
    min-height: auto;
  }
}
</style>