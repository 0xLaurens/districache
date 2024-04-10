server:
	go run cmd/cache/main.go

client:
	go run cmd/client/main.go

test:
	go test -v ./...

bench:
	go test -v ./... -bench=.