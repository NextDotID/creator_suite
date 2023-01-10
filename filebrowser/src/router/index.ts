import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import routes from './routes'
import { websiteTitle } from '@/config'

const router = createRouter({
  // history: createWebHistory(import.meta.env.BASE_URL),
  history: createWebHashHistory(),
  routes,
})

router.afterEach((to) => {
  const { title } = to.meta
  document.title = title ? `${title} - ${websiteTitle}` : websiteTitle
})

export default router
