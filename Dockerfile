# This Dockerfile is used to build a Docker image for a Go application.
# It sets the working directory, copies the application code into the container,
# exposes port 8080, and runs the main.go file using the "go run" command.

FROM golang:lastest

WORKDIR /app

COPY . . 

EXPOSE 1323

CMD ["go", "run", "cmd/delivery/main.go"]