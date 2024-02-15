# Burmese Jewellery Server
Burmese Jewellery Server

---

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

---

## Misc
[Country Codes](https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes)
[Currency Codes](https://en.wikipedia.org/wiki/ISO_4217)

---

## TODO
Implement go-migration
https://github.com/golang-migrate/migrate?tab=readme-ov-file#cli-usage

---

## Tasks

### Admin Screen
- [x] /api/admin/login - POST
>>
- [x] /api/admin/account_admin - GET
- [x] /api/admin/account_admin/[id] - GET
- [x] /api/admin/account_admin/[id] - PUT
- [x] /api/admin/account_admin/[id] - DELETE
>>
- [ ] /api/admin/account - GET
- [ ] /api/admin/account/[id] - GET
>>
- [ ] /api/admin/jewellery - GET
- [ ] /api/admin/jewellery/[id] - GET
- [ ] /api/admin/jewellery/[id] - PUT
- [ ] /api/admin/jewellery/[id] - DELETE
>>
- [ ] /api/admin/currency - GET
- [ ] /api/admin/currency/[id] - GET
- [ ] /api/admin/currency/[id] - PUT
- [ ] /api/admin/currency/[id] - DELETE
>>
- [ ] /api/admin/country - GET
- [ ] /api/admin/country/[id] - GET
- [ ] /api/admin/country/[id] - PUT
- [ ] /api/admin/country/[id] - DELETE
>>
- [ ] /api/admin/faq - GET
- [ ] /api/admin/faq/[id] - GET
- [ ] /api/admin/faq/[id] - PUT
- [ ] /api/admin/faq/[id] - DELETE

### App
- [ ] /api/auth/email/register - POST
- [ ] /api/auth/email/login - POST
- [ ] /api/auth/email/otp - POST
>>
- [ ] /api/auth/phone/login - POST
- [ ] /api/auth/phone/otp - POST
>>
- [x] /api/auth/google/login - POST
- [x] /api/auth/google/callback - POST
>>
- [ ] /api/auth/facebook/login - POST
- [ ] /api/auth/facebook/callback - POST
>>
- [ ] /api/profile - POST - register profile
- [ ] /api/profile - PUT - edit profile
- [ ] /api/profile - GET
>>
- [ ] /api/jewellery - GET
- [ ] /api/jewellery/[id] - GET
>>
- [ ] /api/cart - POST - add to cart
- [ ] /api/cart - GET - cart jewellery list
>>
- [ ] /api/favourite - POST - add to favourite
- [ ] /api/favourite - GET - favourite list
>>
- [ ] /api/order - POST - confirm order
- [ ] /api/order - GET - order list
- [ ] /api/order/[id] - GET - get order info
- [ ] /api/order/cancel - POST - cancel order
>>
- [ ] /api/faq - GET
