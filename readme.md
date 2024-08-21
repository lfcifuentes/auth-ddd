# Simple Auth Services

Example project of an authentication system using Go

## Tech Stack
 - [Go 1.23](https://go.dev/doc/install)
 - [Gin](https://gin-gonic.com/)
 - [Postgres](https://www.postgresql.org)
 - [Docker](https://www.docker.com)
   - [Docker Compose](https://docs.docker.com/compose/)


## Installation

Follow these steps to install the project:

1. Clone the repository:
    ```sh
    git clone https://github.com/lfcifuentes/auth-ddd
    ```
2. Navigate to the project directory:
    ```sh
    cd auth-ddd
    ```
3. Install the dependencies:
    ```sh
    go mod tidy
    ```

# Makefile
Here is a list of the available commands in the Makefile and their description:

Starts the Docker containers in the background.
```sh
make docker_up
```

Stops and removes the Docker containers.
```sh
make docker_down:
```

Generates the API documentation using swag.
```sh
make docs_generate
```

Formats the documentation code using swag.
```sh
make docs_format
```

### Migrations

Run migrations
```sh
make migrate
```

Add new migration
```sh
 make add_migration name=mew_migration_name
```

Rollback all migrations
```sh
make migrate_down
```

Rollback a specific migration
```sh
make migrate_down steps=1
```

