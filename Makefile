.PHONY: run build clean

run:
	go run cmd/bot/main.go

build:
	go build -o bin/salary-bot cmd/bot/main.go

clean:
	rm -rf bin/ data/calendar_cache.json

deps:
	go mod download
	go mod tidy