up:
	docker compose up --build
down:
	docker compose down
lint:
	docker run --rm -v "${PWD}:/app" -w /app golangci/golangci-lint:v2.5.0 golangci-lint run
test:
	docker compose exec app go test ./...
ci: 
	lint test
tidy:
	go mod tidy