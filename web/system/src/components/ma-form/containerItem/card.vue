<template>
  <a-card v-if="(typeof props.component?.display == 'undefined' ||
    props.component?.display === true) &&
    (hasDisplayTrue((props.component?.formList ?? [])) ||
      props.component?.forceShow)" :class="[props.component?.customClass]" :extra="props.component?.extra"
    :bordered="props.component?.bordered" :loading="props.component?.loading" :hoverable="props.component?.hoverable"
    :size="props.component?.size" :header-style="props.component?.headerStyle" :body-style="props.component?.bodyStyle">
    <template #title>
      <slot :name="`cardTitle-${props.component?.dataIndex ?? ''}`">{{ props.component?.title }}</slot>
    </template>
    <template #actions>
      <slot :name="`cardAction-${props.component?.dataIndex ?? ''}`"></slot>
    </template>
    <template #cover>
      <slot :name="`cardCover-${props.component?.dataIndex ?? ''}`"></slot>
    </template>
    <template #extra>
      <slot :name="`cardExtra-${props.component?.dataIndex ?? ''}`"></slot>
    </template>
    <template v-for="(component, componentIndex) in (props.component?.formList ?? [])" :key="componentIndex">
      <component :is="getComponentName(component?.formType ?? 'input')" :component="component">
        <template v-for="slot in Object.keys($slots)" #[slot]="component">
          <slot :name="slot" v-bind="component" />
        </template>
      </component>
    </template>
  </a-card>
</template>

<script setup>
import { onMounted, inject } from 'vue'
import { getComponentName } from '../js/utils.js'
import { runEvent } from '../js/event.js'
const props = defineProps({ component: Object })

const formModel = inject('formModel')
const getColumnService= inject('getColumnService')
const columns = inject('columns')

const hasDisplayTrue = (list) => {
  return list.some(item => item.display ?? true);
}

const rv = async (ev, value = undefined) => await runEvent(props.component, ev, { formModel, getColumnService, columns }, value)

rv('onCreated')
onMounted(() => rv('onMounted'))
</script>
