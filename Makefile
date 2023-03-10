build:
	docker compose build
init-air:
	docker compose run --rm votify-api air init
run:
	docker compose up -d
down:
	docker compose down
stop:
	docker compose --project-name 'votify-api' stop