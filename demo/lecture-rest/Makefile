deps:
	$(info Installing dependencies)
	go mod download && go mod tidy

bin:
	go build -o binary/app cmd/crud/main.go

run:
	go run cmd/crud/main.go