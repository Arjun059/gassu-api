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

dev:
	CompileDaemon --build="go build -o gassu.exe cmd/api/main.go" --command="./gassu.exe" --polling