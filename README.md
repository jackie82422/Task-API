# Task API

This is a sample server for managing tasks. It is built with Go using the Gin framework, and it supports both in-memory storage and MySQL storage.

## Features

- RESTful API for managing tasks
- Supports in-memory and MySQL storage
- OpenAPI documentation with Swagger UI

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Documentation](#api-documentation)

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/jackie82422/Task-API.git
    cd task-api
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Install Docker and Docker Compose if you haven't already.

## Usage

### Running with Docker Compose

1. Build and start the services:

    ```sh
    docker build -f build/docker/api/Dockerfile -t task_api:v1.0.0 .
    docker-compose -f ./build/docker-compose.yml up -d --build
    ```

2. The API server will be running at `http://localhost:8080`.

### Running Locally

1. Set up MySQL containers if you want to use them:

    ```sh
    docker-compose up mysql
    ```

2. Run the API server:

    ```sh
    go run cmd/api/main.go
    ```

3. The API server will be running at `http://localhost:8080`.

### Environment Variables

- `STORAGE_TYPE`: Set to  `mysql` for MySQL storage or use in-memory storage.
- `MYSQL_DSN`: The DSN (Data Source Name) for MySQL connection, e.g., `root:root@tcp(localhost:3306)/TaskDB`.

## API Documentation

### Swagger UI

The API documentation is available at `http://localhost:8080/swagger/index.html`.

### OpenAPI JSON

The OpenAPI JSON document is available at `http://localhost:8080/openapi.json`.

## Project Structure

```plaintext
.
├── README.md
├── build
│   ├── docker
│   │   ├── api
│   │   │   └── Dockerfile
│   │   ├── mysql
│   │   │   └── Dockerfile
│   │   └── redis
│   │       └── Dockerfile
│   └── docker-compose.yml
├── cmd
│   └── api
│       ├── main.go
│       └── openapi.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── injector.go
│   ├── domain
│   │   └── task
│   │       ├── repository.go
│   │       └── task.go
│   ├── handlers
│   │   └── task_handler.go
│   ├── infrastructure
│   │   ├── persistence
│   │   │   ├── memory
│   │   │   │   └── repository.go
│   │   │   └── mysql
│   │   │       └── repository.go
│   │   └── redis
│   ├── mocks
│   │   ├── mock_task_repository.go
│   │   └── mock_task_service.go
│   └── services
│       └── task
│           ├── service.go
│           └── impl.go
├── migrations
│   └── init.sql
