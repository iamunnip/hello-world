# Hello World Programs

This repository contains two minimal web applications, one written in Python and one written in Go. Each application starts a local web server and displays a "Hello, World" message in the browser.

The instructions below are written for Linux.

## Table of contents

- [Overview](#overview)
- [How it works](#how-it-works)
- [Project structure](#project-structure)
- [Python](#python)
  - [Requirements](#requirements)
  - [Setup and run](#setup-and-run)
  - [Run with Docker](#run-with-docker)
  - [Code explanation](#code-explanation)
  - [Dependencies](#dependencies)
  - [Dockerfile](#dockerfile)
- [Go](#go)
  - [Requirements](#requirements-1)
  - [Setup and run](#setup-and-run-1)
  - [Run with Docker](#run-with-docker-1)
  - [Code explanation](#code-explanation-1)
  - [Module file](#module-file)
  - [Dockerfile](#dockerfile-1)

## Overview

| Language | Folder | Port | URL |
|---|---|---|---|
| Python | `src/python` | 8000 | http://localhost:8000 |
| Go | `src/go` | 8080 | http://localhost:8080 |

Each application uses its own port, so both can run at the same time in separate terminals.

## How it works

When you open the URL in a browser, the browser sends a request to the application. The application replies with a short HTML message, and the browser displays it.

A port is a number that identifies a program on your computer. Giving each application its own port keeps their traffic separate.

## Project structure

```
hello-world/
    .gitignore
    README.md
    src/
        go/
            .dockerignore
            Dockerfile
            go.mod
            main.go
        python/
            .dockerignore
            Dockerfile
            main.py
            requirements.txt
```

# Python

## Requirements

- Python 3.14 or newer
- Flask 3.1.3 (installed in the steps below)

Flask is a lightweight Python framework for building web applications.

To check your Python version:

```
python3.14 --version
```

## Setup and run

Run these commands from the root of the repository.

1. Create a virtual environment. This keeps the project's packages separate from the rest of your system.

   ```
   python3.14 -m venv venv
   ```

2. Activate the virtual environment.

   ```
   source venv/bin/activate
   ```

3. Install the dependencies.

   ```
   pip install -r src/python/requirements.txt
   ```

4. Start the application.

   ```
   python src/python/main.py
   ```

5. Open http://localhost:8000 in your browser. You should see "Hello, World from Python!"

To stop the application, press `Ctrl+C` in the terminal. To leave the virtual environment, run `deactivate`.

## Run with Docker

Instead of installing Python and Flask, you can run the application in a Docker container. This needs Docker installed and running.

Run these commands from the root of the repository.

1. Build the image. `-t` gives the image a name, and `src/python` is the folder that contains the `Dockerfile`.

   ```
   docker build -t hello-python src/python
   ```

2. Start a container. `-p 8000:8000` connects port 8000 on your computer to port 8000 in the container, and `--rm` removes the container when it stops.

   ```
   docker run --rm -p 8000:8000 hello-python
   ```

3. Open http://localhost:8000 in your browser.

To stop the container, press `Ctrl+C` in the terminal.

## Code explanation

File: `src/python/main.py`

```python
from flask import Flask

app = Flask(__name__)


@app.route("/")
def hello():
    return "<h1>Hello, World from Python!</h1>"


if __name__ == "__main__":
    app.run(port=8000)
```

| Code | What it does |
|---|---|
| `from flask import Flask` | Imports the `Flask` class from the Flask library. |
| `app = Flask(__name__)` | Creates the web application. `__name__` tells Flask which file the application lives in. |
| `@app.route("/")` | Tells Flask to run the function below it when someone visits the main page (`/`). |
| `def hello():` | Defines a function named `hello` that handles visits to the main page. |
| `return "<h1>Hello, World from Python!</h1>"` | Sends the message to the browser. The `<h1>` tag displays it as a large heading. |
| `if __name__ == "__main__":` | Makes sure the next line only runs when this file is started directly. |
| `app.run(port=8000)` | Starts the web server on port 8000 and waits for requests. |

## Dependencies

File: `src/python/requirements.txt`

| Code | What it does |
|---|---|
| `Flask==3.1.3` | Tells pip to install Flask version 3.1.3. Pip also installs the packages Flask depends on. |

## Dockerfile

File: `src/python/Dockerfile`

```dockerfile
FROM python:3.14-slim
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY main.py ./
EXPOSE 8000
CMD ["flask", "--app", "main", "run", "--host", "0.0.0.0", "--port", "8000"]
```

| Code | What it does |
|---|---|
| `FROM python:3.14-slim` | Starts from an official image that already has Python 3.14 installed. `slim` is a smaller version of the image. |
| `WORKDIR /app` | Sets `/app` as the folder inside the image where the next commands run. |
| `COPY requirements.txt ./` | Copies the dependency list into the image. |
| `RUN pip install --no-cache-dir -r requirements.txt` | Installs Flask. `--no-cache-dir` keeps the image smaller. Docker reuses this step on later builds as long as `requirements.txt` has not changed. |
| `COPY main.py ./` | Copies the application code into the image. |
| `EXPOSE 8000` | Notes that the application listens on port 8000. |
| `CMD [...]` | Starts the application with the `flask` command. `--host 0.0.0.0` makes the server accept connections from outside the container. Without it, the server only listens inside the container and the browser cannot reach it. |

The `src/python/.dockerignore` file lists files that Docker should not send into the build, such as `__pycache__/` and virtual environments.

# Go

## Requirements

- Go 1.27 or newer

The Go application uses only the Go standard library, so no extra packages are needed.

To check your Go version:

```
go version
```

## Setup and run

Run these commands from the root of the repository.

1. Move into the Go folder.

   ```
   cd src/go
   ```

2. Start the application. The `.` means "run the module in the current folder".

   ```
   go run .
   ```

3. Open http://localhost:8080 in your browser. You should see "Hello, World from Go!"

To stop the application, press `Ctrl+C` in the terminal.

## Run with Docker

Instead of installing Go, you can run the application in a Docker container. This needs Docker installed and running.

Run these commands from the root of the repository.

1. Build the image. `-t` gives the image a name, and `src/go` is the folder that contains the `Dockerfile`.

   ```
   docker build -t hello-go src/go
   ```

2. Start a container. `-p 8080:8080` connects port 8080 on your computer to port 8080 in the container, and `--rm` removes the container when it stops.

   ```
   docker run --rm -p 8080:8080 hello-go
   ```

3. Open http://localhost:8080 in your browser.

To stop the container, press `Ctrl+C` in the terminal.

## Code explanation

File: `src/go/main.go`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Go server running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

| Code | What it does |
|---|---|
| `package main` | Marks this file as part of a program that can be run. |
| `import (` | Starts the list of packages the program uses. |
| `"fmt"` | Imports the package for printing and writing text. |
| `"log"` | Imports the package for logging errors. |
| `"net/http"` | Imports Go's built in package for web servers. |
| `)` | Ends the list of packages. |
| `const port = "8080"` | Stores the port number in a constant so it is defined in one place. |
| `func hello(w http.ResponseWriter, r *http.Request) {` | Defines a function named `hello` that handles visits to the page. `w` is used to send the reply and `r` holds details about the request. |
| `fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")` | Sends the message to the browser. The `<h1>` tag displays it as a large heading. |
| `}` | Ends the `hello` function. |
| `func main() {` | Defines the `main` function. Go runs this function first when the program starts. |
| `http.HandleFunc("/", hello)` | Tells Go to run `hello` when someone visits the main page (`/`). |
| `fmt.Println("Go server running at http://localhost:" + port)` | Prints a message in the terminal showing the address of the server. |
| `log.Fatal(http.ListenAndServe(":"+port, nil))` | Starts the web server on port 8080 and waits for requests. `nil` means it uses the default settings. If the server cannot start, for example because the port is already in use, `log.Fatal` prints the error and stops the program. |
| `}` | Ends the `main` function. |

## Module file

File: `src/go/go.mod`

```
module hello-world

go 1.27.1
```

| Code | What it does |
|---|---|
| `module hello-world` | Sets the name of the Go module. A module is a Go project. |
| `go 1.27.1` | Sets the Go version the project uses. |

There is no `go.sum` file because the application only uses Go's standard library. Go creates `go.sum` only when a project uses outside packages.

## Dockerfile

File: `src/go/Dockerfile`

```dockerfile
# Build stage: compile the program
FROM golang:1.27 AS build
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o hello-world .

# Run stage: copy only the compiled program into a minimal image
FROM gcr.io/distroless/static-debian12
COPY --from=build /app/hello-world /hello-world
EXPOSE 8080
CMD ["/hello-world"]
```

This Dockerfile has two stages. The first stage compiles the program, and the second stage keeps only the compiled program. The final image does not include the Go compiler, so it is much smaller.

| Code | What it does |
|---|---|
| `FROM golang:1.27 AS build` | Starts the first stage from an official image that has Go 1.27 installed, and names the stage `build`. |
| `WORKDIR /app` | Sets `/app` as the folder inside the image where the next commands run. |
| `COPY go.mod ./` | Copies the module file into the image. |
| `COPY main.go ./` | Copies the application code into the image. |
| `RUN CGO_ENABLED=0 go build -o hello-world .` | Compiles the program into a file named `hello-world`. `CGO_ENABLED=0` makes the program self-contained, so it runs without any other system libraries. |
| `FROM gcr.io/distroless/static-debian12` | Starts the second stage from a very small image that contains almost nothing except what a self-contained program needs. |
| `COPY --from=build /app/hello-world /hello-world` | Copies the compiled program from the `build` stage. |
| `EXPOSE 8080` | Notes that the application listens on port 8080. |
| `CMD ["/hello-world"]` | Runs the program when the container starts. |

The `src/go/.dockerignore` file tells Docker not to send the `hello-world` binary into the build. That file only exists if you ran `go build` on your computer, and the image compiles its own copy.
