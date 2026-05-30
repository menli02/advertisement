# Advertisement Platform

Classified ads microservices platform — rewritten from a Gin monolith to a production-grade go-zero architecture.

```
┌──────────────────────────────────────────────────────────┐
│  Client (iOS / Android / React)                          │
└───────────────────────┬──────────────────────────────────┘
                        │ HTTP/REST :8080
┌───────────────────────▼──────────────────────────────────┐
│  Gateway (go-zero REST)                                  │
│  services/gateway/api                                    │
└──────────┬────────────────────────┬──────────────────────┘
           │ gRPC :9001             │ gRPC :9002
┌──────────▼──────────┐  ┌──────────▼──────────────────────┐
│  Auth RPC           │  │  Ad RPC                         │
│  services/auth/rpc  │  │  services/ad/rpc                │
└──────────┬──────────┘  └──────────┬────────────────────  ┘
           │                        │
    ┌──────▼──────┐         ┌───────▼──────┐
    │ PostgreSQL  │         │ PostgreSQL   │
    │ Redis (OTP) │         │ Redis (cache)│
    └─────────────┘         └──────────────┘
           │                        │
           └──────────┬─────────────┘
                   etcd
              (service discovery)
```

## Services

| Service     | Port | Description                          |
|-------------|------|--------------------------------------|
| gateway     | 8080 | REST API — public entry point        |
| auth-rpc    | 9001 | OTP + JWT auth service               |
| ad-rpc      | 9002 | Advertisement CRUD + categories      |
| PostgreSQL  | 5432 | Primary database                     |
| Redis       | 6379 | OTP cache + ad view counters         |
| etcd        | 2379 | Service discovery                    |
| Swagger UI  | 7081 | API documentation                    |

## Quick Start

```bash
# 1. Start dependencies
cd server && make dependencies-up

# 2. Run migrations
make migrate-up

# 3. Start all services
make run-all-sleep-mode

# 4. Check status
make check-all-sleep-mode

# 5. Open docs
open http://localhost:7081
```

## API Endpoints

### Auth
| Method | Path                 | Auth | Description           |
|--------|----------------------|------|-----------------------|
| POST   | /v1/auth/send-otp    | —    | Send OTP to phone     |
| POST   | /v1/auth/verify-otp  | —    | Verify OTP → temp token |
| POST   | /v1/auth/sign-in     | —    | Sign in → access + refresh tokens |
| POST   | /v1/auth/renew-token | —    | Rotate tokens         |

### Advertisements
| Method | Path          | Auth | Description        |
|--------|---------------|------|--------------------|
| GET    | /v1/ads       | —    | List (search, filter, paginate) |
| POST   | /v1/ads       | ✓    | Create             |
| GET    | /v1/ads/:id   | —    | Get by ID (increments view) |
| PUT    | /v1/ads/:id   | ✓    | Update (owner only) |
| DELETE | /v1/ads/:id   | ✓    | Soft delete (owner only) |

### Categories
| Method | Path            | Auth | Description  |
|--------|-----------------|------|--------------|
| GET    | /v1/categories  | —    | List all     |
| POST   | /v1/categories  | ✓    | Create       |

## Auth Flow

```
POST /v1/auth/send-otp   → { otp_id, exp_sec }
POST /v1/auth/verify-otp → { temp_token }          (valid 5 min)
POST /v1/auth/sign-in    → { access_token, refresh_token, user_id }
POST /v1/auth/renew-token → { access_token, refresh_token }
```

All protected routes require: `Authorization: Bearer <access_token>`

## Makefile Reference

```bash
make run-all-sleep-mode      # start all services in background
make stop-all-sleep-mode     # stop all services
make check-all-sleep-mode    # colored status table
make build-all               # linux/amd64 binaries → deployment/bin/
make migrate-up              # apply all migrations
make migrate-down            # rollback one
make migrate-create          # create new migration
make dependencies-up         # docker-compose up
make dependencies-down       # docker-compose down
make tidy                    # go mod tidy
```

## DB Schema

```
users                 — id, phone, first_name, last_name
categories            — id, name, slug
advertisements        — id, user_id, category_id, title, description,
                        slug, price, currency, status, view_count, deleted_at
advertisement_images  — id, ad_id, url, position
```

## Config Placeholders

Replace placeholders before running:

| Placeholder            | File                           |
|------------------------|-------------------------------|
| `<DB_PASSWORD>`        | configs/auth.yaml, configs/ad.yaml |
| `<TEMP_JWT_SECRET>`    | configs/auth.yaml              |
| `<JWT_SECRET>`         | configs/auth.yaml, configs/gateway.yaml |
| `<REFRESH_JWT_SECRET>` | configs/auth.yaml              |

## Architecture Improvements Over Original

| Old (Gin monolith)          | New (go-zero microservices)         |
|-----------------------------|-------------------------------------|
| Single service              | Auth RPC + Ad RPC + Gateway         |
| MySQL + GORM                | PostgreSQL + pgx/v5 (no ORM)        |
| No auth                     | OTP → JWT (access + refresh + temp) |
| SQL injection via fmt.Sprintf | Parameterized queries + whitelist  |
| context.Background() everywhere | Request context propagated      |
| No soft delete              | deleted_at TIMESTAMPTZ              |
| No search                   | ILIKE + GIN full-text index         |
| No slug                     | gosimple/slug with uniqueness loop  |
| No view counting            | DB increment + Redis sync           |
| No categories               | categories table with slug          |
| No image management         | advertisement_images with position  |
| No pagination               | page + page_size on all list routes |
