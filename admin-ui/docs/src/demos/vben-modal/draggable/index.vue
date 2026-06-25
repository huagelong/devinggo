<script lang="ts" setup>
import { markRaw } from 'vue';

import { useVbenModal, VbenButton } from '@vben/common-ui';

import ExtraModal from './modal.vue';

const [Modal, modalApi] = useVbenModal({
  // 连接抽离的组件
  connectedComponent: markRaw(ExtraModal),
});

function openModal() {
  modalApi.open();
}
</script>

<template>
  <div>
    <Modal />
    <VbenButton @click="openModal">Open</VbenButton>
  </div>
</template>
