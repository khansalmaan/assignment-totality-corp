# Define variables for the image name and tag
IMAGE_NAME := go-assignment
IMAGE_TAG := latest
DOCKER_FILE := Dockerfile

# Build the Docker image
.PHONY: build
build:
	@echo "Building Docker image..."
	docker build -f $(DOCKER_FILE) -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Run the Docker container
.PHONY: run
run:
	@echo "Running Docker container..."
	docker run -p 50051:50051 $(IMAGE_NAME):$(IMAGE_TAG)

# Build and run the Docker container
.PHONY: dev
dev: build run

# Clean up Docker images and containers
.PHONY: clean
clean:
	@echo "Cleaning up Docker images and containers..."
	docker stop $$(docker ps -q --filter ancestor=$(IMAGE_NAME):$(IMAGE_TAG)) || true
	docker rm $$(docker ps -a -q --filter ancestor=$(IMAGE_NAME):$(IMAGE_TAG)) || true
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG) || true

# Build the Docker image, run the container, and then clean up
.PHONY: full
full: build run clean

# Display the status of Docker images and containers
.PHONY: status
status:
	@echo "Listing Docker images..."
	docker images
	@echo "Listing Docker containers..."
	docker ps -a

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Show test coverage
.PHONY: coverage
coverage:
	@echo "Running test coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out