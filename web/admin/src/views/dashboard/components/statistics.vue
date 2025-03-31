<template>

  <st-welcome />
  <st-count />
  <st-login-chart />

  <div class="block lg:flex">
    <st-announced />
    <st-devinggo />
  </div>
</template>

<script setup>
  import StCount from './components/st-count.vue'
  import StWelcome from './components/st-welcome.vue'
  import StLoginChart from './components/st-loginChart.vue'
  import StDevinggo from './components/st-devinggo.vue'
  import StAnnounced from './components/st-announced.vue'
</script>
