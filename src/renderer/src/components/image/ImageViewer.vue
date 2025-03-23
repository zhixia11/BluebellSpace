<template>
  <div class="image-viewer">
    <div ref="containerRef" class="image-container" @wheel="wheel" @dblclick="reset">
      <img ref="imageRef" @load="onload" @mousedown="mousedown" :src="`local://${props.path}`" :style="style" />
    </div>
    <div class="meta-container" v-if="props.metaVisible">
      <p>
        <span class="meta-label">路径：</span>
        <span class="meta-value">{{ props.path }}</span>
      </p>
      <p>
        <span class="meta-label">分辨率：</span>
        <span class="meta-value">{{ naturalWidth + '*' + naturalHeight }}</span>
      </p>
      <p v-if="meta?.size">
        <span class="meta-label">大小：</span>
        <span class="meta-value">{{ formatSize(meta?.size) }}</span>
      </p>
      <p v-if="meta?.mod_time">
        <span class="meta-label">修改时间：</span>
        <span class="meta-value">{{ meta?.mod_time }}</span>
      </p>
    </div>
    <div class="action-container" v-if="loaded">
      <x-button title="放大" @click="zoomin" :disabled="scale >= 10">
        <template #icon>
          <icon size="small" name="zoomin" />
        </template>
      </x-button>
      <x-button title="缩小" @click="zoomout" :disabled="scale <= 0.1"> <template #icon>
          <icon size="small" name="zoomout" />
        </template></x-button>
      <x-button :title="isOriginal ? '自适应大小' : '图片原始大小'" @click="convert">
        <template #icon>
          <icon size="small" name="restore" v-if="isOriginal" />
          <icon size="small" name="maximize" v-else /></template>
      </x-button>
      <x-button title="旋转" @click="rotate">
        <template #icon>
          <icon size="small" name="goback" />
        </template>
      </x-button>
      <x-button title="定位文件" @click="location">
        <template #icon>
          <icon size="small" name="location" />
        </template>
      </x-button>
      <slot name="ext-action" />
    </div>
    <Transition name="fade">
      <div class="message" v-if="messageVisible">
        {{ Math.round(scale * 100) }}%
      </div>
    </Transition>
  </div>
</template>

<script setup>
import XButton from '../button/XButton.vue'
import Icon from '../icon/Icon.vue'
import { computed, ref, watch, onMounted } from 'vue'
import { info } from '../../api/explorer'
import { formatSize } from '../../utils/utils'
const props = defineProps({
  path: { type: String, required: true },
  metaVisible: { type: Boolean, default: true },
})
const loaded = ref(false)
const naturalWidth = ref(0)
const naturalHeight = ref(0)
const containerRef = ref(null)
const imageRef = ref(null)
const meta = ref({})
const style = computed(() => {
  return {
    transform: `scale(${scale.value}) rotate(-${deg.value}deg) translate(${offsetX.value}px, ${offsetY.value}px)`,
    cursor: isDragging.value ? 'grabbing' : 'grab'
  }
})
const onload = () => {
  loaded.value = true
  naturalWidth.value = imageRef.value.naturalWidth
  naturalHeight.value = imageRef.value.naturalHeight
  auto()
}
// 移动相关操作
const isDragging = ref(false)
const offsetX = ref(0)
const offsetY = ref(0)
const startX = ref(0)
const startY = ref(0)
const initOffsetX = ref(0)
const initOffsetY = ref(0)
const mousedown = (event) => {
  event.preventDefault()
  isDragging.value = true
  initOffsetX.value = offsetX.value
  initOffsetY.value = offsetY.value
  startX.value = event.clientX
  startY.value = event.clientY
  window.addEventListener('mousemove', mousemove)
  window.addEventListener('mouseup', mouseup)

}
const mousemove = (event) => {
  if (isDragging.value) {
    offsetX.value = initOffsetX.value + (event.clientX - startX.value) / scale.value
    offsetY.value = initOffsetY.value + (event.clientY - startY.value) / scale.value
  }
}
const mouseup = () => {
  if (isDragging.value) {
    isDragging.value = false
    window.removeEventListener('mousemove', mousemove);
    window.removeEventListener('mouseup', mouseup);
  }
}
const reset = () => {
  offsetX.value = 0
  offsetY.value = 0
  auto()
}

