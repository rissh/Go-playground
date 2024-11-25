#Movie Mate Crud API's

A simple RESTful API for managing a collection of movies. This project is designed to demonstrate CRUD (Create, Read, Update, Delete) operations using **Go** with the `mux` router.

---

## Features

- **Get All Movies**: Retrieve a list of all movies in the collection.
- **Get a Movie**: Retrieve details of a specific movie by its ID.
- **Create a Movie**: Add a new movie to the collection.
- **Update a Movie**: Modify the details of an existing movie by its ID.
- **Delete a Movie**: Remove a movie from the collection by its ID.

---

## Technologies Used

- **Language**: Go (Golang)
- **Router**: [gorilla/mux](https://github.com/gorilla/mux)
- **JSON Encoding/Decoding**: For handling request and response payloads.

---

## Project Structure

- **`main.go`**: Contains the entire implementation of the Movies API, including:
  - Definition of `Movie` and `Director` structs.
  - Handlers for all CRUD operations.
  - Initial sample data for movies.
  - API routing.

---

## Endpoints

| HTTP Method | Endpoint           | Description                     | Example Payload (for POST/PUT)                                    |
|-------------|--------------------|---------------------------------|-------------------------------------------------------------------|
| **GET**     | `/movies`          | Retrieve all movies            | N/A                                                               |
| **GET**     | `/movies/{id}`     | Retrieve a movie by its ID      | N/A                                                               |
| **POST**    | `/movies`          | Add a new movie                | `{ "isbn": "12345", "title": "New Movie", "director": { "firstname": "John", "lastname": "Doe" } }` |
| **PUT**     | `/movies/{id}`     | Update an existing movie by ID | `{ "isbn": "67890", "title": "Updated Movie", "director": { "firstname": "Jane", "lastname": "Smith" } }` |
| **DELETE**  | `/movies/{id}`     | Delete a movie by its ID        | N/A                                                               |

---

## What This Project Demonstrates

- **CRUD Operations**: The project demonstrates how to perform basic Create, Read, Update, and Delete operations using Go.
- **RESTful Principles**: The API adheres to REST principles, using HTTP methods to interact with resources.
- **JSON Handling**: Shows how to encode and decode JSON payloads in Go.
- **Routing with `mux`**: Demonstrates the use of the `gorilla/mux` package for building REST APIs in Go.
- **Thread Safety**: By introducing `sync.Mutex`, the code ensures safe access to shared resources in a concurrent environment.

---

## How to Run

1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd MovieMateCRUD

2. Install the required dependencies:
   ```bash
   go mod init MovieMateCRUD
   go get -u github.com/gorilla/mux

3. Run the application:
   ```bash
   go run main.go

4. The server will start on port 8000. You can access the API at:
   ```bash
   http://localhost:8000
   
