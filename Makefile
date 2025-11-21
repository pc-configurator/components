run:
	go run ./cmd/app

up:
	docker compose up --build --force-recreate

generate:
	go generate ./...

integration-test:
	go test -count=1 -v -tags=integration ./test/integration