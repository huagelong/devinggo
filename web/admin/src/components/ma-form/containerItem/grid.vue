<template>
  <a-row v-if="(typeof props.component?.display == 'undefined' || props.component?.display === true)"
    :class="[props.component?.customClass]" :gutter="props.component?.gutter" :justify="props.component?.justify"
    :align="props.component?.align" :div="props.component?.div" :wrap="props.component?.wrap">
    <template v-for="(col, colIndex) in (props.component?.cols ?? [])" :key="colIndex">
      <ma-grid-col :component="col">
        <template v-for="slot in Object.keys($slots)" #[slot]="component">
          <slot :name="slot" v-bind="component" />
        </template>
      </ma-grid-col>
    </template>
  </a-row>
</template>

<script setup>
import { onMounted, inject } from 'vue'
import MaGridCol from './grid-col.vue'
import { runEvent } from '../js/event.js'
const props = defineProps({ component: Object })

const formModel = inject('formModel')
const getColumnService= inject('getColumnService')
const columns = inject('columns')
const rv = async (ev, value = undefined) => await runEvent(props.component, ev, { formModel, getColumnService, columns }, value)

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>