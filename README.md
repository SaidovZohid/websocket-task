# WebSocket Service

This project implements a WebSocket service in Go, designed to handle client connections, reconnections, message broadcasting, and message tracking using SQLite.

## Table of Contents

- [Setup and Running Instructions](#setup-and-running-instructions)
- [API Documentation](#api-documentation)
- [Testing](#testing)

## Setup and Running Instructions

### Prerequisites

- Go 1.18 or later
- SQLite for the database
- [github.com/gorilla/websocket](https://github.com/gorilla/websocket) for WebSocket handling
- [github.com/stretchr/testify](https://github.com/stretchr/testify) for testing

### Installation
1. **Clone the repository:**
```bash
   git clone https://github.com/SaidovZohid/websocket-task.git
   cd websocket-task
```

2. **Install dependencies:**
```bash
   go mod tidy && go mod vendor
```

3. **Run the server:**
```bash
go run main.go
```

5. **Connect to the WebSocket:**
Use a WebSocket client (like Postman) to connect to the server at the following URL, providing a user_id in the query parameters:
```bash
ws://localhost:8080/ws?user_id=<YOUR_USER_ID>
```

## API Documentation
### Websocket Endpoint
-  **Endpoint: **ws/
- **Method:** GET
- **Query Parameters:** ```user_id: (string)``` A unique identifier for the client

### Message Broadcasting
- Broadcast: The server broadcasts messages to all connected clients. Each client receives messages only when they are online.

## Testing

### Running Tests
To run the unit tests for the WebSocket service, execute:
```bash
go test ./...
```

## Roadmap

This project was developed as part of a test task for the Golang Backend Developer position at Hyssa Company. The implementation showcases fundamental skills in WebSocket service design, client management, and message tracking, with a focus on performance and reliability.

The code is demonstrates readiness to contribute to real-world backend challenges at Hyssa Company.
