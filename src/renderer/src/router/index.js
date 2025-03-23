import { createRouter, createWebHashHistory } from 'vue-router'
import { useStateStore } from '../store/state'
import Modal from '../components/dialog'
const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/repetitive'
    },
    {
      path: '/repetitive',
      name: 'repetitive',
      component: () => import('../views/RepetitiveFileView.vue')
    },
    {
      path: '/similar',
      name: 'similar',
      component: () => import('../views/SimilarImageView.vue')
    },
    {
      path: '/empty',
      name: 'empty',
      component: () => import('../views/EmptyDirView.vue')
    },
    {
      path: '/analyse',
      name: 'analyse',
      component: () => import('../views/AnalyseView.vue')
    },
    {
      path: '/settings',
      name: 'settings',
      component: () => import('../views/settings/SettingsView.vue')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    }
  ]
})

router.beforeEach((to, from, next) => {
  const store = useStateStore()
  if (!!store.running) {
    Modal.show({
      width: 300,
      height: 160,
      message: '当前任务正在运行中，确认要离开吗',
      onOk: () => {
        store.setRunning(false)
        next()
      }
    })
  } else if (!!store.completed) {
    Modal.show({
      width: 300,
      height: 160,
      message: '当前任务还未处理完毕，确认要离开吗',
      onOk: () => {
        store.setCompleted(false)
        next()
      },
    })
  } else {
    next()
  }
})

export default router
