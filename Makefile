env:
	cp .env.example .env

run:
	go run main.go

test:
	go test -v