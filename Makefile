GO_GENERATE_FILE := $(CURDIR)/generate/generate.go

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## 開発環境セットアップ
	pip install -r requirements.txt
	go install github.com/cosmtrek/air@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

.PHONY: run
run: ## Run server
	make docker-up-db
	sh ./tools/server_run.sh

.PHONY: merge
merge: ## Merge swagger into single file
	cd ./swagger && merger -f index.yml -o swagger.yml

.PHONY: gen
gen: ## Generate go code from Swagger
	make merge
	go generate $(GO_GENERATE_FILE)

.PHONY: boil
boil: ## Generate sqlboiler code (mac)
	sqlboiler --wipe --add-global-variants --add-enum-types --no-tests --no-back-referencing --add-soft-deletes --struct-tag-casing=camel -p boiler psql

# Docker

.PHONY: docker-up
docker-up: ## Docker Up
	docker compose up --build -d

.PHONY: docker-down
docker-down: ## Docker Up
	docker compose down

.PHONY: docker-tail
docker-tail: ## Docker tail
	docker compose logs -f

.PHONY: docker-up-server
docker-up-server: ## Docker Up Server
	docker compose up --build -d server

.PHONY: docker-down-server
docker-down-server: ## Docker Down Server
	docker compose down server

.PHONY: docker-up-db
docker-up-db: ## Docker up DB
	docker compose up --build -d database
	
.PHONY: docker-down-db
docker-down-db: ## Docker Down DB
	docker compose down database

.PHONY: docker-exec-server
docker-exec-server: ## docker compose exec server bash を実行する
	docker compose exec burmese-jewellery bash