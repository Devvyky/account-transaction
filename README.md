# Account Transaction API

This project is a Go-based account transaction api that connects to a PostgreSQL database.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Running Tests](#running-tests)
- [Docker Setup](#docker-setup)

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.21 or higher
- Docker
- Make
- golang-migrate

  ```bash
  ## install golang-migrate
  brew install golang-migrate

  ## check installation
  migrate -version
  ```

- PostgreSQL

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/devvyky/account-transaction.git
   cd account-transaction
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Configuration

Create a configuration file named `app.env` in the root directory of the project. Add the following configuration:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://postgres:postgres@localhost:5432/account_transaction?sslmode=disable
SERVER_ADDRESS=0.0.0.0:2000
```

Alternatively, you can set these environment variables in your shell.

## Running the Application

1. Start up PostgreSQL with docker using the following command below.

   ```bash
    ## start up postgres
    make postgres

    ## create a database
    make createdb
   ```

   Alternatively, ensure you have PostgreSQL running locally and accessible with the credentials provided in `app.env`.

2. Run database migrations:

   ```bash
   make migrateup
   ```

3. Start the application:

   ```bash
   make server
   ```

The server will start on the address specified in `SERVER_ADDRESS`.

## Running Tests

To run tests, ensure that the PostgreSQL database is running and accessible, then use the following command:

```bash
make test
```

## Docker Setup

### Using Docker Compose

1. Build and run the application and PostgreSQL using Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. The application will be accessible at `http://localhost:2000`.

### Stopping the Services

To stop the services, run:

```bash
docker-compose down
```
