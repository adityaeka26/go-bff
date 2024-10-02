# Backend for Frontend (BFF) Implementation in Go

## Overview

Backend for Frontend (BFF) is an architectural pattern where a separate backend service is created specifically for each frontend (or type of frontend) in an application. This allows for a more tailored experience for different types of clients, such as web, mobile, or desktop applications. Each frontend interacts with its respective backend, which in turn communicates with the core backend services or APIs.

## Deployed URL
[https://bff.adityaeka.my.id/web/order/user/1](https://bff.adityaeka.my.id/web/order/user/1)

## API Doc
[Postman Collection](postman-collection.json)

## Installation

### Requirements
- Docker
- Docker Compose

### Steps
1. Clone the repository
    ```
    git clone https://github.com/adityaeka26/go-bff
    ```
2. Change the directory to the project directory
    ```
    cd go-bff
    ```
3. Deploy and run the system using Docker Compose
    ```
    docker compose up -d
    ```
4. Sometimes services exit due to a failure to connect to the database because it is still in the initialization state. In this case, just re-run docker compose

## Tech Stack
- Programming Language: Golang
- Router Lib: Go Fiber
- ORM: GORM
- Logger: zap
- Request Validation: Go Validator
- Communication Protocol: REST for frontend, gRPC for service-to-service
- System Architecture: Microservices, Backend for Frontend (BFF)
- Code Pattern: Clean Architecture
- Database: PostgreSQL
- Container: Docker

## System Architecture
![alt text](<System Architecture.jpg>)

## Contributors
1. Aditya Eka B.