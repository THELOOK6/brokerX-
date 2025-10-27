<template>
  <div class="dashboard">
    <h1 class="title">Dashboard</h1>

    <!-- Stat Cards -->
    <div class="cards">
      <div v-for="card in summary" :key="card.label" class="card">
        <p class="label">{{ card.label }}</p>
        <h2 class="value">{{ card.value }}</h2>
      </div>
    </div>

    <!-- Content Grid -->
    <div class="content-grid">
      <!-- Recent Orders -->
      <div class="panel orders">
        <h3 class="panel-title">Recent Orders</h3>
        <table class="orders-table">
          <thead>
            <tr>
              <th>Order</th>
              <th>Customer</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="o in orders" :key="o.id">
              <td>#{{ o.id }}</td>
              <td>{{ o.customer }}</td>
              <td>{{ o.date }}</td>
              <td>${{ o.amount }}</td>
              <td>
                <span class="status" :class="o.status.toLowerCase()">
                  {{ o.status }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Revenue Chart -->
      <div class="panel chart">
        <h3 class="panel-title">Revenue</h3>
        <canvas id="revenueChart" class="chart-canvas"></canvas>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Chart from 'chart.js/auto'

const summary = ref([
  { label: 'Total Orders', value: '12,304' },
  { label: 'Total Users', value: '3,210' },
  { label: 'Revenue', value: '$56,431' },
  { label: 'Profit', value: '$12,056' },
])

const orders = ref([
  { id: 3210, customer: 'Acme Corp', date: '01/22/2024', amount: '3,200', status: 'Pending' },
  { id: 3252, customer: 'Acme Corp', date: '01/22/2024', amount: '2,900', status: 'Completed' },
  { id: 3211, customer: 'Acme Corp', date: '01/22/2024', amount: '1,750', status: 'Completed' },
  { id: 3212, customer: 'Acme Corp', date: '01/22/2024', amount: '4,100', status: 'Refunded' },
])

onMounted(() => {
  const ctx = document.getElementById('revenueChart')
  new Chart(ctx, {
    type: 'line',
    data: {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
      datasets: [
        {
          label: 'Revenue',
          data: [10000, 20000, 30000, 40000, 48000, 56000],
          borderColor: '#3b82f6',
          backgroundColor: 'rgba(59,130,246,0.1)',
          tension: 0.4,
          fill: true,
        },
      ],
    },
    options: {
      plugins: { legend: { display: false } },
      scales: {
        x: { ticks: { color: '#9ca3af' }, grid: { color: 'rgba(255,255,255,0.05)' } },
        y: { ticks: { color: '#9ca3af' }, grid: { color: 'rgba(255,255,255,0.05)' } },
      },
    },
  })
})
</script>

<style scoped>
.dashboard {
  color: var(--text-color);
}
.title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--accent-color);
  margin-bottom: 1.5rem;
}

/* Stat Cards */
.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 1rem;
  margin-bottom: 2rem;
}
.card {
  background: var(--bg-secondary);
  padding: 1.5rem;
  border-radius: 12px;
  border: 1px solid var(--border-color);
  text-align: left;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.25);
  transition: transform 0.2s;
}
.card:hover {
  transform: translateY(-2px);
}
.label {
  font-size: 0.9rem;
  color: var(--text-muted);
}
.value {
  font-size: 1.8rem;
  font-weight: 600;
  margin-top: 0.3rem;
}

/* Panels */
.content-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}
.panel {
  background: var(--bg-secondary);
  border: 1px solid var(--border-color);
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.25);
}
.panel-title {
  font-size: 1.1rem;
  font-weight: 500;
  margin-bottom: 1rem;
  color: var(--text-color);
}

/* Table */
.orders-table {
  width: 100%;
  border-collapse: collapse;
}
.orders-table th {
  text-align: left;
  font-size: 0.9rem;
  color: var(--text-muted);
  padding-bottom: 0.75rem;
}
.orders-table td {
  padding: 0.6rem 0;
  font-size: 0.95rem;
  color: var(--text-color);
  border-top: 1px solid var(--border-color);
}
.status {
  padding: 0.2rem 0.6rem;
  border-radius: 6px;
  font-size: 0.8rem;
  font-weight: 500;
}
.status.pending {
  background: rgba(59, 130, 246, 0.2);
  color: #3b82f6;
}
.status.completed {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
}
.status.refunded {
  background: rgba(239, 68, 68, 0.2);
  color: #ef4444;
}

/* Chart */
.chart-canvas {
  width: 100%;
  height: 260px;
}
</style>
