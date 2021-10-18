# Builds the app to ./build/goapp
.PHONY: build

build: 
	go build -o build/goapp cmd/goapp/main.go

# Runs the app located in ./build/goapp
.PHONY:
run:
	build/goapp

# Runs unit tests
.PHONY:
test:
	go test ./... -race -coverprofile=coverage.out

# Builds the app docker as a image
.PHONY:
docker-build:
	docker build --tag goapp .

# Builds the local docker compose env
.PHONY:
docker-compose:
	docker-compose -f "deployments/docker-compose.yml" up -d --build