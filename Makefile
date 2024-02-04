APP_NAME := sse-openai-server

build:
	go build -o $(APP_NAME) .

run: build
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

deps:
	go get -u ./...

test:
	go test -v ./...

.DEFAULT_GOAL := run