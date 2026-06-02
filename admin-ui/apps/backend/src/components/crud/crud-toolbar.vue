<script lang="ts" setup>
import type { OptionItem } from '#/types/common';

import { computed } from 'vue';

import { $t } from '@vben/locales';

import {
  DeleteIcon,
  RefreshIcon,
  RollbackIcon,
  SettingIcon,
} from 'tdesign-icons-vue-next';
import { Button, Checkbox, CheckboxGroup, Popup, Tooltip } from 'tdesign-vue-next';

type ColumnOption = OptionItem<string>;

interface Props {
  columnOptions?: ColumnOption[];
  isRecycleBin?: boolean;
  modelValue?: string[];
}

const props = withDefaults(defineProps<Props>(), {
  columnOptions: () => [],
  isRecycleBin: false,
  modelValue: () => [],
});

const emit = defineEmits<{
  (e: 'refresh'): void;
  (e: 'toggle-recycle'): void;
  (e: 'update:modelValue', value: string[]): void;
}>();

const selectedColumns = computed({
  get: () => props.modelValue || [],
  set: (value: string[]) => {
    emit('update:modelValue', value);
  },
});

const allColumnKeys = computed(() =>
  (props.columnOptions || []).map((item) => item.value),
);

const isAllSelected = computed(
  () =>
    selectedColumns.value.length > 0 &&
    selectedColumns.value.length === allColumnKeys.value.length,
);

const isIndeterminate = computed(
  () =>
    selectedColumns.value.length > 0 &&
    selectedColumns.value.length < allColumnKeys.value.length,
);

function toggleSelectAll(checked: boolean) {
  emit('update:modelValue', checked ? [...allColumnKeys.value] : []);
}

function handleSelectAllChange(value: unknown) {
  toggleSelectAll(Boolean(value));
}
</script>

<template>
  <div class="flex items-center gap-2">
    <Tooltip :content="$t('common.refresh')">
      <Button shape="square" variant="outline" @click="emit('refresh')">
        <template #icon><RefreshIcon /></template>
      </Button>
    </Tooltip>

    <Tooltip :content="isRecycleBin ? $t('common.backToList') : $t('common.viewRecycleBin')">
      <Button shape="square" variant="outline" @click="emit('toggle-recycle')">
        <template #icon>
          <RollbackIcon v-if="isRecycleBin" />
          <DeleteIcon v-else />
        </template>
      </Button>
    </Tooltip>

    <Tooltip :content="$t('common.columnConfig')">
      <Popup placement="bottom-right" trigger="click">
        <Button shape="square" variant="outline">
          <template #icon><SettingIcon /></template>
        </Button>

        <template #content>
          <div class="min-w-[140px] p-3">
            <Checkbox
              :checked="isAllSelected"
              :indeterminate="isIndeterminate"
              @change="handleSelectAllChange"
            >
              {{ $t('common.selectAll') }}
            </Checkbox>
            <div class="my-2 h-px bg-gray-100" />
            <CheckboxGroup
              v-model="selectedColumns"
              :options="columnOptions"
              layout="vertical"
            />
          </div>
        </template>
      </Popup>
    </Tooltip>
  </div>
</template>
