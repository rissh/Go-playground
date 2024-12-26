# GoRabbitMQ

GoRabbitMQ is a simple project demonstrating how to use RabbitMQ with Go for message publishing and consumption. The project uses the [RabbitMQ](https://www.rabbitmq.com/) message broker and the `amqp091-go` library to establish connections, send messages, and receive them using Go.

## Features

- Publish messages to a RabbitMQ queue.
- Consume messages from a RabbitMQ queue.
- Example implementation of RabbitMQ's `QueueDeclare` and `Consume`.

---

## Prerequisites

1. Install Docker to run RabbitMQ.
   - [Download Docker](https://www.docker.com/get-started)

2. Install Go (1.20 or later recommended).
   - [Download Go](https://golang.org/dl/)

---
## Project Structure
```yaml
GoRabbitMQ/
├── docker-compose.yml   # Docker setup for RabbitMQ
├── send/
│   └── send.go          # Code to send messages to RabbitMQ
├── receive/
│   └── receive.go       # Code to receive messages from RabbitMQ
├── go.mod               # Go module configuration
└── README.md            # Project documentation
```
---

## Getting Started

Follow the steps below to set up and run the project.

### 1. Set up RabbitMQ with Docker

Run the RabbitMQ service using Docker Compose. This will start RabbitMQ with the management console enabled.

```bash
docker-compose up -d
```

RabbitMQ will now be available at:

- Management UI: http://localhost:15672/
    - Username: guest
    - Password: guest
- Default port for messaging: 5672

To stop the RabbitMQ service:

```bash
docker-compose down
```

### 2. Install Dependencies
Initialize your Go project and install the required dependency github.com/rabbitmq/amqp091-go.

```bash
go mod tidy
```

### 3. Running the Code
**Sending Messages:**
To send a message to RabbitMQ, run the send.go file:

```bash
go run send.go
```
This will send a message to the queue named Hello.

**Receiving Messages:**
To consume messages from RabbitMQ, run the receive.go file:

```bash
go run receive.go
```
This will start consuming messages from the Hello queue and print them to the console.

## Example Output
#### `send.go`
When you run the sender, it will produce output like:

```plaintext
[x] Sent Hello World!
```
#### `receive.go`
When you run the receiver, it will output:

```plaintext
[*] Waiting for messages. To exit press CTRL+C
Received a message: Hello World!
```

## Code Explanation
#### `send.go`
- Establishes a connection with RabbitMQ.
- Declares a queue named Hello.
- Publishes a Hello World! message to the queue.

#### `receive.go`
- Establishes a connection with RabbitMQ.
- Declares the same queue (Hello) to consume messages.
- Listens for incoming messages and prints them to the console.

## Docker Compose Configuration
The docker-compose.yml file sets up a RabbitMQ instance with the following configuration:

```yaml
version: '3'

services:
  rabbitmq:
    image: rabbitmq:3.11-management-alpine
    environment:
      RABBITMQ_DEFAULT_USER: guest
      RABBITMQ_DEFAULT_PASS: guest
    ports:
      - "5672:5672"
      - "15672:15672"
```

- `5672`: Port for RabbitMQ messaging.
- `15672`: Port for RabbitMQ management UI.

## Troubleshooting
Common Issues
RabbitMQ Not Running: Ensure RabbitMQ is running with docker-compose up -d.

Connection Refused: Verify RabbitMQ is running on localhost and ports 5672 and 15672 are accessible.

Module Import Errors: Run go mod tidy to install missing dependencies.

## Contributing
Feel free to fork this repository, open issues, or submit pull requests to improve this project.





