<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';

import Editor from '@tinymce/tinymce-vue';
import tinymce from 'tinymce';

import 'tinymce/icons/default';
import 'tinymce/models/dom';
import 'tinymce/plugins/advlist';
import 'tinymce/plugins/autolink';
import 'tinymce/plugins/charmap';
import 'tinymce/plugins/code';
import 'tinymce/plugins/fullscreen';
import 'tinymce/plugins/link';
import 'tinymce/plugins/lists';
import 'tinymce/plugins/preview';
import 'tinymce/plugins/table';
import 'tinymce/plugins/wordcount';
import 'tinymce/skins/content/default/content.min.css';
import 'tinymce/skins/ui/oxide/skin.min.css';
import 'tinymce/themes/silver';

void tinymce;

const props = defineProps<{
  modelValue?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const innerValue = computed({
  get: () =>
    typeof props.modelValue === 'string'
      ? props.modelValue
      : String(props.modelValue ?? ''),
  set: (value: string) => {
    emit('update:modelValue', value ?? '');
  },
});

const initOptions = {
  height: 260,
  menubar: false,
  branding: false,
  license_key: 'gpl',
  placeholder: $t('common.enterContent'),
  plugins: 'advlist autolink lists link charmap preview code fullscreen table wordcount',
  toolbar:
    'undo redo | blocks | bold italic underline | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link table | removeformat code fullscreen',
  content_style: 'body { font-family: -apple-system, BlinkMacSystemFont, Segoe UI, PingFang SC, Microsoft YaHei, sans-serif; font-size: 14px; }',
};
</script>

<template>
  <div class="config-rich-text-editor">
    <Editor
      v-model="innerValue"
      :init="initOptions"
      output-format="html"
    />
  </div>
</template>

<style scoped>
.config-rich-text-editor {
  width: 100%;
}

.config-rich-text-editor :deep(.tox-tinymce) {
  border-radius: 8px;
  border-color: var(--td-component-border, #dcdcdc);
}
</style>
