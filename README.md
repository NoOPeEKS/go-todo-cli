# Go Todo List CLI Application

This is a simple command-line Todo List application built using Go. It allows users to manage their tasks efficiently through the terminal.

## Features

- **Add Tasks:** Easily add new tasks to your list.
- **List Tasks:** View all pending tasks in your list.
- **Delete Tasks:** Remove unwanted tasks from the list.
- **Persistent Storage:** Tasks are saved between sessions through a database.

## Installation

### Prerequisites

- GoLang installed

### Steps

1. Clone this repository:

   ```bash
   git clone https://github.com/NoOPeEKS/go-todo-cli.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-todo-cli
   ```

3. Build the application:

   ```bash
   go build -o todo
   ```

4. Run the application:

   ```bash
   ./todo
   ```

## Usage

The application supports the following commands:

- `add <task_name>`: Add a new task to the list.
- `list`: Display all pending tasks.
- `delete <task_name>`: Remove a task from the list.
- `help`: Display usage instructions.

## Examples

- To add a task:

  ```bash
  ./todo add "Complete README file"
  ```

- To list all tasks:

  ```bash
  ./todo list
  ```

- To delete a task:
  ```bash
  ./todo delete task_name
  ```

## Contributing

Contributions are welcome! If you'd like to improve this project, feel free to fork the repository and submit a pull request.

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit/).
