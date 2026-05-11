<template>
  <div class="stats-grid">
    <StatCard icon="文" label="文章数" :value="userInfo.article_num" />
    <StatCard icon="粉" label="关注数" :value="userInfo.fans_num" />
    <StatCard icon="赞" label="喜欢数" :value="userInfo.like_num" />
    <StatCard icon="藏" label="收藏数" :value="userInfo.collect_num" />
    <StatCard icon="访" label="访问数" :value="userInfo.visit_num" />
    <StatCard icon="榜" label="排名" :value="userInfo.rank" />
    <StatCard icon="享" label="分享数" :value="userInfo.share_num" />
    <StatCard icon="龄" label="年龄" :value="userInfo.code_age" />
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
  display: grid;
  grid-template-columns: repeat(8, minmax(76px, 1fr));
  align-items: center;
  gap: 8px;
  padding: 4px 0;
  height: 100%;
  min-width: 0;
}

@media (max-width: 1280px) {
  .stats-grid {
    grid-template-columns: repeat(4, minmax(92px, 1fr));
  }
}
</style>
