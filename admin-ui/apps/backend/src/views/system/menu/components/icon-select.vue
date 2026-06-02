<script lang="ts" setup>
import { computed, ref } from 'vue';

import { $t } from '@vben/locales';

import {
  SearchIcon,
} from 'tdesign-icons-vue-next';
import {
  Input,
  Popup,
} from 'tdesign-vue-next';

// Common TDesign icons used in admin menus
const iconList = [
  'app',
  'dashboard',
  'user-circle',
  'lock-on',
  'menu-fold',
  'setting',
  'file',
  'folder',
  'browse',
  'edit',
  'delete',
  'add',
  'search',
  'filter',
  'download',
  'upload',
  'link',
  'image',
  'video',
  'notification',
  'chat',
  'mail',
  'calendar',
  'time',
  'chart-bar',
  'chart-line',
  'chart-pie',
  'server',
  'code',
  'bug',
  'tools',
  'precise-monitor',
  'root-list',
  'assignment',
  'save',
  'print',
  'share',
  'star',
  'heart',
  'thumbup',
  'home',
  'layers',
  'toggle-left',
  'map',
  'map-information',
  'logo-wrench',
  'control-platform',
  ' rolled-back',
];

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
  if (!searchKeyword.value) return iconList;
  const kw = searchKeyword.value.toLowerCase();
  return iconList.filter((name) => name.includes(kw));
});

function handleSelect(name: string) {
  emit('update:modelValue', name);
  popupVisible.value = false;
}

function handleInput(val: string) {
  emit('update:modelValue', val);
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
        @input="handleInput"
        @clear="handleInput('')"
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
            class="flex h-9 cursor-pointer flex-col items-center justify-center rounded hover:bg-blue-50"
            :title="name"
            @click="handleSelect(name)"
          >
            <i :class="`i-lucide:${name}`" class="text-lg" />
          </div>
        </div>
        <div v-if="filteredIcons.length === 0" class="py-4 text-center text-xs text-gray-400">
          {{ $t('common.noMatch') }}
        </div>
      </div>
    </template>
  </Popup>
</template>
