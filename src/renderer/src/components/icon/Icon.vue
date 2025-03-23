<template>
  <svg class="app-icon" :class="{ [`app-icon-${props.size}`]: props.size }">
    <use :xlink:href="symbolId" fill="currentColor" />
  </svg>
</template>

<script setup>
import { computed } from "vue"
const props = defineProps({
  name: {
    type: String,
    required: true,
  },
  size: {
    type: String,
    required: false,
  }
})

const symbolId = computed(() => `#icon-${props.name}`);
</script>

<style lang="less">
.app-icon-nano {
  width: .75rem !important;
  height: .75rem !important;
}

.app-icon-tiny {
  width: .875rem !important;
  height: .875rem !important;
}

.app-icon-mini {
  width: 1rem !important;
  height: 1rem !important;
}

.app-icon-small {
  width: 1.25rem !important;
  height: 1.25rem !important;
}

.app-icon {
  width: 1.5rem;
  height: 1.5rem;
}

.app-icon-large {
  width: 2rem !important;
  height: 2rem !important;
}
</style>
