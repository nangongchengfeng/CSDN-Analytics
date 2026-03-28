<template>
  <div class="stats-grid">
    <StatCard icon="📝" label="文章数" :value="userInfo.article_num" />
    <StatCard icon="👥" label="关注数" :value="userInfo.fans_num" />
    <StatCard icon="❤️" label="喜欢数" :value="userInfo.like_num" />
    <StatCard icon="📥" label="收藏数" :value="userInfo.collect_num" />
    <StatCard icon="👁️" label="访问数" :value="userInfo.visit_num" />
    <StatCard icon="🏅" label="排名" :value="userInfo.rank" />
    <StatCard icon="🔄" label="分享数" :value="userInfo.share_num" />
    <StatCard icon="⏳" label="年龄" :value="userInfo.code_age" />
  </div>
</template>

<script setup lang="ts">
import StatCard from './StatCard.vue'
import type { UserInfo } from '@/types'

interface Props {
  userInfo: UserInfo
}

defineProps<Props>()
</script>

<style scoped>
.stats-grid {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 6px 12px;
  height: 100%;
  flex-wrap: nowrap;
  min-width: 0;
}
</style>