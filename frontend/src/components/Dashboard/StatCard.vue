<template>
  <div class="stat-card">
    <span class="stat-icon">{{ icon }}</span>
    <div class="stat-content">
      <span class="stat-label">{{ label }}</span>
      <span class="stat-value">{{ value }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Props {
  icon: string
  label: string
  value: string | number
}

defineProps<Props>()
</script>

<style scoped>
.stat-card {
  display: inline-flex;
  align-items: center;
  color: var(--text-primary);
  padding: 12px 16px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 14px;
  border: 1px solid rgba(226, 232, 240, 0.85);
  transition: all 0.35s cubic-bezier(0.22, 1, 0.36, 1);
  box-shadow: 0 5px 16px rgba(15, 23, 42, 0.05),
              0 2px 6px rgba(15, 23, 42, 0.03),
              0 1px 0 rgba(255, 255, 255, 1) inset;
  transform: translateY(0);
  flex-shrink: 0;
}

.stat-card:hover {
  background: rgba(255, 255, 255, 0.98);
  transform: translateY(-2px) scale(1.02);
  box-shadow: 0 10px 28px rgba(59, 130, 246, 0.12),
              0 5px 14px rgba(15, 23, 42, 0.06),
              0 1px 0 rgba(255, 255, 255, 1) inset;
  border-color: rgba(59, 130, 246, 0.3);
}

.stat-icon {
  font-size: 1.45rem;
  margin-right: 11px;
  filter: drop-shadow(0 2px 4px rgba(15, 23, 42, 0.1));
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.stat-label {
  font-size: 0.78rem;
  color: var(--text-secondary);
  font-weight: 600;
  letter-spacing: 0.01em;
}

.stat-value {
  font-size: 1.15rem;
  font-weight: 800;
  color: var(--accent-color);
  line-height: 1.15;
  letter-spacing: -0.02em;
}
</style>
