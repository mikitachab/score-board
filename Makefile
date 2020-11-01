build:
	go build .

lint:
	golint -set_exit_status ./...
