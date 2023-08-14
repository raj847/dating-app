# dating-app

# Service Structure Documentation

Welcome to the repository! This README provides an overview of the structure of the service. Our codebase is divided into several packages/folders to ensure a clean separation of concerns and improve maintainability.

## Table of Contents
- [Repository](#repository)
- [Service](#service)
- [Handler](#handler)
  - [Request](#request)
  - [Response](#response)
- [Utils](#utils)
- [Main Package](#main-package)

## Repository
The `Repository` package is where we interact with the database. It contains all the necessary operations for querying and updating the database.

## Service
The `Service` package is responsible for housing the business logic of our application. It acts as an intermediary between the Repository and Handler, processing the data to meet the required specifications.

## Handler
The `Handler` package manages the requests and responses. Within this package, you'll find:

### Request
This sub-package within `Handler` is dedicated to storing the entities related to requests, ensuring that data entering the service is appropriately structured.

### Response
Another sub-package of `Handler`, the `Response` section is where entities and logic related to responses are kept. This ensures a consistent and well-defined output for consumers of the service.

## Utils
The `Utils` package is pivotal for setting up and configuring our service. Here, you'll find configurations for the database, GORM settings, and helper utilities, such as hashing passwords.

## Main Package
Lastly, the main package serves as the execution point for the entire codebase. It integrates all the other packages and ensures the smooth running of the service.

---

Thank you for exploring our service structure! Should you have any questions or suggestions, please feel free to raise an issue or submit a pull request.

# Instructions on How to Run the Service

Follow these steps to set up and run the service on your local environment:

## Table of Contents
- [Download or Clone Repository](#download-or-clone-repository)
- [Install Dependencies](#install-dependencies)
- [Database Setup](#database-setup)
- [Environment Variables Configuration](#environment-variables-configuration)
- [Running the Service](#running-the-service)
- [Additional Resources](#additional-resources)

## Download or Clone Repository
Begin by downloading or cloning the repository to your local machine.

```bash
git clone https://github.com/raj847/dating-app.git
```

## Install Dependencies
Once you have the repository on your machine, navigate to the root directory and install all the necessary packages using:

```bash
go mod download
```

## Database Setup
Ensure that PostgreSQL is running on your machine. If it's not installed yet, you'll need to set it up before proceeding.

## Environment Variables Configuration

Environment variables are crucial for keeping sensitive information and configuration details outside of the main codebase. This not only enhances security but also provides flexibility, as you can modify these variables without affecting the main application logic. 

Below, we break down each environment variable present in the `sample.env` file:

### Table of Contents

- [DATABASE_URL](#database_url)
- [SIGNED_TOKEN](#signed_token)
- [PORT](#port)

### DATABASE_URL

**Description**: 
This variable points to the connection string for the PostgreSQL database. It provides all the essential details required to establish a connection to the database, such as the user, password, database host, port, and database name.

**Format**: 
```
postgres://[USER]:[PASSWORD]@[HOST]:[PORT]/[DATABASE_NAME]
```

**Sample Value**: 
```
postgres://postgres:postgres@localhost:5432/user
```

### SIGNED_TOKEN

**Description**: 
The `SIGNED_TOKEN` is used for token generation and verification, often applied for JWT (JSON Web Tokens) or other authentication mechanisms. By keeping this value secret and outside the main code, you ensure that your tokens are secure and cannot be tampered with.

**Format**: 
```
[Any string that acts as a secret key]
```

**Sample Value**: 
```
secret
```

### PORT

**Description**: 
This variable defines the port on which your service will run. By externalizing this configuration, you can easily switch ports when required without altering the code.

**Format**: 
```
:[PORT_NUMBER]
```

**Sample Value**: 
```
:1323
```

---

Always ensure that your `.env` file (where you'd place actual values for these variables) is included in your `.gitignore` to prevent accidental pushes of sensitive information to public repositories. If you need to share how to set up these variables, always use the `sample.env` or similar as a template.

## Running the Service
With everything set up, you can now run the service with the following command:

```bash
go run main.go
```

## Additional Resources

- For **Linting Jobs**: Visit [this link](https://github.com/raj847/dating-app/actions/runs/5848212167/job/15855261943).
  
- For **Automatic Testing/Unit Tests**: Access [this link](https://github.com/raj847/dating-app/actions/runs/5848212167/job/15855261943).

- **Postman Test Collection**: You can view and import the Postman test collection from [here](https://documenter.getpostman.com/view/16407134/2s9Xy5LqMr).

---

We hope this guide helps you get started with the service. If you encounter any issues or have suggestions, please raise an issue or contribute by submitting a pull request.