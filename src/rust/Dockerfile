# Start from a small image that already has Rust installed.
FROM rust:1.99.1-slim

# Use /app as the working folder inside the container.
WORKDIR /app

# Copy the program and the list of add-ons.
COPY Cargo.toml Cargo.lock ./
COPY src ./src

# Download the add-ons and turn the code into a program.
RUN cargo build --release

# Start the program.
CMD ["./target/release/hello-world"]
