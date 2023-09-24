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
#### ER diagram
```
cd database/plantuml
planter "host=localhost port=5432 user=postgres password=postgres dbname=burmese_jewellery sslmode=disable" -T burmese_jewellery -o ./database/plantuml/burmese_jewellery.uml
java -jar plantuml.jar -verbose example.uml
```
```
go build . &&  ./planter "host=localhost port=5432 user=postgres password=postgres dbname=burmese_jewellery sslmode=disable" -T burmese_jewellery -o ./burmese_jewellery.uml
cp ./planter /Users/naylinhtet/.goenv/shims  
```

---

## Git Branch Strategy
![Git Branch Strategy](https://github.com/Nlhmmh/howto/blob/master/git_branch_strategy.png)

