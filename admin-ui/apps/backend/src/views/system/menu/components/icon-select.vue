<script lang="ts" setup>
import { computed, ref } from 'vue';

import { menuIconNames } from '@vben/icons';
import { $t } from '@vben/locales';

import {
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Input,
  Popup,
} from 'tdesign-vue-next';

const props = defineProps<{
  modelValue?: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const popupVisible = ref(false);
const searchKeyword = ref('');

const filteredIcons = computed(() => {
  if (!searchKeyword.value) return menuIconNames;
  const kw = searchKeyword.value.toLowerCase();
  return menuIconNames.filter((name) => name.includes(kw));
});

function handleSelect(name: string) {
  emit('update:modelValue', name);
  popupVisible.value = false;
}

</script>

<template>
  <Popup
    v-model="popupVisible"
    trigger="click"
    placement="bottom-left"
    :overlay-style="{ width: '320px' }"
  >
    <div class="flex items-center gap-2">
      <Input
        :value="modelValue"
        :placeholder="placeholder || $t('ui.placeholder.input')"
        clearable
        readonly
        @clear="emit('update:modelValue', '')"
      />
      <span
        v-if="modelValue"
        class="flex h-8 w-8 shrink-0 items-center justify-center rounded border border-gray-200"
      >
        <i :class="`i-lucide:${modelValue}`" class="text-base" />
      </span>
    </div>
    <template #content>
      <div class="flex flex-col gap-2 p-2">
        <Input
          v-model="searchKeyword"
          :placeholder="$t('ui.placeholder.input')"
          size="small"
          clearable
        >
          <template #prefix><SearchIcon /></template>
        </Input>
        <div class="grid max-h-48 grid-cols-6 gap-1 overflow-y-auto">
          <div
            v-for="name in filteredIcons"
            :key="name"
            class="flex h-9 cursor-pointer flex-col items-center justify-center rounded hover:bg-primary/10"
            :title="name"
            @click="handleSelect(name)"
          >
            <i :class="`i-lucide:${name}`" class="text-lg" />
          </div>
        </div>
        <div v-if="filteredIcons.length === 0" class="py-4 text-center text-xs text-muted-foreground/80">
          {{ $t('common.noMatch') }}
        </div>
      </div>
    </template>
  </Popup>
</template>
