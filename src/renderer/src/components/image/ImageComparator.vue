<template>
  <Teleport to="#app-content">
    <Transition name="fade">
      <div class="image-comparator" v-if="visible">
        <flex-layout headerSize="1.75rem">
          <template #header>
            <div class="comparator-options">
              <x-button type="text" size="small" @click="cancel">
                <template #icon>
                  <icon name="left" size="mini" />
                </template>
                退出</x-button>
              <x-button type="text" size="small" @click="metaVisible = !metaVisible">
                <template #icon>
                  <icon name="info" size="mini" />
                </template>图像信息</x-button>
              <x-button type="text" size="small" @click="selectorVisble = !selectorVisble">
                <template #icon>
                  <icon name="exchange" size="mini" />
                </template>
                {{ selectorVisble ? '隐藏列表' : '显示列表' }}
              </x-button>
              <div class="comparator-status">
                {{ list.length }}正比较/{{ archive.length }}待比较/{{ data.length - archive.length }}已忽略/{{ data.length }}总图像
              </div>
            </div>
          </template>
          <div class="view-container">
            <div class="view-wrapper" v-for="path in list" :key="path">
              <image-viewer :path="path" :metaVisible="metaVisible">
                <template #ext-action>
                  <x-button title="忽略图像" @click="ignore(path)">
                    <template #icon>
                      <icon size="small" name="hide" />
                    </template>
                  </x-button>
                  <x-button title="删除图像" @click="remove(path)">
                    <template #icon>
                      <icon size="small" name="delete" />
                    </template>
                  </x-button>
                </template>
              </image-viewer>
            </div>
          </div>
          <Transition name="fade">
            <div class="selector-container" v-if="selectorVisble" @click="selectorVisble = false">
              <div class="selector-wrapper">
                <div class="selector-item" :class="{ 'selector-item-selected': list.includes(path) }"
                  v-for="path in archive" :key="path" @click.stop.prevent="handleClickItem(path)">
                  <img :src="`local://${path}`">
                </div>
              </div>
            </div>
          </Transition>
        </flex-layout>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup>
import FlexLayout from '../layout/FlexLayout.vue'
import XButton from '../button/XButton.vue'
import Icon from '../icon/Icon.vue'
import ImageViewer from './ImageViewer.vue'
import { ref, watch } from 'vue'
import { deleteFile } from '../../api/explorer'
import Modal from '../dialog'
//比较器是否显示
const visible = defineModel('visible', { type: Boolean, default: false })
//是否展示元数据信息
const metaVisible = ref(true)
//选择器列表是否显示
const selectorVisble = ref(false)
//最大可以同时比较的图片数量
const max = ref(2)
//图片数据
const data = defineModel('data', { type: Array, default: () => [] })
//正在比较的图片路径列表
const list = ref([])
//待比较的图片路径列表
const archive = ref([])
// 监听data的变化，更新正在比较的图片路径列表
watch([data, visible], ([newData, newVisible]) => {
  if (newVisible) {
    archive.value = newData
    const array = []
    for (let i = 0; i < archive.value.length && i < max.value; i++) {
      const item = archive.value[i]
      array.push(item)
    }
    list.value = array
  }
}, { immediate: true })
// 选中或取消选中某个文件
const handleClickItem = (path) => {
  //如果list中存在path，就移除，否则就加入
  if (list.value.includes(path)) {
    list.value = list.value.filter((item) => item !== path)
  } else {
    list.value.push(path)
  }
}
// 忽略
const ignore = (path) => {
  //移除list中的path
  list.value = list.value.filter((item) => item !== path)
  // 移除data中的path
  archive.value = archive.value.filter((item) => item !== path)
}
// 删除
const remove = (path) => {
  Modal.show({
    message: '确定删除该图像吗',
    width: 280,
    height: 140,
    onOk: () => {
      deleteFile(path).then((errs) => {
        if (!!errs) {
          console.log('删除失败', errs)
          return
        }
        ignore(path)
        data.value = data.value.filter((item) => item !== path)
      })
    }
  })
}

const cancel = () => {
  visible.value = false
}
</script>

<style lang="less">
.image-comparator {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: var(--color-bg-1);
  z-index: 1000;

  .comparator-options {
    display: flex;
    flex-flow: row wrap;
    align-items: center;
    padding: 0 .375rem;
    gap: .25rem;
  }

  .comparator-status {
    height: fit-content;
    display: flex;
    padding-left: .375rem;
    border-left: 1px solid var(--color-text-2);
  }

  .view-container {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    overflow: hidden;
    position: relative;

    .view-wrapper {
      flex-grow: 1;
    }

    .view-wrapper:not(:first-child) {
      border-left: 1px solid var(--color-mask-bg);
    }
  }

  .selector-container {
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    backdrop-filter: blur(5px);
    display: flex;
    justify-content: center;
    align-items: center;
    box-sizing: content-box;
    padding: 2rem 5rem;

    .selector-wrapper {
      width: fit-content;
      display: flex;
      justify-content: center;
      flex-flow: row wrap;
      gap: .375rem;
    }

    .selector-item {
      border-radius: .25rem;
      box-shadow: 0 0 1rem rgba(0, 0, 0, 0.5);
      overflow: hidden;
      transition: all .32s cubic-bezier(0.4, 0, 0.2, 1);
      box-sizing: border-box;
      border: 2px solid transparent;

      img {
        max-width: 10rem;
        max-height: 10rem;
        display: block;
        object-fit: cover;
        user-select: none;
        transition: all .32s cubic-bezier(0.4, 0, 0.2, 1);
      }

      &:hover {
        box-shadow: 0 0 1rem var(--color-primary-1);

        img {
          transform: scale(1.2);
        }
      }
    }

    .selector-item-selected {
      border-color: var(--color-primary-1);
      box-shadow: 0 0 1rem var(--color-primary-1);
    }
  }
}
</style>
