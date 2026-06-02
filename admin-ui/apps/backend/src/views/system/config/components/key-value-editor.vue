<script lang="ts" setup>
import { DeleteIcon, PlusIcon } from 'tdesign-icons-vue-next';
import { Button, Input, Space } from 'tdesign-vue-next';
import { computed } from 'vue';

import { $t } from '@vben/locales';

interface KeyValueItem {
  key: string;
  value: string;
}

const props = defineProps<{
  modelValue?: KeyValueItem[];
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: KeyValueItem[]): void;
}>();

const items = computed({
  get: () => props.modelValue ?? [{ key: '', value: '' }],
  set: (value: KeyValueItem[]) => emit('update:modelValue', value),
});

function handleAdd() {
  items.value = [...items.value, { key: '', value: '' }];
}

function handleRemove(index: number) {
  const next = [...items.value];
  next.splice(index, 1);
  items.value = next.length > 0 ? next : [{ key: '', value: '' }];
}

function handleUpdate(index: number, field: keyof KeyValueItem, value: string) {
  const next = items.value.map((item, idx) =>
    idx === index ? { ...item, [field]: value } : item,
  );
  items.value = next;
}
</script>

<template>
  <div class="key-value-editor flex flex-col gap-3">
    <div
      v-for="(item, index) in items"
      :key="index"
      class="kv-row"
    >
      <Input
        v-model="item.key"
        :placeholder="$t('common.key')"
        class="kv-input"
        @change="(val) => handleUpdate(index, 'key', val as string)"
      />
      <Input
        v-model="item.value"
        :placeholder="$t('common.value')"
        class="kv-input"
        @change="(val) => handleUpdate(index, 'value', val as string)"
      />
      <Button
        variant="outline"
        shape="square"
        theme="danger"
        @click="handleRemove(index)"
      >
        <template #icon><DeleteIcon /></template>
      </Button>
    </div>
    <Space>
      <Button variant="outline" @click="handleAdd">
        <template #icon><PlusIcon /></template>
        {{ $t('common.addKeyValue') }}
      </Button>
    </Space>
  </div>
</template>

<style scoped>
.kv-row {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(0, 1fr) auto;
  gap: 8px;
  align-items: center;
}

.kv-input {
  width: 100%;
}

@media (max-width: 767px) {
  .kv-row {
    grid-template-columns: minmax(0, 1fr);
  }

  .kv-row :deep(.t-button) {
    width: 100%;
  }
}
</style>
