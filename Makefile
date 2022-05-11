.PHONY: gen
gen:
	protoc --go_out=pkg --go_opt=paths=source_relative --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/v1/api.proto

repomock:
	minimock -i gitlab.ozon.dev/emilgalimov/homework-2/internal/app.Repository -o ./internal/app/repository_mock_test.go -n RepositoryMock