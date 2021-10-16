.PHONY: build swagger

build: 
	go build -o build/application cmd/golang-framework/main.go

run:
	build/application

check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	GO111MODULE=on go mod vendor && GO111MODULE=off swagger generate spec -o ./api/swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger api/swagger.yaml

docker-build:
	docker build --tag golang-framework .