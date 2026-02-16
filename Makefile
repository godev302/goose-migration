DB_DRIVER=postgres
# Extract values from docker-compose.yml
# In a real scenario, we might use a more robust way to parse YAML, 
# but for this tutorial, we'll keep it simple or just hardcode based on what we know.
# Port is 5433, User is postgres, DB is postgres, Password is postgres
DB_DSN="host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
MIGRATION_DIR=migrations

#.PHONY: up up-to down down-to status create run-go

up:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) up

up-to:
	@read -p "Enter version to up-to: " version; \
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) up-to $$version

down:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) down

down-to:
	@read -p "Enter version to down-to: " version; \
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) down-to $$version

status:
	goose -dir $(MIGRATION_DIR) $(DB_DRIVER) $(DB_DSN) status

create:
	@read -p "Enter migration name: " name; \
	goose -s -dir $(MIGRATION_DIR) create $$name sql

run-go:
	go run main.go
