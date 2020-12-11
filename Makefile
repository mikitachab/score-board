build:
	go build .

lint:
	golint -set_exit_status ./...

test:
	go test ./server -v

mocks: FORCE
	mockgen -destination=mocks/mock_repository.go -package=mocks github.com/mikitachab/score-board/db RepositoryInterface
	mockgen -destination=mocks/mock_template.go -package=mocks github.com/mikitachab/score-board/templateloader TemplateInterface

FORCE:
