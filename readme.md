<h1 align="center">Golang REST API Template</h1>

<p align="center">
  <a href="https://opensource.org/licenses/mit-license.php"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103" alt="Report"></a>
  <a href="#"><img src="https://img.shields.io/badge/godoc-reference-brightgreen.svg" alt="Docs"></a>
  <a href="#"><img src="https://img.shields.io/badge/version-0.0.1-brightgreen.svg" alt="Version"></a>
</p>


Template for building a REST API in Go

- Project structure follows to https://github.com/golang-standards/project-layout
- HTTP REST server
- CRUDL Users
- Authentication
- Dockerized using docker-compose
- Chi Router
- PostgreSQL
- SQLX db driver
- Database Migrations
- Logging
- Config and env var management
- Postman collection

## Information
- database migrations can be found in the `/schemas` dir
- database migrations will run automatically if db does not exist
- `config/secrets.json` will be loaded by app on init
- if `config/secrets.json` does not exist the app will try pull them from AWS Secreets Mannager service

## To build and run go app and postgres in docker containers

```$ make docker-compose```

## How to use API
The API can be accessed from: http://127.0.0.1:8080

Request:
```
curl --location --request GET "127.0.0.1:8080/v1/health"
```
