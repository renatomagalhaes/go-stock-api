build:
	@docker-compose up -d --build

down:
	@docker-compose down

test:
	@docker run --rm -v ${PWD}:/app -w /app golang:1.22 go test ./...