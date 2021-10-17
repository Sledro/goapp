<h1 align="center">Golang REST API Template</h1>

<p align="center">
  <a href="https://opensource.org/licenses/mit-license.php"><img src="https://badges.frapsoft.com/os/mit/mit.svg?v=103" alt="Report"></a>
  <a href="#"><img src="https://img.shields.io/badge/godoc-reference-brightgreen.svg" alt="Docs"></a>
  <a href="#"><img src="https://img.shields.io/badge/version-0.0.1-brightgreen.svg" alt="Version"></a>
</p>


Template for building a REST API in Go

- Project structure adheres to https://github.com/golang-standards/project-layout
- Authentication
- Database Migrations
- Logging
- http REST server
- Gorm and gin frameworks avoided for full control
- Config and env var management

## To build and run go app and postgres in docker containers

```$ make docker-compose```

Database migrations will run automatically if db does not exist

## How to use API
The API can be accessed from: http://127.0.0.1:8080

Request:
```
curl --location --request GET "127.0.0.1:8080/v1/api/health"
```