// 缩放相关操作
const scale = ref(1)
const standard = ref(1)
const isOriginal = ref(true)
//放大
const zoomin = () => {
  let val = scale.value
  val *= 1.2
  val = Math.round(val * 100) / 100
  scale.value = Math.min(val, 10)
  afterZoom()
}
// 缩小
const zoomout = () => {
  let val = scale.value
  val *= 0.9
  val = Math.round(val * 100) / 100
  scale.value = Math.max(val, 0.05)
  afterZoom()
}
//缩放后处理
const afterZoom = () => {
  if (scale.value === 1) {
    isOriginal.value = true
  }
  showMessage(`${Math.round(scale.value * 100)}%}`)
}
//自适应和原始大小之间切换
const convert = () => {
  offsetX.value = 0
  offsetY.value = 0
  if (isOriginal.value) {
    auto()
  } else {
    original()
  }
  showMessage(`${Math.round(scale.value * 100)}%}`)
}
//原始大小
const original = () => {
  scale.value = 1
  isOriginal.value = true
}
//自适应大小
const auto = () => {
  const { width: containerWidth, height: containerHeight } = containerRef.value.getBoundingClientRect()
  const scaleWidth = containerWidth / naturalWidth.value
  const scaleHeight = containerHeight / naturalHeight.value
  scale.value = Math.min(Math.min(scaleWidth, scaleHeight), 1)
  isOriginal.value = false
  standard.value = scale.value
}
//根据鼠标滚动方向缩放
const wheel = (event) => {
  if (event.deltaY > 0) {
    zoomout()
  } else {
    zoomin()
  }
}
//旋转
const deg = ref(0)
const rotate = () => {
  let val = deg.value
  val += 90
  deg.value = Math.abs(val % 360)
}

// 消息提示相关
let timer = null
const message = ref('')
const messageVisible = ref(false)
const showMessage = (msg) => {
  message.value = msg
  messageVisible.value = true
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    messageVisible.value = false
  }, 1000)
}
//加载元数据
const loadMeta = () => {
  if (Object.keys(meta.value).length > 0) {
    return
  }
  info(props.path).then((res) => {
    meta.value = res
  })
}
watch(props, (newValue, oldValue) => {
  if (newValue.path !== oldValue?.path) {
    meta.value = {}
  }
  if (newValue.metaVisible) {
    loadMeta()
  }
})
onMounted(() => {
  loadMeta()
})

// 在资源管理器中打开
const location = () => {
  window.api.openInExplorer(props.path)
}
</script>

<style lang="less">
.image-viewer {
  width: 100%;
  height: 100%;
  min-width: 240px;
  min-height: 300px;
  background-color: #000;
  position: relative;
  container-type: inline-size;
  overflow: hidden;
  display: flex;
  justify-content: center;
}

.image-container {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.meta-container {
  position: absolute;
  top: .5rem;
  left: .5rem;
  max-width: calc(100% - 1rem);
  background-color: rgba(255, 255, 255, .75);
  backdrop-filter: blur(5px);
  padding: 0.5rem;
  box-sizing: border-box;
  border-radius: .25rem;
  box-shadow: 0 0 .5rem rgba(0, 0, 0, 0.2);
  line-height: 1.25rem;

  .meta-value {
    user-select: text;
  }
}

.action-container {
  background-color: rgba(255, 255, 255, .75);
  backdrop-filter: blur(5px);
  position: absolute;
  bottom: 5%;
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  align-items: center;
  border-radius: .25rem;
  gap: .5rem;
  margin: 0 1rem;
  padding: .25rem .5rem;

  .x-button {
    padding: .5rem;
  }
}

.message {
  position: absolute;
  padding: .875rem 2rem;
  border-radius: .25rem;
  top: 50%;
  transform: translateY(-100%);
  background-color: rgba(0, 0, 0, .5);
  backdrop-filter: blur(5px);
  color: #ffffff;
}

@container (max-width: 420px) {
  .action-container {
    padding: .25rem .375rem;
    gap: .25rem;

    .x-button {
      padding: .375rem;

      svg {
        width: 1rem !important;
        height: 1rem !important;
      }
    }
  }
}

@container (max-width: 300px) {
  .action-container {
    .x-button {
      padding: .25rem;
    }
  }
}
</style>
