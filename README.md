# Go Tutorial - First App

A simple "Hello, World!" application built with Go to demonstrate basic Go development workflow.

## Project Overview

This project contains a basic Go application that prints "Hello, World!" to the console. It serves as an introduction to Go programming and demonstrates the fundamental steps of creating, building, and running a Go application.

## Project Structure

```
.
├── README.md       # This documentation file
├── app.go         # Main Go source code file
├── go.mod         # Go module file
└── first-app      # Compiled binary executable
```

## Steps Followed to Create This Project

### 1. Initialize Go Module
First, we initialized a new Go module to manage dependencies and define the project:

```bash
go mod init example.com/first-app
```

This created the `go.mod` file with:
- Module name: `example.com/first-app`
- Go version: `1.25.2`

### 2. Create the Main Application File
We created `app.go` with a simple "Hello, World!" program:

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, World!")
}
```

Key components:
- `package main`: Declares this as an executable program
- `import "fmt"`: Imports the formatting package for output
- `func main()`: Entry point of the program
- `fmt.Print()`: Prints text to the console

### 3. Build the Application
We compiled the Go source code into an executable binary:

```bash
go build -o first-app app.go
```

This command:
- Compiles the `app.go` source file
- Creates an executable named `first-app`
- The `-o` flag specifies the output filename

### 4. Run the Application
The application can be run in multiple ways:

**Option 1: Run the compiled binary**
```bash
./first-app
```

**Option 2: Run directly with Go**
```bash
go run app.go
```

Both methods will output: `Hello, World!`

## Development Environment Setup

To work with this project, you need:

1. **Go Installation**: Go version 1.25.2 or compatible
2. **Text Editor**: Any text editor or IDE (VS Code, GoLand, vim, etc.)
3. **Terminal**: Command line interface for running Go commands

## Basic Go Commands Used

- `go mod init <module-name>`: Initialize a new Go module
- `go build`: Compile the Go program
- `go run <file.go>`: Compile and run Go program directly
- `go build -o <output-name> <file.go>`: Build with custom output name

## Next Steps

This basic project demonstrates:
- ✅ Go module initialization
- ✅ Basic Go syntax and structure
- ✅ Building and running Go applications
- ✅ Package management with go.mod

Potential improvements for learning:
- Add command-line arguments
- Create functions and organize code into packages
- Add error handling
- Implement user input
- Add unit tests
- Explore Go's standard library

## Running the Project

To run this project on your machine:

1. Ensure Go is installed on your system
2. Clone or download this project
3. Navigate to the project directory
4. Run one of these commands:
   ```bash
   # Run the compiled binary
   ./first-app
   
   # Or run directly with Go
   go run app.go
   ```

Expected output: `Hello, World!`