<template>
  <div>
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <a-divider v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
        :class="[props.component.customClass]" :margin="props.component.margin" :direction="props.component.direction"
        :orientation="props.component.orientation" :type="props.component.type" :size="props.component.size">
        {{ props.component?.title ?? '' }}
      </a-divider>
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