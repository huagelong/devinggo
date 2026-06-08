<script lang="ts" setup>
import { computed } from 'vue';

const props = defineProps<{
  modelValue?: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const value = computed({
  get: () => props.modelValue || '',
  set: (val) => emit('update:modelValue', val),
});
</script>

<template>
  <div class="code-editor-wrapper">
    <textarea
      v-model="value"
      :placeholder="placeholder"
      class="code-editor"
      spellcheck="false"
    />
  </div>
</template>

<style scoped>
.code-editor-wrapper {
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  overflow: hidden;
}

.code-editor {
  width: 100%;
  min-height: 180px;
  padding: 12px;
  font-family: 'Consolas', 'Monaco', 'Courier New', monospace;
  font-size: 14px;
  line-height: 1.6;
  color: #333;
  background-color: #fafafa;
  border: none;
  outline: none;
  resize: vertical;
}

.code-editor:focus {
  background-color: #fff;
}

.code-editor::placeholder {
  color: #bfbfbf;
}
</style>
