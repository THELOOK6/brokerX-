import { createRouter, createWebHistory } from 'vue-router'
import DefaultLayout from '../layouts/DefaultLayout.vue'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/DashboardView.vue'
import OrdersView from '../views/OrdersView.vue'
import MetricsView from '../views/MetricsView.vue'

const routes = [
  {
    path: '/',
    component: DefaultLayout,
    children: [
      { path: '', name: 'home', component: DashboardView },
      { path: '/metrics', name: 'metrics', component: MetricsView },
      { path: '/orders', name: 'orders', component: OrdersView },
      { path: '/login', name: 'login', component: LoginView },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
