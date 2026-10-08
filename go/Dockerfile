# Build stage: compile the program
FROM golang:1.27 AS build
WORKDIR /app
COPY go.mod ./
COPY src/ ./src/
RUN CGO_ENABLED=0 go build -o hello-world ./src

# Run stage: copy only the compiled program into a minimal image
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /app/hello-world /hello-world
EXPOSE 8080
CMD ["/hello-world"]
