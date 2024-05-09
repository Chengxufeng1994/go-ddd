.PHONY: postgres run swagger

postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=P@ssw0rd -d postgres:latest

run:
	@go run ./cmd/main.go

swagger:
	@swag init --parseInternal -g ./cmd/main.go --output ./docs

docker-build:
	docker build -t go-ddd .

docker-up:
	docker run -d --name go-ddd -p 3030:3030 go-ddd

docker-down:
	docker rm -f go-ddd