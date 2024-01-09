build: 
	@go build -o bin/goSeeds

run: build
	@./bin/goSeeds
test:
	@go test -v ./...