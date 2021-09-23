SHELL := /bin/bash
.DEFAULT_GOAL := help

# https://gist.github.com/tadashi-aikawa/da73d277a3c1ec6767ed48d1335900f3
.PHONY: $(shell grep -h -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://')

start: ## start server
	make -C demo_app/
	docker-compose up --build --force-recreate -d

start_debug: ## start with debug
	make -C demo_app/
	docker-compose up --build --force-recreate

open: ## open browser
	firefox https://127.0.0.1/

clean: check_clean ## Clean docker containers, images, and volumes
	docker rm $(docker ps -a -q)
	docker image rm $(docker image ls -q)
	docker volume rm $(docker volume ls -q)

check_clean:
	@echo -n -e "\033[0;33m" "Are you sure to clean docker containers, images, and volumes? [y/N] " && read ans && [ $${ans:-N} = y ]

# https://postd.cc/auto-documented-makefile/
help: ## Show help message
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
