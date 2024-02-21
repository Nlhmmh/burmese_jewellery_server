GO_GENERATE_FILE := $(CURDIR)/generate/generate.go

.PHONY: help
help: ## Help
	@echo 'Usage:'
	@echo "  make [command]"
	@echo 'Commands:'
	@grep -E '^[a-zA-Z_0-9%-]+:.*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?##"}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Prepare Environment
	pip install -r requirements.txt
	go install github.com/cosmtrek/air@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	brew install swagger-codegen

.PHONY: run
run: ## Run server
	make dk-up-db
	sh ./tools/server_run.sh

.PHONY: lint
lint: ## Lint
	golangci-lint run

.PHONY: precommit
precommit: ## precommit
	pre-commit run

.PHONY: merge
merge: ## Merge swagger into single file
	cd ./swagger && merger -f index.yml -o swagger.yml

.PHONY: gen
gen: ## Generate go code from Swagger
	make merge
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go generate $(GO_GENERATE_FILE)
	go mod tidy
	make gen-swagger

.PHONY: gen-er
gen-er: ## Generate temporary plantuml from db
	./tools/planter "host=localhost port=5432 user=postgres password=postgres dbname=burmese_jewellery sslmode=disable" -T burmese_jewellery -o ./database/plantuml/burmese_jewellery.pu.tmp

.PHONY: gen-swagger
gen-swagger: ## Generate html from Swagger
	cd ./swagger && swagger-codegen generate -i swagger.yml -l html -o static

.PHONY: boil
boil: ## Generate sqlboiler code (mac)
	sqlboiler --wipe --add-global-variants --add-enum-types --no-tests --no-back-referencing --add-soft-deletes --struct-tag-casing=camel psql

.PHONY: copy-env
copy-env: ## Copy .env.local as .env
	cp .env.local .env

# Docker

.PHONY: dk-up
dk-up: ## Docker Up
	docker compose up --build -d

.PHONY: dk-down
dk-down: ## Docker Down
	docker compose down

.PHONY: dk-tail-%
dk-tail-%: ## Docker tail (server, db, pgadmin)
	docker compose logs ${@:dk-tail-%=%} -f

.PHONY: dk-reload-%
dk-reload-%: ## Docker Reload (server, db, pgadmin)
	make dk-down-${@:dk-reload-%=%}
	make dk-up-${@:dk-reload-%=%}

.PHONY: dk-up-%
dk-up-%: ## Docker Up (server, db, pgadmin)
	docker compose up --build -d ${@:dk-up-%=%}

.PHONY: dk-down-%
dk-down-%: ## Docker Down (server, db, pgadmin)
	docker compose down ${@:dk-down-%=%}

.PHONY: dk-exec-%
dk-exec-%: ## docker compose exec (server, db, pgadmin) bash
	docker compose exec ${@:dk-exec-%=%} bash

.PHONY: dk-connect-db
dk-connect-db: ## Docker Connect DB
	docker compose exec db psql -U postgres -p 5432 -d burmese_jewellery
