<template>
  <ma-form-item v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
    :component="props.component" :custom-field="props.customField">
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <a-slider v-model="value" :size="props.component.size" :allow-clear="props.component.allowClear"
        :disabled="props.component.disabled" :step="props.component.step" :show-tooltip="props.component.showTooltip"
        :range="props.component.range" :direction="props.component.direction" :max="props.component.max"
        :min="props.component.min" :marks="props.component.marks" :show-input="props.component.showInput"
        :show-ticks="props.component.showTicks" @change="rv('onChange', $event)">
      </a-slider>
    </slot>
  </ma-form-item>
</template>

<script setup>
import { ref, inject, onMounted, watch } from 'vue'
import { get, set } from 'lodash'
import MaFormItem from './form-item.vue'
import { runEvent } from '../js/event.js'
const props = defineProps({
  component: Object,
  customField: { type: String, default: undefined }
})

const formModel = inject('formModel')
const getColumnService= inject('getColumnService')
const columns = inject('columns')
const rv = async (ev, value = undefined) => await runEvent(props.component, ev, { formModel, getColumnService, columns }, value)
const index = props.customField ?? props.component.dataIndex
const value = ref(get(formModel.value, index))

watch( () => get(formModel.value, index), vl => value.value = vl )
watch( () => value.value, v => {
  set(formModel.value, index, v)
  index.indexOf('.') > -1 && delete formModel.value[index]
} )

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>