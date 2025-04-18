# URL Shortener

## Description
A URL shortener service built with Go. This application allows users to shorten long URLs and retrieve the original URLs using the shortened versions. It uses PostgreSQL for persistent storage and Redis for caching.

## Features
- Shorten long URLs into 8-character short URLs.
- Retrieve the original URL using the short URL.
- Rate limiting to prevent abuse.
- Caching with Redis for faster lookups.
- Graceful shutdown of the server.

## Prerequisites
- Docker and Docker Compose installed.
- Go 1.24 or higher installed (if running locally without Docker).

## Installation

### Using Docker
1. Clone the repository:
   ```bash
   git clone https://github.com/mojtabamovahedi/url-shorter.git
   cd url-shorter
   ```
2. Build and start the services:
   ```bash
   docker-compose up --build
   ```
3. The application will be available at `http://localhost:8080`.

### Running Locally
1. Clone the repository:
   ```bash
   git clone https://github.com/mojtabamovahedi/url-shorter.git
   cd url-shorter
   ```
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Start PostgreSQL and Redis services (e.g., using Docker Compose):
   ```bash
   docker-compose up postgres redis
   ```
4. Run the application:
   ```bash
   go run ./cmd/main.go
   ```

## Usage

### Shorten a URL
Send a POST request to `/new` with the following JSON body:
```json
{
  "url": "https://example.com"
}
```
Response:
```json
{
  "message": "success",
  "short": "abcdefgh"
}
```

### Redirect to the Original URL
Send a GET request to `/:short` (e.g., `/abcdefgh`). The server will redirect to the original URL.

## Configuration
The application is configured using the `config.yaml` file. Update the file to change database, Redis, or server settings.

## Project Structure
- `cmd/main.go`: Entry point of the application.
- `internal/`: Contains core business logic, services, and repositories.
- `api/handler/http/`: HTTP handlers and middleware.
- `config/`: Configuration-related files.
- `pkg/`: Utility packages (e.g., database and cache).


## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
