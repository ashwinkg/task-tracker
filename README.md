# Task Tracker CLI

A simple command-line task management application built with Go. Track your tasks with ease by adding, updating, deleting, and marking tasks as in-progress or done.

## Features

- ✅ Add new tasks with descriptions
- 📝 Update existing task descriptions
- 🗑️ Delete tasks
- 🔄 Mark tasks as in-progress or done
- 📋 List all tasks or filter by status (todo, in-progress, done)
- 💾 Persistent storage using JSON file
- 🕒 Automatic timestamp tracking for creation and updates

## Installation

### Prerequisites

- Go 1.23.3 or later

### Build from source

1. Clone the repository:
   ```bash
   git clone https://github.com/ashwinkg/task-tracker.git
   cd task-tracker
   ```

2. Build the application:
   ```bash
   go build -o task-cli ./app
   ```

3. (Optional) Move the binary to a directory in your PATH:
   ```bash
   sudo mv task-cli /usr/local/bin/
   ```

## Usage

The application uses a JSON file (`tasks.json`) to store tasks in the current directory.

### Commands

#### Add a new task
```bash
./task-cli add "Buy groceries"
```

#### Update a task
```bash
./task-cli update <id> "Updated task description"
```

#### Delete a task
```bash
./task-cli delete <id>
```

#### Mark a task as in-progress
```bash
./task-cli mark-in-progress <id>
```

#### Mark a task as done
```bash
./task-cli mark-done <id>
```

#### List all tasks
```bash
./task-cli list
```

#### List tasks by status
```bash
./task-cli list todo
./task-cli list in-progress
./task-cli list done
```

### Examples

```bash
# Add some tasks
./task-cli add "Write documentation"
./task-cli add "Implement feature X"
./task-cli add "Test the application"

# List all tasks
./task-cli list

# Mark first task as in-progress
./task-cli mark-in-progress 1

# Update second task
./task-cli update 2 "Implement feature Y with improvements"

# Mark first task as done
./task-cli mark-done 1

# List only done tasks
./task-cli list done

# Delete a task
./task-cli delete 3
```

### Task Statuses

- **todo**: Default status for new tasks
- **in-progress**: Task is currently being worked on
- **done**: Task has been completed

### Data Storage

Tasks are stored in a `tasks.json` file in the current working directory. The file is created automatically when you first run the application. Each task includes:

- Unique ID
- Description
- Status
- Creation timestamp
- Last update timestamp

## Error Handling

The application includes basic error handling for:
- Invalid task IDs
- Empty task descriptions
- Unknown operations
- File I/O errors

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the terms specified in the LICENSE file.