# Hello World Programs

This repository contains two minimal web applications, one written in Python and one written in Go. Each application starts a local web server and displays a "Hello, World" message in the browser.

The instructions below are written for Linux.

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
    README.md
    src/
        go/
            go.mod
            main.go
        python/
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

## Code explanation

File: `src/go/main.go`

```go
package main

import (
	"fmt"
	"net/http"
)

const port = "8080"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("Go server running at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
```

| Code | What it does |
|---|---|
| `package main` | Marks this file as part of a program that can be run. |
| `import (` | Starts the list of packages the program uses. |
| `"fmt"` | Imports the package for printing and writing text. |
| `"net/http"` | Imports Go's built in package for web servers. |
| `)` | Ends the list of packages. |
| `const port = "8080"` | Stores the port number in a constant so it is defined in one place. |
| `func hello(w http.ResponseWriter, r *http.Request) {` | Defines a function named `hello` that handles visits to the page. `w` is used to send the reply and `r` holds details about the request. |
| `fmt.Fprint(w, "<h1>Hello, World from Go!</h1>")` | Sends the message to the browser. The `<h1>` tag displays it as a large heading. |
| `}` | Ends the `hello` function. |
| `func main() {` | Defines the `main` function. Go runs this function first when the program starts. |
| `http.HandleFunc("/", hello)` | Tells Go to run `hello` when someone visits the main page (`/`). |
| `fmt.Println("Go server running at http://localhost:" + port)` | Prints a message in the terminal showing the address of the server. |
| `http.ListenAndServe(":"+port, nil)` | Starts the web server on port 8080 and waits for requests. `nil` means it uses the default settings. |
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
