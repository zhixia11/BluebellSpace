<template>
  <div class="app-title">
    <div class="right-fixed">
      <button class="app-title-btn" @click="minimize">
        <icon name="app-minimize" />
      </button>
      <button class="app-title-btn" @click="maximize" v-if="!isMax">
        <icon name="app-maximize" size="small" />
      </button>
      <button class="app-title-btn" @click="maximize" v-else>
        <icon name="app-restore" size="small" />
      </button>
      <button class="app-title-btn" @click="close">
        <icon name="app-close" />
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Icon from '../icon/Icon.vue'

const isMax = ref(false)
onMounted(() => {
  // 监听窗口最大化和最小化事件
  window.api.onWindowMaximizeUnmaximize((isMaximize) => {
    isMax.value = isMaximize
  })
})

const minimize = () => {
  window.api.minimize()
}

const maximize = () => {
  window.api.maximize()
}

const close = () => {
  window.api.close()
}
</script>

<style lang="less">
.app-title {
  -webkit-app-region: drag;
  height: 2rem;
  display: flex;
  align-items: center;
  position: relative;
  z-index: 1000;
}

.right-fixed {
  -webkit-app-region: no-drag;
  height: 100%;
  padding: 0 .125rem;
  position: absolute;
  right: 0;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.app-title-btn {
  -webkit-app-region: no-drag;
  height: 1.5rem;
  width: 1.5rem;
  padding: 0;
  margin: 0 .25rem;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #313131;

  &:hover {
    color: #c396ed;
  }
}
</style>
