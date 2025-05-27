up:
	@echo"Starting docker containers"
	docker-compose up -d
up_build:
	@echo "Building auth-service"
	cd ../auth-service && GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd
	@echo "Built auth-service successfully"
	@echo "Stopping containers"
	$(MAKE) down
	@echo "Building container..."
	docker-compose up -d --build
down:
	@echo "Stopping containers"
	docker-compose down