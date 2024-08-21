# Simple Auth Services

Example project of an authentication system using Go

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
```
make docs_format
```
