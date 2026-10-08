# Hello World Programs

This repository contains two minimal web applications, one written in Python and one written in Go. Each application starts a local web server and displays a "Hello, World" message in the browser.

The instructions below are written for Linux.

## Table of contents

- [Overview](#overview)
- [How it works](#how-it-works)
- [Project structure](#project-structure)
- [Continuous integration](#continuous-integration)
  - [Lint checks](#lint-checks)
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
- [Interview preparation](#interview-preparation)
  - [Project structure questions](#project-structure-questions)
  - [Python questions](#python-questions)
  - [Go questions](#go-questions)
  - [Docker questions](#docker-questions)
  - [CI questions](#ci-questions)

## Overview

| Language | Folder | Port | URL |
|---|---|---|---|
| Python | `python` | 8000 | http://localhost:8000 |
| Go | `go` | 8080 | http://localhost:8080 |

Each application uses its own port, so both can run at the same time in separate terminals.

## How it works

When you open the URL in a browser, the browser sends a request to the application. The application replies with a short HTML message, and the browser displays it.

A port is a number that identifies a program on your computer. Giving each application its own port keeps their traffic separate.

## Project structure

```
hello-world
├── .github
│   ├── actions
│   │   ├── go
│   │   │   ├── build
│   │   │   │   └── action.yml
│   │   │   └── lint
│   │   │       └── action.yml
│   │   └── python
│   │       ├── build
│   │       │   └── action.yml
│   │       └── lint
│   │           └── action.yml
│   └── workflows
│       ├── go.yml
│       └── python.yml
├── .gitignore
├── go
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── go.mod
│   └── src
│       └── main.go
├── python
│   ├── Dockerfile
│   ├── .dockerignore
│   ├── requirements.txt
│   └── src
│       └── main.py
└── README.md
```

## Continuous integration

The repository uses GitHub Actions. Each application has its own workflow:

| Workflow | File | Runs when these files change |
|---|---|---|
| Python | `.github/workflows/python.yml` | `python/`, `python.yml`, `.github/actions/python/` |
| Go | `.github/workflows/go.yml` | `go/`, `go.yml`, `.github/actions/go/` |

A change to only the Python files runs only the Python workflow, and the same for Go.

The workflow files only decide when to run and in what order. The steps themselves are in composite actions, which are reusable groups of steps in `.github/actions/`:

| Action | Used by | What it does |
|---|---|---|
| `.github/actions/python/lint` | Python workflow | Runs the Python lint checks. |
| `.github/actions/go/lint` | Go workflow | Runs the Go lint checks. |
| `.github/actions/python/build` | Python workflow | Builds the Python Docker image, and pushes it to Docker Hub when that is turned on. |
| `.github/actions/go/build` | Go workflow | Builds the Go Docker image, and pushes it to Docker Hub when that is turned on. |

Each workflow runs:

- **On a pull request into `main`:** runs the lint checks below, then builds the Docker image to check that it still works. If a lint check fails, the build does not run.
- **When a pull request is merged into `main`:** builds the Docker image. Pushing the image to Docker Hub is not enabled yet. The push steps are already written in each language's Docker build action, and each workflow has commented-out lines that turn them on.
- **By hand:** open the **Actions** tab on GitHub, choose the workflow, and click **Run workflow**.

The lint checks do not run on merge, because the code already passed them in the pull request.

The image is tagged `hello-world:<language>.<run number>`, for example `hello-world:python.12`. The run number goes up by one every time the workflow runs, so it replaces the build number you type yourself when building locally. Each workflow has its own run number.

### Lint checks

Lint tools read the code and point out mistakes and formatting problems without running it.

| Workflow | Tool | What it checks | Run it yourself |
|---|---|---|---|
| Python | ruff | Code mistakes, such as unused imports | `ruff check python` |
| Python | ruff | Code formatting | `ruff format --check python` |
| Go | gofmt | Code formatting | `gofmt -l go/src` |
| Go | go vet | Common code mistakes | `cd go && go vet ./src` |
| Both | hadolint | Dockerfile mistakes and best practices | `hadolint python/Dockerfile go/Dockerfile` |

To fix formatting automatically, run `ruff format python` for Python or `gofmt -w go/src` for Go.

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
   pip install -r python/requirements.txt
   ```

4. Start the application.

   ```
   python python/src/main.py
   ```

5. Open http://localhost:8000 in your browser. You should see "Hello, World from Python!"

To stop the application, press `Ctrl+C` in the terminal. To leave the virtual environment, run `deactivate`.

## Run with Docker

Instead of installing Python and Flask, you can run the application in a Docker container. This needs Docker installed and running.

Run these commands from the root of the repository.

1. Build the image. `-t` names the image `hello-world` and gives it the tag `python.1`, and `python` is the folder that contains the `Dockerfile`.

   The number in the tag counts the builds. Increase it by one each time you build a new version: `python.2`, `python.3`, and so on. This keeps older images available, so you can still run or compare them.

   ```
   docker image build -t hello-world:python.1 python
   ```

2. Start a container. `-p 8000:8000` connects port 8000 on your computer to port 8000 in the container, and `--rm` removes the container when it stops. Use the tag of the image you want to run.

   ```
   docker container run --rm -p 8000:8000 hello-world:python.1
   ```

3. Open http://localhost:8000 in your browser.

To stop the container, press `Ctrl+C` in the terminal.

## Code explanation

File: `python/src/main.py`

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

File: `python/requirements.txt`

| Code | What it does |
|---|---|
| `Flask==3.1.3` | Tells pip to install Flask version 3.1.3. Pip also installs the packages Flask depends on. |

## Dockerfile

File: `python/Dockerfile`

```dockerfile
FROM python:3.14-slim
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY src/main.py ./
EXPOSE 8000
CMD ["flask", "--app", "main", "run", "--host", "0.0.0.0", "--port", "8000"]
```

| Code | What it does |
|---|---|
| `FROM python:3.14-slim` | Starts from an official image that already has Python 3.14 installed. `slim` is a smaller version of the image. |
| `WORKDIR /app` | Sets `/app` as the folder inside the image where the next commands run. |
| `COPY requirements.txt ./` | Copies the dependency list into the image. |
| `RUN pip install --no-cache-dir -r requirements.txt` | Installs Flask. `--no-cache-dir` keeps the image smaller. Docker reuses this step on later builds as long as `requirements.txt` has not changed. |
| `COPY src/main.py ./` | Copies the application code into the image. |
| `EXPOSE 8000` | Notes that the application listens on port 8000. |
| `CMD [...]` | Starts the application with the `flask` command. `--host 0.0.0.0` makes the server accept connections from outside the container. Without it, the server only listens inside the container and the browser cannot reach it. |

The `python/.dockerignore` file lists files that Docker should not send into the build, such as `__pycache__/` and virtual environments.

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
   cd go
   ```

2. Start the application. `./src` tells Go where the code is.

   ```
   go run ./src
   ```

3. Open http://localhost:8080 in your browser. You should see "Hello, World from Go!"

To stop the application, press `Ctrl+C` in the terminal.

## Run with Docker

Instead of installing Go, you can run the application in a Docker container. This needs Docker installed and running.

Run these commands from the root of the repository.

1. Build the image. `-t` names the image `hello-world` and gives it the tag `go.1`, and `go` is the folder that contains the `Dockerfile`.

   The number in the tag counts the builds. Increase it by one each time you build a new version: `go.2`, `go.3`, and so on. This keeps older images available, so you can still run or compare them.

   ```
   docker image build -t hello-world:go.1 go
   ```

2. Start a container. `-p 8080:8080` connects port 8080 on your computer to port 8080 in the container, and `--rm` removes the container when it stops. Use the tag of the image you want to run.

   ```
   docker container run --rm -p 8080:8080 hello-world:go.1
   ```

3. Open http://localhost:8080 in your browser.

To stop the container, press `Ctrl+C` in the terminal.

## Code explanation

File: `go/src/main.go`

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

File: `go/go.mod`

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

File: `go/Dockerfile`

```dockerfile
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
```

This Dockerfile has two stages. The first stage compiles the program, and the second stage keeps only the compiled program. The final image does not include the Go compiler, so it is much smaller.

| Code | What it does |
|---|---|
| `FROM golang:1.27 AS build` | Starts the first stage from an official image that has Go 1.27 installed, and names the stage `build`. |
| `WORKDIR /app` | Sets `/app` as the folder inside the image where the next commands run. |
| `COPY go.mod ./` | Copies the module file into the image. |
| `COPY src/ ./src/` | Copies the folder with the application code into the image. |
| `RUN CGO_ENABLED=0 go build -o hello-world ./src` | Compiles the code in `src` into a file named `hello-world`. `CGO_ENABLED=0` makes the program self-contained, so it runs without any other system libraries. |
| `FROM gcr.io/distroless/static-debian12:nonroot` | Starts the second stage from a very small image that contains almost nothing except what a self-contained program needs. The `nonroot` tag runs the program as a regular user instead of the administrator (root), which is safer. |
| `COPY --from=build /app/hello-world /hello-world` | Copies the compiled program from the `build` stage. |
| `EXPOSE 8080` | Notes that the application listens on port 8080. |
| `CMD ["/hello-world"]` | Runs the program when the container starts. |

The `go/.dockerignore` file tells Docker not to send the `hello-world` binary into the build. That file only exists if you ran `go build` on your computer, and the image compiles its own copy.

# Interview preparation

These questions are based on the code and setup in this repository. Click a question to show the answer.

## Project structure questions

<details>
<summary>Why does each language have its own folder with its own Dockerfile?</summary>

Each folder is a complete, independent project. It can be built, run and tested on its own, and a change to one language does not affect the other. It also lets CI run only the workflow for the language that changed.

</details>

<details>
<summary>Why are <code>go.mod</code> and <code>requirements.txt</code> outside the <code>src</code> folder?</summary>

They describe the whole project, not a single code file. Go treats the folder that contains `go.mod` as the root of the module, and pip and the Dockerfile expect `requirements.txt` at the project root. Only code files go in `src`.

</details>

<details>
<summary>What is a monorepo, and what are its pros and cons?</summary>

A monorepo keeps several projects in one repository, like the Python and Go apps here.

- Pros: one place for code, CI and documentation, and shared tooling and conventions apply to every project.
- Cons: CI needs path filters so every change does not run every workflow, and the repository grows as projects are added.

</details>

## Python questions

<details>
<summary>Why does the Dockerfile start Flask with <code>--host 0.0.0.0</code> instead of running <code>python main.py</code>?</summary>

`app.run(port=8000)` listens on `127.0.0.1` by default. Inside a container, `127.0.0.1` means the container itself, so connections from your computer through `-p 8000:8000` never reach the app. `0.0.0.0` makes the app listen on all network interfaces of the container, so the port mapping works.

</details>

<details>
<summary>Is the Flask built-in server suitable for production?</summary>

No. It is a development server, built for convenience rather than performance or security. In production, Flask apps run behind a WSGI server such as Gunicorn, often with a reverse proxy like Nginx in front.

</details>

<details>
<summary>Why is Flask pinned to an exact version (<code>Flask==3.1.3</code>)?</summary>

So every install gets the same version. Without a pin, a new Flask release could change behavior or break the app, and builds from different days could produce different results.

</details>

<details>
<summary>What does <code>if __name__ == "__main__":</code> do?</summary>

The code under it runs only when the file is started directly, such as `python main.py`. It does not run when the file is imported, for example when `flask --app main run` loads the app in the Dockerfile.

</details>

<details>
<summary>Why use a virtual environment?</summary>

It keeps the project's packages separate from the system Python and from other projects, so different projects can use different versions of the same package without conflicts.

</details>

## Go questions

<details>
<summary>Why is <code>http.ListenAndServe</code> wrapped in <code>log.Fatal</code>?</summary>

`ListenAndServe` returns an error if the server cannot start, for example when the port is already in use. Without `log.Fatal`, the error is ignored and the program exits without saying why. `log.Fatal` prints the error and exits with a non-zero status.

</details>

<details>
<summary>Which URLs does <code>http.HandleFunc("/", hello)</code> respond to?</summary>

All of them. The pattern `/` matches every path that no other pattern matches, so `/`, `/about` and `/anything` all return the hello message.

</details>

<details>
<summary>Why is there no <code>go.sum</code> file?</summary>

`go.sum` stores checksums of outside modules. This app uses only the Go standard library, so there is nothing to record.

</details>

<details>
<summary>Why does the build use <code>go build -o hello-world ./src</code>?</summary>

The code is in `src`, not next to `go.mod`, so the build has to point at `./src`. Without `-o`, Go names the binary after the folder, so it would be called `src`.

</details>

<details>
<summary>What is the difference between <code>gofmt</code> and <code>go vet</code>?</summary>

`gofmt` checks only formatting, such as indentation and spacing. `go vet` looks for likely bugs, such as wrong arguments to `fmt.Printf` or unreachable code.

</details>

## Docker questions

<details>
<summary>What is a multi-stage build, and why does the Go Dockerfile use one?</summary>

A multi-stage build uses several `FROM` lines in one Dockerfile. The Go Dockerfile compiles the program in a stage with the full Go toolchain, then copies only the compiled program into a very small final image. The final image is much smaller and contains no compiler or shell, so it has less that can be attacked.

</details>

<details>
<summary>Why does <code>CGO_ENABLED=0</code> matter with a distroless static image?</summary>

`CGO_ENABLED=0` makes Go produce a fully static binary that does not need the C library at runtime. The `distroless/static` image does not include a C library, so a binary that needed one would fail to start.

</details>

<details>
<summary>Why does the Python Dockerfile copy <code>requirements.txt</code> before <code>main.py</code>?</summary>

Docker caches each step. If only `main.py` changes, Docker reuses the cached `pip install` step and only redoes `COPY main.py` and the steps after it. If `main.py` were copied first, every code change would reinstall Flask.

</details>

<details>
<summary>What does <code>.dockerignore</code> do?</summary>

It lists files that Docker does not send to the build. This keeps builds fast and stops local files, like `__pycache__` or a locally built binary, from ending up in the image.

</details>

<details>
<summary>Does <code>EXPOSE 8080</code> make the app reachable from your computer?</summary>

No. `EXPOSE` only documents which port the app uses. The port is opened with `-p` when the container starts, for example `docker container run -p 8080:8080 hello-world:go.1`.

</details>

<details>
<summary>Why does the Go image use the <code>nonroot</code> tag?</summary>

It runs the program as a regular user instead of root. If someone breaks into the app, they get fewer permissions inside the container. The app uses port 8080, which does not need root.

</details>

<details>
<summary>What is the difference between <code>docker build</code> and <code>docker image build</code>?</summary>

They do the same thing. `docker image build` is the management command form, which groups commands by what they act on (`docker image ...`, `docker container ...`). It makes commands easier to read and find.

</details>

<details>
<summary>In <code>hello-world:python.1</code>, which part is the name and which is the tag? Why not use <code>latest</code>?</summary>

`hello-world` is the image name and `python.1` is the tag. A numbered tag points to one specific build, so you can run or roll back to an exact version. `latest` changes every time a new image is pushed, so it does not say which build you are running.

</details>

## CI questions

<details>
<summary>What do the <code>paths</code> filters in the workflows do?</summary>

They make a workflow run only when files in certain folders change. The Python workflow watches:

- `python/**`: the Python code, Dockerfile and dependencies
- `.github/workflows/python.yml`: the workflow itself
- `.github/actions/python/**`: the Python lint and build actions

The Go workflow watches the same three paths for Go. The two lists do not overlap, so a change to Python files never runs the Go workflow, and the other way around. A pull request that changes both runs both. Manual runs (`workflow_dispatch`) ignore the filters.

</details>

<details>
<summary>Why is lint skipped on merge into <code>main</code>?</summary>

The code already passed lint in the pull request, so running it again on merge would repeat the same check.

</details>

<details>
<summary>Why does the build job need <code>if: !cancelled() && (needs.lint.result == 'success' || needs.lint.result == 'skipped')</code>?</summary>

By default, a job is skipped when a job it depends on (`needs`) is skipped. On merge, lint is skipped, so without this condition the build would be skipped too. The condition lets the build run when lint passed or was skipped, but not when it failed or the run was cancelled.

</details>

<details>
<summary>What is the difference between a composite action and a reusable workflow?</summary>

- A composite action is a group of steps. It runs inside a job of the calling workflow and can be stored in any folder, like `.github/actions/python/lint`.
- A reusable workflow contains whole jobs. It is called with `uses:` at the job level and must be stored directly in `.github/workflows`.

</details>

<details>
<summary>Why does each job check out the code before using a local action?</summary>

A local action like `./.github/actions/python/build` is a file in the repository. Each job starts on a fresh machine, so the file does not exist until `actions/checkout` downloads the repository.

</details>

<details>
<summary>How are Docker Hub credentials given to the composite action?</summary>

Composite actions cannot read secrets directly. The workflow reads the secrets (`secrets.DOCKERHUB_USERNAME` and `secrets.DOCKERHUB_TOKEN`) and passes them to the action as inputs.

</details>

<details>
<summary>Why do pull requests build the image but never push it?</summary>

Pull requests contain code that has not been reviewed or merged yet. Pushing would publish unapproved images. Also, pull requests from forks cannot read the repository's secrets, so a push would fail anyway.

</details>

<details>
<summary>What is <code>github.run_number</code>?</summary>

A number that goes up by one every time a workflow runs. Each workflow has its own counter, so the Python and Go images get separate build numbers.

</details>

<details>
<summary>What is the problem with making path-filtered workflows required checks?</summary>

If a workflow is required before merging but its path filter means it never starts, the pull request waits for that check forever. For example, a Go-only change would never get the Python check.

</details>

<details>
<summary>What did hadolint find in this project, and how was it fixed?</summary>

It reported DL3006: the Go Dockerfile used `gcr.io/distroless/static-debian12` without a tag. Without a tag, Docker uses `latest`, which can change at any time. It was fixed by pinning the image to the `nonroot` tag.

</details>
