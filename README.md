# CRUD Storage API

This is a RESTful API that provides CRUD (Create, Read, Update, Delete) operations for articles, categories, and items, with persistent storage in a database and/or file system.


## File Structure

```
crud-storage-api/
├── cmd
|   └── main.go
├── config
|   └── env.go
├── internal
|   ├── article
|   |   ├── apis
|   |   |   └── article.go
|   |   ├── models
|   |   |   ├── orms
|   |   |   |   └── article.go
|   |   |   ├── requests
|   |   |   |   └── article.go
|   |   |   └── responses
|   |   |       └── article.go
|   |   └── services
|   |       └── article.go
|   ├── category
|   |   ├── apis
|   |   |   └── category.go
|   |   ├── models
|   |   |   ├── orms
|   |   |   |   └── category.go
|   |   |   ├── requests
|   |   |   |   └── category.go
|   |   |   └── responses
|   |   |       └── category.go
|   |   └── services
|   |       └── category.go
|   ├── item
|   |   ├── apis
|   |   |   └── item.go
|   |   ├── models
|   |   |   ├── orms
|   |   |   |   └── item.go
|   |   |   ├── requests
|   |   |   |   └── item.go
|   |   |   └── responses
|   |   |       └── item.go
|   |   └── services
|   |       └── item.go
|   └── server
|       ├── router.go
|       └── server.go
├── scripts
|   ├── build.sh
|   ├── env.sh
|   └── run.sh
├── shared
|   ├── api
|   |   ├── api.go
|   |   ├── request.go
|   |   └── response.go
|   ├── service
|   |   └── service.go
|   └── storage
|       ├── models
|       |   └── orm.go
|       ├── db.go
|       ├── dynamic.go
|       ├── file.go
|       └── storage.go
├── .env
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── go.sum

```

## File Structure Details

The app's file structure is organized as follows:

- `cmd/main.go:` the main entry point of the app, which starts the server and sets up the environment.
- `config/env.go:` a configuration file that loads environment variables from .env file(s) and sets default values.
- `internal/article:` a package that handles operations related to articles, including HTTP API endpoints, models, and services.
- `internal/category:` a package that handles operations related to categories, including HTTP API endpoints, models, and services.
- `internal/item:` a package that handles operations related to items, including HTTP API endpoints, models, and services.
- `internal/server:` a package that handles the HTTP server and router configuration.
- `scripts/build.sh:` a shell script that builds the app and generates a binary executable file.
- `scripts/env.sh:` a shell script that sets environment variables from .env file(s) and exports them.
- `scripts/run.sh:` a shell script that runs the app locally using the binary executable file.
- `shared/api:` a package that defines the API request and response structures and interfaces.
- `shared/service:` a package that defines the service interfaces.
- `shared/storage:` a package that provides storage interfaces and implementations for different types of storage, such as database and file system.
- `.env:` a file that stores environment variables that are loaded by config/env.go.
- `docker-compose.yml:` a Docker Compose file that defines the app's containerized environment.
- `Dockerfile:` a Dockerfile that builds a Docker image of the app.
- `go.mod, go.sum:` files that store the Go module dependencies.

## Usage

To run the app locally, you can use the following steps:

1. Set up the environment variables by creating a .env file in the root directory and setting the variables as key-value pairs (see .env.sample for an example).
3. Run the shell script `./scripts/run.sh -b` to build and run the app.
4. The app should now be running and listening on the default port 8080. You can access the API endpoints via a web browser or a HTTP client tool, such as cURL or Postman.

To build and run the app using Docker, you can use the following steps:

1. Set up the environment variables by creating a .env file in the root directory and setting the variables as key-value pairs (see .env.sample for an example).
2. Run the command docker-compose up in the root directory to build and run the Docker containers.
3. The app should now be running and listening on the default port 8080. You can access the API endpoints via a web browser or a HTTP client tool, such as cURL or Postman.
