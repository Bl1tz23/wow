PROJECT_PATH=$(CURDIR)

lint:
	docker run --rm --volume="${PROJECT_PATH}:/go/src/wow" -w /go/src/wow golangci/golangci-lint:v1.50-alpine golangci-lint run -E gofmt --skip-dirs=./vendor --deadline=10m

build_server:
	@docker build -f deploy/Dockerfile.server -t wow_server .

build_client:
	@docker build -f deploy/Dockerfile.client -t wow_client .

run_server:
	@docker run -d --name wow_server_1 --rm wow_server

run_client:
	@docker run --name wow_client_1 --network="container:wow_server_1" --rm wow_client

run: build_server build_client run_server run_client

.PHONY: lint build_server build_client run_server run_client run