# Order System

## Introduction

This is a Go project that implements an order system. The system allows users to create, edit, and delete orders.

## Prerequisites

* Docker installed
* Docker Compose installed
* Go installed (you can download it from golang.org)

## Setting Up the Environment

1. **Starting the Containers**

   In the project's root directory, run the following command to start the required containers:

   ```bash
   docker-compose up
   ```

   This command will start the database services needed for the project.

2. **Running the Migrations**

   Once the containers are up and running, execute the following command to apply the database migrations:

   ```bash
   make migrate
   ```

  This will create the necessary tables in the database.

## Running the Main Program

In the cmd/ordersystem directory, run the following command to start the main program:
   ```bash
   go run main.go wire_gen.go
   ```

This will start the Order System application.

## Testing the API Routes

In the `api` directory, you'll find the test files for the API routes.

## License

This project is licensed under the MIT license.
