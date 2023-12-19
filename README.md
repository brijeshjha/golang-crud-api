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

## Setup

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/brijeshjha/golang-crud-api.git
   cd golang-crud-api
   go mod download
   go run .
   
2. **create product table in postgres**
   ```bash
   CREATE TABLE products (
       id SERIAL PRIMARY KEY,
       product_name VARCHAR(255) NOT NULL,
       price DOUBLE PRECISION NOT NULL,
       description TEXT
   );

3. **API Endpoints**

     ```bash
      POST /products: Create a new product.
      GET /products/:id: Retrieve product details by ID.
      PUT /products/:id: Update product details.
      DELETE /products/:id: Delete a product.
