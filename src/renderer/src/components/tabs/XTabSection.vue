<template>
  <div class="x-tab-section" v-show="isActive">
    <slot></slot>
  </div>
</template>

<script setup>
import { inject, computed } from 'vue'
const props = defineProps({
  name: {
    type: String,
    required: true
  },
  title: {
    type: String,
    required: true
  }
})
const activeName = inject('activeTabName', null)
const isActive = computed(() => activeName.value === props.name)
</script>

<style lang="less"></style>
