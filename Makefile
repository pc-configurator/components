run:
	go run ./cmd/app

generate:
	go generate ./...

integration-test:
	go test -count=1 -v -tags=integration ./test/integration