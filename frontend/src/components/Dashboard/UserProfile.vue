<template>
  <div class="user-profile">
    <a href="https://blog.csdn.net/heian_99" target="_blank" rel="noopener noreferrer">
      <img
        src="https://www.ownit.top/img/avatar_hu227367ba8544f2fc7811ed9508937bec_102665_300x0_resize_box_3.png"
        alt="头像"
        class="avatar"
      />
    </a>
    <div class="user-info">
      <a href="https://blog.csdn.net/heian_99" target="_blank" rel="noopener noreferrer">
        <h1 class="user-name">{{ userInfo.author_name }} 博客洞察</h1>
      </a>
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
  justify-content: center;
  padding: 16px;
  height: 100%;
}

.avatar {
  width: 70px;
  height: 70px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 14px;
  border: 3px solid rgba(59, 130, 246, 0.2);
  box-shadow: 0 4px 15px rgba(59, 130, 246, 0.15);
}

.user-info {
  flex: 0;
}

.user-name {
  font-size: 1.6rem;
  color: var(--accent-color);
  font-weight: 700;
  margin: 0;
  white-space: nowrap;
  letter-spacing: -0.02em;
}

a {
  text-decoration: none;
}

a:hover .user-name {
  opacity: 0.85;
}
</style>
