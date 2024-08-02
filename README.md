# Clothing Inventory Management API

## Description

This API is designed to manage the inventory for a clothing store. It allows for creating, reading, updating, and deleting clothing items, as well as managing stock levels. The system supports querying and filtering based on color and size, adding and reducing stock, and retrieving information about items with low or empty stock.

## Tech Stack

This project is developed using the following technologies:

- **Go:** The primary programming language used for building the API. Go provides strong performance and concurrency features, making it ideal for scalable backend systems.

- **Gin:** A web framework for Go that is used to handle HTTP requests and routes. Gin is known for its high performance and ease of use, making it a popular choice for developing RESTful APIs.

- **PostgreSQL:** The relational database management system used to store clothing items and their attributes. PostgreSQL is known for its robustness and support for advanced features.

- **GORM:** An Object-Relational Mapping (ORM) library for Go that interacts with the PostgreSQL database. GORM simplifies database operations and provides an easy-to-use interface for managing database records.


## Database Schema

### Table: `clothes`

- **id** (Primary Key): Integer, auto-incremented
- **color**: VARCHAR(50), not null
- **size**: VARCHAR(10), not null (Enum: S, M, L, XL, XXL)
- **price**: NUMERIC(10, 2), not null
- **stock**: Integer, not null

## Endpoints

### Create Clothing Item

- **Endpoint:** `POST /api/v1/clothes`
- **Description:** Create a new clothing item.
- **Payload :**
  ```json
  {
      "color": "string", // Required
      "size": "string",  // Required, Enum: S, M, L, XL, XXL
      "price": 0.0, // Required
      "stock": 0 // Required
  }
  ```
- **Response:**
  - **Success (201 Created):** The item has been successfully created.
  - **Failure :**
    - **400 Bad Request:** Payload doesn't pass validation
    - **500 Internal Server Error:** Unexpected condition was encountered on the server side

### Fetch Clothing Items

- **Endpoint:** `GET /api/v1/clothes`
- **Description:** Retrieve a list of clothing items.
- **Query Params (optional):**
    - **color** : Filter by color
    - **size** : Filter by size (Enum: S, M, L, XL, XXL) 
- **Response:**
    - **Success (200 Created):** The item has been fetched successfully.
        ```json
        [
          {
              "id": uuid,
              "color": "string",
              "size": "string",
              "price": 0.0,
              "stock": 0
          }
        ]
        ```
    - **400 Bad Request:** Size query params doesn't pass validation (enum)
    - **500 Internal Server Error:** Unexpected condition was encountered on the server side
Note : If no query parameters are provided or no items match the query, an empty array is returned with a 200 OK status.

### Update Clothing Item

- **Endpoint:** `PUT /api/v1/clothes/{id}`
- **Description:** Update an existing clothing item by its ID.
- **Payload:**
  ```json
  {
      "color": "string",
      "size": "string",  // Enum: S, M, L, XL, XXL
      "price": 0.0,
      "stock": 0
  }
  ```
- **Response:**
  - **Success (200 OK):** The item has been successfully updated.
    - **Failure :**
      - **400 Bad Request:** Payload doesn't pass validation
      - **404 Not Found:** The item with the specified ID does not exist.
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side
  
### Delete Clothing Item

- **Endpoint:** `DELETE /api/v1/clothes/{id}`
- **Description:** Delete a clothing item by its ID.
- **URL Parameters:**
  - `id` (required): The unique identifier of the clothing item to be deleted.
- **Response:**
  - **Success (200 OK):** The item has been successfully deleted.
  - **Failure :**
      - **404 Not Found:** The item with the specified ID does not exist.
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side

### Add Stock

- **Endpoint:** `POST /api/v1/clothes/{id}/stock/add`
- **Description:** Add stock to an existing clothing item.
- **URL Parameters:**
  - `id` (required): The unique identifier of the clothing item to which stock should be added.
- **Payload:**
  ```json
  {
      "quantity": 0
  }
  ```
- **Response:**
  - **Success (200 OK):** The item stock has been successfully updated.
  - **Failure :**
      - **400 Bad Request:** Payload doesn't pass validation
      - **404 Not Found:** The item with the specified ID does not exist.
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side

### Reduce Stock

- **Endpoint:** `POST /api/v1/clothes/{id}/stock/reduce`
- **Description:** Decrease stock to an existing clothing item.
- **URL Parameters:**
  - `id` (required): The unique identifier of the clothing item to which stock should be reduce.
- **Payload:**
  ```json
  {
      "quantity": 0
  }
  ```
- **Response:**
  - **Success (200 OK):** The item stock has been successfully updated.
  - **Failure :**
      - **400 Bad Request:** Payload doesn't pass validation
      - **404 Not Found:** The item with the specified ID does not exist.
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side

### Get Items with Empty Stock

- **Endpoint:** `GET /api/v1/clothes/stock/empty`
- **Description:** Retrieve a list of clothing items with zero stock.
- **Response:**
  - **Success (200 OK):** Returns a list of clothing items that have zero stock.
  - **Response Format:**
    ```json
    [
        {
            "id": uuid,
            "color": "string",
            "size": "string",
            "price": 0.0,
            "stock": 0
        }
    ]
    ```
  - **Failure :**
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side

### Get Items with Empty Stock

- **Endpoint:** `GET /api/v1/clothes/stock/low`
- **Description:** Retrieve a list of clothing items with low stock (<5).
- **Response:**
  - **Success (200 OK):** Returns a list of clothing items that have less than 5 stock.
  - **Response Format:**
    ```json
    [
        {
            "id": uuid,
            "color": "string",
            "size": "string",
            "price": 0.0,
            "stock": 0
        }
    ]
    ```
  - **Failure :**
      - **500 Internal Server Error:** Unexpected condition was encountered on the server side
