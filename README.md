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