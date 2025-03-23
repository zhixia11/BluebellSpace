<template>
  <div class="settings-view">
    <x-tabs default-active-name="common" v-model="activeName">
      <x-tab-section name="common" title="通用">
        <common-settings />
      </x-tab-section>
      <x-tab-section name="paths" title="扫描路径">
        <scan-path-settings />
      </x-tab-section>
    </x-tabs>
  </div>
</template>

<script setup>
import { XTabs, XTabSection } from '../../components/tabs'
import CommonSettings from './CommonSettings.vue'
import ScanPathSettings from './ScanPathSettings.vue'
import { ref } from 'vue'
const activeName = ref('common')
</script>

<style lang="less">
.settings-view {
  width: 100%;
  height: 100%;
  .x-tabs-content{
    padding-right: .5rem;
    overflow-y: auto;
  }
}
</style>
