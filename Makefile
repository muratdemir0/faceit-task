run:
	APP_ENV=dev go run ./cmd/faceit-task/main.go
build:
	go build ./cmd/faceit-task/main.go
unit-test:
	go test ./... -v

generate-mocks:
	mockgen -destination=mocks/user/service.go -package mocks github.com/muratdemir0/faceit-task/internal/user Service
	mockgen -destination=mocks/user/store.go -package mocks github.com/muratdemir0/faceit-task/internal/user Store