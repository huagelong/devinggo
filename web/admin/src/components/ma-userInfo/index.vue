<template>
  <a-input v-model="val" :label="props.title" disabled class="w-full" />
</template>

<script setup>
import { ref, watch } from 'vue'
import { useUserStore } from '@/store'

const user = useUserStore().user
const val = ref()

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: [ String, Number ],
  title: { type: String, default: '用户信息'},
  field: { type: String, default: 'id'},
})

val.value = user[props.field] ? user[props.field].toString() : user.id.toString()

watch(
  () => val.value,
  vl => emit('update:modelValue', vl),
  { immediate: true }
)

</script>

<style scoped>
:deep(.arco-select-option-content) {
  width: 100%;
}
</style>
