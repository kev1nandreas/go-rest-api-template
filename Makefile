include .env
export

setup:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/server/main.go -o ./docs
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files

build-docker:
	docker compose build --no-cache

run-local:
	docker compose up -d db postgres redis
	go run cmd/server/main.go

up:
	docker compose up

down:
	docker compose down

restart:
	docker compose restart

build:
	go build -v ./...

test:
	go test -v ./... -race -cover

seed:
	go run pkg/database/seeders/main.go up

seed-clear:
	go run pkg/database/seeders/main.go down

clean:
	docker compose down db postgres redis
	rm -rf .dbdata
