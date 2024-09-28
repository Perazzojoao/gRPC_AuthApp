AUTH_BINARY=authService.exe
MAIL_BINARY=mailService.exe

#1 ------------------------- Docker commands -------------------------
## up: starts all containers in the background without forcing build
up:
	@echo Starting Docker images...
	docker compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: 
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

## clear: clear dungling docker images
clear:
	@echo Clearing dungling docker images...
	docker image prune
	@echo Done!

#2 ------------------------- Auth build commands -------------------------
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

## run_auth: runs auth binary
run_auth: build_auth
	@echo Running authentication binary...
	cd auth-service/bin && ${AUTH_BINARY}
	@echo Done!


#3 ------------------------- Mail build commands -------------------------
## proto_mail: generates mail proto files
proto_mail:
	@echo Generating mail proto...
	cd mail-service && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/server.proto
	@echo Done!

## build_mail: builds mail binary
build_mail:
	@echo Building mail binary...
	cd mail-service && go build -o bin/${MAIL_BINARY} .
	@echo Done!

## run_mail: runs mail binary
run_mail: build_mail
	@echo Running mail binary...
	cd mail-service/bin && ${MAIL_BINARY}
	@echo Done!


#4 ------------------------- Test commands -------------------------
## test_auth: runs auth tests
test_auth:
	@echo Running auth tests...
	cd auth-service && go test -timeout 60s -v ./...
	@echo Done!

## test_all: runs all tests
test_all: test_auth 