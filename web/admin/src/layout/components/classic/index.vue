<template>
  <a-layout class="layout flex justify-between h-full">

    <ma-classic-slider class="ma-ui-slider" />

    <a-layout-content class="flex flex-col">
      <ma-classic-header class="ma-ui-header" />
      <ma-worker-area />
    </a-layout-content>

  </a-layout>
</template>

<script setup>
import { ref } from 'vue'

import MaClassicSlider from './ma-classic-slider.vue'
import MaClassicHeader from './ma-classic-header.vue'
import MaWorkerArea from '../ma-workerArea.vue'

</script>
