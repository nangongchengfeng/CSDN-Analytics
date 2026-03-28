<template>
  <div class="mac-window">
    <div class="window-content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
// 移除 MacTrafficLights 导入
// 移除 title prop
</script>

<style scoped>
.mac-window {
  background-color: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--border-radius);
  box-shadow: var(--shadow);
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  transition: box-shadow 0.2s;
}

.mac-window:hover {
  box-shadow: var(--shadow-hover);
}

.window-content {
  flex: 1;
  overflow: auto;
  padding: 16px;
}
</style>
