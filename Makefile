build:
	go build .

lint:
	golint -set_exit_status ./...

test:
	go test ./... -v

mocks: FORCE
	mockgen -destination=mocks/mock_repository.go -package=mocks github.com/mikitachab/score-board/db RepositoryInterface

FORCE:
