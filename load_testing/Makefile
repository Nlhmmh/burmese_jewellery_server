.PHONY: help
help: ## Help
	@echo 'Usage:'
	@echo "  make [command]"
	@echo 'Commands:'
	@grep -E '^[a-zA-Z_0-9%-]+:.*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?##"}; {printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Prepare Environment
	pip3 install locust

.PHONY: run
run: ## Run Locust
	locust -f ./locustfile.py --processes 4 --users 30 --spawn-rate 5 -H http://local.8077

.PHONY: run-headless
run-headless: ## Run Locust Headless
	locust -f ./locustfile.py --processes 4 --headless --users 30 --spawn-rate 5 -H http://local.8077
