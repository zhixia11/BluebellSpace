<template>
  <div class="scan-path-settings">
    <form-group label="重复文件清理">
      <template #ext>
        <x-button type="text" size="mini" @click="handleNewPath('repetitive')">新增</x-button>
      </template>
      <data-list v-model="paths.repetitive" />
    </form-group>
    <form-group label="相似图片清理">
      <template #ext>
        <x-button type="text" size="mini" @click="handleNewPath('similar')">新增</x-button>
      </template>
      <data-list v-model="paths.similar" />
    </form-group>
    <form-group label="空文件夹清理">
      <template #ext>
        <x-button type="text" size="mini" @click="handleNewPath('empty_dir')">新增</x-button>
      </template>
      <data-list v-model="paths.empty_dir" />
    </form-group>
    <form-group label="存储分析">
      <template #ext>
        <x-button type="text" size="mini" @click="handleNewPath('analyse')">{{ !!paths.analyse ? '修改' : '新增'
          }}</x-button>
      </template>
      <data-list v-model="paths.analyse" />
    </form-group>
  </div>
</template>

<script setup>
import FormGroup from '../../components/form/FormGroup.vue'
import XButton from '../../components/button/XButton.vue'
import DataList from '../../components/list/DataList.vue'
import { useSettingsStore } from '../../store/settings'
const settingsStore = useSettingsStore()
const paths = settingsStore.settings.paths

const handleNewPath = async (type) => {
  const path = await window.api.selectFolder()
  if (!path) {
    return
  }
  if (type === 'analyse') {
    paths.analyse = path
    return
  }
  const list = paths[type]
  if (list.includes(path)) {
    return
  }
  list.push(path)
}
</script>

<style lang="less"></style>
