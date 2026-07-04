include .env


migrations:
	@goose -dir=internal/migrations/schema create ${name} sql

mu: migrate/up
migrate/up:
	@echo "Running up migrations..."
	@goose -dir=internal/migrations/schema postgres ${DB_URL} up

md: migrate/down
migrate/down:
	@echo "Rolling back migrations..."
	@goose -dir=internal/migrations/schema postgres ${DB_URL} down

mf: migrate/fresh
migrate/fresh:
	@echo "Dropping..."
	@goose -dir=internal/migrations/schema postgres ${DB_URL} reset
	@echo "Migrating..."
	@$(MAKE) --no-print-directory migrate/up
