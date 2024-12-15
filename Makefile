# Define your application entry point
run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

start: build
	./bin/api

test:
	go test ./...

clean:
	rm -rf bin/
