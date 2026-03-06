import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/Home.vue'
import Chat from '../components/Chat.vue'
import Login from '../components/Login.vue'
import Settings from '../components/Settings.vue'

const routes = [
  { path: '/login', component: Login },
  { path: '/', component: Home },
  { path: '/chat', component: Chat },
  { path: '/c/:id', component: Chat, name: 'ChatSession' },
  { path: '/settings', component: Settings },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('admin_token')
  if (to.path !== '/login' && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router