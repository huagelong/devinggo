<template>
  <a-cascader v-model="value" :placeholder="props.component.searchPlaceholder || `请选择${props.component.title}`"
    allow-clear allow-search :expand-trigger="props.component.trigger || 'click'"
    :options="dicts[props.component.dataIndex]" :multiple="props.component.multiple" />
</template>

<script setup>
import { ref, inject, watch } from 'vue'
import { get, set } from 'lodash'
const props = defineProps({
  component: Object,
})
const searchForm = inject('searchForm')
const dicts = inject('dicts')

const value = ref(get(searchForm.value, props.component.dataIndex, props.component.searchDefaultValue))
set(searchForm.value, props.component.dataIndex, value.value)

watch( () => get(searchForm.value, props.component.dataIndex), vl => value.value = vl )
watch( () => value.value, v => set(searchForm.value, props.component.dataIndex, v) )
</script>