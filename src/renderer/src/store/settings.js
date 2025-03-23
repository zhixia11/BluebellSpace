import { defineStore } from 'pinia'
import { reactive, watch, toRaw } from 'vue'

const DEFAULT_SETTINGS = {
  common: {
    theme: 'light',
    language: 'zh-CN',
    layout: 'expand'
  },
  paths: {
    repetitive: [],
    similar: [],
    empty_dir: [],
    analyse: ''
  }
}
const STORAGE_KEY = 'app'
const STORAGE_SETTINGS = await window.api.getConfig(STORAGE_KEY)

export const useSettingsStore = defineStore('settings', () => {
  const settings = reactive({ ...DEFAULT_SETTINGS, ...STORAGE_SETTINGS })
  watch(settings, (value) => {
    window.api.setConfig(STORAGE_KEY, toRaw(value))
  }, { deep: true })
  return { settings }
})