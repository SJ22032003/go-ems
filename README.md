# Go-EMS

Go-EMS is a REST API project for running an event management server. It is built using Go and the Gin framework.

## Project Structure

The project is structured into several packages:

- `controllers`: Contains the handlers for the HTTP endpoints.
- `db`: Contains the database connection and queries.
- `middlewares`: Contains the middleware functions for the HTTP handlers.
- `models`: Contains the data models.
- `routes`: Contains the route definitions.
- `services`: Contains the business logic.
- `utils`: Contains utility functions.

## Dependencies

This project uses several dependencies:

- [Gin](https://github.com/gin-gonic/gin): A HTTP web framework written in Go.
- [Golang JWT](https://github.com/golang-jwt/jwt): A Go implementation of JSON Web Tokens (JWT).
- [Go's database/sql](https://golang.org/pkg/database/sql/): Go's standard library for SQL database interaction.

## Getting Started

To run this project, you need to have Go installed on your machine. Then, you can clone this repository and run `go run app.go`.

