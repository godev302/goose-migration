# Pressly Goose Migration Tutorial

This tutorial demonstrates how to use `goose` for database migrations with a PostgreSQL database running in Docker.

## Prerequisites

- `goose` CLI tool installed.
- `docker` and `docker compose` installed.
- `make` installed.

## Setup

### 1. Start the Database

The project includes a `docker-compose.yml` file. Start the PostgreSQL and Adminer containers:

```bash
docker compose up -d
```

### 2. Database Administration with Adminer

Once the containers are running, you can access **Adminer** (a lightweight database management tool) to view your database:

- **URL**: `http://localhost:8080`
- **System**: PostgreSQL
- **Server**: `user-postgres.diwakar`
- **Username**: `postgres`
- **Password**: `postgres`
- **Database**: `postgres`

### 3. Configuration

The `Makefile` is configured to read database connection details. It uses the following connection string:
`host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable`

## Goose Commands

We have created 4 migrations in the `migrations/` directory:
1. `00001_create_users_table.sql`
2. `00002_add_email_to_users.sql`
3. `00003_create_posts_table.sql`
4. `00004_add_title_to_posts.sql`

### Check Status

View the current state of migrations:

```bash
make status
```

### Apply Migrations (Up)

Apply all pending migrations:

```bash
make up
```

### Rollback (Down)

Roll back the last applied migration:

```bash
make down
```

### Migrate Up to a Specific Version

Migrate up to a specific version (e.g., version 2):

```bash
make up-to
# When prompted, enter: 2
```

### Rollback to a Specific Version

Roll back to a specific version (e.g., version 1):

```bash
make down-to
# When prompted, enter: 1
```

To roll back all migrations, use version `0`:
```bash
make down-to
# When prompted, enter: 0
```

### Create a New Migration

Generate a new SQL migration file:

```bash
make create
# When prompted, enter: name_of_migration
```

## Running Migrations via Go Code

This project also demonstrates how to run migrations directly using Go code with the `github.com/pressly/goose/v3` package.

### 1. Build and Run the Tool

The `main.go` file contains the implementation for running migrations. You can simply run it with:

```bash
go run main.go
```

To change which command is executed (e.g., `up`, `down`, `status`), open `main.go` and uncomment the desired function call in the `main()` function.

### 2. Manual Execution via Makefile

You can also use the `Makefile` shortcut:

```bash
make run-go
```

The Go implementation uses `github.com/pressly/goose/v3` and `github.com/lib/pq` to connect to the database and manage migrations programmatically.

## Summary of Goose CLI Usage

If you want to run `goose` directly without the Makefile, you can set environment variables:

```bash
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"

goose -dir migrations status
goose -dir migrations up
goose -dir migrations down
goose -dir migrations up-to 2
goose -dir migrations down-to 1
```
