# WebSocket Room Management Library

A flexible Go library for managing WebSocket connections and rooms, built on top of the `gorilla/websocket` package. This library is designed for use cases such as real-time multiplayer games, chat applications, and more.

## Features

- **Room Management:** Create and manage rooms where WebSocket connections can join, leave, and communicate with each other.
- **Connection Handling:** Safely manage WebSocket connections with built-in concurrency handling.
- **Broadcast Messaging:** Broadcast messages to all connections within a room.
- **Extensible Architecture:** Easily extend the library with custom features and integrate it into your applications.

## Installation

To install the library, use `go get`:

```bash
go get github.com/kevbeltrao/websocket
```

## Usage

### Basic Example

Here’s a basic example of how to use the library:

```go
package main

import (
	"net/http"

	"github.com/kevbeltrao/websocket"
)

func main() {
	room := websocket.NewRoom("exampleRoom")
	go room.Run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.NewUpgrader()
		connection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}
		clientConnection := websocket.NewConnection(connection)
		room.Register <- clientConnection
	})

	http.ListenAndServe(":8000", nil)
}
```

In this example:
- A new room called `exampleRoom` is created.
- WebSocket connections are handled, and each connection is added to the room.
- The room broadcasts messages to all connected clients.

## Makefile Commands

This project includes a Makefile to streamline common tasks. Here are the available commands:

```bash
make          # Build the application
make test     # Run tests
make clean    # Clean the build
make fmt      # Format the code
make deps     # Install dependencies
make prepare  # Set up Git pre-commit hooks
make help     # Display this help message
```

### Building the Application

To build the application:

```bash
make build
```

### Running Tests

To run the tests:

```bash
make test
```

### Formatting the Code

To format the code:

```bash
make fmt
```

### Cleaning Up

To clean up the build artifacts:

```bash
make clean
```

### Installing Dependencies

To install the necessary dependencies:

```bash
make deps
```

### Setting Up Pre-Commit Hooks

This project includes a pre-commit hook that automatically runs tests before allowing a commit. To set up the pre-commit hook:

```bash
make prepare
```

## Testing

The library includes a comprehensive test suite using `testing`, `gomock`, and `testify`. To run the tests:

```bash
make test
```

<!-- ## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) file for more details on how to contribute to this project. -->

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gorilla WebSocket](https://github.com/gorilla/websocket) for providing a robust WebSocket implementation in Go.
- [Testify](https://github.com/stretchr/testify) and [Gomock](https://github.com/golang/mock) for their excellent testing tools.


