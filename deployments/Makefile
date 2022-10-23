build:
	go mod tidy
	go build -v ../cmd/httpserver

swagger:
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-model

docker-compose:
	docker-compose up --build app

docker-run:
	docker-compose up app

.DEFAULT_GOAL := build