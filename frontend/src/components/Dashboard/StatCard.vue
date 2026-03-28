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
  margin: 0 12px;
  color: var(--text-primary);
  padding: 8px 12px;
  background-color: var(--bg-card);
  border-radius: 8px;
  transition: all 0.3s;
}

.stat-card:hover {
  background-color: var(--bg-secondary);
  transform: translateY(-2px);
  box-shadow: var(--shadow-hover);
}

.stat-icon {
  font-size: 1.5rem;
  margin-right: 8px;
}

.stat-content {
  display: flex;
  flex-direction: column;
}

.stat-label {
  font-size: 0.85rem;
  color: var(--text-secondary);
}

.stat-value {
  font-size: 1.1rem;
  font-weight: 600;
  color: var(--accent-color);
}
</style>