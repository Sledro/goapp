# Builds the app to ./build/goapp
.PHONY: build

build: 
	go build -o build/goapp cmd/goapp/main.go

# Runs the app located in ./build/goapp
.PHONY:
run:
	build/goapp

# Builds the app docker as a image
.PHONY:
docker-build:
	docker build --tag goapp .