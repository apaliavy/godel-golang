DB_USERNAME ?= postgres
DB_PASSWORD ?= mysecretpassword
DB_HOST ?= localhost
DB_PORT ?= 5432

deps:
	$(info Installing dependencies)
	GO111MODULE=on go mod download

###########################
### MIGRATION UTILITIES ###
###########################
goose-status:
	goose postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" status

goose-migrate:
	goose -dir tools/goose/migrations postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

goose-degrade:
	goose -dir tools/goose/migrations postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down

goose-create:
	goose -dir tools/goose/migrations create $(FILENAME) sql

goose-fixtures:
	goose -dir tools/goose/fixtures postgres "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up
############################


############################
# INTEGRATION TESTS
############################
build_integration_tests_containers:
	$(info building containers for integration testing)
	docker-compose build
	docker-compose up -d

seed_db_for_integration_tests:
	# users database
	goose -dir tools/goose/migrations postgres "postgresql://postgres:mysecretpassword@localhost:5433/users?sslmode=disable" up
	goose --no-versioning -dir tools/goose/fixtures postgres "postgresql://postgres:mysecretpassword@localhost:5433/users?sslmode=disable" up
