# Kanban Ticket Management Application BACKEND

This is a Kanban ticket management web application with a **Golang** backend (using **Gin**), a **PostgreSQL** database

## Features

- **Create Ticket**: Add a new ticket to the Kanban board.
- **View Tickets**: Retrieve and view all tickets.
- **Persistent Storage**: Tickets are stored in a PostgreSQL database.

## Prerequisites

- **Go** (version 1.16+)
- **PostgreSQL** (configured and running)

## Project Structure

- **repositories/**: Contains database interaction functions.
- **handlers/**: Contains HTTP handlers for API endpoints.


## I. Installation and Setup

### Clone the Repository

```bash
git clone <repository_url>
cd  testbackend
```

## 2.Set Up the API

### 1. Install dependencies:
```bash
go mod download
```

### 3. Configure Environment Variables:

*Create a `.env` file in the `server/` directory (or use defaults). Example:

```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=kanban
DB_HOST=localhost
DB_PORT=5432
```

## 3. Run Database Migrations:
* Ensure that the tickets table exists in PostgreSQL:

```
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    status VARCHAR(50),
    priority VARCHAR(50)
);

```

## 4. Start the Backend Server:
```bash
go run main.go
```

## API Endpoint

### 1. Get All Tickets
* 
```bash 
curl -X GET http://localhost:8080/tickets 
```


* URL: /tickets

* Method: GET

* Description: Retrieve a list of all tickets.

* Response:
```
[
  {
    "id": 1,
    "title": "Set up database",
    "status": "To Do",
    "priority": "High"
  },
  {
    "id": 2,
    "title": "Create API endpoints",
    "status": "In Progress",
    "priority": "Medium"
  }
]

```

### 2. Creating a Ticket:

```bash
curl -X POST http://localhost:8080/tickets -H "Content-Type: application/json" -d '{"title":"New Feature","status":"To Do","priority":"High"}'
```

Response:
```
{
  "id": 3,
  "title": "New Feature",
  "status": "To Do",
  "priority": "High"
}

```