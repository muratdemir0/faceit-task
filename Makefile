unit-test:
	go test ./... -v

generate-mocks:
	mockgen -destination=mocks/user/service.go -package mocks github.com/muratdemir0/faceit-task/internal/user Service
	mockgen -destination=mocks/user/repo.go -package mocks github.com/muratdemir0/faceit-task/internal/user Repository