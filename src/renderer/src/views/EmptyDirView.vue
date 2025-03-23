<template>
  <div class="empty-dir-view">
    <flex-layout>
      <template #header>
        <x-progress @click="handleClick" :running="running" :percent="percent" :log="log"/>
      </template>
      <collapse v-for="(list, key) in data" :title="`${key}`">
        <div class="empty-dir-item" v-for="(path, index) in list" :key="index">
          <div class="empty-dir-item-path">{{ path }}</div>
        </div>
      </collapse>
    </flex-layout>
  </div>
</template>

<script setup>
import FlexLayout from '../components/layout/FlexLayout.vue'
import XProgress from '../components/progress/XProgress.vue'
import Collapse from '../components/collapse/Collapse.vue'
import { reactive } from 'vue'

const data = reactive({
  'C:\\': ['C:\\Users\\Administrator\\Desktop','C:\\Users\\Administrator\\AppData\\Local\\Temp','C:\\Users\\Windows']
})
</script>

<style lang="less"></style>
