<template>
  <div class="flex-layout" :class="{ [`flex-layout-{{ props.direction }}`]: !!props.direction }">
    <div class="flex-layout-header" :style="headerStyle">
      <slot name="header"></slot>
    </div>
    <div class="flex-layout-content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
const props = defineProps({
  direction: { type: String },
  headerSize: { type: String, default: '' }
})

const headerStyle = computed(() => {
  return props.direction === 'column' ? { width: props.headerSize } : { height: props.headerSize }
})

</script>

<style lang="less">
.flex-layout {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: column;
  .flex-layout-header{
    display: flex;
    flex-flow: row;
  }
}

.flex-layout-row {
  flex-flow: row;
  .flex-layout-header{
    flex-flow: column;
  }
}

.flex-layout-content {
  flex: 1;
  overflow: auto;
  position: relative;
}
</style>
