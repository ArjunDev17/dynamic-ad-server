# dynamic-ad-server
A dynamic ad server built with Go using the Gin framework. This server fetches ad templates and data from Redis, dynamically renders ad content based on incoming requests, and serves HTML/JavaScript or XML responses. It is designed to handle various ad formats efficiently and support different environments (development, staging, production).
# dynamic-ad-server Project with Gin Framework

## Overview
This project is a RESTful API developed using the Go programming language and the Gin web framework. The API includes a single endpoint (`/bid`) that accepts JSON payloads, validates input, checks authorization via a Bearer token, and interacts with a Redis database to fetch and replace data.

## Features
- **Authentication**: The API uses a Bearer token for simple authentication.
- **Validation**: JSON payloads are validated for required fields and data types.
- **Redis Integration**: The API interacts with a Redis database to fetch and replace content based on incoming requests.
- **Dynamic Response**: Returns dynamically generated JavaScript or XML responses based on input parameters.

## Project Structure
Here’s an overview of the project structure and what each directory/file is responsible for:


## Getting Started

### Prerequisites
- **Go**: Ensure Go is installed on your machine. You can download it from [golang.org](https://golang.org).
- **Redis**: Redis should be running locally or accessible via the network. You can download Redis from [redis.io](https://redis.io).

### Installation

1. **Clone the Repository:**
    ```bash
    git clone https://github.com/ArjunDev17/dynamic-ad-server
    cd go-api
    ```

2. **Install Dependencies:**
    ```bash
    go mod tidy
    ```

3. **Configure Redis:**
    Update the Redis connection settings in `config/config.go` if your Redis instance is not running on `localhost:6379`.

4. **Run the Application:**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on `http://localhost:8080`.

## Usage

### API Endpoint
- **Endpoint**: `/bid`
- **Method**: `POST`
- **Content-Type**: `application/json`
- **Authorization**: Bearer token in the header (`Authorization: Bearer <token>`)

### Sample Request
```bash
curl --location --request POST 'http://localhost:8080/bid' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer test1234' \
--data-raw '{
    "id": "id_123",
    "width": 600,
    "height": 328,
    "banner": {
        "type": 1
    }
}'
