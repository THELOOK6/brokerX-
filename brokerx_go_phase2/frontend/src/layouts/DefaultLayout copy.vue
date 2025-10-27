<template>
  <div class="layout">
    <aside class="sidebar">
      <Sidebar />
    </aside>
    <main class="content">
      <!-- dynamic page will be rendered here -->
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue'
</script>

<style scoped>
.layout {
  display: flex;
  height: 100vh;
}
.sidebar {
  width: 220px;
  border-right: 1px solid #e0e0e0;
  background-color: #fff;
}
.content {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}
</style>
