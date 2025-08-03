# Todo-Go Application

A dual-interface todo list application built with Go, featuring both CLI and REST API capabilities.

## Features

- Add, list, update and delete todo items
- Mark todos as complete
- Command-line interface using Cobra
- REST API using Fiber framework
- Data persistence using SQLite

## Installation

1. Clone the repository
    ```bash
    git clone https://github.com/Nexain/todo-go.git
    ```

2. Navigate to project directory
    ```bash
    cd todo-go
    ```

3. Install dependencies
    ```bash
    go mod download
    ```

4. Build the application
    ```bash
    go build
    ```

## CLI Usage

### Add a new todo
```bash
./todo-go add "Your todo task"
```

### List all todos
```bash
./todo-go list
```

### Complete a todo
```bash
./todo-go complete <id>
```

### Delete a todo
```bash
./todo-go delete <id>
```

### Update a todo
```bash
./todo-go update <id> "New task description"
```

## API Usage

Start the API server:
```bash
./todo-go serve
```

The server runs on `http://localhost:3000` by default.

### API Endpoints

#### Get all todos
```http
GET /todos
```

#### Create a todo
```http
POST /todos
Content-Type: application/json

{
    "task": "Your todo task"
}
```

#### Update a todo
```http
PUT /todos/:id
Content-Type: application/json

{
    "task": "Updated task"
}
```

#### Complete a todo
```http
PATCH /todos/:id/complete
```

#### Delete a todo
```http
DELETE /todos/:id
```

## Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/gofiber/fiber/v2` - Web framework

## Development

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request