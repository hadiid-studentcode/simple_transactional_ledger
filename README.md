# Simple Transactional Ledger

This is a simple transactional ledger application built with Go, using the standard `net/http` package and MySQL as the database. It provides an API for managing financial accounts and their associated transactional entries.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Database Setup](#database-setup)
  - [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Accounts](#accounts)
  - [Entries](#entries)
- [Contributing](#contributing)
- [License](#license)

## Features

-   Create, Read, Update, and Delete (CRUD) operations for accounts.
-   Create, Read, Update, and Delete (CRUD) operations for transactional entries.
-   Database schema management through SQL migration files.
-   Environment variable-based configuration for database connection.

## Technologies Used

-   **Go**: Programming language.
-   **MySQL**: Relational database.
-   **`net/http`**: Go's standard library for HTTP servers.
-   **`github.com/joho/godotenv`**: For loading environment variables from a `.env` file.
-   **`github.com/go-sql-driver/mysql`**: MySQL driver for Go's `database/sql` package.

## Getting Started

Follow these instructions to set up and run the project locally.

### Prerequisites

Before you begin, ensure you have the following installed:

-   Go (version 1.25.5 or higher)
-   MySQL (version 5.7 or higher recommended)

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/simple_transactional_ledger.git
    cd simple_transactional_ledger
    ```

2.  **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Create a `.env` file:**
    Copy the `.env.example` file and rename it to `.env`. Update the placeholder values with your actual database credentials and application settings.

    ```bash
    cp .env.example .env
    ```

    Example `.env` content:
    ```
    DB_USER=root
    DB_PASSWORD=
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=simple_transaction_ledger_db
    APP_PORT=8080
    APP_URL=http://localhost
    ```

### Database Setup

1.  **Create the MySQL database:**
    Make sure your MySQL server is running. Then, create a database named `simple_transaction_ledger_db` (or whatever you configured in `DB_NAME` in your `.env` file).

    ```sql
    CREATE DATABASE simple_transaction_ledger_db;
    ```

2.  **Run migrations:**
    Apply the SQL migration files to set up the `accounts` and `entries` tables. You can execute these SQL files manually using a MySQL client (e.g., MySQL Workbench, DBeaver, or the `mysql` command-line client).

    -   `migrations/001_create_accounts_table.sql`
    -   `migrations/002_create_entries_table.sql`

    Example using `mysql` command-line client:
    ```bash
    mysql -u your_db_user -p simple_transaction_ledger_db < migrations/001_create_accounts_table.sql
    mysql -u your_db_user -p simple_transaction_ledger_db < migrations/002_create_entries_table.sql
    ```
    (Replace `your_db_user` with your MySQL username and enter your password when prompted.)

### Running the Application

1.  **Start the Go application:**
    ```bash
    go run main.go
    ```
    The server will start on the port specified in your `.env` file (default: `8080`). You should see output similar to:
    ```
    Server is running on port http://localhost:8080
    ```

## API Endpoints

The API base URL is `http://localhost:8080` (or `APP_URL:APP_PORT` from your `.env`).

### Accounts

Manage financial accounts.

-   **`GET /accounts`**
    -   **Description:** Get a list of all accounts.
    -   **Response:** `200 OK` with an array of account objects.
    -   **Example Response:**
        ```json
        [
            {
                "id": 1,
                "name": "Checking Account",
                "balance": 1000.50,
                "create_at": "2023-01-01T10:00:00Z",
                "update_at": "2023-01-01T10:00:00Z"
            }
        ]
        ```

-   **`GET /accounts/{id}`**
    -   **Description:** Get details of a specific account by ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `200 OK` with an account object. `400 Bad Request` if `id` is not an integer. `500 Internal Server Error` if account not found or other database error.
    -   **Example Response:**
        ```json
        {
            "id": 1,
            "name": "Checking Account",
            "balance": 1000.50,
            "create_at": "2023-01-01T10:00:00Z",
            "update_at": "2023-01-01T10:00:00Z"
        }
        ```

-   **`POST /accounts/create`**
    -   **Description:** Create a new account.
    -   **Request Body (Form Data):**
        -   `name` (string, required): Name of the account (must be unique).
        -   `balance` (float, required): Initial balance of the account.
    -   **Response:** `200 OK` with "Account created successfully" message. `400 Bad Request` if `balance` is invalid. `500 Internal Server Error` if there's a database error (e.g., duplicate name).
    -   **Example `curl`:**
        ```bash
        curl -X POST -d "name=Savings Account&balance=500.00" http://localhost:8080/accounts/create
        ```

-   **`PUT /accounts/update/{id}`**
    -   **Description:** Update an existing account by ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Request Body (Form Data):**
        -   `name` (string, required): New name for the account.
        -   `balance` (float, required): New balance for the account.
    -   **Response:** `200 OK` with "Account updated successfully" message. `400 Bad Request` if `id` is invalid. `500 Internal Server Error` on database error.
    -   **Example `curl`:**
        ```bash
        curl -X PUT -d "name=Updated Savings&balance=750.25" http://localhost:8080/accounts/update/1
        ```

-   **`DELETE /accounts/delete/{id}`**
    -   **Description:** Delete an account by ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `204 No Content` on successful deletion. `400 Bad Request` if `id` is invalid. `500 Internal Server Error` on database error.
    -   **Example `curl`:**
        ```bash
        curl -X DELETE http://localhost:8080/accounts/delete/1
        ```

### Entries

Manage transactional entries (transactions) associated with accounts.

-   **`GET /entries`**
    -   **Description:** Get a list of all entries, including associated account details.
    -   **Response:** `200 OK` with an array of entry objects.
    -   **Example Response:**
        ```json
        [
            {
                "id": 1,
                "account_id": 1,
                "amount": 100.00,
                "create_at": "2023-01-01T10:05:00Z",
                "update_at": "2023-01-01T10:05:00Z",
                "name": "Checking Account",
                "balance": 1000.50
            }
        ]
        ```

-   **`GET /entries/{id}`**
    -   **Description:** Get details of a specific entry by ID, including associated account details.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `200 OK` with an entry object. `400 Bad Request` if `id` is invalid. `500 Internal Server Error` if entry not found or other database error.
    -   **Example Response:**
        ```json
        {
            "id": 1,
            "account_id": 1,
            "amount": 100.00,
            "create_at": "2023-01-01T10:05:00Z",
            "update_at": "2023-01-01T10:05:00Z",
            "name": "Checking Account",
            "balance": 1000.50
        }
        ```

-   **`POST /entries/create`**
    -   **Description:** Create a new entry (transaction).
    -   **Request Body (Form Data):**
        -   `account_id` (integer, required): The ID of the account the entry belongs to.
        -   `amount` (float, required): The transaction amount.
    -   **Response:** `201 Created` with "Entry created successfully with ID {id}" message. `400 Bad Request` if `account_id` or `amount` are invalid. `500 Internal Server Error` on database error.
    -   **Example `curl`:**
        ```bash
        curl -X POST -d "account_id=1&amount=50.75" http://localhost:8080/entries/create
        ```

-   **`PUT /entries/update/{id}`**
    -   **Description:** Update an existing entry by ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Request Body (Form Data):**
        -   `account_id` (integer, required): The new account ID for the entry.
        -   `amount` (float, required): The new transaction amount.
    -   **Response:** `204 No Content` on successful update. `400 Bad Request` if `id`, `account_id`, or `amount` are invalid. `404 Not Found` if entry not found. `500 Internal Server Error` on database error.
    -   **Example `curl`:**
        ```bash
        curl -X PUT -d "account_id=1&amount=150.00" http://localhost:8080/entries/update/1
        ```

-   **`DELETE /entries/delete/{id}`**
    -   **Description:** Delete an entry by ID.
    -   **Parameters:** `id` (path parameter, integer)
    -   **Response:** `204 No Content` on successful deletion. `400 Bad Request` if `id` is invalid. `500 Internal Server Error` on database error.
    -   **Example `curl`:**
        ```bash
        curl -X DELETE http://localhost:8080/entries/delete/1
        ```

## Contributing

Feel free to fork the repository, open issues, and submit pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
