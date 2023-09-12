
# Go Gin REST API with GORM and PostgreSQL

This is a simple RESTful API project built with Go, Gin, GORM, and PostgreSQL for managing "todo" items. It provides endpoints to create, read, update, and delete todos.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

Follow these steps to get the project up and running on your local machine.

### Prerequisites

Before you begin, make sure you have the following prerequisites installed:

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Git](https://git-scm.com/)

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <project-folder>
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

### Configuration

1. Create a PostgreSQL database and note down the connection details.

2. Configure the database connection in the `InitDB` function within `main.go`. Update the `dsn` variable with your database credentials:

   ```go
   dsn := "user=postgres password=your-password dbname=your-dbname host=your-hostname port=5432 sslmode=disable"
   ```

## Usage

Start the Go application by running:

```bash
go run main.go
```

### API Endpoints

- **Create a new todo:**
  
  ```
  POST /todos
  ```
  
  Example request body:
  ```json
  {
      "Title": "Buy groceries",
      "Completed": false
  }
  ```

- **Get all todos:**

  ```
  GET /todos
  ```

- **Get a single todo by ID:**

  ```
  GET /todos/:id
  ```

- **Update a todo by ID:**

  ```
  PUT /todos/:id
  ```

  Example request body:
  ```json
  {
      "Title": "Buy milk",
      "Completed": true
  }
  ```

- **Delete a todo by ID:**

  ```
  DELETE /todos/:id
  ```

## Testing

You can test the API using tools like `curl` or Postman. Refer to the API Endpoints section for sample requests.

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow the standard GitHub workflow:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes.
4. Push your changes to your fork.
5. Open a pull request to the main repository.

## License

This project is licensed under the [MIT License](LICENSE).
```

You can copy and paste this Markdown content into your README.md file in your project repository, making sure to replace `<repository-url>` and `<project-folder>` with the actual Git repository URL and the folder where you cloned the project.
