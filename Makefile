start:
	go run cmd/main.go

build:
	docker build -t dan-dawson-onboarding .

build-start:
	make build && docker compose up -d