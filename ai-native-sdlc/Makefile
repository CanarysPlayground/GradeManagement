.PHONY: run

run:
	cd server/grades-service && go run ./cmd/grades-service

compose-up:
	docker compose -f infra/docker-compose.yml up --build
