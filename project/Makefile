AUTH_BINARY=authApp
USER_BINARY=userApp
GATEWAY_BINARY=gateway
EXPENSE_BINARY=expenseApp
up:
	@echo "Starting docker containers"
	docker-compose up -d

up_build: build_auth build_user build_gateway build_expense
	@echo "Stopping containers"
	$(MAKE) down
	@echo "Building container..."
	docker-compose up -d --build

bd_user: build_user
	docker-compose stop user-service
	docker-compose rm -f user-service
	docker-compose build user-service
	docker-compose up -d user-service

bd_auth: build_auth
	docker-compose stop auth-service
	docker-compose rm -f auth-service
	docker-compose build auth-service
	docker-compose up -d auth-service

bd_gateway: build_gateway
	docker-compose stop api-gateway
	docker-compose rm -f api-gateway
	docker-compose build api-gateway
	docker-compose up -d api-gateway

bd_expense: build_expense
	docker-compose stop expense-service
	docker-compose rm -f expense-service
	docker-compose build expense-service
	docker-compose up -d expense-service

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

build_expense:
	@echo "Building expense-service"
	cd ../expense-service && GOOS=linux CGO_ENABLED=0 go build -o ${EXPENSE_BINARY} ./cmd
	@echo "Built expense binary successfully"

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