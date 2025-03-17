<template>
  <a-layout-content class="flex flex-col">
    <statistics v-if="userStore.user && userStore.user.dashboard === 'statistics'" />
    <work-panel v-else-if="userStore.user" />
  </a-layout-content>
</template>

<script setup>
import { useUserStore } from '@/store'
import Statistics from './components/statistics.vue'
import WorkPanel from './components/work-panel.vue'
const userStore = useUserStore()

</script>
<script>
export default { name: 'dashboard' }
</script>
