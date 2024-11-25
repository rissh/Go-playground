# PodGreet: A Simple Go Web Server Application

**PodGreet** is a basic web server built using **Go**, designed to run in a containerized environment. It serves two HTTP routes: one to display a greeting with the URL path, and another with a friendly "Hi buddy!" message. 

The project is designed to showcase basic **Go** HTTP handling and containerization using **Podman** or **Docker**.

## Features
- **Go-based Web Server**: A lightweight Go web server using the `net/http` package.
- **Simple Routes**: Two routes:
  - **`/`** – Returns a greeting with the URL path.
  - **`/hi`** – Returns a friendly "Hi buddy!".
- **Containerized Application**: Easily deployable using Docker or Podman.
- **Port**: The server listens on port `8081`.

## How to Build and Run

### 1. Clone the repository:
 ```bash
    git https://github.com/rissh/Go-playground.git
    cd PodGreet
 ```

### 2. Build the Go application:
 ```bash
    go build -o main .
 ```

### 3. Run the application locally:
 ```bash
    ./main
 ```

Your server will be running at `http://localhost:8081`.

## Build and Run the Image with Docker
### 1. Build the Docker Image
Navigate to the directory where the Dockerfile is located and run the following command to build the image:
 ```bash
    docker build -t podgreet .
 ```

### 2. Run the Docker Container
Once the image is built, you can run the container using the following command:
 ```bash
    docker run -p 8081:8081 podgreet
 ```
This will start the application and map port `8081` on your local machine to port `8081` in the container.

## Build and Run the Image with Podman
### 1. Build the Podman Image
To build the image with Podman, run the following command in the directory containing the Dockerfile:

 ```bash
    podman build -t podgreet .
 ```

### 2. Run the Podman Container
Once the image is built with Podman, run the container using the following command:
 ```bash
    podman run -p 8081:8081 podgreet
 ```
This will start the application and map port `8081` on your local machine to port `8081` in the container.

## Docker or Podman Image Usage
Once you’ve built the image, you can push it to a container registry like **Docker Hub** or **Quay.io** for easy access.
To run the image from the registry, use the following commands:

### Quay.io:
 ```bash
    podman pull quay.io/rh-ee-rjagadal/podgreet:latest
    podman run -p 8081:8081 quay.io/rh-ee-rjagadal/podgreet:latest
 ```

## Access the Web Server

Once the application is running, visit the following URLs in your browser:

- `http://localhost:8081` – Displays a greeting with the URL path.
- `http://localhost:8081/hi` – Displays "Hi buddy!".
