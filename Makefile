up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_gateway
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"
	docker image prune -f

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_gateway:
	@echo "Building gateway binary..."
	cd ./Gateway && env GOOS=linux CGO_ENABLED=0 go build -o gatewayApp ./cmd
	@echo "Done!"