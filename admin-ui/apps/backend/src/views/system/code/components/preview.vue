<script lang="ts" setup>
import type { PreviewCodeRow } from '../model';

import { ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { message } from '#/adapter/tdesign';
import { logger } from '#/utils/logger';

import { previewCode } from '#/api/system/generate';
import { CodeIcon } from 'tdesign-icons-vue-next';

import { Button, TabPanel, Tabs } from 'tdesign-vue-next';

const emit = defineEmits<{
  success: [];
}>();

const loading = ref(false);
const previewList = ref<PreviewCodeRow[]>([]);
const activeTab = ref('0');

async function open(id: number) {
  loading.value = true;
  previewList.value = [];
  try {
    const response = await previewCode(id);
    previewList.value = response.data || [];
    if (previewList.value.length > 0) {
      activeTab.value = '0';
    }
    modalApi.setState({ title: $t('system.code.previewTitle') });
    modalApi.open();
  } catch (error) {
    logger.error(error);
    message.error($t('common.previewFailed'));
  } finally {
    loading.value = false;
  }
}

function handleCopy(code: string) {
  navigator.clipboard.writeText(code).then(() => {
    message.success($t('common.copySuccess'));
  }).catch(() => {
    message.error($t('common.copyFailed'));
  });
}

const [Modal, modalApi] = useVbenModal({
  footer: false,
  class: 'w-[1000px]',
});

defineExpose({ open });
</script>

<template>
  <Modal>
    <div class="flex flex-col gap-3">
      <div v-if="loading" class="flex items-center justify-center py-8">
        {{ $t('common.loading') }}
      </div>

      <div v-else-if="previewList.length === 0" class="py-8 text-center text-gray-500">
        {{ $t('common.noPreviewData') }}
      </div>

      <template v-else>
        <Tabs v-model:value="activeTab">
          <TabPanel
            v-for="(item, index) in previewList"
            :key="item.name"
            :value="String(index)"
            :label="item.tab_name"
          >
            <div class="flex justify-end mb-2">
              <Button
                size="small"
                @click="handleCopy(item.code)"
              >
                <template #icon><CodeIcon /></template>
                {{ $t('common.copyCode') }}
              </Button>
            </div>
            <pre
              class="max-h-[500px] overflow-auto rounded bg-gray-900 p-4 text-sm text-gray-100"
            >{{ item.code }}</pre>
          </TabPanel>
        </Tabs>
      </template>
    </div>
  </Modal>
</template>
