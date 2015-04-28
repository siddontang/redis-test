all: build

build:
	go build -o ./bin/redis-test ./...

clean:
	rm -rf ./bin