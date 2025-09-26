# GoFrame Backend API Demo

A comprehensive backend API demo built with GoFrame framework, showcasing modern Go backend development practices.

## Features

- **RESTful API**: Complete CRUD operations for user management
- **Database Integration**: SQLite database with automatic table creation
- **Input Validation**: Request validation using GoFrame's built-in validator
- **Error Handling**: Comprehensive error handling and response formatting
- **CORS Support**: Cross-origin resource sharing enabled
- **Logging**: Structured logging with configurable levels
- **Pagination**: Built-in pagination support for list endpoints
- **API Documentation**: Self-documenting API endpoints

## Project Structure

```
go-quickstart/
├── config/
│   └── config.yaml          # Application configuration
├── internal/
│   ├── controller/          # HTTP handlers
│   │   ├── user.go         # User controller
│   │   └── route.go        # Route registration
│   ├── dao/                # Data access objects
│   │   └── user.go         # User DAO
│   ├── model/              # Data models
│   │   └── user.go         # User models and DTOs
│   └── service/            # Business logic
│       ├── user.go         # User service
│       └── service.go      # Service initialization
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
└── README.md              # This file
```

## Quick Start

### Prerequisites

- Go 1.25.1 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd go-quickstart
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8000`

## API Endpoints

### User Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/users` | Create a new user |
| GET | `/api/v1/users` | List users with pagination |
| GET | `/api/v1/users/{id}` | Get user by ID |
| PUT | `/api/v1/users/{id}` | Update user |
| DELETE | `/api/v1/users/{id}` | Delete user |

### System Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message |
| GET | `/api/v1/health` | Health check |
| GET | `/api/v1/docs` | API documentation |

## API Examples

### Create User
```bash
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Get User List
```bash
curl "http://localhost:8000/api/v1/users?page=1&size=10&search=john"
```

### Get User by ID
```bash
curl http://localhost:8000/api/v1/users/1
```

### Update User
```bash
curl -X PUT http://localhost:8000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john_updated",
    "email": "john.updated@example.com",
    "status": 1
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8000/api/v1/users/1
```

## Configuration

The application uses `config/config.yaml` for configuration. Key settings:

- **Server**: Port, address, and static file serving
- **Database**: SQLite connection settings
- **Logger**: Log levels and output destinations
- **CORS**: Cross-origin request settings

## Database

The application uses SQLite by default for simplicity. The database file (`data.db`) will be created automatically when the application starts.

### Database Schema

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    status INTEGER DEFAULT 1,
    create_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Development

### Adding New Features

1. **Models**: Add new data models in `internal/model/`
2. **DAO**: Create data access objects in `internal/dao/`
3. **Service**: Implement business logic in `internal/service/`
4. **Controller**: Add HTTP handlers in `internal/controller/`
5. **Routes**: Register new routes in `internal/controller/route.go`

### Code Structure

- **Controller**: Handles HTTP requests and responses
- **Service**: Contains business logic and validation
- **DAO**: Manages database operations
- **Model**: Defines data structures and validation rules

## Production Deployment

For production deployment:

1. Update database configuration to use PostgreSQL/MySQL
2. Set up proper logging configuration
3. Configure environment variables
4. Use a reverse proxy (nginx)
5. Set up SSL/TLS certificates
6. Configure monitoring and health checks

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Resources

- [GoFrame Documentation](https://goframe.org/)
- [GoFrame GitHub](https://github.com/gogf/gf)
- [GoFrame Examples](https://github.com/gogf/gf/tree/master/example)
