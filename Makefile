docker-up: 
	docker-compose -f docker/docker-compose.yaml up -d

docker-down:
	docker-compose -f docker/docker-compose.yaml down

migration-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose up

migration-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose down

.PHONY: docker-up docker-down migration-up migration-down
