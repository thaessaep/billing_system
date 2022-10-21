build:
	go build -v ./cmd/httpserver

swagger:
	GO111MODULE=off swagger generate spec -o ./docs/swagger.yaml --scan-model

.DEFAULT_GOAL := build