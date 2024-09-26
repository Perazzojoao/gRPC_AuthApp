AUTH_BINARY=authApp.exe

## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_auth
	@echo Stopping docker images (if running...)
	docker compose down
	@echo Building (when required) and starting docker images...
	docker compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker compose down
	@echo Done!

## proto_auth: generates auth proto files
proto_auth:
	@echo Generating auth proto...
	cd auth-service && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/server.proto
	@echo Done!

## build_auth: builds auth binary
build_auth:
	@echo Building authentication binary...
	cd auth-service && go build -o bin/${AUTH_BINARY} .
	@echo Done!

## start: starts the authApp
start: build_auth
	@echo Starting authApp
	cd auth-service/bin && ${AUTH_BINARY}

## test_auth: runs auth tests
test_auth:
	@echo Running auth tests...
	cd auth-service && go test -timeout 30s -v ./...
	@echo Done!

## test_all: runs all tests
test_all: test_auth 
