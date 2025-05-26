up_build:
	@echo "Building auth-service"
	cd ../auth-service && GOOS=linux CGO_ENABLED=0 go build -o authApp ./cmd
	@echo "Built auth-service successfully"
	@echo "Building container..."
	cd ../auth-service && docker build -f auth-service.dockerfile -t auth-service:latest . && docker run -d  --name auth-service -p 80:8082 auth-service:latest
stop:
	@echo "Stopping containers"
	docker stop auth-service
	docker rm auth-service