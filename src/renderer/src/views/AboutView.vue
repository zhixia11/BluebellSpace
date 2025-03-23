<template>
  <div class="version-view">
    <div class="version-group-title">版本</div>
    <div class="product-name">
      <a-tag color="purple">Bluebell v{{ versions.bluebell }}</a-tag>
    </div>
    <div class="version-group-title">架构</div>
    <a-space wrap>
      <a-tag color="red">Electron v{{ versions.electron }}</a-tag>
      <a-tag color="blue">Node v{{ versions.node }}</a-tag>
      <a-tag color="green">Vue v{{ versions.vue }}</a-tag></a-space>
    <div class="version-group-title">关于</div>
    <a-space direction="vertical">
      <div class="copyright">Copyright ©{{ year }} BLUEBELL. All Rights Reserved.</div>
    </a-space>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
const versions = reactive({ ...window.electron.process.versions, ...window.api.versions })
const year = new Date().getFullYear()
</script>

<style lang="less">
.product-name {
  font-size: 1.25rem;
  margin-bottom: .375rem;
  display: flex;
  align-items: end;
}

.version-group-title {
  font-size: .875rem;
  margin-bottom: .375rem;
}
</style>
