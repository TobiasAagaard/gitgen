build: 
	@go build -o bin/gitgen
run: build
	@./bin/gitgen
test: build
	@go test -v  ./...