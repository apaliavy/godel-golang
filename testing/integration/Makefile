TOOLS += github.com/maxbrunsfeld/counterfeiter/v6

.PHONY: $(TOOLS)
$(TOOLS): %:
	GOBIN=$(GOBIN) go install $*

.PHONY: tools
tools: deps $(TOOLS)

mocks:
	@go generate ./...

test:
	go test --cover --race ./client/...

############################
# INTEGRATION TESTS
############################
build_integration_tests_containers:
	$(info building containers for integration testing)
	docker-compose build
	docker-compose up -d

seed_db_for_integration_tests:
	goose -dir tools/goose/migrations/prices postgres "postgresql://postgres:mysecretpassword@localhost:5432/prices?sslmode=disable" up
	goose --no-versioning -dir tools/goose/fixtures/prices postgres "postgresql://postgres:mysecretpassword@localhost:5432/prices?sslmode=disable" up

integration_tests: build_integration_tests_containers seed_db_for_integration_tests
	go test -v -tags=integration ./client/...
	docker-compose down
############################
