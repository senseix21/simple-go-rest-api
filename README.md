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

- Go (https://golang.org/)
- PostgreSQL (https://www.postgresql.org/)
- Git (https://git-scm.com/)

### Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <project-folder>
Install Go dependencies:

bash
Copy code
go mod tidy
Configuration
Create a PostgreSQL database and note down the connection details.

Configure the database connection in the InitDB function within main.go. Update the dsn variable with your database credentials:

go
Copy code
dsn := "user=postgres password=your-password dbname=your-dbname host=your-hostname port=5432 sslmode=disable"
Usage
Start the Go application by running:

bash
Copy code
go run main.go
API Endpoints
Create a new todo:

bash
Copy code
POST /todos
Example request body:

json
Copy code
{
    "Title": "Buy groceries",
    "Completed": false
}
Get all todos:

bash
Copy code
GET /todos
Get a single todo by ID:

bash
Copy code
GET /todos/:id
Update a todo by ID:

bash
Copy code
PUT /todos/:id
Example request body:

json
Copy code
{
    "Title": "Buy milk",
    "Completed": true
}
Delete a todo by ID:

bash
Copy code
DELETE /todos/:id
Testing
You can test the API using tools like curl or Postman. Refer to the API Endpoints section for sample requests.

Contributing
Contributions are welcome! If you'd like to contribute to this project, please follow the standard GitHub workflow:

Fork the repository.
Create a new branch for your feature or bug fix.
Commit your changes.
Push your changes to your fork.
Open a pull request to the main repository.
License
This project is licensed under the MIT License.


Replace `<repository-url>` and `<project-folder>` with the actual Git repository URL and the folder where you cloned the project.

This README.md provides an overview of your project, how to set it up, how to use the API, and how to contribute to the project. Make sure to customize it further with specific details, and you can also add sections like "Deployment" or "Troubleshooting" if needed.
