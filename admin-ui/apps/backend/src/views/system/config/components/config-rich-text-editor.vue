<script lang="ts" setup>
import { computed } from 'vue';

import { $t } from '@vben/locales';
import { usePreferences } from '@vben/preferences';

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
import 'tinymce/themes/silver';

void tinymce;

const { isDark } = usePreferences();

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

const editorKey = computed(() => `tinymce-${isDark.value ? 'dark' : 'light'}`);

function applyEditorTheme(editor: any) {
  setTimeout(() => {
    const iframe = editor.getContainer()?.querySelector('iframe');
    if (!iframe) return;
    const doc = iframe.contentDocument || iframe.contentWindow?.document;
    if (!doc) return;

    const body = doc.body;
    const html = doc.documentElement;
    if (!body || !html) return;

    if (isDark.value) {
      html.style.backgroundColor = '#1e1e1e';
      body.style.backgroundColor = '#1e1e1e';
      body.style.color = '#e0e0e0';
      body.setAttribute('data-mce-theme', 'dark');
    } else {
      html.style.backgroundColor = '#ffffff';
      body.style.backgroundColor = '#ffffff';
      body.style.color = '#333333';
      body.removeAttribute('data-mce-theme');
    }
  }, 100);
}

const initOptions = computed(() => ({
  height: 260,
  menubar: false,
  branding: false,
  license_key: 'gpl',
  placeholder: $t('common.enterContent'),
  plugins: 'advlist autolink lists link charmap preview code fullscreen table wordcount',
  toolbar:
    'undo redo | blocks | bold italic underline | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | link table | removeformat code fullscreen',
  skin: isDark.value ? 'oxide-dark' : 'oxide',
  content_style: isDark.value
    ? 'body.mce-content-body { background-color: #1e1e1e !important; color: #e0e0e0 !important; }'
    : '',
  setup: (editor: any) => {
    editor.on('init', () => {
      applyEditorTheme(editor);
    });
  },
  init_instance_callback: (editor: any) => {
    applyEditorTheme(editor);
  },
}));
</script>

<template>
  <div class="config-rich-text-editor">
    <Editor
      :key="editorKey"
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
