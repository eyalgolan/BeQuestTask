install: ## build docker containers.
	docker-compose -f deployment/docker-compose.yml up --build
start: ## start docker containers.
	docker-compose -f deployment/docker-compose.yml up --no-build
stop:  ## Stop docker containers.
	docker-compose -f deployment/docker-compose.yml down
clean: ## clear artefacts.
	docker system prune