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
  background: rgba(255, 255, 255, 0.96);
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid rgba(226, 232, 240, 0.9);
  border-radius: 18px;
  box-shadow: 0 10px 40px rgba(15, 23, 42, 0.1),
              0 4px 12px rgba(15, 23, 42, 0.06),
              0 1px 0 rgba(255, 255, 255, 1) inset;
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  transition: all 0.4s cubic-bezier(0.22, 1, 0.36, 1);
  transform: translateY(0);
}

.mac-window:hover {
  box-shadow: 0 20px 60px rgba(15, 23, 42, 0.15),
              0 8px 24px rgba(15, 23, 42, 0.1),
              0 1px 0 rgba(255, 255, 255, 1) inset;
  transform: translateY(-4px);
}

.window-content {
  flex: 1;
  overflow: auto;
  padding: 18px;
}
</style>
