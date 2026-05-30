DB_URL ?= postgres://postgres:password@127.0.0.1:5432/advertisement?sslmode=disable
MIGRATIONS_DIR := .

.PHONY: up down create force goto rolldown rolldown1

up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down

create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

force:
	@read -p "Version: " v; \
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$v

goto:
	@read -p "Version: " v; \
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" goto $$v

rolldown:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down -all

rolldown1:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version
