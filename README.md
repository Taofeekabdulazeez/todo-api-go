# 📝 Todo API in Go

This is a simple RESTful Todo API built with **Golang**. It provides basic CRUD operations for managing todo items. The project is designed to demonstrate how to build a lightweight API using Go's built-in libraries.

---

## 🚀 Features

- Create a new todo
- Retrieve all todos
- Get a single todo by ID
- Update a todo by ID
- Delete a todo by ID

---

## 📦 Technologies Used

- **Go (Golang)** – Core language
- **net/http** – HTTP server and routing

---

## 📂 Project Structure

```
todo-api-go/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   ├── handler/             # HTTP request handlers
│   │   ├── middleware/          # HTTP middleware
│   │   ├── request/             # Request/response models
│   │   └── routes/              # Route definitions
│   ├── core/
│   │   ├── model/               # Domain models
│   │   └── service/             # Business logic
│   └── web/                     # Web-related components
├── pkg/
│   ├── config/                  # Configuration management
│   ├── database/                # Database connections
│   └── utils/                   # Utility functions
├── .air.toml                    # Air live reload config
├── .env                         # Environment variables
├── .gitignore                   # Git ignore file
├── Dockerfile                   # Docker configuration
├── go.mod                       # Go module file
├── go.sum                       # Go module checksums
├── Makefile                     # Build automation
├── README.md                    # Project documentation
└── request.http                 # HTTP request examples
```

---

## ⚙️ Environment Setup

The application requires several environment variables to function properly. The `.env` file is excluded from version control for security reasons.

### Required Environment Variables

Create a `.env` file in the project root with the following variables:

```bash
# Application
APP_PORT=8080
APP_ENV=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=todo_db
DB_SSLMODE=disable

# Session Management
SESSION_SECRET=your_super_secret_session_key_here
SESSION_MAX_AGE=86400

# Google OAuth
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
GOOGLE_CALLBACK_URL=http://localhost:8080/api/v1/auth/google/callback
GOOGLE_PROVIDER_SCOPES=email profile
```

### Setting Up Google OAuth

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the Google+ API
4. Create OAuth 2.0 credentials:
   - Go to "Credentials" → "Create Credentials" → "OAuth 2.0 Client ID"
   - Select "Web application"
   - Add authorized redirect URI: `http://localhost:8080/api/v1/auth/google/callback`
5. Copy the Client ID and Client Secret to your `.env` file

### Database Setup

1. Install PostgreSQL
2. Create a database:
   ```sql
   CREATE DATABASE todo_db;
   ```
3. Create a user (optional):
   ```sql
   CREATE USER todo_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE todo_db TO todo_user;
   ```

---

## 🛠️ Getting Started

### Prerequisites

- Go installed (v1.18+ recommended)
- PostgreSQL database (for production)
- Google OAuth credentials (for authentication)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Taofeekabdulazeez/todo-api-go.git
   cd todo-api-go
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   # Development mode
   go run ./cmd/server
   
   # Or use Air for hot reload
   air
   ```

5. **Build for production**
   ```bash
   go build -o todo-api ./cmd/server
   ./todo-api
   ```

---

## 📚 API Documentation

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication
The API uses Google OAuth for authentication. Users must authenticate via Google to access protected endpoints.

### Endpoints

#### Health Check
```bash
GET /api/v1/health
```

**Response:**
```json
{
  "status": "OK"
}
```

#### Todo Operations

**Get All Todos**
```bash
GET /api/v1/todos
Authorization: Bearer <session-token>
```

**Get Single Todo**
```bash
GET /api/v1/todos/{id}
Authorization: Bearer <session-token>
```

**Create Todo**
```bash
POST /api/v1/todos
Authorization: Bearer <session-token>
Content-Type: application/json

{
  "title": "Learn Go",
  "description": "Complete the Go tutorial",
  "completed": false
}
```

**Update Todo**
```bash
PUT /api/v1/todos/{id}
Authorization: Bearer <session-token>
Content-Type: application/json

{
  "title": "Learn Go",
  "description": "Complete the Go tutorial",
  "completed": true
}
```

**Delete Todo**
```bash
DELETE /api/v1/todos/{id}
Authorization: Bearer <session-token>
```

#### Authentication

**Google OAuth Login**
```bash
GET /api/v1/auth/google
```

**Google OAuth Callback**
```bash
GET /api/v1/auth/google/callback
```

**Logout**
```bash
POST /api/v1/auth/logout
Authorization: Bearer <session-token>
```

### Example Usage with curl

```bash
# 1. Authenticate with Google
curl -X GET "http://localhost:8080/api/v1/auth/google"

# 2. After authentication, create a todo
curl -X POST "http://localhost:8080/api/v1/todos" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-session-token>" \
  -d '{
    "title": "Complete project",
    "description": "Finish the todo API project",
    "completed": false
  }'

# 3. Get all todos
curl -X GET "http://localhost:8080/api/v1/todos" \
  -H "Authorization: Bearer <your-session-token>"
```

---

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

### How to Contribute

1. **Fork the repository**
   ```bash
   git clone https://github.com/your-username/todo-api-go.git
   cd todo-api-go
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add tests for new functionality
   - Update documentation as needed

4. **Run tests**
   ```bash
   go test ./...
   ```

5. **Commit your changes**
   ```bash
   git commit -m "feat: add your feature description"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**
   - Provide a clear description of your changes
   - Link any relevant issues
   - Ensure all tests pass

### Code Style Guidelines

- Follow Go's standard formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Handle errors appropriately

### Testing

- Write unit tests for new features
- Ensure all tests pass before submitting
- Aim for good test coverage
- Test edge cases and error conditions

### Issues

- Use the GitHub issue tracker for bug reports and feature requests
- Provide detailed information about bugs
- Include steps to reproduce issues
- Suggest solutions when possible

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### MIT License Summary

```
MIT License

Copyright (c) 2024 Taofeek Abdulazeez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## 📞 Contact

- **GitHub:** [@Taofeekabdulazeez](https://github.com/Taofeekabdulazeez)
- **Project Link:** [https://github.com/Taofeekabdulazeez/todo-api-go](https://github.com/Taofeekabdulazeez/todo-api-go)

---

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin) for the HTTP router
- [GORM](https://gorm.io/) for the ORM
- [Goth](https://github.com/markbates/goth) for OAuth authentication
- The Go community for excellent tools and libraries
