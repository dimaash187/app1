build:
	go build -o main main.go

build-docker:
	docker build . -t app1 --no-cache

run:
	go run main.go

test: ## run our tests [not implemeted yet]
	go test -v