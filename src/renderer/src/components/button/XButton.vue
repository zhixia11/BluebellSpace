<template>
  <button class="x-button" type="button"
    :class="{ 'x-button-active': isActive, 'x-button-only-icon': onlyIcon, [`x-button-${props.type}`]: !!props.type, [`x-button-${props.size}`]: !!props.size }"
    @click="handleClick">
    <span class="x-button-icon" v-if="!!slots.icon">
      <slot name="icon"></slot>
    </span>
    <span class="x-button-content">
      <slot></slot>
    </span>
  </button>
</template>

<script setup>
import { computed, inject, useSlots } from 'vue'
const props = defineProps({
  name: { type: String, default: '' },
  active: { type: Boolean, default: false },
  type: { type: String, default: '' },
  size: { type: String, default: '' },
  icon: { type: String }
})
const slots = useSlots()
const onlyIcon = computed(() => {
  return slots.default == undefined
})
const emit = defineEmits(['click'])
const activeName = inject('activeButtonName', null);
const updateActiveName = inject('updateActiveButtonName', null);
const isActive = computed(() => {
  if (!activeName) {
    return props.active;
  } else {
    return activeName.value === props.name;
  }
})
const handleClick = () => {
  if (updateActiveName) {
    updateActiveName(props.name)
  } else {
    emit('click')
  }
}
</script>

<style lang="less">
.x-button {
  color: #4E5969;
  padding: .25rem;
  border-radius: .25rem;
  line-height: 1;
  cursor: pointer;
  display: flex;
  align-items: center;
  box-sizing: border-box;
  overflow: hidden;
  transition: all 0.32s;

  &:not(:disabled):not(.x-button-active):hover {
    background-color: #e6ddfa;
  }

  &:not(:disabled):active {
    background-color: #ffffff;
    color: #c3a6f8 !important;
  }

  &:disabled {
    color: #c9cdd4;
    cursor: not-allowed;
  }
}

.x-button-icon {
  display: flex;
  align-items: center;
  margin-right: .25rem;
}

.x-button-only-icon {
  .x-button-icon {
    margin-right: 0;
  }
}

.x-button-active {
  background-color: #c3a6f8;
  color: #ffffff;
}

.x-button-text {
  background-color: transparent;

  &:hover,
  &:active {
    color: #c3a6f8;
    background-color: transparent !important;
  }
}

.x-button-outline {
  border-radius: 0;
  background-color: transparent;
  border-bottom: 2px solid transparent;

  &:hover {
    background-color: transparent !important;
    // color: #e6ddfa;
    border-color: #e6ddfa;
  }

  &.x-button-active {
    background-color: transparent;
    color: #c3a6f8;
    border-color: #c3a6f8;
  }
}

.x-button-small {
  font-size: .875rem;

  .x-button-icon {
    margin-right: .125rem !important;
  }
}

.x-button-mini {
  font-size: .75rem;

  .x-button-icon {
    margin-right: .125rem !important;
  }
}
</style>
