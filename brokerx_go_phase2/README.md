
# BrokerX+ Phase 2 (Go/Fiber) – Scaffold

Aligned with the Cahier des charges: UC-01..UC-08, JWT auth, idempotent orders, SSE quotes, Postgres, Redis, NGINX LB, Kong gateway, Prometheus/Grafana, Vue 3 frontend.

## Run
```bash
docker compose up --build
```
- Frontend: http://localhost:8080
- API via LB: http://localhost:8080/api/v1/...
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (admin/admin)

## Notes
- Enable Postgres extension `pgcrypto` for `gen_random_uuid()` (or switch to app-side UUIDs).
- Replace JWT secret in production.
- k6: `BASE=http://localhost:8080/api/v1 k6 run k6/test.js`
