run:
	go run internal/main.go

build:
	go build -o bin/go-rest-api internal/main.go

deps:
	go get -d ./internal

swagger.validate:
	swagger validate package/swagger/swagger.yml

swagger.gen:
	GOFLAGS=-mod=mod go generate ./internal ./package/swagger
