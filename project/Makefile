AUTH_BINARY=authApp
USER_BINARY=userApp
GATEWAY_BINARY=gateway
up:
	@echo "Starting docker containers"
	docker-compose up -d

up_build: build_auth build_user build_gateway
	@echo "Stopping containers"
	$(MAKE) down
	@echo "Building container..."
	docker-compose up -d --build


build_gateway:
	@echo "Building api-gateway"
	cd ../api-gateway && GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd
	@echo "Built auth-service successfully"

build_auth:
	@echo "Building auth-service"
	cd ../auth-service && go mod download && GOOS=linux CGO_ENABLED=0 go build -o $(AUTH_BINARY) ./cmd
	@echo "Built auth-service successfully"

build_user:
	@echo "Building user-service"
	cd ../user-service && GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd
	@echo "Built user binary successfully"

down:
	@echo "Stopping containers"
	docker-compose down

down_clean:
	$(MAKE) down
	$(MAKE) cleanup
cleanup:
	@echo removing binaries
	rm -rf ../auth-service/$(AUTH_BINARY)
	rm -rf ../user-service/$(USER_BINARY)