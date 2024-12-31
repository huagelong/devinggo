<template>
  <slot :name="`form-${props.component.dataIndex}`" v-bind="props.component">
    <a-link :status="props.component.status" :hoverable="props.component.hoverable" :disabled="props.component.disabled"
      :loading="props.component.loading" :href="props.component.href" @click="rv('onClick')">
      <template #icon v-if="props.component.icon">
        <component :is="props.component.icon" />
      </template>
      {{ props.component.title ?? 'link' }}
    </a-link>
  </slot>
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