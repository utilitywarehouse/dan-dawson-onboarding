start:
	go run cmd/main.go

stop:
	docker ps -aq | xargs docker stop | xargs docker rm

build:
	docker build -t dan-dawson-onboarding .

docker-start:
	make build && docker compose up -d