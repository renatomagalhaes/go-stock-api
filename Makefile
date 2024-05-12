include .env
export $(shell sed 's/=.*//' .env)

build:
	docker-compose up -d --build

down:
	docker-compose down

test:
	docker run --rm -v ${PWD}:/app -w /app golang:1.22 go test ./...

seed:
	docker run --rm -v ${PWD}:/app -w /app golang:1.22 go run ./seeders/*.go

migrate:
	docker run --rm -v ${PWD}/migrations:/migrations -w /migrations migrate/migrate -path=/migrations -database="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up

rollback:
	docker run --rm -v ${PWD}/migrations:/migrations -w /migrations migrate/migrate -path=/migrations -database="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" down 1