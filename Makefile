start:
	go run cmd/main.go

test:
	go test ./handler -v

build:
	docker build -t dan-dawson-onboarding .

docker-start:
	make build && docker compose up -d

docker-stop:
	docker ps -aq | xargs docker stop | xargs docker rm
