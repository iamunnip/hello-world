# Start from an image that already has Go installed.
FROM golang:1.27.3

# Use /app as the working folder inside the container.
WORKDIR /app

# Copy the program.
COPY go.mod main.go ./

# Turn the code into a program called "hello-world".
RUN go build -o hello-world .

# Start the program.
CMD ["./hello-world"]
