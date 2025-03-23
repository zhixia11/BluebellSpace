<template>
  <div class="x-select" :style="{ width: `${witdh}px` }">
    <div class="x-select-selected" @click="handlePopupVisible">
      {{ selectedLabel }}
      <div class="x-select-arrow"
        :class="{ 'x-select-arrow-expand': popupVisible, 'x-select-arrow-collapse': !popupVisible }">
        <icon name="bottom" size="nano" />
      </div>
    </div>
    <Transition name="slide-dynamic-origin">
      <div v-show="popupVisible" class="x-select-dropdown">
        <slot></slot>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import { ref, provide } from 'vue'
const props = defineProps({
  witdh: { ype: Number },
})
const selectedValue = defineModel()
const popupVisible = ref(false)
const selectedLabel = ref('')
const handlePopupVisible = (event) => {
  if (popupVisible.value) {
    popupVisible.value = false
  } else {
    document.body.click()
    popupVisible.value = true
    event.stopPropagation()
    document.body.addEventListener('click', () => {
      popupVisible.value = false
    }, { once: true })
  }
}
const updateSelectedValue = (value, label) => {
  selectedValue.value = value
  selectedLabel.value = label
  popupVisible.value = false
}
const registerOption = (value, label) => {
  if (selectedValue.value === value) {
    selectedLabel.value = label
  }
}
provide('selectSelectedValue', selectedValue)
provide('updateSelecteSelectedValue', updateSelectedValue)
provide('registerSelectOption', registerOption)
</script>

<style lang="less">
.x-select {
  width: 10rem;
  height: 1.5rem;
  border: 1px solid var(--color-primary-2);
  border-radius: 0.25rem;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  position: relative;
  cursor: pointer;
}

.x-select-selected {
  width: 100%;
  height: 100%;
  position: relative;
  display: flex;
  align-items: center;
  padding: 0 .5rem;
}

.x-select-arrow {
  height: 100%;
  display: flex;
  align-items: center;
  position: absolute;
  top: 0;
  right: .25rem;
}

.x-select-arrow-expand {
  svg {
    transform: rotate(180deg);
  }
}

.x-select-dropdown {
  width: calc(100% + 2px);
  padding: .375rem;
  box-sizing: border-box;
  border-radius: .25rem;
  position: absolute;
  left: -1px;
  top: calc(100% + .325rem);
  background-color: var(--color-bg-3);
  box-shadow: rgba(0, 0, 0, 0.2) 0 0 .5rem;
  z-index: 1000;

  &>.x-select-option:not(:last-child) {
    margin-bottom: .125rem;
  }
}
</style>
