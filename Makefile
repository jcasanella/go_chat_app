MAIN_PACKAGE_PATH := .
BIN_NAME := chatApp
BIN_FOLDER := bin
TMP_FOLDER := tmp

pre: docker-up migration-up
all: clean test build copy/config run


# ==================================================================================== #
# CONTAINERS up - down
# ==================================================================================== #

.PHONY: docker-up
docker-up: 
	docker-compose -f docker/docker-compose.yaml up -d

.PHONY: docker-down
docker-down:
	docker-compose -f docker/docker-compose.yaml down

.PHONY: docker-rm-volume
docker-rm-volume: docker-down
	docker volume rm docker_postgres_data

# ==================================================================================== #
# MIGRATIONS up - down
# ==================================================================================== #

.PHONY: migration-up
migration-up: docker-up
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose up

.PHONY: migration-down
migration-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5435/chat_development?sslmode=disable" -verbose down

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

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=${TMP_FOLDER}/coverage.out ./...
	go tool cover -html=${TMP_FOLDER}/coverage.out

## build: build the application
.PHONY: build
build: test
	@echo "Build app into ${BIN_FOLDER}/${BIN_NAME}"
	go build -o=${BIN_FOLDER}/${BIN_NAME} ${MAIN_PACKAGE_PATH}

## copy config file
.PHONY: copy/config
copy/config:
	@echo "Copy confing.yaml into ${BIN_FOLDER}"
	cp config.yaml ${BIN_FOLDER}/.

## run: run the  application
.PHONY: run
run: build
	@echo "Run binary from ${BIN_FOLDER}/${BIN_NAME}"
	${BIN_FOLDER}/${BIN_NAME}

.PHONY: clean
clean:
	@echo "Removing ${BIN_FOLDER}"
	rm -rf ${BIN_FOLDER}

