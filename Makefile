BIN := "./bin/primes"

build:
	go build -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/...

run: build
	HTTP_HOST="localhost" HTTP_PORT="8000" $(BIN)

test:
	go test -v ./internal/...

up:
	sudo docker-compose -f ./deployments/docker-compose.yaml up --build

down:
	sudo docker-compose -f ./deployments/docker-compose.yaml down

integration-tests:
	set -e ;\
	sudo docker-compose -f ./deployments/docker-compose.test.yaml up --build -d ;\
	test_status_code=0 ;\
	sudo docker-compose -f ./deployments/docker-compose.test.yaml run integration-tests go test -v || test_status_code=$$? ;\
	sudo docker-compose -f ./deployments/docker-compose.test.yaml down ;\
	exit $$test_status_code ;\

.PHONY: build run test integration-tests