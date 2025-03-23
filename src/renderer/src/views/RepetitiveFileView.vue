<template>
  <div class="repetitive-view">
    <x-progress @click="handleClick" :running="running" :percent="percent" :log="log" />
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Modal from '../components/dialog'
import XProgress from '../components/progress/XProgress.vue'
const running = ref(false)
const percent = ref(0)
const log = ref('')

const handleClick = () => {
  if (!running.value) {
    running.value = true
  } else {
    Modal.show({
      message: '确认取消任务吗',
      width: 280,
      height: 140,
      onOk: () => {
        running.value = false
      }
    })
  }
}
</script>

<style lang="less"></style>
