cleanup:
	rm -rf build
build: cleanup
	go build -o build/explorer cmd/main.go
test:
	go test -v -cover ./...