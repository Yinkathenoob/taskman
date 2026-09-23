# taskman

A simple command-line task manager written in Go. Add tasks, list them, mark them done, and your tasks persist between sessions via a local JSON file.

## Features

- Add tasks with auto-incrementing IDs
- List all tasks with completion status
- Mark tasks as done
- Persistent storage — tasks are saved to `tasks.json` and reloaded on startup

## Requirements

- Go 1.21 or later

## Running

Clone the repo and run:

```bash
go run .
```

## Usage

The program presents a menu:

--- Task Manager ---

Add a task
List tasks
Mark a task done
Quit
Choose an option:


Example session:

Choose an option: 1
Enter task title: Learn Go
Added task #1: Learn Go

Choose an option: 2
Your tasks:
[ ] #1 Learn Go

Choose an option: 3
Enter task ID to mark done: 1
Marked task #1 as done.

Tasks are saved automatically when you quit and restored the next time you run the program.

## What I learned

This project was built to practice core Go concepts: structs, slices, JSON serialization, file I/O, error handling, and building an interactive CLI loop.