# Go File Operations Demo

This project demonstrates basic file operations in Go, including creating, reading, writing, and deleting files. It also includes error handling for each operation.

## Project Structure

```
go-file-demo
├── src
│   ├── main.go        # Entry point of the application
│   └── fileops.go     # Contains file operation functions
├── tests
│   └── fileops_test.go # Unit tests for file operations
├── go.mod             # Module definition
└── README.md          # Project documentation
```

## Getting Started

To run this project, ensure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

### Running the Application

1. Navigate to the project directory:
   ```
   cd go-file-demo/src
   ```

2. Run the application:
   ```
   go run main.go
   ```

### Running Tests

To run the tests for the file operations, navigate to the tests directory and execute the following command:

```
cd go-file-demo/tests
go test
```

This will execute the unit tests defined in `fileops_test.go` and report any failures.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.