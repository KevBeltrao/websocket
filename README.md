# WebSocket Room Management Library

A simple, efficient Go library built on top of the `gorilla/websocket` package that provides room management for WebSocket connections. This library is designed to be flexible enough for a variety of use cases, such as real-time multiplayer games, chat applications, and more.

## Features

- **Room Management:** Create and manage rooms where WebSocket connections can join, leave, and communicate with each other.
- **Connection Handling:** Safely manage WebSocket connections with built-in concurrency handling.
- **Broadcast Messaging:** Broadcast messages to all connections within a room.
- **Extensible Architecture:** Easily extend the library with custom features and integrate it into your applications.
- **Mocking Support:** Built-in support for testing with `gomock` for connection interfaces.

## Installation

To install the library, use `go get`:

```bash
go get github.com/kevbeltrao/websocket
```

## WIP
