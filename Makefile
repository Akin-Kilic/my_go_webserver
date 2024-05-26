.PHONY: build hot down

build:
	docker compose -f docker-compose-hot.yml build
hot:
	docker compose -f docker-compose-hot.yml up --build
down:
	docker compose -f docker-compose-hot.yml down