import './assets/main.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
import 'virtual:svg-icons-register'

const app = createApp(App)
app.use(router).use(createPinia()).use(ArcoVue).mount('#app')
