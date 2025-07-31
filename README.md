# Todo-Go CLI Application

A simple command-line todo list application built with Go using Cobra CLI framework.

## Features

- Add new todo items
- List all todo items
- Mark todo items as complete 
- Delete todo items
- Update todo items

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

## Usage

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

## Dependencies
- `github.com/spf13/cobra` - CLI framework

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a new Pull Request