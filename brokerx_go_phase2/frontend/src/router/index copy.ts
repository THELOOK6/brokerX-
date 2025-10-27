import { createRouter, createWebHistory , RouterView} from 'vue-router'
import { useRoute } from 'vue-router';

import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import OrdersView from '../views/OrdersView.vue'
import MetricsView from '../views/MetricsView.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: DashboardView
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/orders',
      name: 'orders',
      component: OrdersView
    },
    {
      path: '/metrics',
      name: 'Metrics',
      component: MetricsView
    }
  ]
})

export default router
