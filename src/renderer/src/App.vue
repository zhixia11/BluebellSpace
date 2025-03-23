<template>
  <app-layout>
    <div class="main-container">
      <div class="left-wrapper" :style="{ width: collapse ? '' : '20%' }">
        <button-group default-active-name="repetitive" :active="route.name" :collapse="collapse" @change="handleChange">
          <x-button name="repetitive">
            <template #icon>
              <icon name="copy" width="20" height="20" />
            </template>
            重复文件清理
          </x-button>
          <x-button name="similar">
            <template #icon>
              <icon name="image" width="20" height="20" />
            </template>
            相似图片清理
          </x-button>
          <x-button name="empty">
            <template #icon>
              <icon name="box" width="20" height="20" />
            </template>
            空文件夹清理
          </x-button>
          <x-button name="analyse">
            <template #icon>
              <icon name="pie" width="20" height="20" />
            </template>
            存储分析
          </x-button>
          <x-button name="settings">
            <template #icon>
              <icon name="settings" width="20" height="20" />
            </template>
            设置
          </x-button>
          <x-button name="about">
            <template #icon>
              <icon name="bookmark" width="20" height="20" />
            </template>
            关于
          </x-button>
        </button-group>
      </div>
      <div class="main-content">
        <router-view />
      </div>
    </div>
  </app-layout>
</template>

<script setup>
import AppLayout from './components/layout/AppLayout.vue'
import ButtonGroup from './components/button/ButtonGroup.vue'
import XButton from './components/button/XButton.vue'
import Icon from './components/icon/Icon.vue'
import { useRouter, useRoute } from 'vue-router'
import { computed, onMounted, watchEffect } from 'vue'
import { useSettingsStore } from './store/settings'
const settingsStore = useSettingsStore()
const router = useRouter()
const route = useRoute()
const handleChange = (name) => {
  router.push(name)
}

const collapse = computed(() => {
  return settingsStore.settings.common.layout === 'collapse'
})
watchEffect(() => {
  let theme = settingsStore.settings.common.theme
  if (theme === 'auto') {
    theme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light'
  }
  document.body.setAttribute('theme', theme)
})

onMounted(() => {
  window.api.onWindowBlur(() => {
    document.body.click()
  })
})
</script>

<style lang="less">
.main-container {
  height: 100%;
  display: flex;
  flex-flow: row;
}

.left-wrapper {
  height: 100%;
  display: flex;
  flex-flow: column;
  margin: 0 1.25rem;
}

.main-content {
  flex: 1;
  min-width: 0;
}
</style>