build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o dist/matheus-darwin-amd64 cmd/cli/main.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o dist/matheus-darwin-arm64 cmd/cli/main.go

# build-linux-386:
# 	GOOS=linux GOARCH=386 go build -o dist/matheus cmd/cli/main.go

# build-linux-amd64:
# 	GOOS=linux GOARCH=amd64 go build -o dist/matheus cmd/cli/main.go

# build-linux-arm:
# 	GOOS=linux GOARCH=arm go build -o dist/matheus cmd/cli/main.go

# build-linux-arm64:
# 	GOOS=linux GOARCH=arm64 go build -o matheus cmd/cli/main.go

.PHONY: go
