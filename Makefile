.PHONY: network postgres pgadmin run-server swagger

network:
	@docker network create go-ddd

postgres:
	@docker run --network go-ddd \
		--name postgres \
		-p 5432:5432 \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=P@ssw0rd \
		-e POSTGRES_DB=postgres \
		-d postgres:latest

pgadmin:
	@docker run --network go-ddd \
		--name pgadmin \
		-p 5050:80 \
		-e "PGADMIN_DEFAULT_EMAIL=root@example.com" \
    -e "PGADMIN_DEFAULT_PASSWORD=P@ssw0rd" \
		-d dpage/pgadmin4:latest

run-server:
	@go run ./cmd/main.go

swagger:
	@swag init --parseInternal -g ./cmd/main.go --output ./docs

docker-build:
	@docker build -t go-ddd .

docker-up:
	@docker run -d --name go-ddd -p 3030:3030 go-ddd

docker-down:
	@docker rm -f go-ddd

dev-configmap:
	@kubectl create configmap go-ddd-config --from-file=config/config.yaml --from-file=config/model.conf

skaffold-dev:
	@skaffold dev -p dev --wait-for-deletions=true --tail