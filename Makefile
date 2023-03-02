generate:
	go mod tidy
	rm -f ./internal/generated/resolver.go
	go run github.com/99designs/gqlgen generate --config ./configs/gqlgen.yml


dev:
	go run cmd/main.go