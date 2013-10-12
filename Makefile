all: get-deps test

get-deps:
	@go get -d -v ./...

test:
	@go test -v ./...
