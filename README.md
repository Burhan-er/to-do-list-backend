# To-Do-List CLI

A simple and lightweight command-line interface (CLI) application for managing your to-do tasks. Built with Go, this tool allows you to add, list, complete, and remove tasks directly from your terminal. The application also includes a `Dockerfile` and `docker-compose.yml` for easy containerization.

---

### Features

* **Add Tasks**: Quickly add new tasks with a single command.
* **List Tasks**: View all your tasks, including their status (completed or not).
* **Complete Tasks**: Mark a task as done by its index number.
* **Remove Tasks**: Delete a task from your list.
* **Persistent Storage**: Tasks are saved to a JSON file in your home directory, ensuring data persistence.
* **Docker Support**: Easily build and run the application using Docker and Docker Compose.

---

### Prerequisites

To run this application, you need to have the following tools installed on your system:

* **[Go](https://go.dev/doc/install)** (version 1.24.5 or higher)
* **[Docker](https://docs.docker.com/get-docker/)**
* **[Docker Compose](https://docs.docker.com/compose/install/)**

---

### Getting Started

#### 1. Clone the Repository

```sh
git clone [https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git](https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git)
cd YOUR_REPO_NAME
2. Run with Go (Without Docker)
If you prefer to run the application directly on your machine, first initialize the Go module and install dependencies.

Bash

go mod tidy
Then, run the application using the go run command.

Bash

go run .
3. Run with Docker Compose
For a containerized environment, use Docker Compose. This will build the Docker image and start the container in one step.

Bash

docker-compose up --build
Usage
The application's executable is your project folder itself. You can run commands by typing go run . or, if using Docker, by running commands within the container.

Using Docker Compose:
To run commands with Docker Compose, you can use docker-compose run.

Add a Task:

Bash

docker-compose run --rm go-app add "Buy groceries from the market"
List All Tasks:

Bash

docker-compose run --rm go-app list
Complete a Task:

Bash

docker-compose run --rm go-app done 1
Remove a Task:

Bash

docker-compose run --rm go-app remove 1
Using the Go Executable:
If you prefer to run it without Docker, use go run.

Add a Task:

Bash

go run . add "Go to the gym"
List All Tasks:

Bash

go run . list
How It Works
The application stores your to-do tasks in a simple JSON file named .todo.json within your home directory (~). The Dockerfile creates a lightweight, production-ready image by using a multi-stage build process. This ensures that the final image contains only the compiled application and a minimal base OS, keeping the size small and efficient.