start:
	go run cmd/main.go

build:
	docker build -t dan-dawson-onboarding .

docker-start:
	make build && docker compose up -d

docker-stop:
	docker ps -aq | xargs docker stop | xargs docker rm

all:
	go mod tidy && make test
