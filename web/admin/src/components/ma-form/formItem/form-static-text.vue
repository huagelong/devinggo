<template>
  <div>
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <div v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
        :class="['static-text', props.component.customClass]" :style="props.component.style">{{ props.component.title
          ?? '' }}</div>
    </slot>
  </div>
</template>

<script setup>
import { onMounted, inject } from 'vue'
import { runEvent } from '../js/event.js'
const props = defineProps({
  component: Object,
})

const formModel = inject('formModel')
const getColumnService= inject('getColumnService')
const columns = inject('columns')
const rv = async (ev, value = undefined) => await runEvent(props.component, ev, { formModel, getColumnService, columns }, value)

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>