<template>
  <div class="app-layout">
    <div class="app-header">
      <app-title />
    </div>
    <div id="app-content" class="app-content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import AppTitle from './AppTitle.vue'
</script>

<style lang="less">
.app-layout {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: column;
}

.app-header {
  height: 2rem;
}

.app-content {
  flex: 1;
  overflow: hidden;
  position: relative;
}
</style>
