<template>
  <div class="mac-traffic-lights">
    <div class="traffic-light red"></div>
    <div class="traffic-light yellow"></div>
    <div class="traffic-light green"></div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.mac-traffic-lights {
  display: flex;
  gap: 8px;
  padding: 8px 12px;
}

.traffic-light {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  opacity: 0.8;
  transition: opacity 0.2s;
}

.traffic-light:hover {
  opacity: 1;
}

.traffic-light.red {
  background-color: #ff5f57;
}

.traffic-light.yellow {
  background-color: #ffbd2e;
}

.traffic-light.green {
  background-color: #28ca42;
}
</style>