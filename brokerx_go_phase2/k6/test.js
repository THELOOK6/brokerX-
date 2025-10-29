
import http from 'k6/http';
import { sleep, check } from 'k6';
export const options = {
  stages: [{ duration: '30s', target: 20 }, { duration: '1m', target: 50 }, { duration: '30s', target: 0 }],
};
const BASE = __ENV.BASE || 'http://localhost:8080/api/v1';
export default function () {
  let login = http.post(`${BASE}/auth/login`, JSON.stringify({ email: "user@example.com", password: "password" }), { headers: { 'Content-Type': 'application/json' } });
  if (login.status === 200) {
    const token = JSON.parse(login.body).token;
    let headers = { Authorization: `Bearer ${token}`, 'Content-Type': 'application/json' }
    let q = http.get(`${BASE}/quotes?symbol=AAPL`, { headers });
    check(q, { 'quote 200': (r) => r.status === 200 });
  }
  sleep(1);
}
