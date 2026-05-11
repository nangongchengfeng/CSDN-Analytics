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
  position: relative;
  background:
    linear-gradient(145deg, rgba(255, 255, 255, 0.98), rgba(248, 250, 252, 0.86)),
    var(--bg-card);
  backdrop-filter: blur(24px);
  -webkit-backdrop-filter: blur(24px);
  border: 1px solid var(--border-color);
  border-radius: 22px;
  box-shadow: var(--shadow);
  overflow: hidden;
  height: 100%;
  display: flex;
  flex-direction: column;
  transition: border-color 0.25s ease, box-shadow 0.25s ease, transform 0.25s ease;
  transform: translateY(0);
}

.mac-window::before {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: inherit;
  background:
    linear-gradient(120deg, rgba(255, 255, 255, 0.85), transparent 34%),
    radial-gradient(circle at 85% 0%, rgba(59, 130, 246, 0.12), transparent 28%);
  pointer-events: none;
}

.mac-window::after {
  content: '';
  position: absolute;
  inset: 0;
  border-radius: inherit;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.58);
  pointer-events: none;
}

.mac-window:hover {
  border-color: rgba(59, 130, 246, 0.36);
  box-shadow: var(--shadow-hover);
  transform: translateY(-2px);
}

.window-content {
  position: relative;
  z-index: 1;
  flex: 1;
  overflow: auto;
  padding: 20px;
}
</style>
