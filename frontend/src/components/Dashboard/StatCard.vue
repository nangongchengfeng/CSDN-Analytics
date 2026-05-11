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
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  min-width: 0;
  min-height: 68px;
  color: var(--text-primary);
  padding: 10px 11px;
  background:
    linear-gradient(145deg, rgba(255, 255, 255, 0.95), rgba(248, 250, 252, 0.76));
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border-radius: 16px;
  border: 1px solid rgba(203, 213, 225, 0.7);
  transition: border-color 0.25s ease, box-shadow 0.25s ease, transform 0.25s ease;
  box-shadow: 0 10px 24px rgba(15, 23, 42, 0.055),
              0 1px 0 rgba(255, 255, 255, 0.9) inset;
  transform: translateY(0);
  overflow: hidden;
}

.stat-card::after {
  content: '';
  position: absolute;
  inset: auto 10px 0;
  height: 2px;
  background: linear-gradient(90deg, var(--accent-secondary), var(--accent-highlight));
  opacity: 0.6;
}

.stat-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 14px 34px rgba(30, 64, 175, 0.12),
              0 5px 14px rgba(15, 23, 42, 0.06);
  border-color: rgba(59, 130, 246, 0.36);
}

.stat-icon {
  position: absolute;
  top: 8px;
  right: 8px;
  display: inline-flex;
  width: 24px;
  height: 24px;
  margin-right: 0;
  align-items: center;
  justify-content: center;
  flex: 0 0 auto;
  border-radius: 9px;
  background: linear-gradient(135deg, rgba(30, 64, 175, 0.12), rgba(59, 130, 246, 0.08));
  color: var(--accent-color);
  font-size: 0.72rem;
  font-weight: 800;
  box-shadow: inset 0 0 0 1px rgba(59, 130, 246, 0.16);
}

.stat-content {
  display: flex;
  flex-direction: column;
  gap: 3px;
  width: 100%;
  min-width: 0;
}

.stat-label {
  max-width: calc(100% - 26px);
  font-size: 0.68rem;
  color: var(--text-secondary);
  font-weight: 700;
  letter-spacing: 0.04em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.stat-value {
  max-width: 100%;
  font-size: clamp(0.95rem, 0.95vw, 1.12rem);
  font-weight: 800;
  color: var(--accent-color);
  line-height: 1.15;
  letter-spacing: -0.02em;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
