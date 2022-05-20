BINARY_NAME=dragon_spider

build: update
	@go mod vendor
	@echo "Building Dragon Spider..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Dragon Spider built!"

run: build
	@echo "Starting Dragon Spider..."
	@./tmp/${BINARY_NAME} &
	@echo "Dragon Spider started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start_compose:
	docker-compose up -d

stop_compose:
	docker-compose down

start: run

stop:
	@echo "Stopping Dragon Spider..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Dragon Spider!"

restart: stop start