<template>
  <flex-layout>
    <template #header>
      <x-progress @click="handleClick" :running="running" :percent="percent" :log="log">
        <x-button @click="deleteSelectPath" v-if="Object.keys(selectPathMap).length > 0">
          <template #icon>
            <icon name="delete-fill" size="small" />
          </template>
        </x-button>
      </x-progress>
    </template>
    <div class="similar-content">
      <template v-for="(list, hash) in data" :key="hash">
        <collapse v-if="list.length > 1" :title="`${getFileName(list[0])}等${list.length}个相似图像`">
          <template #options>
            <div class="similar-options">
              <x-button type="text" size="small" @click="ignoreGroup(hash)">
                <template #icon>
                  <icon name="hide" size="mini" />
                </template>
                忽略
              </x-button>
              <x-button type="text" size="small" @click="compare(hash)">
                <template #icon>
                  <icon name="compare" size="mini" />
                </template>
                比对
              </x-button>
            </div>
          </template>
          <div class="similar-list">
            <div class="similar-item" :class="{ 'similar-item-selected': selectPathMap[path] }"
              v-for="(path, index) in list" :key="path" @click="handleClickItem(path)">
              <img :src="`local://${path}`">
              <div class="similar-check-container"></div>
            </div>
          </div>
        </collapse>
      </template>
    </div>
    <image-comparator v-model:visible="compareVisible" v-model:data="data[selectedHash]" />
  </flex-layout>
</template>

<script setup>
import FlexLayout from '../components/layout/FlexLayout.vue'
import Modal from '../components/dialog'
import XProgress from '../components/progress/XProgress.vue'
import XButton from '../components/button/XButton.vue'
import ImageComparator from '../components/image/ImageComparator.vue'
import Collapse from '../components/collapse/Collapse.vue'
import Icon from '../components/icon/Icon.vue'
import { ref, reactive, onUnmounted } from 'vue'
import { submit, progressEventSource, cancel } from '../api/worker'
import { deleteFile } from '../api/explorer'
import { getFileName } from '../utils/utils'
import { useSettingsStore } from '../store/settings'
import { Notification } from '@arco-design/web-vue'
const running = ref(false)
const percent = ref(0)
const log = ref('')
const id = ref('')
const data = reactive({})

//忽略组别
const ignoreGroup = (hash) => {
  //移除选中的，避免误删
  const list = data[hash]
  for (let i = 0; i < list.length; i++) {
    const path = list[i]
    if (selectPathMap.value[path]) {
      delete selectPathMap.value[path]
    }
  }
  //移除group
  delete data[hash]
}

// 比较文件
const selectedHash = ref(false)
const compareVisible = ref(false)
const compare = (hash) => {
  selectedHash.value = hash
  compareVisible.value = true
}

// 启动按钮点击事件
const handleClick = () => {
  if (!running.value) {
    if (Object.keys(data).length > 0) {
      Modal.show({
        message: '重新开始任务会清空当前数据，确认重新开始吗',
        width: 280,
        height: 140,
        onOk: () => {
          submitTask()
        }
      })
    } else {
      submitTask()
    }
  } else {
    Modal.show({
      message: '确认取消任务吗',
      width: 280,
      height: 140,
      onOk: () => {
        cancel(id.value)
      }
    })
  }
}

// 提交任务
const settingsStore = useSettingsStore()
const submitTask = () => {
  const paths = settingsStore.settings?.paths?.similar
  if (paths.length <= 0) {
    Notification.info({
      content: '请先设置相似度比较路径'
    })
    running.value = false
    return
  }
  Object.assign(data, {})
  submit('similar', {
    path: paths
  }).then(data => {
    running.value = true
    id.value = data
    handleProgress()
  })
}

// EventSource
const es = ref(null)
// 处理进度日志
const handleProgress = () => {
  es.value = progressEventSource(id.value)
  es.value.onmessage = (event) => {
    const message = JSON.parse(event.data)
    const type = message.type
    const data = message.data
    if (type === 'info') {
      percent.value = message.percent
      log.value = data
    } else if (type === 'data') {
      handleNewData(data)
    } else if (type === 'error') {
      Notification.error({
        content: `${data}`
      })
      running.value = false
    }
  }
  es.value.onclose = () => {
    running.value = false
    es.value.close()
    es.value = null
  }
  es.value.onerror = () => {
    running.value = false
    es.value.close()
    es.value = null
  }
}

//处理新增数据
const handleNewData = (message) => {
  const hash = message.hash
  const image = message.image
  if (!data[hash]) {
    data[hash] = []
  }
  data[hash].push(image.path)
}

// 卸载组件的时候停止事件监听
onUnmounted(() => {
  if (!!es.value) {
    es.value.close()
    es.value = null
  }
})

//选中图像
const selectPathMap = ref({})
const handleClickItem = (path) => {
  if (selectPathMap.value[path]) {
    delete selectPathMap.value[path]
  } else {
    selectPathMap.value[path] = true
  }
}

// 删除选中的文件
const deleteSelectPath = () => {
  const paths = Object.keys(selectPathMap.value)
  if (paths.length > 0) {
    Modal.show({
      message: '确认删除选中的文件吗',
      width: 280,
      height: 140,
      onOk: () => {
        for (const key in selectPathMap.value) {
          const paths = []
          if (!!selectPathMap.value[key]) {
            paths.push(key)
          }
        }
        deleteFile(paths).then((errs) => {
          const map = {}
          for (const index in paths) {
            const path = paths[index]
            if (!!errs && errs.includes(path)) {
              // 没删掉的
              continue
            } else {
              delete selectPathMap.value[path]
              map[path] = true
            }
          }
          //移除重复组别
          for (const key in data) {
            const list = data[key]
            const dist = []
            for (const index in list) {
              const image = list[index]
              if (map[image.path]) {
                //移除
                continue
              }
              dist.push(image)
            }
            if (dist.length > 1) {
              data[key] = dist
            } else {
              delete data[key]
            }
          }
        })
      }
    })
  }
}
</script>

<style lang="less">
.similar-content {
  padding-right: .5rem;
}

.similar-options {
  display: flex;
  flex-flow: row nowrap;
  gap: .25rem;
}

.similar-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: left;
  gap: 0.125rem;
  border-radius: .25rem;
}

.similar-item {
  width: 9rem;
  height: 9rem;
  overflow: hidden;
  border-radius: .25rem;
  position: relative;

  &:hover {
    img {
      transform: scale(1.2);
    }
  }

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    object-position: center;
    transition: all .3s ease;
  }
}

.similar-check-container {
  position: absolute;
  right: .5rem;
  top: .5rem;
  width: .75rem;
  height: .75rem;
  border: 2px solid #ffffff;
  background-color: var(--color-primary-1);
  z-index: 100;
  border-radius: 2px;
  display: none;

  &::after {
    content: " ";
    width: .25rem;
    height: .5rem;
    border: solid white;
    border-width: 0 2px 2px 0;
    transform: rotate(45deg);
    position: absolute;
    left: 3px;
  }
}

.similar-item-selected {
  &::after {
    content: " ";
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(255, 255, 255, .5);
  }

  .similar-check-container {
    display: block;
  }
}
</style>
