# Product API

Product CRUD API

## Getting Started

### Prerequisites
- [Go](https://go.dev/dl/) (version 1.23 or higher recommended)

### Clone the repository

```bash
git clone https://github.com/Mzzqq/soulcode-product-api.git
cd soulcode-product-api
```

### Running the Application

```bash
# Install dependencies
go mod tidy

# Run the server
go run main.go
```
The server will be available at `http://localhost:3000`.

### Running Tests
To execute the test suite:
```bash
go test ./... -v
```

## API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/products` | Create a new product |
| `GET` | `/products` | List all products |
| `GET` | `/products/:id` | Get product details by ID |
| `PUT` | `/products/:id` | Update an existing product |
| `DELETE` | `/products/:id` | Remove a product |
| `GET` | `/swagger` | Swagger documentation |


## Design Notes

The project is structured into domain, usecase, and handler layers.

The handler layer handles HTTP requests and responses. The usecase layer contains business logic. The domain layer contains the Product entity and abstractions such as repository and cache interfaces.

An in-memory repository is used to keep the project simple and avoid external dependencies and it's used to demonstrate how product data can be cached when retrieved by ID.

Unit tests focus on the usecase layer because it contains the main business logic.