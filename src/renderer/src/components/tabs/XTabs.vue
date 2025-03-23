<template>
  <div class="x-tabs">
    <div class="x-tabs-header">
      <div class="x-tabs-nav-pre" v-if="overlay">
        <x-button>
          <template #icon>
            <icon name="left" size="mini" />
          </template>
        </x-button>
      </div>
      <div class="x-tabs-nav-container" ref="navContainer">
        <div class="x-tabs-nav-wrapper" ref="navWrapper">
          <x-button v-for="tab in tabs" :key="tab.name" :name="tab.name" class="x-tabs-header-item" type="outline"
            :active="activeName === tab.name" @click="setActiveTab(tab.name)">
            {{ tab.title }}
          </x-button>
        </div>
      </div>
      <div class="x-tabs-nav-next" v-if="overlay">
        <x-button>
          <template #icon>
            <icon name="right" size="mini" />
          </template>
        </x-button>
      </div>
    </div>
    <div class="x-tabs-content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import { useSlots, computed, provide, ref } from 'vue'
import XButton from '../button/XButton.vue'
const props = defineProps({
  defaultActiveName: {
    type: String,
    default: ''
  }
})
const emit = defineEmits(['onChange'])
const activeName = defineModel({ type: String, default: '' })
const slots = useSlots()
const tabs = computed(() => {
  return slots.default().map((tab) => ({
    name: tab.props.name,
    title: tab.props.title
  }))
})
const setActiveTab = (name) => {
  emit('onChange', name)
  activeName.value = name
}
provide('activeTabName', activeName)

//todo 导航栏超出处理
const overlay = ref(false)
</script>

<style lang="less">
.x-tabs {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: column;
}

.x-tabs-header {
  height: 1.5rem;
  display: flex;
  flex-flow: row nowrap;

  .x-tabs-nav-container {
    flex: 1;
    overflow: hidden;
  }

  .x-tabs-nav-wrapper {
    width: max-content;
    display: flex;
    flex-flow: row nowrap;
  }

  .x-button {
    height: 1.5rem;
    padding: .25rem .75rem;
  }
}

.x-tabs-content {
  flex: 1;
  min-height: 0;
  margin-top: .375rem;
  overflow: auto;
}
</style>
