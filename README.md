# Burmese Jewellery Server
Burmese Jewellery Server

---

## Development
- Golang HTTP Server

#### Server
- Framework - Gin Server - https://github.com/gin-gonic/gin
- ORM - SQL Boiler - https://github.com/volatiletech/sqlboiler
- Swagger OpenAPI generator - https://github.com/deepmap/oapi-codegen
- Hot Reload - air - https://github.com/cosmtrek/air
#### API Documentation
- Swagger
#### Database
- Postgres
- PlantUML ER Diagram
#### Container
- Docker
#### Infra
- AWS?
#### Cerfificate SSL
```
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout key.pem -out cert.pem
```
#### Healthcheck
```
curl --insecure https://localhost:8077/api/health_check
```

---

## Git Branch Strategy
![Git Branch Strategy](https://raw.githubusercontent.com/Nlhmmh/burmese_jewellery/main/git_branch_strategy.png?token=GHSAT0AAAAAACGEKCVQVCVNYBXKUASSR2L4ZHHBDZQ)

