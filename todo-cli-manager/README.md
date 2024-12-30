# Todo CLI Manager

This is a personal project written in Go to learn more about the file and I/O system in Go. It is a simple command-line interface (CLI) application to manage a todo list.

## Features

- Add a new todo item
- List all todo items
- Mark a todo item as completed
- Delete a todo item

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/kbbansi/learning-go.git
    ```
2. Navigate to the project directory:
    ```sh
    cd todo-cli-manager
    ```
3. Build the application:
    ```sh
    go build
    ```

## Usage

### Add a new todo item
```sh
./todo-cli-manager add "Buy groceries"
```

### List all todo items
```sh
./todo-cli-manager list
```

### Mark a todo item as completed
```sh
./todo-cli-manager complete 1
```

### Delete a todo item
```sh
./todo-cli-manager delete 1
```

## Contributing

Feel free to fork this repository and submit pull requests. Any contributions are welcome!

## License

This project is licensed under the MIT License.