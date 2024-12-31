<template>
  <ma-form-item v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
    :component="props.component" :custom-field="props.customField">
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <a-switch v-model="value" :size="props.component.size" :disabled="props.component.disabled"
        :loading="props.component.loading" :type="props.component.type" :checked-value="props.component.checkedValue"
        :unchecked-value="props.component.uncheckedValue" :checked-color="props.component.checkedColor"
        :unchecked-color="props.component.uncheckedColor" :before-change="props.component.beforeChange"
        @change="rv('onChange', $event)" @focus="rv('onFocus')" @blur="rv('onBlur')">
        <template #checked>
          <slot :name="`switchChecked-${props.component.dataIndex}`" />
        </template>
        <template #unchecked>
          <slot :name="`switchUnchecked-${props.component.dataIndex}`" />
        </template>
        <template #checked-icon>
          <slot :name="`switchChecked-${props.component.dataIndex}`" />
        </template>
        <template #unchecked-icon>
          <slot :name="`switchUncheckedIcon-${props.component.dataIndex}`" />
        </template>
      </a-switch>
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