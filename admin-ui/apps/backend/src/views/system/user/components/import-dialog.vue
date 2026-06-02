<script lang="ts" setup>
import { $t } from '@vben/locales';

import { DownloadIcon, UploadIcon } from 'tdesign-icons-vue-next';
import { Button, Dialog } from 'tdesign-vue-next';

interface Props {
  visible: boolean;
  importLoading: boolean;
  templateLoading: boolean;
}

interface Emits {
  (e: 'update:visible', value: boolean): void;
  (e: 'download-template'): void;
  (e: 'import'): void;
  (e: 'cancel'): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

function handleUpdateVisible(value: boolean) {
  emit('update:visible', value);
}

function handleDownloadTemplate() {
  emit('download-template');
}

function handleImport() {
  emit('import');
}

function handleCancel() {
  emit('cancel');
}
</script>

<template>
  <Dialog
    :visible="visible"
    width="420px"
    :header="$t('common.import')"
    destroy-on-close
    :close-on-overlay-click="true"
    @update:visible="handleUpdateVisible"
  >
    <div class="flex flex-col gap-4">
      <p class="text-sm text-text-2">
        {{ $t('common.importDialogDescription') }}
      </p>
      <div class="flex flex-col gap-3">
        <Button
          variant="outline"
          :loading="templateLoading"
          @click="handleDownloadTemplate"
        >
          <template #icon><DownloadIcon /></template>
          {{ $t('common.importTemplate') }}
        </Button>
        <Button
          theme="primary"
          :loading="importLoading"
          @click="handleImport"
        >
          <template #icon><UploadIcon /></template>
          {{ $t('common.import') }}
        </Button>
      </div>
    </div>
    <template #footer>
      <div class="flex justify-end">
        <Button theme="default" @click="handleCancel">
          {{ $t('common.cancel') }}
        </Button>
      </div>
    </template>
  </Dialog>
</template>