build-app:
	@go build -o ./bin/app ./...
	@chmod +x ./bin/app

app: build-app
	@./bin/app

# "..." : A special wildcard token in the Go 
# toolchain meaning "match all subdirectories 
# recursively"
format: # automatically formats every go file in your codebase
	@go format ./... 

vet: # scans the entire project for structural and semantic errors.
	@go vet ./...



test:
	@go test ./...


test-app-race:
	@go clean -testcache
	@go test -race -v ./...