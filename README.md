# Golang CRUD Project with PostgreSQL

## Overview

This Golang project is a simple CRUD (Create, Read, Update, Delete) application that utilizes PostgreSQL as its database. The project defines a `Product` struct to represent the data model for products.

## Technologies Used

- **Golang:** The project is built using the Go programming language.
- **PostgreSQL:** Data storage and retrieval are handled by PostgreSQL, a powerful open-source relational database.
- **Docker:** postgres image is run using docker

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/brijeshjha/golang-crud-api.git
   cd golang-crud-api
   go mod download
   go run .

2. **Download docker to use this postgres image**
   ```bash
   docker run -d -p 5432:5432 --name my-postgres -e POSTGRES_PASSWORD=postgres postgres
   
3. **Create product table in postgres**
   ```bash
   CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       product_name VARCHAR(255) NOT NULL,
       price DOUBLE PRECISION NOT NULL,
       description TEXT
   );

4. **API Endpoints**

     ```bash
      POST /products: Create a new product.
      GET /products/:id: Retrieve product details by ID.
      PUT /products/:id: Update product details.
      DELETE /products/:id: Delete a product.
