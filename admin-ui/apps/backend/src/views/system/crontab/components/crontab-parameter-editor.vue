<script lang="ts" setup>
import type { CrontabParameterFormValue, KeyValueRow } from '../utils/crontab-parameter';

import { computed, ref, watch } from 'vue';

import { DeleteIcon, PlusIcon } from 'tdesign-icons-vue-next';
import { Button, Input, Select, Switch, Textarea } from 'tdesign-vue-next';

import {
  buildCrontabParameter,
  parseCrontabParameter,
} from '../utils/crontab-parameter';

const URL_TASK_TYPE = 3;

const props = defineProps<{
  modelValue?: string;
  taskType?: number;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const methodOptions = [
  { label: 'GET', value: 'get' },
  { label: 'POST', value: 'post' },
  { label: 'PUT', value: 'put' },
  { label: 'DELETE', value: 'delete' },
];

const mode = computed(() => (props.taskType === URL_TASK_TYPE ? 'url' : 'common'));
const parameterValue = ref<CrontabParameterFormValue>(
  parseCrontabParameter(props.modelValue, mode.value),
);

watch(
  () => [props.modelValue, mode.value] as const,
  ([modelValue, currentMode]) => {
    parameterValue.value = parseCrontabParameter(modelValue, currentMode);
  },
);

watch(
  parameterValue,
  (value) => {
    emit('update:modelValue', buildCrontabParameter(value));
  },
  { deep: true },
);

function addRow(rows: KeyValueRow[]) {
  rows.push({ key: '', value: '' });
}

function removeRow(rows: KeyValueRow[], index: number) {
  rows.splice(index, 1);
  if (rows.length === 0) {
    rows.push({ key: '', value: '' });
  }
}

function handleRawJsonChange(useRawJson: boolean) {
  if (useRawJson) {
    parameterValue.value.rawJson = buildCrontabParameter({
      ...parameterValue.value,
      useRawJson: false,
    });
  }
  parameterValue.value.useRawJson = useRawJson;
}
</script>

<template>
  <div class="crontab-parameter-editor">
    <div class="editor-toolbar">
      <span class="editor-tip">
        {{ mode === 'url' ? '填写请求信息，系统会自动生成 JSON 参数。' : '填写参数名和参数值，系统会自动生成 JSON 参数。' }}
      </span>
      <label class="advanced-switch">
        <span>高级 JSON</span>
        <Switch
          :model-value="parameterValue.useRawJson"
          @update:model-value="handleRawJsonChange"
        />
      </label>
    </div>

    <Textarea
      v-if="parameterValue.useRawJson"
      v-model="parameterValue.rawJson"
      placeholder='例如：{"name":"name1","value":"shuju1"}'
      :autosize="{ minRows: 6, maxRows: 10 }"
    />

    <template v-else-if="mode === 'url'">
      <div class="url-grid">
        <Input
          v-model="parameterValue.url.url"
          class="url-input"
          placeholder="请输入请求地址，例如 https://example.com/api"
        />
        <Select
          v-model="parameterValue.url.method"
          :options="methodOptions"
          class="method-select"
        />
      </div>

      <div class="group-title">Header</div>
      <div
        v-for="(item, index) in parameterValue.url.headers"
        :key="`header-${index}`"
        class="kv-row"
      >
        <Input v-model="item.key" placeholder="名称，例如 Authorization" />
        <Input v-model="item.value" placeholder="值，例如 Bearer token" />
        <Button
          shape="square"
          theme="danger"
          variant="outline"
          @click="removeRow(parameterValue.url.headers, index)"
        >
          <template #icon><DeleteIcon /></template>
        </Button>
      </div>
      <Button variant="outline" @click="addRow(parameterValue.url.headers)">
        <template #icon><PlusIcon /></template>
        新增 Header
      </Button>

      <div class="group-title">Params</div>
      <div
        v-for="(item, index) in parameterValue.url.params"
        :key="`param-${index}`"
        class="kv-row"
      >
        <Input v-model="item.key" placeholder="参数名，例如 page" />
        <Input v-model="item.value" placeholder="参数值，例如 1" />
        <Button
          shape="square"
          theme="danger"
          variant="outline"
          @click="removeRow(parameterValue.url.params, index)"
        >
          <template #icon><DeleteIcon /></template>
        </Button>
      </div>
      <Button variant="outline" @click="addRow(parameterValue.url.params)">
        <template #icon><PlusIcon /></template>
        新增 Params
      </Button>
    </template>

    <template v-else>
      <div
        v-for="(item, index) in parameterValue.commonParams"
        :key="index"
        class="kv-row"
      >
        <Input v-model="item.key" placeholder="参数名，例如 name" />
        <Input v-model="item.value" placeholder="参数值，例如 name1" />
        <Button
          shape="square"
          theme="danger"
          variant="outline"
          @click="removeRow(parameterValue.commonParams, index)"
        >
          <template #icon><DeleteIcon /></template>
        </Button>
      </div>
      <Button variant="outline" @click="addRow(parameterValue.commonParams)">
        <template #icon><PlusIcon /></template>
        新增参数
      </Button>
    </template>
  </div>
</template>

<style scoped>
.crontab-parameter-editor {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 100%;
}

.editor-toolbar {
  display: flex;
  gap: 12px;
  align-items: center;
  justify-content: space-between;
}

.editor-tip {
  color: var(--td-text-color-secondary);
  font-size: 12px;
  line-height: 20px;
}

.advanced-switch {
  display: inline-flex;
  flex-shrink: 0;
  gap: 8px;
  align-items: center;
  color: var(--td-text-color-secondary);
  font-size: 12px;
}

.url-grid {
  display: grid;
  grid-template-columns: minmax(0, 1fr) 120px;
  gap: 8px;
}

.kv-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(0, 1fr) auto;
  gap: 8px;
  align-items: center;
}

.group-title {
  margin-top: 4px;
  color: var(--td-text-color-primary);
  font-size: 13px;
  font-weight: 500;
}

@media (max-width: 767px) {
  .editor-toolbar,
  .url-grid,
  .kv-row {
    grid-template-columns: minmax(0, 1fr);
  }

  .editor-toolbar {
    align-items: flex-start;
  }

  .kv-row :deep(.t-button) {
    width: 100%;
  }
}
</style>
