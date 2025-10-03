# Go Data API Microservice

[Português](README.pt-br.md) | [English](README.md)


![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub%20Actions-2671E5?style=for-the-badge&logo=githubactions&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

<p align="center">
  <img src="./docs/hero_image.png" alt="Go Data API Microservice Hero Image">
</p>

## Overview

This repository contains a simple Go microservice that exposes a data API. It demonstrates a professional project structure, functional code with unit tests, and comprehensive bilingual documentation.

## Features

- **Professional Structure**: Organized with `src/`, `tests/`, `docs/`, and `config/` folders.
- **Functional Code**: A basic Go HTTP server with a data endpoint.
- **Unit Tests**: Comprehensive unit tests for API handlers.
- **Bilingual Documentation**: `README.md` in English and Portuguese.
- **Visual Elements**: Architecture diagrams and badges.

## Architecture

The microservice architecture is designed for scalability and maintainability.

```mermaid
graph TD
    A[Client] --> B(Load Balancer)
    B --> C{API Gateway}
    C --> D[Go Data API Microservice]
    D --> E[Database]
```

## Getting Started

### Prerequisites

- Go (version 1.18 or higher)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/galafis/go-data-api-microservices-new.git
   cd go-data-api-microservices
   ```
2. Compile the application:
   ```bash
   go build -o bin/server src/main.go
   ```

### Running the Application

```bash
./bin/server
```

The server will start on `http://localhost:8080`.

### API Endpoints

- `GET /`: Returns a welcome message.
- `GET /data`: Returns a JSON array of sample data.

Example usage with `curl`:

```bash
curl http://localhost:8080/
# Expected output: Welcome to the Go Microservices API!

curl http://localhost:8080/data
# Expected output: [
#   {
#     "id": "1",
#     "name": "Item 1"
#   },
#   {
#     "id": "2",
#     "name": "Item 2"
#   }
# ]
```

## Running Tests

To run the unit tests, execute the following command:

```bash
go test ./tests
```

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Gabriel Demetrios Lafis**

