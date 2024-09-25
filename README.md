# CLI Chat Application

This is a command-line interface (CLI) chat application built using Go, PostgreSQL, and socket programming with `go-socket.io`. The application allows users to register, log in, and chat in real-time with other logged-in users. All user data and chat history are stored in a PostgreSQL database.

## Features

- **User Registration**: Users can register with a username, email, and password.
- **User Login**: Registered users can log in using their username and password.
- **Real-time Chat**: Logged-in users can send and receive messages in real-time.
- **PostgreSQL Integration**: User data and chat history are persisted in a PostgreSQL database.
- **Socket Programming**: Real-time communication is handled via `go-socket.io` for both the server and client.

## Technologies Used

- [Go](https://golang.org/): The primary language for building the CLI and server components.
- [PostgreSQL](https://www.postgresql.org/): Used to store user data and chat history.
- [go-socket.io](https://github.com/googollee/go-socket.io): Socket library for real-time communication between clients and the server.
- [pgx](https://github.com/jackc/pgx): PostgreSQL driver for Go.
- [spf13/cobra](https://github.com/spf13/cobra): A library to create a powerful CLI.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [DBeaver](https://dbeaver.io/download/) or another PostgreSQL management tool (optional)

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/cli-chat-app.git
   cd cli-chat-app
   ```

2. **Install dependencies**:

   Ensure you have Go installed and set up in your `$PATH`.

3. **Set up PostgreSQL**:

   - Install PostgreSQL.
   - Create a new database, and add a `.env` file to the project root with the following variables:

     ```bash
     DB_USERNAME=your_postgres_username
     DB_PASSWORD=your_postgres_password
     DB_NAME=your_database_name
     ```

4. **Run database migrations**:

   Ensure the required tables for users and chat history are set up in PostgreSQL.

   You can manually create the tables or use a tool like DBeaver to execute SQL commands.

5. **Run the application**:

   ```bash
   go run main.go
   ```

## Usage

The CLI application provides commands to register users, log in, and start chatting.

### 1. Register a new user

```bash
go run main.go register <username> <email> <password>
```

### 2. Login with an existing user

```bash
go run main.go login <userame> <password>
```

### 3. Start Chatting

After logging in, you can send and receive messages from other users logged into the chat application. Messages are broadcast to all connected users in real-time.

### 4. Running the server

```bash
go run main.go server
```

This will start the Socket.io server at localhost:8000
