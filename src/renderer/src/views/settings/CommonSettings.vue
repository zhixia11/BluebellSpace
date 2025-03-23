<template>
  <div class="common-settings">
    <form-group label="基本设置">
      <form-item label="语言">
        <x-select v-model="data.language">
          <x-option value="zh-CN">中文</x-option>
          <x-option value="en-US">English</x-option>
          <x-option value="zh-TW">繁體中文</x-option>
        </x-select>
      </form-item>
    </form-group>
    <form-group label="外观设置">
      <form-item label="主题色">
        <x-select v-model="data.theme">
          <x-option value="light">亮色</x-option>
          <x-option value="dark">暗色</x-option>
          <x-option value="auto">跟随系统</x-option>
        </x-select>
      </form-item>
      <form-item label="侧边栏">
        <x-select v-model="data.layout">
          <x-option value="expand">展开</x-option>
          <x-option value="collapse">收起</x-option>
        </x-select>
      </form-item>
    </form-group>
  </div>
</template>

<script setup>
import FormGroup from '../../components/form/FormGroup.vue'
import FormItem from '../../components/form/FormItem.vue'
import XSelect from '../../components/select/XSelect.vue'
import XOption from '../../components/select/XOption.vue'
import { useSettingsStore } from '../../store/settings'
const store = useSettingsStore()
const data = store.settings.common
</script>

<style lang="less"></style>
