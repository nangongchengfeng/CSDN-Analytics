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
  background: rgba(255, 255, 255, 0.75);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border: 1px solid rgba(226, 232, 240, 0.6);
  border-radius: var(--border-radius);
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06),
              0 1px 3px rgba(0, 0, 0, 0.04),
              inset 0 1px 0 rgba(255, 255, 255, 0.9);
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.mac-window:hover {
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.1),
              0 4px 6px rgba(0, 0, 0, 0.06),
              inset 0 1px 0 rgba(255, 255, 255, 0.9);
  transform: translateY(-2px);
}

.window-content {
  flex: 1;
  overflow: auto;
  padding: 16px;
}
</style>
