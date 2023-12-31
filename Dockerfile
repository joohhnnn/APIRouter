# Use the Golang base image
FROM golang:1.16 as builder

# Set the working directory
WORKDIR /app

# Copy the source code of the Golang service
COPY src/*.go .

# Compile the Golang service
RUN go build -o my_service

# Use the Nginx base image
FROM nginx:alpine

# Copy the compiled Golang service
COPY --from=builder /app/my_service /usr/local/bin/

# Copy Nginx configuration
COPY nginx/nginx.conf /etc/nginx/nginx.conf
COPY nginx/conf.d/ /etc/nginx/conf.d/

# Start Nginx and the Golang service
CMD ["sh", "-c", "my_service & nginx -g 'daemon off;'"]
