# Golang CRUD Project with PostgreSQL

## Overview

This Golang project is a simple CRUD (Create, Read, Update, Delete) application that utilizes PostgreSQL as its database. The project defines a `Product` struct to represent the data model for products.

## Features

- **Create:** Add new products to the database.
- **Read:** Retrieve product information from the database.
- **Update:** Modify existing product details.
- **Delete:** Remove products from the database.

## Technologies Used

- **Golang:** The project is built using the Go programming language.
- **PostgreSQL:** Data storage and retrieval are handled by PostgreSQL, a powerful open-source relational database.

## Project Structure

The project follows a standard Go project layout:


- **cmd:** Contains the main application entry point.
- **internal/app:** Houses the application-specific logic, including handlers, services, and repositories.
- **db/migrations:** Stores database migration files for managing database schema changes.
- **models:** Defines the data models, including the `Product` struct.

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-username/your-repo.git
   cd your-repo
   
go mod download

go run cmd/main.go

API Endpoints
POST /products: Create a new product.
GET /products/:id: Retrieve product details by ID.
PUT /products/:id: Update product details.
DELETE /products/:id: Delete a product.
