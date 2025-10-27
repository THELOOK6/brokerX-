
<template>
  <div>
    <h2>Orders</h2>
    <input v-model="accId" placeholder="Account ID"/>
    <input v-model="symbol" placeholder="Symbol"/>
    <select v-model="side"><option>BUY</option><option>SELL</option></select>
    <input v-model.number="qty" type="number" placeholder="Qty"/>
    <input v-model.number="price" type="number" step="0.01" placeholder="Price"/>
    <button @click="place">Place</button>
    <pre v-if="order">{{ order }}</pre>
  </div>
</template>
<script setup>
import axios from 'axios'
import { ref } from 'vue'
const API = import.meta.env.VITE_API_BASE ?? '/api'
const accId = ref('')
const symbol = ref('AAPL')
const side = ref('BUY')
const qty = ref(1)
const price = ref(100.0)
const order = ref(null)
function headers(){ const t=localStorage.getItem('jwt'); return t?{Authorization:`Bearer ${t}`}:{}}
async function place(){
  const key = crypto.randomUUID()
  const { data } = await axios.post(`${API}/v1/orders`, { account_id: accId.value, symbol: symbol.value, side: side.value, qty: qty.value, price: price.value }, { headers: { ...headers(), "Idempotency-Key": key } })
  order.value = data
}
</script>
