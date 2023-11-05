# ==================================================================================== #
# CONTAINERS up - down
# ==================================================================================== #

.PHONY: docker-up
docker-up: 
	docker-compose -f docker/docker-compose.yaml up -d

.PHONY: docker-down
docker-down:
	docker-compose -f docker/docker-compose.yaml down

# ==================================================================================== #
# MIGRATIONS up - down
# ==================================================================================== #

.PHONY: migration-up
migration-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose up

.PHONY: migration-down
migration-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose down

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...
