BINARY_DIR := ../deployment/bin
LOG_DIR    := /tmp/advertisement-logs

# ─── Service targets ──────────────────────────────────────────────────────────

.PHONY: run-auth run-ad run-gateway
.PHONY: run-all-sleep-mode stop-all-sleep-mode check-all-sleep-mode
.PHONY: build-all
.PHONY: migrate-up migrate-down migrate-create migrate-version
.PHONY: dependencies-up dependencies-down
.PHONY: generate-auth generate-ad

run-auth:
	cd services/auth/rpc && go run auth.go -f configs/auth.yaml

run-ad:
	cd services/ad/rpc && go run ad.go -f configs/ad.yaml

run-gateway:
	cd services/gateway/api && go run gateway.go -f configs/gateway.yaml

# ─── Sleep-mode helpers ───────────────────────────────────────────────────────

define _stop-service
	@PID=$$(lsof -ti :$(1) 2>/dev/null); \
	if [ -n "$$PID" ]; then \
		kill -SIGTERM $$PID 2>/dev/null; sleep 1; \
		kill -9 $$PID 2>/dev/null || true; \
		echo "  stopped $(2) (pid $$PID)"; \
	else \
		echo "  $(2) not running"; \
	fi
endef

define _check-service
	@PORT=$(1); NAME=$(2); \
	if lsof -ti :$$PORT > /dev/null 2>&1; then \
		printf "  \033[32m●\033[0m %-20s (:%s)\n" "$$NAME" "$$PORT"; \
	else \
		printf "  \033[31m○\033[0m %-20s (:%s)\n" "$$NAME" "$$PORT"; \
	fi
endef

run-all-sleep-mode:
	@mkdir -p $(LOG_DIR)
	@nohup caffeinate -s sh -c 'cd services/auth/rpc    && go run auth.go    -f configs/auth.yaml'    > $(LOG_DIR)/auth.log    2>&1 & echo "auth    → $(LOG_DIR)/auth.log"
	@sleep 2
	@nohup caffeinate -s sh -c 'cd services/ad/rpc      && go run ad.go      -f configs/ad.yaml'      > $(LOG_DIR)/ad.log      2>&1 & echo "ad      → $(LOG_DIR)/ad.log"
	@sleep 2
	@nohup caffeinate -s sh -c 'cd services/gateway/api && go run gateway.go  -f configs/gateway.yaml' > $(LOG_DIR)/gateway.log 2>&1 & echo "gateway → $(LOG_DIR)/gateway.log"
	@echo "All services started in sleep mode."

stop-all-sleep-mode:
	$(call _stop-service,9001,auth-rpc)
	$(call _stop-service,9002,ad-rpc)
	$(call _stop-service,8080,gateway)
	@echo "All services stopped."

check-all-sleep-mode:
	@echo ""
	@echo "  ┌─────────────────────────────────────────┐"
	@echo "  │        Advertisement — Service Status   │"
	@echo "  └─────────────────────────────────────────┘"
	$(call _check-service,9001,auth-rpc)
	$(call _check-service,9002,ad-rpc)
	$(call _check-service,8080,gateway)
	$(call _check-service,5432,postgres)
	$(call _check-service,6379,redis)
	$(call _check-service,2379,etcd)
	@echo ""

# ─── Build ────────────────────────────────────────────────────────────────────

build-all:
	@mkdir -p $(BINARY_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_DIR)/auth    ./services/auth/rpc/auth.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_DIR)/ad      ./services/ad/rpc/ad.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BINARY_DIR)/gateway ./services/gateway/api/gateway.go
	@echo "Binaries → $(BINARY_DIR)/"

# ─── Migrations ───────────────────────────────────────────────────────────────

DB_URL ?= postgres://postgres:password@127.0.0.1:5432/advertisement?sslmode=disable

migrate-up:
	migrate -path ../migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ../migrations -database "$(DB_URL)" down 1

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir ../migrations -seq $$name

migrate-version:
	migrate -path ../migrations -database "$(DB_URL)" version

# ─── Dependencies ─────────────────────────────────────────────────────────────

dependencies-up:
	docker compose -f ../deployment/dependencies/docker-compose.yaml up -d
	@echo "Waiting for services to be ready..."
	@sleep 5

dependencies-down:
	docker compose -f ../deployment/dependencies/docker-compose.yaml down

# ─── Tidy ─────────────────────────────────────────────────────────────────────

tidy:
	go mod tidy
