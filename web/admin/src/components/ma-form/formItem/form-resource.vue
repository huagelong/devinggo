<template>
  <ma-form-item v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
    :component="props.component" :custom-field="props.customField">
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <ma-resource v-if="(props.component.type ?? 'preview') == 'preview'" v-model="value"
        :multiple="props.component.multiple" :onlyData="props.component.onlyData"
        :returnType="props.component.returnType" />
      <ma-resource-button v-else v-model="value" :multiple="props.component.multiple"
        :onlyData="props.component.onlyData" :returnType="props.component.returnType" />
    </slot>
  </ma-form-item>
</template>

<script setup>
import { ref, inject, onMounted, watch } from 'vue'
import { get, set } from 'lodash'
import MaResource from '@/components/ma-resource/index.vue'
import MaResourceButton from '@/components/ma-resource/button.vue'
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

if (props.component.multiple && ! value.value) {
  value.value = []
}

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>