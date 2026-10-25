# Enterprise Production-Ready Docker Suite

[![Docker](https://img.shields.io/badge/Docker-Compose-blue?logo=docker)](https://www.docker.com/)
[![Nginx](https://img.shields.io/badge/Reverse_Proxy-Nginx-green?logo=nginx)](https://nginx.org/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL_16-336791?logo=postgresql)](https://www.postgresql.org/)

An enterprise-ready, containerized microservice foundation designed for production deployments on VPS (AWS EC2, DigitalOcean, Hetzner, Linode).

```mermaid
graph TD
    Client([Internet Traffic]) -->|Port 80| Nginx[Nginx Reverse Proxy\nRate Limiting + Security Headers]
    Nginx -->|Route /| Frontend[SPA Frontend Container\nNginx Alpine]
    Nginx -->|Route /api/| Backend[Microservice API\nNon-Root Golang]
    Backend -->|Internal Bridge| DB[(PostgreSQL 16\nPersistent Named Volume)]
    BackupService[Automated Backup Cron] -.->|pg_dump + gzip| DB
    BackupService -.->|Store 7-day archive| BackupVolume[(Backup Volume)]

Key Deliverables & Features
Edge Reverse Proxy: Nginx configured with DDoS rate limiting, Gzip compression, and strict security headers.
Minimal Attack Surface: All microservices run using multi-stage builds and non-root users (UID 10001).
Automated Disaster Recovery: Self-contained backup service running automated compressed pg_dump with automated 7-day file retention.
Zero-Loss Data Persistence: Production-grade named Docker volumes for transactional data and backups.
Getting Started
Initialize environment configuration:
cp .env.example .env
Start the production stack:
docker compose up -d
Trigger an on-demand database backup:
docker compose up backup-service
Endpoints
Web App UI: http://localhost/
Healthcheck API: http://localhost/api/v1/healths