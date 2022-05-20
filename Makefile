.PHONY: gen
gen:
	buf generate

repomock:
	minimock -i gitlab.ozon.dev/emilgalimov/homework-2/internal/app.Repository -o ./internal/app/repository_mock_test.go -n RepositoryMock

run:
	go mod tidy
	go run cmd/server/main.go