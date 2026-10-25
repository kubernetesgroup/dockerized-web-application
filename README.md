code
Markdown
# Enterprise Production-Ready Docker Suite

[![Docker](https://img.shields.io/badge/Docker-Compose-blue?logo=docker)](https://www.docker.com/)
[![Nginx](https://img.shields.io/badge/Reverse_Proxy-Nginx-green?logo=nginx)](https://nginx.org/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL_16-336791?logo=postgresql)](https://www.postgresql.org/)

An enterprise-ready, containerized microservice foundation designed for production deployments on VPS (AWS EC2, DigitalOcean, Hetzner, Linode).

## Architecture Flow

```text
[Internet Traffic] --> (Port 80)
       |
       v
[Nginx Reverse Proxy] (DDoS Rate Limiting + Security Headers)
       |
       +---> [Route /]     --> [SPA Frontend Container] (Alpine Nginx)
       |
       +---> [Route /api/] --> [Microservice API] (Non-Root Golang)
                                    |
                                    v
                               [PostgreSQL 16 DB] <--- [Automated Backup Service] (Cron pg_dump)
Key Deliverables & Features
Edge Reverse Proxy: Nginx configured with DDoS rate limiting, Gzip compression, and strict security headers.
Minimal Attack Surface: All microservices run using multi-stage builds and non-root users (UID 10001).
Automated Disaster Recovery: Self-contained backup service running automated compressed pg_dump with automated 7-day file retention.
Zero-Loss Data Persistence: Production-grade named Docker volumes for transactional data and backups.
Getting Started
Initialize environment configuration:
code
Bash
cp .env.example .env
Start the production stack:
code
Bash
docker compose up -d
Trigger an on-demand database backup:
code
Bash
docker compose up backup-service
Endpoints
Web App UI: http://localhost/
Healthcheck API: http://localhost/api/v1/health
code
Code
