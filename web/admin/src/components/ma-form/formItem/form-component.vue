<template>
  <ma-form-item v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
    :component="props.component" :custom-field="props.customField">
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <component :is="props.component.dataIndex" :component="props.component" :customField="props.customField" />
    </slot>
  </ma-form-item>
</template>

<script setup>
import { onMounted, getCurrentInstance, watch, inject } from 'vue'
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

const app = getCurrentInstance().appContext.app

if (props.component.formType === 'component' && props.component.component && !app._context.components[props.component.dataIndex]) {
  app.component(props.component.dataIndex, props.component.component)
}

runEvent('onCreated', 'handleCommonEvent')
onMounted(() => {
  runEvent('onMounted', 'handleCommonEvent')
})
</script>