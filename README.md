# APIRouter

This project provides a Golang service for handling `sendRawTransactionOptional` requests with dynamic request rate limiting, proxied through Nginx.

## Structure

- `src/`: Source code for the Golang service.
- `nginx/`: Nginx configuration files.

## Running the Project

### Without Docker

1. Start the Golang service: `go run src/main.go`
2. Start Nginx with the provided configuration.

### With Docker

1. Build the Docker image: `docker build -t my_project .`
2. Run the Docker container: `docker run -p 80:80 my_project`
