<template>
  <div class="user-profile">
    <div class="avatar-ring">
      <a href="https://blog.csdn.net/heian_99" target="_blank" rel="noopener noreferrer">
        <img
          src="https://www.ownit.top/img/avatar_hu227367ba8544f2fc7811ed9508937bec_102665_300x0_resize_box_3.png"
          alt="头像"
          class="avatar"
        />
      </a>
    </div>
    <div class="user-info">
      <span class="eyebrow">CSDN Analytics</span>
      <a href="https://blog.csdn.net/heian_99" target="_blank" rel="noopener noreferrer">
        <h1 class="user-name">{{ userInfo.author_name }} 博客洞察</h1>
      </a>
      <p class="user-subtitle">创作数据总览 · 阅读趋势 · 内容类型洞察</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { UserInfo } from '@/types'

interface Props {
  userInfo: UserInfo
}

defineProps<Props>()
</script>

<style scoped>
.user-profile {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 10px 4px;
  height: 100%;
  min-width: 0;
}

.avatar-ring {
  position: relative;
  width: 84px;
  height: 84px;
  margin-right: 18px;
  flex: 0 0 auto;
  border-radius: 50%;
  background:
    conic-gradient(from 130deg, var(--accent-secondary), var(--accent-highlight), #22d3ee, var(--accent-secondary));
  box-shadow: 0 14px 34px rgba(30, 64, 175, 0.18);
}

.avatar-ring::after {
  content: '';
  position: absolute;
  inset: 5px;
  border-radius: inherit;
  background: #ffffff;
}

.avatar {
  position: relative;
  z-index: 1;
  width: 74px;
  height: 74px;
  border-radius: 50%;
  object-fit: cover;
  margin: 5px;
  border: 3px solid rgba(255, 255, 255, 0.95);
}

.user-info {
  flex: 1;
  min-width: 0;
}

.eyebrow {
  display: inline-flex;
  margin-bottom: 7px;
  padding: 5px 10px;
  border-radius: 999px;
  background: rgba(30, 64, 175, 0.08);
  color: var(--accent-color);
  font-size: 0.72rem;
  font-weight: 800;
  letter-spacing: 0.12em;
  text-transform: uppercase;
}

.user-name {
  font-size: clamp(1.35rem, 1.65vw, 1.9rem);
  color: var(--accent-color);
  font-weight: 850;
  margin: 0;
  white-space: nowrap;
  letter-spacing: -0.04em;
  line-height: 1.05;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-subtitle {
  margin-top: 8px;
  color: var(--text-secondary);
  font-size: 0.86rem;
  font-weight: 600;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

a {
  text-decoration: none;
}

a:hover .user-name {
  opacity: 0.85;
}
</style>
