<template>
  <div class="button-group" :class="{ 'button-group-collapse': collapse }">
    <slot></slot>
  </div>
</template>
<script setup>
import { provide, computed } from 'vue'
const props = defineProps({
  defaultActiveName: { type: String, default: '' },
  collapse: { type: Boolean, default: false },
  active: { type: String }
})
const activeName = computed(() => {
  return props.active || props.defaultActiveName
})
const emit = defineEmits(['change'])
provide('activeButtonName', activeName)
provide('updateActiveButtonName', (val) => {
  emit('change', val)
})
</script>
<style lang="less">
.button-group {
  display: flex;
  flex-flow: column;

  .x-button {
    height: 2.25rem;
    padding: .25rem .5rem;

    &:not(:last-child) {
      margin-bottom: .5rem;
    }
  }
}

.button-group-collapse {
  .x-button-icon {
    margin-right: 0;
  }

  .x-button-content {
    width: 0;
    overflow: hidden;
  }
}
</style>