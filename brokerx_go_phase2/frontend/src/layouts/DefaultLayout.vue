<template>
  <div class="layout">
    <aside class="sidebar">
      <Sidebar />
    </aside>
    <main class="content">
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
  background-color: #0f1117;
  color: #e5e7eb;
  font-family: 'Inter', 'Segoe UI', Roboto, sans-serif;
}

/* sidebar already fixed in Sidebar.vue */
.content {
  flex: 1;
  margin-left: 220px; /* adjust dynamically if you later sync collapse */
  padding: 2rem;
  overflow-y: auto;
  transition: margin-left 0.25s ease;
  background-color: #0f1117;
}
</style>
