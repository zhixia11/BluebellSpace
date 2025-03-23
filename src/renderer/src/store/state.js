import { defineStore } from 'pinia'
export const useStateStore = defineStore('state', {
  state: () => ({
    running: false,
    completed: false
  }),
  actions: {
    setRunning(value) {
      this.running = value
    },
    setCompleted(value) {
      this.completed = value
    }
  }
})