codegen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        app/auth/api/auth.proto app/users/api/users.proto

build: codegen
	docker-compose build

run: build
	docker-compose up -d

stop:
	docker-compose down --remove-orphans