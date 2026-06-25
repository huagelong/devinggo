<script lang="ts" setup>
import { nextTick, onMounted, onUnmounted, ref, watch } from 'vue';

import { defaultKeymap, indentWithTab } from '@codemirror/commands';
import { json } from '@codemirror/lang-json';
import { EditorState } from '@codemirror/state';
import { oneDark } from '@codemirror/theme-one-dark';
import { EditorView, keymap, lineNumbers } from '@codemirror/view';

const props = defineProps<{
  modelValue?: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const editorRef = ref<HTMLDivElement>();
let view: EditorView | null = null;

function createEditor() {
  if (!editorRef.value) return;

  const startState = EditorState.create({
    doc: props.modelValue || '',
    extensions: [
      lineNumbers(),
      json(),
      oneDark,
      keymap.of([...defaultKeymap, indentWithTab]),
      EditorView.updateListener.of((update) => {
        if (update.docChanged) {
          emit('update:modelValue', update.state.doc.toString());
        }
      }),
      EditorView.theme({
        '&': {
          fontSize: '14px',
          height: '180px',
        },
        '.cm-content': {
          fontFamily: "'Consolas', 'Monaco', 'Courier New', monospace",
        },
      }),
    ],
  });

  view = new EditorView({
    state: startState,
    parent: editorRef.value,
  });
}

watch(
  () => props.modelValue,
  (newValue) => {
    if (view && newValue !== view.state.doc.toString()) {
      view.dispatch({
        changes: { from: 0, to: view.state.doc.length, insert: newValue || '' },
      });
    }
  },
);

onMounted(() => {
  nextTick(() => {
    setTimeout(() => {
      createEditor();
    }, 300);
  });
});

onUnmounted(() => {
  view?.destroy();
  view = null;
});
</script>

<template>
  <div ref="editorRef" class="code-editor-wrapper"></div>
</template>

<style scoped>
.code-editor-wrapper {
  display: flex;
  flex-direction: column;
  border: 1px solid #d9d9d9;
  border-radius: 4px;
  overflow: hidden;
  width: 100%;
  min-height: 180px;
  height: 180px;
}

.code-editor-wrapper :deep(.cm-editor) {
  border-radius: 4px;
  width: 100% !important;
  height: 100% !important;
}

.code-editor-wrapper :deep(.cm-focused) {
  outline: none;
}

.code-editor-wrapper :deep(.cm-scroller) {
  min-height: 180px;
}
</style>
