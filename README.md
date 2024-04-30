# Cerebro: API Gateway for Microservices

Cerebro is a Golang project focused on developing an API gateway, connected to microservices using gRPC connections. It acts as a central entry point for client requests, providing functionalities such as routing, authentication, and load balancing to the underlying microservices.

## Features

- API routing and forwarding
- Authentication and authorization
- Load balancing and service discovery
- Logging and monitoring

## Structure

```
    cerebro/
        ├── api/
        │    ├── core/
        │    │    ├── controllers/
        │    │    └── grpc/
        │    ├── routes/
        ├── cmd/
        │    ├── main.go
        ├── internal/
        │    ├── config/
        │    ├── log/
        │    ├── model/
        ├── pkg/
        │    ├── utils/
        ├── scripts/
        ├── .dockerignore
        ├── config.json
        ├── go.mod
        ├── go.sum
        ├── README.md
        └── Dockerfile
```

* `api/` This directory contains the API-related code.
    * core/: Core functionality of the API.
        * controllers/: Controllers for handling HTTP requests.
        * grpc/: gRPC related code.
    * routes/: API routes definition.

* `cmd/` Main application entry point.

* `internal/`: Internal package containing application-specific code.
    * config/: Configuration related code.
    * log/: Logging functionality.
    * model/: Data models used by the application.

* `pkg/`: Package directory containing reusable code.
    * utils/: Utility functions and libraries.

* `scripts/`: Directory containing scripts for building, testing, and running the application.

* `Dockerfile`: File used to build a Docker image for the project.
* `.dockerignore`: File specifying paths to exclude when building Docker images.
* `config.json`: Configuration file for the application.
* `go.mod`, `go.sum`: Go module files specifying dependencies for the project.
* README.md: Markdown file containing project information and instructions.

This structure organizes the project into separate directories for different components, making it easier to manage and maintain the codebase.

## Frameworks & Libraries

The project relies on the following frameworks and libraries:

- [Fiber v3](https://github.com/gofiber/fiber/v3) v3.0.0-beta.2
- [Logrus](https://github.com/sirupsen/logrus) v1.9.3
- [Viper](https://github.com/spf13/viper) v1.18.2
- [gRPC](https://google.golang.org/grpc) v1.63.2
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf) v1.33.0


## Usage

### Building Docker Images

To build the Docker images for Cerebro, follow these steps:

1. Clone the Cerebro repository:

   ```bash
   git clone https://github.com/HouseCham/cerebro.git
   ```

2. Navigate to the project directory:
    ```bash
    cd cerebro
    ```
3. Build the docker image using the provided Dockerfile:
    ```bash
    docker build -t cerebro:[version] .
    ```
### Running Docker Containers

Once you have built the Docker image, you can run containers for Cerebro using the following commands (remove the `-it` tag in case you do not want the container to be interactive) :

```bash
sudo docker run -d \
--name cerebro-api-gateway \
-p 3000:3000 \
-it \
--restart unless-stopped \
cerebro-service:0
```

Replace `3000:3000` with the desired port mapping if needed. Remember to change that configuration within the config.json file

```json
{
    "app": {
        "port": 3000,
        "host": "localhost"
    }
}
```

### Finally, in case you leaved the `-it` tag so the container is interactive, you will see the next CLI.

```bash
INFO[2024-04-30T21:57:32Z] Setting up config file
INFO[2024-04-30T21:57:32Z] Setting up Fiber instance with CORS middleware and routes
INFO[2024-04-30T21:57:32Z] Setting up gRPC connections
INFO[2024-04-30T21:57:32Z] Server is running on http://localhost:3000

    _______ __
   / ____(_) /_  ___  _____
  / /_  / / __ \/ _ \/ ___/
 / __/ / / /_/ /  __/ /
/_/   /_/_.___/\___/_/          v3.0.0-beta.2
--------------------------------------------------
INFO Server started on:         http://127.0.0.1:3000 (bound on host 0.0.0.0 and port 3000)
INFO Total handlers count:      3
INFO Prefork:                   Disabled
INFO PID:                       1
INFO Total process count:       1
```