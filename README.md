# Hello World

A small web page that says "Hello, World", written in different programming languages.

- Each language has its own folder inside `src/`.
- Each program starts a small web server on your computer. A web server is a program that sends pages to your browser.
- You can run one program, or all of them at the same time.
- To stop a program, press `Ctrl+C` in the terminal.
- The easiest way to run them is with Docker. Docker runs a program in its own box, so you only need Docker installed, not the language.

## Table of Contents

- [Python](#python)
- [Go](#go)
- [Rust](#rust)

## Python

- Folder: `src/python/`
- It uses FastAPI, a popular tool for building websites in Python.
- The add-ons and their exact versions are listed in `requirements.txt`.

Run it with Docker:

```sh
cd src/python
docker build -t hello-python .
docker run --rm --init -p 8000:8000 hello-python
```

- The `docker build` line packs the program and everything it needs into an image.
- The `docker run` line starts the program and connects port 8000 to your computer.
- `--init` makes `Ctrl+C` work.

Check that it works:

- Open http://127.0.0.1:8000 in your browser, or run this in a second terminal:

```sh
curl http://127.0.0.1:8000
```

- You should see: `<h1>Hello, World from Python!</h1>`

Run it without Docker:

- You need: Python 3.15

```sh
cd src/python
python3 -m venv .venv
.venv/bin/pip install -r requirements.txt
.venv/bin/uvicorn main:app --port 8000
```

- The `venv` line makes a private space for this program's add-ons.
- The `pip` line installs the add-ons.
- The `uvicorn` line starts the program.
- Then open http://127.0.0.1:8000 in your browser.

## Go

- Folder: `src/go/`
- It uses only the tools that come with Go.
- There are no add-ons to install.

Run it with Docker:

```sh
cd src/go
docker build -t hello-go .
docker run --rm --init -p 8001:8001 hello-go
```

- The `docker build` line packs the program and everything it needs into an image.
- The `docker run` line starts the program and connects port 8001 to your computer.
- `--init` makes `Ctrl+C` work.

Check that it works:

- Open http://127.0.0.1:8001 in your browser, or run this in a second terminal:

```sh
curl http://127.0.0.1:8001
```

- You should see: `<h1>Hello, World from Go!</h1>`

Run it without Docker:

- You need: Go 1.27

```sh
cd src/go
go run .
```

- Then open http://127.0.0.1:8001 in your browser.

## Rust

- Folder: `src/rust/`
- It uses Axum, a popular tool for building websites in Rust.
- The add-ons are listed in `Cargo.toml`, and their exact versions are saved in `Cargo.lock`.

Run it with Docker:

```sh
cd src/rust
docker build -t hello-rust .
docker run --rm --init -p 8002:8002 hello-rust
```

- The `docker build` line packs the program and everything it needs into an image.
- The `docker run` line starts the program and connects port 8002 to your computer.
- `--init` makes `Ctrl+C` work.

Check that it works:

- Open http://127.0.0.1:8002 in your browser, or run this in a second terminal:

```sh
curl http://127.0.0.1:8002
```

- You should see: `<h1>Hello, World from Rust!</h1>`

Run it without Docker:

- You need: Rust 1.99

```sh
cd src/rust
cargo run
```

- The first run takes a little longer because Rust downloads and builds the add-ons.
- Then open http://127.0.0.1:8002 in your browser.
