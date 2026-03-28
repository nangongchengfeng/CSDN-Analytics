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
  padding: 16px;
  height: 100%;
}

.avatar {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  object-fit: cover;
  margin-right: 16px;
  border: 2px solid var(--border-color);
}

.user-info {
  flex: 1;
}

.user-name {
  font-size: 1.5rem;
  color: var(--accent-color);
  font-weight: 600;
  margin: 0;
  white-space: nowrap;
}

a {
  text-decoration: none;
}

a:hover .user-name {
  opacity: 0.8;
}
</style>