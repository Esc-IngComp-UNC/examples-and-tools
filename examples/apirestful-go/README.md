# RESTful API in Go

## Description

This project is an example of a RESTful API developed in Go, managing basic user operations. It uses **Chi** as an HTTP router and **MongoDB** as the database. The API allows creating, reading, updating, and deleting users, offering a modular and scalable structure.

## Technologies Used

- [Go](https://golang.org/) - Primary programming language.
- [Chi](https://github.com/go-chi/chi) - HTTP router.
- [MongoDB](https://www.mongodb.com/) - NoSQL database.
- [Git](https://git-scm.com/) - Version control.
- [Godotenv](https://github.com/joho/godotenv) - Environment variable management.
- [Validator](https://github.com/go-playground/validator) - Data validation.
- [Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo) - MongoDB driver.

## Project Structure

```plaintext
apirestful-go/
├── cmd/
│   └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── dtos/
│   │   └── user_dto.go
│   ├── handlers/
│   │   └── user_handler.go
│   ├── mappers/
│   │   └── user_mapper.go
│   ├── models/
│   │   └── user.go
│   ├── repository/
│   │   └── user_repository.go
│   └── services/
│       └── user_service.go
├── pkg/
│   ├── errors/
│   │   └── errors.go
│   └── helpers/
│       └── message.go
├── .gitignore
├── LICENSE
└── README.md
```

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [MongoDB](https://www.mongodb.com/try/download/community) (installed and running)
- [Git](https://git-scm.com/downloads)

### Steps

1. **Clone the repository**

   ```bash
   git clone https://github.com/your-username/apirestful-go.git
   cd apirestful-go
   go mod init apirestful-go
   ```

2. **Configure environment variables**

   Create a `.env` file in the root directory of the project with the following content:

   ```env
   SERVER_ADDR=:8080
   DB_CONN_STRING=mongodb://localhost:27017
   DATABASE_NAME=apirestful
   ```

   Make sure to adjust the values according to your configuration.

3. **Install dependencies**

   ```bash
   go mod tidy
   ```

## Usage

### Run the API

From the root directory of the project, execute:

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`.

### Available Endpoints

#### Users

- **Get all users**

  ```
  GET /api/v1/users
  ```

- **Create a new user**

  ```
  POST /api/v1/users
  ```

  **Request Body:**

  ```json
  {
    "username": "johndoe",
    "name": "John Doe",
    "email": "johndoe@example.com",
    "age": 30,
    "dni": "12345678",
    "phone": "1234567890",
    "country": "Country",
    "state": "State",
    "city": "City",
    "address": "Address",
    "postal_code": "1234",
    "password": "securepassword"
  }
  ```

- **Get a user by ID**

  ```
  GET /api/v1/users/{id}
  ```

- **Update a user**

  ```
  PUT /api/v1/users/{id}
  ```

  **Request Body:** *(similar to creation)*

- **Delete a user**

  ```
  DELETE /api/v1/users/{id}
  ```

#### Ping

- **Check server status**

  ```
  GET /api/v1/ping
  ```

  **Response:**

  ```json
  {
    "data": "pong"
  }
  ```

## Contribution

Contributions are welcome! Please follow these steps:

1. **Fork** the project.
2. Create a branch for your feature: `git checkout -b feature/new-feature`.
3. **Commit** your changes: `git commit -m 'Add new feature'`.
4. **Push** to the branch: `git push origin feature/new-feature`.
5. Open a **Pull Request**.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).