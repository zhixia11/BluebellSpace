<template>
  <div class="dialog-container" v-show="visible">
    <Transition name="fade-dialog">
      <div class="dialog-mask" v-if="visible" @click="handleCancel">
      </div>
    </Transition>
    <Transition name="zoom-dialog">
      <div class="dialog-wrapper" :style="{ width: `${data.width}px`, height: `${data.height}px` }" v-if="visible">
        <icon class="dialog-icon" @click="handleCancel" name="close" size="mini" />
        <div class="dialog-header">
          <div class="dialog-title">{{ data.title }}</div>
        </div>
        <div class="dialog-content">
          <slot></slot>
          <div class="dialog-message">{{ data.message }}</div>
        </div>
        <div class="dialog-footer">
          <button class="dialog-footer-btn" @click="handleCancel" v-if="data.onCancel">取消</button>
          <button class="dialog-footer-btn dialog-footer-btn-ok" @click="handleOk">确认</button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup>
import Icon from '../icon/Icon.vue'
import { ref } from 'vue'
const emit = defineEmits(['onOk', 'onCancel'])
const props = defineProps({
  title: {
    type: String
  },
  width: {
    type: Number,
    default: 360
  },
  height: {
    type: Number,
    default: 240
  }
})
const data = ref({ ...props })
const visible = defineModel({
  type: Boolean,
  default: false
})

const handleOk = () => {
  if (!!data.value.onOk) {
    data.value.onOk()
  } else {
    emit('onOk')
  }
  visible.value = false
}

const handleCancel = () => {
  if (!!data.value.onCancel) {
    data.value.onCancel()
  } else {
    emit('onCancel')
  }
  visible.value = false
}

const show = (options) => {
  data.value = { ...options }
  visible.value = true
}

defineExpose({
  show
})
</script>

<style lang="less">
.dialog-container {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  z-index: 1001;
  display: flex;
  align-items: center;
  justify-content: center;
}

.dialog-mask {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: var(--color-mask-bg);
}

.dialog-wrapper {
  position: relative;
  z-index: 1002;
  background-color: var(--color-bg-3);
  border-radius: .25rem;
  padding: .5rem;
  box-sizing: border-box;
  display: flex;
  flex-flow: column;
}

.dialog-header {
  height: 1rem;
  padding-bottom: .375rem;
}

.dialog-content {
  flex: 1;
  min-height: 0;
  overflow: overlay;
  display: flex;
  justify-content: center;
  align-items: center;
}

.dialog-footer {
  display: flex;
  justify-content: center;
  padding: .375rem 0;

  &>button {
    margin-left: .5rem;
  }
}

.dialog-icon {
  position: absolute;
  top: .375rem;
  right: .375rem;
  cursor: pointer;

  &:hover {
    color: #C396ED;
  }
}

.dialog-message {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: .875rem;
}

.dialog-footer-btn {
  height: 1.25rem;
  box-sizing: content-box;
  padding: .25rem 1rem;
  border-radius: .25rem;
  background-color: #E5E6EB;
  color: #4E5969;
  cursor: pointer;

  &:hover {
    background-color: #F2F3F5;
  }
}

.dialog-footer-btn-ok {
  background-color: #C396ED;
  color: #ffffff;

  &:hover {
    background-color: #DDBEF6;
  }
}

@keyframes fadeIn {
  from {
    transform: scale(0.6);
    opacity: 0;
  }

  to {
    transform: scale(1);
    opacity: 1;
  }
}

@keyframes fadeOut {
  from {
    transform: scale(1);
    opacity: 1;
  }

  to {
    transform: scale(0.6);
    opacity: 0;
  }
}

.fade-dialog-enter-active,
.fade-dialog-leave-active {
  transition: opacity 0.4s cubic-bezier(0.3, 1.3, 0.3, 1);
}

.fade-dialog-enter-to,
.fade-dialog-leave-from {
  opacity: 1;
}

.fade-dialog-enter-from,
.fade-dialog-leave-to {
  opacity: 0;
}

.zoom-dialog-enter-active {
  transition: opacity 0.4s cubic-bezier(0.3, 1.3, 0.3, 1),
    transform 0.4s cubic-bezier(0.3, 1.3, 0.3, 1);
}

.zoom-dialog-enter-from {
  transform: scale(0.5);
  opacity: 0;
}

.zoom-dialog-enter-to {
  transform: scale(1, 1);
  opacity: 1;
}

.zoom-dialog-leave-active {
  transition: opacity 0.4s cubic-bezier(0.3, 1.3, 0.3, 1),
    transform 0.4s cubic-bezier(0.3, 1.3, 0.3, 1);
}

.zoom-dialog-leave-from {
  transform: scale(1, 1);
  opacity: 1;
}

.zoom-dialog-leave-to {
  transform: scale(0.5);
  opacity: 0;
}
</style>
