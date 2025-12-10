build: 
	@go build -o bin/gitgen

run: build
	@./bin/gitgen

test:
	@go test -v ./...