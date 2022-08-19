run:
	docker-compose up

dev:
	go run main.go

test:
	go test ./...

coverage:
	go test -cover ./...