<template>
  <ma-form-item v-if="(typeof props.component.display == 'undefined' || props.component.display === true)"
    :component="props.component" :custom-field="props.customField">
    <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
      <component :is="getComponentName()" v-model="value"
        :placeholder="props.component.formType === 'range' ? props.component.placeholder ?? ['请选择开始时间', '请选择结束时间'] : props.component.placeholder ?? `请选择${props.component.title}`"
        :hide-trigger="props.component.hideTrigger" :allow-clear="props.component.allowClear ?? true"
        :format="props.component.format" :size="props.component.size" :shortcuts="props.component.shortcuts"
        :shortcuts-position="props.component.shortcutsPosition" :position="props.component.position"
        :popup-visible="props.component.popupVisible" :default-popup-visible="props.component.defaultPopupVisible"
        :trigger-props="props.component.triggerProps" :unmount-on-close="props.component.unmountOnClose"
        :disabled="props.component.disabled" :disabled-input="props.component.disabledInput"
        :disabled-date="props.component.disabledDate" :disabled-time="props.component.disabledTime"
        :value-format="props.component.valueFormat" :readonly="props.component.readonly" :error="props.component.error"
        :show-time="props.component.showTime" :preview-shortcut="props.component.previewShortcut"
        :show-confirm-btn="props.component.showConfirmBtn"
        :type="props.component.range ? props.component.formType === 'time' ? 'time-range' : 'range' : ''"
        :time-picker-props="props.component.formType == 'range' ? { defaultValue: ['00:00:00', '23:59:59'] } : {}"
        :separator="props.component.separator" :mode="props.component.mode" style="width: 100%;"
        @change="handlePickerChangeEvent" @select="handlePickerSelectEvent" @ok="handlePickerOkEvent"
        @clear="rv('onClear')" @popup-visible-change="rv('onvVisibleChange', $event)"
        @select-shortcut="rv('onSelectShortcut', $event)" @picker-value-change="handlePickerValueChangeEvent" />
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

const getComponentName = () => {
  if (['date', 'month', 'year', 'week', 'quarter', 'range', 'time'].includes(props.component.formType)) {
    return `a-${props.component.formType}-picker`
  }
}

const handlePickerChangeEvent = (value, date, dateString) => {
  rv('onPickerChange', { value, date, dateString })
}

const handlePickerSelectEvent = (value, date, dateString) => {
  rv('onSelect', { value, date, dateString })
}

const handlePickerValueChangeEvent = (value, date, dateString) => {
  rv('onPickerValueChange', { value, date, dateString })
}

const handlePickerOkEvent = (value, date, dateString) => {
  rv('onOk', { value, date, dateString })
}

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>