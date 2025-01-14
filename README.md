Sure! I'll include information on `golint` and how to run it in the README. Here's the updated version:

---

# Rest-API with GO

This project is a Go server that handles user registration, authentication, and management of user favorites. The server uses a SQLite database to store user and company data.

## Table of Contents
1. [Project Structure](#project-structure)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Configuration](#configuration)
5. [Running the Server](#running-the-server)
6. [Populating the Database](#populating-the-database)
7. [Example Requests](#example-requests)
8. [Testing](#testing)
9. [Linting](#linting)
10. [Frameworks Used](#frameworks-used)


## Project Structure
```plaintext
.
├── app.log
├── cmd
│   └── main.go
├── company.db
├── config.yaml
├── docker-compose.yaml
├── Dockerfile
├── go.mod
├── go.sum
├── init.sql
├── internal
│   ├── context
│   │   └── keys.go
│   ├── database
│   │   └── database.go
│   ├── handlers
│   │   ├── company_handler.go
│   │   ├── company_handlers_test.go
│   │   └── user_handler.go
│   ├── logs
│   │   └── logs.go
│   ├── middleware
│   │   ├── auth.go
│   │   └── auth_test.go
│   ├── mocks
│   │   ├── mock_company_repository.go
│   │   ├── mock_company_services.go
│   │   ├── mock_user_repository.go
│   │   └── mock_user_service.go
│   ├── models
│   │   ├── claims.go
│   │   ├── company.go
│   │   └── user.go
│   ├── repository
│   │   ├── company_repository.go
│   │   ├── company_repository_test.go
│   │   ├── user_repository.go
│   │   └── user_repository_test.go
│   ├── server
│   │   └── server.go
│   ├── services
│   │   ├── company_service.go
│   │   ├── company_service_test.go
│   │   ├── user_service.go
│   │   └── user_service_test.go
│   └── utils
│       ├── jwt_utils.go
│       └── jwt_utils_test.go
├── scripts
│   ├── company_id.txt
│   ├── requests.sh
│   └── token.txt
└── tree
```

## Prerequisites
- Go 1.22.4
- Docker
- Docker Compose
- `golint` for linting Go code

## Installation

### Clone the repository:
```sh
git clone git@github.com:alkosmas92/xm-golang.git
cd xm-golang
```

## Configuration

Edit the `config.yaml` file to set your JWT secret keys:
```yaml
jwt:
  secret_key: "your_secret_key"
  old_secret_key: "your_old_secret_key"
sqlite:
  file_name: "./company.db"
```

## Running the Server

Start the server using Docker Compose:
```sh
docker-compose up --build
```
The server will start on [http://localhost:8080](http://localhost:8080).

## Populating the Database

The database is automatically populated when the Docker container is initialized using the `init.sql` script defined in `docker-compose.yaml`.

## Example Requests

Use the `scripts/requests.sh` script to interact with the server. The commands available in `requests.sh` are:

- **Create a new user**:
  ```sh
  ./scripts/requests.sh register
  ```
- **Login for a user**:
  ```sh
  ./scripts/requests.sh login
  ```

- **Create a new company**:
  ```sh
  ./scripts/requests.sh create_company
  ```

- **Get all companies**:
  ```sh
  ./scripts/requests.sh get_companies
  ```

- **Update a company**:
  ```sh
  ./scripts/requests.sh update_company 
  ```

- **Delete a company**:
  ```sh
  ./scripts/requests.sh delete_company 
  ```

## Testing

Run the tests with the following command:
```sh
go test ./...
```

## Linting

To lint your code with `golint`, run:
```sh
golint ./...
```

## Frameworks Used

This project uses the following frameworks and libraries:
- `github.com/golang-jwt/jwt/v4` - For JWT authentication.
- `github.com/stretchr/testify` - For testing utilities.
- `github.com/mattn/go-sqlite3` - For SQLite database integration.
- `github.com/sirupsen/logrus` - For logging.
- `github.com/golang/mock` - For mocking in tests.
- `github.com/onsi/ginkgo` and `github.com/onsi/gomega` - For BDD testing.
- `github.com/google/uuid` - For UUID generation.

