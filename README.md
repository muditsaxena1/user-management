# User Management Service
This project is a user management service built with Go, using the Gin framework. It allows you to store user details, retrieve a user by ID, and list all users. Basic authentication is required to interact with the API.

## Features
* __Create User__: Store user details, including ID, name, and signup time.
* __Retrieve User__: Get details of a specific user by their ID.
* __List Users__: Retrieve a list of all users stored in the system.
* __Basic Authentication__: All operations require basic authentication.

## Project Structure
The project is organized as follows:
```
url-shortener/
│
├── cmd/server/
│   └── main.go             # Entry point of the application
│
├── internal/
│   ├── api/                # Handlers, routes, and middlewares
│   ├── errors/             # Custom error logic
│   ├── models/             # Data structures and validation logic
│   ├── storage/            # Storage interface and in-memory implementation
│
└── go.mod                  # Go module file
```
## Prerequisites
Go 1.23 or higher
Docker (for containerization)

# Getting Started
## 1. Clone the repository
```
git clone https://github.com/muditsaxena1/user-management.git
cd user-management
```
## 2. Build and Run Locally
You can use the provided Makefile to manage the project.
### Format, Vendor, and Tidy Go Modules
```
make tidy
```
### Run the Application
```
make run PORT=8080
```
If no port is specified it defaults to 8080.
### Run Tests
```
make test
```
### Test Coverage
```
make coverage
```
## 3. Build and Run with Docker
### Build the Docker Image
```
make docker-build
```
### Run the Docker Container
```
make docker-run PORT=8080
```
# API Endpoints
## Create a User
Request:
```
curl --location 'http://localhost:8080/v1/user' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--data '{
    "id": "1234",
    "name": "Jon Doe",
    "signupTime": 1622548800000
}'
```
Response:
```
{
    "status": "user set successfully"
}
```
## Get User by ID
Request:
```
curl --location 'http://localhost:8080/v1/user/123' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='
```
Response:
```
{
    "id": "123",
    "name": "Jon Doe",
    "signupTime": 1622548800000
}
```
## List all users
Request:
```
curl --location 'http://localhost:8080/v1/users' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='

```
Response:
```
[
    {
        "id": "1234",
        "name": "John Wick",
        "signupTime": 1622548800000
    },
    {
        "id": "5678",
        "name": "Jane Doe",
        "signupTime": 1622635200000
    }
]

```
# Future Improvements
### 1. Database Integration:
Use a database like PostgreSQL or MySQL for persistent storage. This will allow the service to retain shortened URLs and metrics across restarts.
Include a docker-compose.yml file to easily set up the database along with the service.
### 2. Improved Logging:
Implement logging to a file for better traceability and debugging.
Use structured logging libraries like logrus or zap.
### 3. Authentication:
Improve authentication (e.g., API keys, OAuth) for URL shortening to restrict usage.
### 4. Enhanced Comments:
Improve inline documentation and comments to make the codebase more understandable for new contributors. 

# License
This project is licensed under the MIT License.
