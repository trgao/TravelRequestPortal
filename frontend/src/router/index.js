import { useAuthStore } from '@/stores/AuthStore'
import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import RegisterView from '@/views/RegisterView.vue'
import OnboardView from '@/views/OnboardView.vue'
import RequestFormView from '@/views/RequestFormView.vue'
import RequestsView from '@/views/RequestsView.vue'
import DashboardView from '@/views/DashboardView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView,
    },
    {
      path: '/onboard',
      name: 'onboard',
      component: OnboardView,
    },
    {
      path: '/form',
      name: 'form',
      component: RequestFormView,
    },
    {
      path: '/',
      name: 'requests',
      component: RequestsView,
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: DashboardView,
    },
  ],
})

router.beforeEach(async (to) => {
  const authStore = useAuthStore()

  const publicPages = ['/login', '/register']
  const adminPages = ['/onboard', '/dashboard']
  const requiredAuth = !publicPages.includes(to.path)
  const requiredAdmin = adminPages.includes(to.path)

  if (!authStore.isAuthenticated && requiredAuth) {
    return '/login'
  } else if (
    authStore.isAuthenticated &&
    (!requiredAuth || (requiredAdmin && !authStore.isAdmin))
  ) {
    return '/'
  }
})

export default router
