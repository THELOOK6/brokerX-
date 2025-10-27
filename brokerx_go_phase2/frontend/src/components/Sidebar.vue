<template>
  <div class="sidebar" :class="{ collapsed }">
    <!-- Brand -->
    <div class="logo-area" @click="toggleCollapse">
      <span class="logo-text" v-if="!collapsed">BrokerX</span>
      <el-icon v-else><component :is="collapsed ? Expand : Fold" /></el-icon>
    </div>

    <!-- Navigation -->
    <el-menu
      :default-active="activePath"
      class="menu"
      :collapse="collapsed"
      background-color="transparent"
      text-color="#a0a0b0"
      active-text-color="#409eff"
      router
    >
      <el-menu-item index="/">
        <el-icon><House /></el-icon>
        <span>Dashboard</span>
      </el-menu-item>

      <el-menu-item index="/orders">
        <el-icon><ShoppingCart /></el-icon>
        <span>Orders</span>
      </el-menu-item>

      <el-menu-item index="/metrics">
        <el-icon><DataAnalysis /></el-icon>
        <span>Metrics</span>
      </el-menu-item>

      <el-menu-item index="/account">
        <el-icon><User /></el-icon>
        <span>Account</span>
      </el-menu-item>

      <el-menu-item index="/login">
        <el-icon><User /></el-icon>
        <span>Login</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import {
  Fold,
  Expand,
  House,
  ShoppingCart,
  DataAnalysis,
  User,
} from '@element-plus/icons-vue'

const collapsed = ref(false)
const toggleCollapse = () => (collapsed.value = !collapsed.value)

const route = useRoute()
const activePath = route.path
</script>

<style scoped>
.sidebar {
  position: fixed;
  left: 0;
  top: 0;
  bottom: 0;
  height: 100vh;
  background: #0f1117;
  border-right: 1px solid #1b1d25;
  transition: width 0.25s ease;
  display: flex;
  flex-direction: column;
  z-index: 1000;
  overflow: hidden;
}

.sidebar:not(.collapsed) {
  width: 220px;
}

.sidebar.collapsed {
  width: 72px;
}

/* Logo */
.logo-area {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 64px;
  border-bottom: 1px solid #1b1d25;
  cursor: pointer;
  color: #409eff;
  font-weight: 600;
  font-size: 20px;
  transition: all 0.3s ease;
}
.logo-area:hover {
  background-color: #1b1d25;
}
.logo-text {
  letter-spacing: 0.5px;
}

/* Menu */
.menu {
  border-right: none;
  padding-top: 12px;
  background: transparent;
}
.el-menu-item {
  font-size: 15px;
  border-radius: 8px;
  margin: 4px 8px;
  transition: background-color 0.2s ease;
}
.el-menu-item:hover {
  background-color: #1b1d25 !important;
}
.el-menu-item.is-active {
  background-color: #1b1d25 !important;
  color: #409eff !important;
}
</style>
