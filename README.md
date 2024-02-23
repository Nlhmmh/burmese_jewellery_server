# Burmese Jewellery Server
Burmese Jewellery Server

---

## Documentation
https://naylinhtet.atlassian.net/wiki/spaces/BJ1/pages/4980744/Backend
https://orange-moon-6960.postman.co/workspace/Burmese-Jewellery~fc558fc1-9818-4d2a-8307-4cbfb453797a/overview

## Development
- Golang HTTP Server
```
make dk-up-db
make dk-up-pgadmin
make run

go mod edit -go=1.21
```

#### Server
- Framework - Gin Server - https://github.com/gin-gonic/gin
- ORM - SQL Boiler - https://github.com/volatiletech/sqlboiler
- Swagger OpenAPI generator - https://github.com/deepmap/oapi-codegen
- Hot Reload - air - https://github.com/cosmtrek/air
- Load Testing - https://locust.io/
#### API Documentation
- Swagger
- https://swagger.io/docs/specification/basic-structure/
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
#### Load Testing
[Run with Makefile](load_testing/Makefile)

---

## Git Branch Strategy
![Git Branch Strategy](https://github.com/Nlhmmh/howto/blob/master/git_branch_strategy.png)

---

## Misc
[Country Codes](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes)
[Currency Codes](https://en.wikipedia.org/wiki/ISO_4217)

---

## TODO
Implement go-migration
https://github.com/golang-migrate/migrate?tab=readme-ov-file#cli-usage

Test Coverage
