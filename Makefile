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

.PHONY: run
run: ## Run server
	make dk-up-db
	sh ./tools/server_run.sh

.PHONY: merge
merge: ## Merge swagger into single file
	cd ./swagger && merger -f index.yml -o swagger.yml

.PHONY: gen
gen: ## Generate go code from Swagger
	make merge
	go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.4
	go generate $(GO_GENERATE_FILE)

.PHONY: boil
boil: ## Generate sqlboiler code (mac)
	sqlboiler --wipe --add-global-variants --add-enum-types --no-tests --no-back-referencing --add-soft-deletes --struct-tag-casing=camel -p boiler psql

# Docker

.PHONY: dk-up
dk-up: ## Docker Up
	docker compose up --build -d

.PHONY: dk-down
dk-down: ## Docker Down
	docker compose down

.PHONY: dk-tail-%
dk-tail-%: ## Docker tail (server, db)
	docker compose logs ${@:dk-tail-%=%} -f

.PHONY: dk-up-%
dk-up-%: ## Docker Up (server, db)
	docker compose up --build -d ${@:dk-up-%=%}

.PHONY: dk-down-%
dk-down-%: ## Docker Down (server, db)
	docker compose down ${@:dk-down-%=%}

.PHONY: dk-exec-%
dk-exec-%: ## docker compose exec (server, db) bash
	docker compose exec ${@:dk-exec-%=%} bash

.PHONY: dk-connect-db
dk-connect-db: ## Docker Connect DB
	docker compose exec db psql -U postgres -p 5432 -d burmese_jewellery