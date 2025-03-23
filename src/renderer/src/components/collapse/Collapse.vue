<template>
  <div class="collapse">
    <div class="collapse-header" @click="collapse = !collapse">
      <div class="collapse-icon">
        <icon name="right" size="tiny" :style="{ transform: collapse ? '' : 'rotate(90deg)' }" />
      </div>
      <div class="collapse-title">
        {{ title }}
      </div>
      <div class="collapse-options" @click.stop v-if="!collapse">
        <slot name="options" />
      </div>
    </div>
    <div class="collapse-content" v-if="!collapse">
      <slot />
    </div>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import { ref } from 'vue'
const props = defineProps({
  title: {
    type: String,
    default: ''
  }
})
const collapse = ref(true)
</script>

<style lang="less">
.collapse {
  padding: .5rem;
  border-radius: .25rem;
  transition: all .3s ease-out;

  &:hover {
    background-color: var(--color-primary-3);
  }
}

.collapse-header {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  height: 1.75rem;
  line-height: 100%;
}

.collapse-icon {
  display: flex;
  align-items: center;
  margin-right: .25rem;

  svg {
    transition: all .2s ease;
  }
}

.collapse-title {
  flex-grow: 1;
}

.collapse-options {
  margin-right: .25rem;
}

.collapse-content {
  overflow: hidden;
}
</style>
