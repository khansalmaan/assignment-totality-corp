# assignment-totality-corp

This repository contains a gRPC server assignment for Totality Corp.

## Overview

This project is a gRPC server implementation designed for a specific assignment at Totality Corp. It demonstrates how to build and run a gRPC server using Go, along with instructions for both Docker and non-Docker environments.

## Features

- Implements a gRPC server in Go
- Default port configuration: 50051
- Provides various gRPC services and endpoints as per the assignment requirements

## Prerequisites

Before running the server, ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.20 or later)
- [Docker](https://www.docker.com/get-started) (for Docker-based setup)

## Installation

To set up the project, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/khansalmaan/assignment-totality-corp.git
cd assignment-totality-corp
```

## Usage

### Using Docker

- **To create a Docker image from the app:**

    ```bash
    make build
    ```

- **To run the server:**

    ```bash
    make run
    ```

    The server will run on PORT: 50051 by default.

### Without Docker

- **To run the server without Docker:**

    ```bash
    go run cmd/server/main.go
    ```

    The server will run on PORT: 50051 by default.

## Testing

- **To run tests:**

    ```bash
    make test
    ```

- **To run tests with coverage:**

    ```bash
    make coverage
    ```

- **To test using a client, start the server and run**

    ```bash
    go run client/main.go
    ```
