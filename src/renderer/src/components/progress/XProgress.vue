<template>
  <div class="progress">
    <x-button v-if="props.running" @click="handleClick">
      <template #icon>
        <icon name="pause-fill" size="small" />
        <icon class="loading-icon" name="loading" size="small" />
      </template>
    </x-button>
    <x-button v-else @click="handleClick">
      <template #icon>
        <icon name="play-fill" size="small" />
      </template>
    </x-button>
    <div class="log" v-if="props.running" :title="props.log">{{ props.log }}
    </div>
    <div class="slot" v-else>
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import { watch } from 'vue'
import XButton from '../button/XButton.vue'
import { useStateStore } from '../../store/state'
const props = defineProps({
  running: {
    type: Boolean,
    default: false
  },
  percent: {
    type: Number,
    default: 0
  },
  log: {
    type: String,
    default: ''
  }
})
const emit = defineEmits(['click'])

const handleClick = () => {
  emit('click')
}

const stateStore = useStateStore()
watch(() => props.running, (val) => {
  stateStore.setRunning(val)
})

</script>

<style lang="less">
.progress {
  width: 100%;
  height: 100%;
  font-size: .875rem;
  display: flex;
  align-items: center;
  line-height: 100%;
  flex-flow: row;
  overflow: hidden;

  .loading-icon {
    animation: loading-rotate 1s linear infinite;
  }

  .log {
    margin-left: .25rem;
    flex: 1;
    overflow: hidden;
    white-space: nowrap;
  }

  .slot {
    margin-left: .25rem;
  }
}
</style>
